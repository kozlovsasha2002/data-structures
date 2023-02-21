package hashtable

type HashTable interface {
	Put(key string, value interface{}) (bool, error)
	Replace(key string, value interface{}) error
	Find(key string) (bool, error)
	Get(key string) *interface{}
	Remove(key string) error
	Size() int
	Clear()
	Print()
}
