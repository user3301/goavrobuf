package goavrobuf

type NodeType int

const (
	Record NodeType = iota
	Primitive
	Repeated
	Enum
	Unknown
)
