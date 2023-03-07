package stack

type node struct {
	data interface{}
	prev *node
}

func NewNode(data interface{}) *node {
	return &node{data: data, prev: nil}
}
