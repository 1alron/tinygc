package gc

func MarkAll(vm *VM) {
	for i := 0; i < vm.StackSize; i++ {
		Mark(vm.Stack[i])
	}
}

func Mark(obj Object) {
	if obj.IsMarked() {
		return
	}
	obj.Mark()
	if obj.Type() == ObjPair {
		pairObj := obj.(*PairObject)
		Mark(pairObj.Head)
		Mark(pairObj.Tail)
	}
}
