package linkedlist

import (
	"errors"
	"fmt"
	"reflect"
)

type node struct {
	data interface{}
	next *node
	prev *node
}

func newNode(data interface{}) *node {
	return &node{data: data, next: nil, prev: nil}
}

type LinkedList struct {
	first *node
	last  *node
	size  int
}

func New() *LinkedList {
	return &LinkedList{
		first: nil,
		last:  nil,
		size:  0,
	}
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Sprint() string {
	if l.size == 0 {
		return "[]"
	}

	if l.size == 1 {
		return fmt.Sprintf("[%v]", l.first.data)
	}

	result := "["
	var node *node
	for node = l.first; node != nil; node = node.next {
		if node == l.last {
			result += fmt.Sprint(node.data, "]")
			break
		}
		result += fmt.Sprint(node.data, ", ")
	}

	return result
}

func (l *LinkedList) AddFront(data interface{}) error {
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

func (l *LinkedList) Add(data interface{}) (interface{}, error) {
	if !l.checkDataType(data) {
		return nil, errors.New("invalid data type for adding")
		//log.Fatal("invalid data type for adding")
	}

	if l.size == 0 {
		l.addToEmptyList(data)
		return l.last.data, nil
	}

	node := newNode(data)
	node.prev = l.last
	l.last.next = node
	l.last = node
	l.size++
	return l.last.data, nil
}

func (l *LinkedList) addToEmptyList(data interface{}) {
	node := newNode(data)
	l.first = node
	l.last = node
	l.size++
}

func (l *LinkedList) AddElements(elements ...interface{}) error {
	reflect.TypeOf(elements)
	if len(elements) == 0 {
		return errors.New("input parameters were not passed to the function")
	}

	for _, elem := range elements {
		_, err := l.Add(elem)
		if err != nil {
			return errors.New("invalid data type")
		}
	}
	return nil
}

func (l *LinkedList) AddList(list *LinkedList) (bool, error) {
	if list.size == 0 {
		return false, errors.New("passed list is empty")
	}

	if l.size != 0 {
		if reflect.TypeOf(l.first.data) != reflect.TypeOf(list.first.data) {
			return false, errors.New("lists contain different types of data")
		}
	}

	for item := list.first; item != nil; item = item.next {
		_, err := l.Add(item.data)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (l *LinkedList) SetByIndex(data interface{}, index int) error {
	if !l.checkDataType(data) {
		return errors.New("invalid data type for adding")
	}

	if l.size <= index || index < 0 {
		return errors.New("invalid index. Out of range")
	}

	item := l.first
	for i := 0; i < index; i++ {
		item = item.next
	}
	item.data = data

	return nil
}

func (l *LinkedList) checkDataType(data interface{}) bool {
	if l.size == 0 {
		return true
	}
	if reflect.TypeOf(data) != reflect.TypeOf(l.first.data) {
		return false
	}
	return true
}

func (l *LinkedList) Remove() error {
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

func (l *LinkedList) RemoveFirst() error {
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

func (l *LinkedList) RemoveByIndex(index int) error {
	if l.size == 0 {
		return errors.New("list is empty")
	}

	if l.size <= index || index < 0 {
		return errors.New("invalid index")
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

func (l *LinkedList) Clear() error {
	for l.size != 0 {
		err := l.Remove()
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *LinkedList) removeEmptyListOrWithOneElement() (bool, error) {
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

func (l *LinkedList) Insert(data interface{}, index int) (bool, error) {
	if !l.checkDataType(data) {
		return false, errors.New("invalid data type for adding")
	}

	if l.size < index || index < 0 {
		return false, errors.New("invalid index. Out of range")
	}

	if l.size == 0 {
		l.addToEmptyList(data)
		return true, nil
	}

	if index == 0 {
		err := l.AddFront(data)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	if index == l.size {
		_, err := l.Add(data)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	node := newNode(data)
	curIndex := 0
	left := newNode(-1)
	right := newNode(-1)
	for item := l.first; curIndex <= index-1; curIndex++ {
		left = item
		right = item.next
		item = item.next
	}

	left.next = node
	node.prev = left
	right.prev = node
	node.next = right
	l.size++
	return true, nil
}

func (l *LinkedList) First() (interface{}, error) {
	if l.size == 0 {
		return nil, errors.New("list is empty")
	}
	return l.first.data, nil
}

func (l *LinkedList) Last() (interface{}, error) {
	if l.size == 0 {
		return nil, errors.New("list is empty")
	}
	return l.last.data, nil
}

func (l *LinkedList) Contains(elem interface{}) (bool, error) {
	if l.size == 0 {
		return false, nil
	}

	if reflect.TypeOf(l.first.data) != reflect.TypeOf(elem) {
		return false, errors.New("list contains data of a different type")
	}

	if l.last.data == elem {
		return true, nil
	}
	for node := l.first; node != nil; node = node.next {
		if node.data == elem {
			return true, nil
		}
	}
	return false, nil
}

func (l *LinkedList) Clone() *LinkedList {
	list := New()

	for item := l.first; item != nil; item = item.next {
		list.Add(item.data)
	}

	return list
}
