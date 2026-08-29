package gc

/*
#include <stdlib.h>
*/
import "C"
import "unsafe"

type ObjectType int

const (
	ObjInt ObjectType = iota
	ObjPair
)

type Object interface {
	Type() ObjectType
}

type IntObject struct {
	Value int
}

type PairObject struct {
	Head Object
	Tail Object
}

func (o *IntObject) Type() ObjectType { return ObjInt }

func (o *PairObject) Type() ObjectType { return ObjPair }

func NewObject(vm *VM, ot ObjectType) Object {
	var obj Object
	if ot == ObjInt {
		obj = (*IntObject)(C.malloc(C.size_t(unsafe.Sizeof(IntObject{}))))
	} else {
		obj = (*PairObject)(C.malloc(C.size_t(unsafe.Sizeof(PairObject{}))))
	}
	return obj
}
