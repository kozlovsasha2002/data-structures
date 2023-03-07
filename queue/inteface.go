package queue

type Queue interface {
	Offer(data interface{}) (bool, error)
	Peek() (interface{}, error)
	Poll() interface{}
	Size() int
	IsEmpty() bool
}
