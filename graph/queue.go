package graph

type node struct {
	data string
	next *node
}

func newNode(data string) *node {
	return &node{data: data, next: nil}
}

type queue struct {
	first *node
	last  *node
	size  int
}

func newQueue() *queue {
	return &queue{first: nil, last: nil, size: 0}
}

func (q *queue) push(data string) (string, error) {
	node := newNode(data)

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

func (q *queue) pop() string {
	newFirst := q.first
	q.first = newFirst.next
	q.size--
	return newFirst.data
}
