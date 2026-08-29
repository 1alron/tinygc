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

	Mark()
	IsMarked() bool
}

type IntObject struct {
	Value int

	Marked bool
}

type PairObject struct {
	Head Object
	Tail Object

	Marked bool
}

func (o *IntObject) Type() ObjectType { return ObjInt }
func (o *IntObject) Mark()            { o.Marked = true }
func (o *IntObject) IsMarked() bool   { return o.Marked }

func (o *PairObject) Type() ObjectType { return ObjPair }
func (o *PairObject) Mark()            { o.Marked = true }
func (o *PairObject) IsMarked() bool   { return o.Marked }

func NewObject(vm *VM, ot ObjectType) Object {
	var obj Object
	if ot == ObjInt {
		obj = (*IntObject)(C.malloc(C.size_t(unsafe.Sizeof(IntObject{}))))
	} else {
		obj = (*PairObject)(C.malloc(C.size_t(unsafe.Sizeof(PairObject{}))))
	}
	return obj
}
