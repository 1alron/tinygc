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
	Unmark()
	IsMarked() bool

	GetNext() Object
	SetNext(next Object)

	Free()
}

type IntObject struct {
	Value int

	Marked bool

	Next Object
}

type PairObject struct {
	Head Object
	Tail Object

	Marked bool

	Next Object
}

func (o *IntObject) Type() ObjectType    { return ObjInt }
func (o *IntObject) Mark()               { o.Marked = true }
func (o *IntObject) Unmark()             { o.Marked = false }
func (o *IntObject) IsMarked() bool      { return o.Marked }
func (o *IntObject) GetNext() Object     { return o.Next }
func (o *IntObject) SetNext(next Object) { o.Next = next }
func (o *IntObject) Free()               { C.free(unsafe.Pointer(o)) }

func (o *PairObject) Type() ObjectType    { return ObjPair }
func (o *PairObject) Mark()               { o.Marked = true }
func (o *PairObject) Unmark()             { o.Marked = false }
func (o *PairObject) IsMarked() bool      { return o.Marked }
func (o *PairObject) GetNext() Object     { return o.Next }
func (o *PairObject) SetNext(next Object) { o.Next = next }
func (o *PairObject) Free()               { C.free(unsafe.Pointer(o)) }

func NewObject(vm *VM, ot ObjectType) Object {
	var obj Object
	if ot == ObjInt {
		obj = (*IntObject)(C.malloc(C.size_t(unsafe.Sizeof(IntObject{}))))
	} else {
		obj = (*PairObject)(C.malloc(C.size_t(unsafe.Sizeof(PairObject{}))))
	}
	obj.SetNext(vm.FirstObject)
	vm.FirstObject = obj
	vm.NumObjects++
	return obj
}
