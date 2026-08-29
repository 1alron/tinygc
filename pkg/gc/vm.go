package gc

/*
#include <stdlib.h>
*/
import "C"
import "unsafe"

const StackMax = 256

type VM struct {
	Stack     [StackMax]*Object
	StackSize int
}

func NewVM() *VM {
	ptr := C.malloc(C.size_t(unsafe.Sizeof(VM{})))
	if ptr == nil {
		panic("Failed to allocate memory for VM!")
	}
	vm := (*VM)(ptr)
	vm.StackSize = 0
	return vm
}

func FreeVM(vm *VM) {
	C.free(unsafe.Pointer(vm))
}

func (vm *VM) Push(value *Object) {
	if vm.StackSize >= StackMax {
		panic("VM's stack overflow!")
	}
	vm.Stack[vm.StackSize] = value
	vm.StackSize++
}

func (vm *VM) Pop() *Object {
	if vm.StackSize <= 0 {
		panic("VM's stack underflow!")
	}
	res := vm.Stack[vm.StackSize-1]
	vm.StackSize--
	return res
}

func PushInt(vm *VM, intValue int) {
	obj := NewObject(vm, ObjInt)
	intObj := obj.(*IntObject)
	intObj.Value = intValue
	vm.Push(&obj)
}

func PushPair(vm *VM) Object {
	obj := NewObject(vm, ObjPair)
	pairObj := obj.(*PairObject)
	pairObj.Tail = *vm.Pop()
	pairObj.Head = *vm.Pop()
	vm.Push(&obj)
	return obj
}
