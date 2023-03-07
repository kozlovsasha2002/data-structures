package stack

import (
	"errors"
	"fmt"
	"reflect"
)

type stack struct {
	last *node
	size int
}

func New() *stack {
	return &stack{
		last: nil,
		size: 0,
	}
}

func (s *stack) Size() int {
	return s.size
}

func (s *stack) IsEmpty() bool {
	if s.size == 0 {
		return true
	}
	return false
}

func (s *stack) Peek() (interface{}, error) {
	if s.IsEmpty() == true {
		return nil, errors.New("stack is empty")
	}
	return s.last.data, nil
}

func (s *stack) Pop() (interface{}, error) {
	if s.IsEmpty() == true {
		return nil, errors.New("stack is empty")
	}
	lastElement := s.last
	s.last = lastElement.prev
	s.size--
	return lastElement.data, nil
}

func (s *stack) Push(data interface{}) error {
	if !s.checkElementType(data) {
		return errors.New("invalid data type")
	}
	node := NewNode(data)
	if s.size == 0 {
		s.last = node
		s.size++
		return nil
	}
	node.prev = s.last
	s.last = node
	s.size++
	return nil
}

func (s *stack) Print() {
	temp := s.last
	for i := 0; i < s.size; i++ {
		fmt.Print(temp.data, " ")
		temp = temp.prev
	}
	fmt.Println()
}

func (s *stack) checkElementType(data interface{}) bool {
	if s.size == 0 {
		return true
	}
	if reflect.TypeOf(s.last.data) == reflect.TypeOf(data) {
		return true
	}
	return false
}
