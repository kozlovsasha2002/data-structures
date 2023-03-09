package stack

//go:generate mockgen -source=interface.go -destination=mocks/mock.go

type Stack interface {
	IsEmpty() bool
	Peek() (interface{}, error)
	Pop() (interface{}, error)
	Push(value interface{}) error
	Print()
	Size() int
}
