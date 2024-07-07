package _5_stack

type Stack interface {
	Push(interface{})
	Pop() interface{}
	IsEmpty() bool
	Flush()
}
