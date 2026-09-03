package gc

import (
	"testing"
)

func TestPreservingObjects(t *testing.T) {
	const preservedObjectsAmount = 2

	vm := NewVM()

	PushInt(vm, 1)
	PushInt(vm, 2)

	MarkAndSweep(vm)
	if vm.NumObjects != preservedObjectsAmount {
		t.Errorf("Expected %d preserved objects, got %d.\n", preservedObjectsAmount, vm.NumObjects)
	}

	FreeVM(vm)
}

func TestCollectingUnreachedObjects(t *testing.T) {
	vm := NewVM()

	PushInt(vm, 1)
	PushInt(vm, 2)

	vm.Pop()
	vm.Pop()

	MarkAndSweep(vm)
	if vm.NumObjects != 0 {
		t.Errorf("Expected no currently allocated objects, got %d.\n", vm.NumObjects)
	}

	FreeVM(vm)
}

func TestReachingNestedObjects(t *testing.T) {
	const reachedObjects = 7

	vm := NewVM()

	PushInt(vm, 1)
	PushInt(vm, 2)

	PushPair(vm)

	PushInt(vm, 3)
	PushInt(vm, 4)

	PushPair(vm)
	PushPair(vm)

	MarkAndSweep(vm)

	if vm.NumObjects != reachedObjects {
		t.Errorf("Expected %d reached objects, got %d.\n", reachedObjects, vm.NumObjects)
	}

	FreeVM(vm)
}

func TestHandlingCycles(t *testing.T) {
	const reachedObjects = 4
	
	vm := NewVM()

	PushInt(vm, 1)
	PushInt(vm, 2)
	
	firstPair := PushPair(vm)

	PushInt(vm, 3)
	PushInt(vm, 4)

	secondPair := PushPair(vm)

	firstPair.(*PairObject).Tail = secondPair
	secondPair.(*PairObject).Tail = firstPair

	MarkAndSweep(vm)

	if vm.NumObjects != reachedObjects {
		t.Errorf("Expected %d reached objects, got %d.\n", reachedObjects, vm.NumObjects)
	}

	FreeVM(vm)
}

func BenchmarkPushPop(b *testing.B) {
	vm := NewVM()

	for i := range b.N {
		for range 20 {
			PushInt(vm, i)
		}

		for range 20 {
			vm.Pop()
		}
	}

	FreeVM(vm)
}