package hashtable

type node struct {
	key   string
	value interface{}
}

func newNode(key string, value interface{}) *node {
	return &node{key: key, value: value}
}
