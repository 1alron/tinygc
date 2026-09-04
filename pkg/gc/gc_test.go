package gc

import (
	"testing"
)

func TestPreservingObjects(t *testing.T) {
	const preservedObjectsAmount = 2

	vm := NewVM()
	defer FreeVM(vm)

	vm.PushInt(1)
	vm.PushInt(2)

	MarkAndSweep(vm)
	if vm.NumObjects != preservedObjectsAmount {
		t.Errorf("expected %d preserved objects, got %d.\n", preservedObjectsAmount, vm.NumObjects)
	}
}

func TestCollectingUnreachedObjects(t *testing.T) {
	vm := NewVM()
	defer FreeVM(vm)

	vm.PushInt(1)
	vm.PushInt(2)

	vm.Pop()
	vm.Pop()

	MarkAndSweep(vm)
	if vm.NumObjects != 0 {
		t.Errorf("expected no currently allocated objects, got %d\n", vm.NumObjects)
	}
}

func TestReachingNestedObjects(t *testing.T) {
	const reachedObjects = 7

	vm := NewVM()
	defer FreeVM(vm)

	vm.PushInt(1)
	vm.PushInt(2)

	vm.PushPair()

	vm.PushInt(3)
	vm.PushInt(4)

	vm.PushPair()
	vm.PushPair()

	MarkAndSweep(vm)

	if vm.NumObjects != reachedObjects {
		t.Errorf("expected %d reached objects, got %d\n", reachedObjects, vm.NumObjects)
	}
}

func TestHandlingCycles(t *testing.T) {
	const reachedObjects = 4

	vm := NewVM()
	defer FreeVM(vm)

	vm.PushInt(1)
	vm.PushInt(2)

	firstPair := vm.PushPair()

	vm.PushInt(3)
	vm.PushInt(4)

	secondPair := vm.PushPair()

	firstPair.(*PairObject).Tail = secondPair
	secondPair.(*PairObject).Tail = firstPair

	MarkAndSweep(vm)

	if vm.NumObjects != reachedObjects {
		t.Errorf("expected %d reached objects, got %d\n", reachedObjects, vm.NumObjects)
	}
}

func BenchmarkPushPop(b *testing.B) {
	vm := NewVM()
	defer FreeVM(vm)

	for i := range b.N {
		for range 20 {
			vm.PushInt(i)
		}

		for range 20 {
			vm.Pop()
		}
	}
}
