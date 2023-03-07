package queue

import (
	"errors"
	"reflect"
)

type queue struct {
	first *node
	size  int
}

func New() *queue {
	return &queue{first: nil, size: 0}
}

func (q *queue) Offer(data interface{}) (bool, error) {
	if !q.checkDataType(data) {
		return false, errors.New("invalid data type")
	}
	node := NewNode(data)
	node.next = q.first
	q.first = node
	q.size++
	return true, nil
}

func (q *queue) Peek() (interface{}, error) {
	if q.IsEmpty() == true {
		return nil, errors.New("queue is empty")
	}
	return q.first.data, nil
}

func (q *queue) Poll() interface{} {
	if q.IsEmpty() == true {
		return nil
	}
	newFirst := q.first
	q.first = newFirst.next
	q.size--
	return newFirst.data
}

func (q *queue) Size() int {
	return q.size
}

func (q *queue) IsEmpty() bool {
	if q.size == 0 {
		return true
	}
	return false
}

func (q *queue) checkDataType(data interface{}) bool {
	if q.size == 0 {
		return true
	}
	if reflect.TypeOf(q.first.data) == reflect.TypeOf(data) {
		return true
	}
	return false
}
