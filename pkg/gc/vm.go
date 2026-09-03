package gc

/*
#include <stdlib.h>
*/
import "C"
import "unsafe"

const (
	StackMax      = 256
	InitObjNumMax = 8
)

type VM struct {
	Stack     [StackMax]Object
	StackSize int

	FirstObject Object

	NumObjects int
	MaxObjects int
}

func NewVM() *VM {
	ptr := C.malloc(C.size_t(unsafe.Sizeof(VM{})))
	if ptr == nil {
		panic("failed to allocate memory for VM")
	}
	vm := (*VM)(ptr)
	vm.NumObjects = 0
	vm.StackSize = 0
	vm.FirstObject = nil
	vm.MaxObjects = InitObjNumMax
	return vm
}

func FreeVM(vm *VM) {
	C.free(unsafe.Pointer(vm))
}

func (vm *VM) Pop() Object {
	if vm.StackSize <= 0 {
		panic("VM's stack underflow")
	}
	res := vm.Stack[vm.StackSize-1]
	vm.StackSize--
	return res
}

func (vm *VM) PushInt(intValue int) {
	obj := newObject(vm, ObjInt)
	intObj := obj.(*IntObject)
	intObj.Value = intValue
	vm.push(obj)
}

func (vm *VM) PushPair() Object {
	obj := newObject(vm, ObjPair)
	pairObj := obj.(*PairObject)
	pairObj.Tail = vm.Pop()
	pairObj.Head = vm.Pop()
	vm.push(obj)
	return obj
}

func (vm *VM) push(value Object) {
	if vm.StackSize >= StackMax {
		panic("VM's stack overflow")
	}
	vm.Stack[vm.StackSize] = value
	vm.StackSize++
}
