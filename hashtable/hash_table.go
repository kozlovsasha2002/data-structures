package hashtable

import (
	"errors"
	"fmt"
)

const (
	defaultSize  = 10
	rehashFactor = 0.7
	step         = 2
)

type hashtable struct {
	capacity int
	size     int
	array    []*node
}

func New() *hashtable {
	return &hashtable{capacity: defaultSize, size: 0, array: make([]*node, defaultSize)}
}

func (h *hashtable) Put(key string, value interface{}) (bool, error) {
	if value == nil {
		return false, errors.New("value is not nil")
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

func (h *hashtable) resize() error {
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

func (h *hashtable) fixCollision(node *node) bool {
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

func (h *hashtable) hash(key string) int {
	runes := []rune(key)

	hash := 0
	for index, item := range runes {
		hash += int(item-'0')*index - 11
	}

	hash %= h.capacity
	return hash
}

func (h *hashtable) Replace(key string, value interface{}) error {
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

func (h *hashtable) Find(key string) (bool, error) {
	hash := h.hash(key)
	if h.array[hash] == nil {
		return false, errors.New("unable to find value with given key")
	}

	if h.array[hash].key == key {
		return true, nil
	} else {
		for i := hash; i < len(h.array)-step; i += step {
			if h.array[i] != nil && h.array[i].key == key {
				return true, nil
			}
		}
		return false, errors.New("unable to find value with given key")
	}
}

func (h *hashtable) Get(key string) *interface{} {
	hash := h.hash(key)
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

func (h *hashtable) Remove(key string) error {
	var item = h.Get(key)
	if item == nil {
		return errors.New("element with given key does not exist")
	}
	item = nil
	return nil
}

func (h *hashtable) Size() int {
	return h.size
}

func (h *hashtable) Clear() {
	h.array = h.array[:0]
	h.capacity = 0
	h.size = 0
}

func (h *hashtable) Print() {
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
