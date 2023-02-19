package linkedlist

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type linkedList struct {
	first *node
	last  *node
	size  int
}

func New() *linkedList {
	return &linkedList{
		first: nil,
		last:  nil,
		size:  0,
	}
}

func (l *linkedList) Size() int {
	return l.size
}

func (l *linkedList) Print() {
	if l.size == 0 {
		fmt.Println("[]")
		return
	}

	if l.size == 1 {
		fmt.Printf("[%v]\n", l.first.data)
		return
	}

	fmt.Print("[")
	var node *node
	for node = l.first; node != nil; node = node.next {
		if node == l.last {
			fmt.Print(node.data, "]\n")
			break
		}
		fmt.Print(node.data, ", ")
	}
}

func (l *linkedList) AddFront(data interface{}) error {
	if !l.checkDataType(data) {
		return errors.New("invalid data type for adding")
	}

	if l.size == 0 {
		l.addToEmptyList(data)
		return nil
	}

	node := newNode(data)
	l.first.prev = node
	node.next = l.first
	l.first = node
	l.size++
	return nil
}

func (l *linkedList) Add(data interface{}) error {
	if !l.checkDataType(data) {
		return errors.New("invalid data type for adding")
		//log.Fatal("invalid data type for adding")
	}

	if l.size == 0 {
		l.addToEmptyList(data)
		return nil
	}

	node := newNode(data)
	node.prev = l.last
	l.last.next = node
	l.last = node
	l.size++
	return nil
}

func (l *linkedList) addToEmptyList(data interface{}) {
	node := newNode(data)
	l.first = node
	l.last = node
	l.size++
}

func (l *linkedList) AddElements(elements ...interface{}) error {
	reflect.TypeOf(elements)
	if len(elements) == 0 {
		return errors.New("input parameters were not passed to the function")
		//log.Fatal("input parameters were not passed to the function")
	}

	for _, elem := range elements {
		err := l.Add(elem)
		if err != nil {
			return errors.New("invalid data type")
		}
	}
	return nil
}

func (l *linkedList) AddList(list *linkedList) error {
	if list.size == 0 {
		return errors.New("passed list is empty")
	}

	if l.size != 0 && reflect.TypeOf(l.first) != reflect.TypeOf(list.first) {
		return errors.New("lists contain different types of data")
	}

	for item := list.first; item != nil; item = item.next {
		err := l.Add(item.data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *linkedList) SetByIndex(data interface{}, index int) error {
	if !l.checkDataType(data) {
		return errors.New("invalid data type for adding")
	}

	if l.size <= index || index < 0 {
		return errors.New("invalid index")
	}

	item := l.first
	for i := 0; i < index; i++ {
		item = item.next
	}
	item.data = data

	return nil
}

func (l *linkedList) checkDataType(data interface{}) bool {
	if l.size == 0 {
		return true
	}
	if reflect.TypeOf(data) != reflect.TypeOf(l.first.data) {
		return false
	}
	return true
}

func (l *linkedList) Remove() error {
	res, err := l.removeEmptyListOrWithOneElement()
	if err != nil {
		return err
	}
	if res == true {
		return nil
	}

	var node = l.last
	l.last = nil
	l.last = node.prev
	l.size--
	return nil
}

func (l *linkedList) RemoveFirst() error {
	res, err := l.removeEmptyListOrWithOneElement()
	if err != nil {
		return err
	}
	if res == true {
		return nil
	}

	var node = l.first
	l.first = nil
	l.first = node.next
	l.size--
	return nil
}

func (l *linkedList) RemoveByIndex(index int) error {
	if l.size <= index || index < 0 {
		return errors.New("invalid index")
	}

	if l.size == 0 {
		return errors.New("list is empty")
	}

	if index == 0 {
		err := l.RemoveFirst()
		if err != nil {
			return err
		}
		return nil
	}

	if index == l.size-1 {
		err := l.Remove()
		if err != nil {
			return err
		}
		return nil
	}

	item := l.first
	for i := 1; i <= index; i++ {
		item = item.next
	}
	left := item.prev
	right := item.next
	left.next = right
	right.prev = left
	l.size--
	return nil
}

func (l *linkedList) Clear() error {
	for l.size != 0 {
		err := l.Remove()
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *linkedList) removeEmptyListOrWithOneElement() (bool, error) {
	if l.size == 0 {
		return false, errors.New("the list is empty")
	}

	if l.size == 1 {
		l.first = nil
		l.last = nil
		l.size--
		return true, nil
	}
	return false, nil
}

func (l *linkedList) Insert(data interface{}, index int) error {
	if !l.checkDataType(data) {
		return errors.New("invalid data type for adding")
	}

	if l.size < index || index < 0 {
		return errors.New("invalid index")
	}

	if l.size == 0 {
		l.addToEmptyList(data)
		return nil
	}

	if index == 0 {
		err := l.AddFront(data)
		if err != nil {
			return err
		}
		return nil
	}

	if index == l.size {
		err := l.Add(data)
		if err != nil {
			return err
		}
		return nil
	}

	node := newNode(data)
	curIndex := 0
	for item := l.first; curIndex <= index-1; curIndex++ {
		left := item
		right := item.next
		left.next = node
		node.prev = left
		right.prev = node
		node.next = right
		l.size++
	}
	return nil
}

func (l *linkedList) First() interface{} {
	return l.first.data
}

func (l *linkedList) Last() interface{} {
	return l.last.data
}

func (l *linkedList) Contains(elem interface{}) bool {
	if l.size == 0 {
		return false
	}

	if reflect.TypeOf(l.first.data) != reflect.TypeOf(elem) {
		log.Fatal("list contains data of a different type")
	}

	for node := l.first; node != nil; node = node.next {
		if node.data == elem {
			return true
		}
	}
	return false
}

func (l *linkedList) Clone() *linkedList {
	list := New()

	for item := l.first; item != nil; item = item.next {
		list.Add(item.data)
	}

	return list
}
