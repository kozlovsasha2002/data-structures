package prefixtree

import (
	"errors"
	"fmt"
)

type PrefixTree struct {
	head                  *vertex
	amountOfAddedElements int
}

func New() *PrefixTree {
	return &PrefixTree{head: newVertex(), amountOfAddedElements: 0}
}

type vertex struct {
	nodes      map[string]*vertex
	data       string
	isTerminal bool
}

func newVertex() *vertex {
	return &vertex{nodes: make(map[string]*vertex), data: "", isTerminal: false}
}

func (p *PrefixTree) Add(str string) bool {
	if str == "" {
		return false
	}

	current := p.head
	for _, ch := range str {
		char := string(ch)

		if current.nodes[char] == nil {
			current.nodes[char] = &vertex{nodes: make(map[string]*vertex), data: "", isTerminal: false}
		}
		current = current.nodes[char]
	}
	current.isTerminal = true
	current.data = str
	p.amountOfAddedElements++
	return true
}

func (p *PrefixTree) AddAll(strings []string) error {
	if len(strings) == 0 {
		return errors.New("can't add an empty slice")
	}

	for _, str := range strings {
		isSuccess := p.Add(str)
		if !isSuccess {
			return errors.New("can't insert empty string")
		}
	}
	return nil
}

func (p *PrefixTree) Contains(str string) bool {
	current := p.head
	for _, ch := range str {
		char := string(ch)

		if current.nodes[char] == nil {
			return false
		}
		current = current.nodes[char]
	}
	if current.isTerminal == false {
		return false
	}
	return true
}

func (p *PrefixTree) Remove(str string) bool {
	current := p.head
	for _, ch := range str {
		char := string(ch)

		if current.nodes[char] == nil {
			return false
		}
		current = current.nodes[char]
	}

	if current.isTerminal == false {
		return false
	} else {
		current.isTerminal = false
		current.data = ""
		p.amountOfAddedElements--
		return true
	}
}

func (p *PrefixTree) GetAllByPrefix(prefix string) ([]string, error) {
	result := make([]string, 0)

	if len(p.head.nodes) == 0 {
		return nil, errors.New("tree is empty")
	}

	current := p.head
	for _, ch := range prefix {
		char := string(ch)

		if current.nodes[char] == nil {
			err := fmt.Sprintf("no words found beginning with prefix = \"%s\"", prefix)
			return nil, errors.New(err)
		}
		current = current.nodes[char]
	}

	if current.isTerminal == true {
		return append(result, prefix), nil
	} else {
		result = addTerminalStrings(result, current)
	}
	return result, nil
}

func (p *PrefixTree) AmountOfAddedElements() int {
	return p.amountOfAddedElements
}

func addTerminalStrings(result []string, v *vertex) []string {
	for _, str := range v.nodes {
		if str != nil && str.isTerminal == true {
			result = append(result, str.data)
		}
		if str != nil && str.isTerminal == false || len(str.nodes) != 0 {
			result = addTerminalStrings(result, str)
		}
	}
	return result
}
