package queue

import (
	"errors"
	"reflect"
)

type node struct {
	data interface{}
	next *node
}

func NewNode(data interface{}) *node {
	return &node{data: data, next: nil}
}

type Queue struct {
	first *node
	last  *node
	size  int
}

func New() *Queue {
	return &Queue{first: nil, last: nil, size: 0}
}

func (q *Queue) Push(data interface{}) (interface{}, error) {
	if !q.checkDataType(data) {
		return false, errors.New("invalid data type")
	}
	node := NewNode(data)

	if q.size == 0 {
		q.first = node
		q.last = q.first
		q.size++
		return q.first.data, nil
	}
	temp := q.last
	temp.next = node
	q.last = temp.next
	q.size++
	return q.last.data, nil
}

func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() == true {
		return nil, errors.New("queue is empty")
	}
	return q.first.data, nil
}

func (q *Queue) Pop() interface{} {
	if q.IsEmpty() == true {
		return nil
	}
	newFirst := q.first
	q.first = newFirst.next
	q.size--
	return newFirst.data
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	if q.size == 0 {
		return true
	}
	return false
}

func (q *Queue) checkDataType(data interface{}) bool {
	if q.size == 0 {
		return true
	}
	if reflect.TypeOf(q.first.data) == reflect.TypeOf(data) {
		return true
	}
	return false
}
