package gc

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
