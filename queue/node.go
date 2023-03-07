package queue

type node struct {
	data interface{}
	next *node
}

func NewNode(data interface{}) *node {
	return &node{data: data, next: nil}
}
