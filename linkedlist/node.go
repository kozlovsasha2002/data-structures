package linkedlist

type node struct {
	data interface{}
	next *node
	prev *node
}

func newNode(data interface{}) *node {
	return &node{data: data, next: nil, prev: nil}
}
