package goavrobuf

type NodeType int

const(
	Root NodeType = iota
	Record
	Primitive
)

//func (n NodeType) String() string {
//	return [...]{""}
//}