package stack

type Stack interface {
	IsEmpty() bool
	Peek() (interface{}, error)
	Pop() (interface{}, error)
	Push(value interface{}) error
	Print()
	Size() int
}
