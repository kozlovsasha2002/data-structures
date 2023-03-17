package hashtable

import (
	"errors"
	"fmt"
)

type node struct {
	key   string
	value interface{}
}

func newNode(key string, value interface{}) *node {
	return &node{key: key, value: value}
}

const (
	defaultSize  = 10
	rehashFactor = 0.7
	step         = 2
)

type Hashtable struct {
	capacity int
	size     int
	array    []*node
}

func New() *Hashtable {
	return &Hashtable{capacity: defaultSize, size: 0, array: make([]*node, defaultSize)}
}

func (h *Hashtable) Put(key string, value interface{}) (bool, error) {
	if value == nil {
		return false, errors.New("value is nil")
	}

	node := newNode(key, value)
	currentFactor := float64(h.size) / float64(h.capacity)
	if currentFactor > rehashFactor {
		err := h.resize()
		if err != nil {
			return false, err
		}
	}

	hash := h.hash(key)
	if h.array[hash] == nil {
		h.array[hash] = node
		h.size++
		return true, nil
	}

	for true {
		if !h.fixCollision(node) {
			err := h.resize()
			if err != nil {
				return false, err
			}
		} else {
			return true, nil
		}
	}

	return false, errors.New("failed to add element")
}

func (h *Hashtable) resize() error {
	h.capacity *= 2
	h.size = 0
	prevArray := h.array
	h.array = make([]*node, h.capacity)
	for _, node := range prevArray {
		if node != nil {
			result, _ := h.Put(node.key, node.value)
			if result == false {
				return errors.New("failed to add element")
			}
		}
	}
	return nil
}

func (h *Hashtable) fixCollision(node *node) bool {
	hash := h.hash(node.key)
	if h.array[hash] == nil {
		h.array[hash] = newNode(node.key, node.value)
		h.size++
		return true
	}
	if h.array[hash] != nil {
		for i := hash; i < len(h.array)-step; i += step {
			if h.array[i] == nil {
				h.array[i] = newNode(node.key, node.value)
				h.size++
				return true
			}
		}
	}
	return false
}

func (h *Hashtable) hash(key string) int {
	runes := []rune(key)

	hash := 0
	for index, item := range runes {
		hash += int(item-'0')*index - 11
	}

	if hash < 0 {
		hash *= -1
	}
	hash %= h.capacity
	return hash
}

func (h *Hashtable) Replace(key string, value interface{}) error {
	if value == nil {
		return errors.New("value is not nil")
	}

	res := h.Get(key)
	if res != nil {
		*res = value
		return nil
	}
	return errors.New("value cannot be replaced because element with given key does not exist")
}

func (h *Hashtable) Find(key string) bool {
	hash := h.hash(key)

	if h.size == 0 {
		return false
	}

	if h.array[hash] == nil {
		return false
	}

	if h.array[hash].key == key {
		return true
	} else {
		for i := hash; i < len(h.array)-step; i += step {
			if h.array[i] != nil && h.array[i].key == key {
				return true
			}
		}
		return false
	}
}

func (h *Hashtable) Get(key string) *interface{} {
	hash := h.hash(key)
	if h.size == 0 {
		return nil
	}

	if h.array[hash] == nil {
		return nil
	}

	if h.array[hash].key == key {
		return &h.array[hash].value
	} else {
		for i := hash; i < len(h.array)-step; i += step {
			if h.array[i] != nil && h.array[i].key == key {
				return &h.array[i].value
			}
		}
		return nil
	}
}

func (h *Hashtable) Remove(key string) error {
	hash := h.hash(key)
	if h.array[hash] == nil {
		return errors.New("element with given key does not exist")
	}
	h.array[hash] = nil
	h.size--
	return nil
}

func (h *Hashtable) Size() int {
	return h.size
}

func (h *Hashtable) RemoveAll() {
	h.array = make([]*node, defaultSize)
	h.capacity = defaultSize
	h.size = 0
}

func (h *Hashtable) Print() {
	if len(h.array) == 0 {
		fmt.Println("[]")
		return
	}

	var count = h.size - 1
	fmt.Print("[")
	for index, _ := range h.array {
		if h.array[index] != nil {
			if count == 0 {
				fmt.Printf("%v]\n", *h.array[index])
			} else {
				fmt.Print(*h.array[index], ", ")
				count--
			}
		}
	}
}
