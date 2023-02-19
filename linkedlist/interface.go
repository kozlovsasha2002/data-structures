package linkedlist

type LinkedList interface {
	Size() int
	Print()
	AddFront(data interface{}) error
	Add(data interface{}) error
	AddElements(elements ...interface{}) error
	AddList(list *linkedList) error
	SetByIndex(data interface{}, index int) error
	Remove() error
	RemoveFirst() error
	RemoveByIndex(index int) error
	Clear() error
	Insert(data interface{}, index int) error
	First() interface{}
	Last() interface{}
	Contains(elem interface{}) bool
	Clone() *linkedList
}
