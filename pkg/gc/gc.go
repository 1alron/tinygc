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

func Sweep(vm *VM) {
	obj := &vm.FirstObject
	for *obj != nil {
		if !(*obj).IsMarked() {
			unreached := *obj
			*obj = unreached.GetNext()
			(*obj).Free()
			vm.NumObjects--
		} else {
			(*obj).Unmark()
			next := (*obj).GetNext()
			obj = &next
		}
	}
}

func gc(vm *VM) {
	MarkAll(vm)
	Sweep(vm)

	if vm.NumObjects == 0 {
		vm.MaxObjects = InitObjNumMax
	} else {
		vm.MaxObjects = vm.NumObjects * 2
	}
}
