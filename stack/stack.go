package stack

import (
	"errors"
	"fmt"
	"reflect"
)

type node struct {
	data interface{}
	prev *node
}

func NewNode(data interface{}) *node {
	return &node{data: data, prev: nil}
}

type Stack struct {
	last *node
	size int
}

func New() *Stack {
	return &Stack{
		last: nil,
		size: 0,
	}
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	if s.size == 0 {
		return true
	}
	return false
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() == true {
		return nil, errors.New("Stack is empty")
	}
	return s.last.data, nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() == true {
		return nil, errors.New("Stack is empty")
	}
	lastElement := s.last
	s.last = lastElement.prev
	s.size--
	return lastElement.data, nil
}

func (s *Stack) Push(data interface{}) (interface{}, error) {
	if !s.checkElementType(data) {
		return nil, errors.New("invalid data type")
	}
	node := NewNode(data)
	if s.size == 0 {
		s.last = node
		s.size++
		return data, nil
	}
	node.prev = s.last
	s.last = node
	s.size++
	return data, nil
}

func (s *Stack) Print() {
	temp := s.last
	for i := 0; i < s.size; i++ {
		fmt.Print(temp.data, " ")
		temp = temp.prev
	}
	fmt.Println()
}

func (s *Stack) checkElementType(data interface{}) bool {
	if s.size == 0 {
		return true
	}
	if reflect.TypeOf(s.last.data) == reflect.TypeOf(data) {
		return true
	}
	return false
}
