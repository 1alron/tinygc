package gc

func MarkAndSweep(vm *VM) {
	markAll(vm)
	sweep(vm)

	if vm.NumObjects == 0 {
		vm.MaxObjects = InitObjNumMax
	} else {
		vm.MaxObjects = vm.NumObjects * 2
	}
}

func markAll(vm *VM) {
	for i := 0; i < vm.StackSize; i++ {
		mark(vm.Stack[i])
	}
}

func mark(obj Object) {
	if obj.IsMarked() {
		return
	}
	obj.Mark()
	if obj.Type() == ObjPair {
		pairObj := obj.(*PairObject)
		mark(pairObj.Head)
		mark(pairObj.Tail)
	}
}

func sweep(vm *VM) {
	var previous Object

	for current := vm.FirstObject; current != nil; {
		next := current.GetNext()
		if !current.IsMarked() {
			if previous == nil {
				vm.FirstObject = next
			} else {
				previous.SetNext(next)
			}
			current.Free()
			vm.NumObjects--
		} else {
			current.Unmark()
			previous = current
		}
		current = next
	}
}
