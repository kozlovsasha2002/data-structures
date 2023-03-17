package binarytree

import (
	"errors"
	"fmt"
	"math"
)

type treenode struct {
	value int
	left  *treenode
	right *treenode
}

func Add(tree *treenode, value int) *treenode {
	if tree == nil {
		tree = create(value)
	} else if value < tree.value {
		tree.left = Add(tree.left, value)
	} else {
		tree.right = Add(tree.right, value)
	}
	return tree
}

func AddAll(inputElements []int) *treenode {
	if len(inputElements) == 0 {
		return nil
	}

	tree := Add(nil, inputElements[0])
	for i, _ := range inputElements {
		if i > 0 {
			tree = Add(tree, inputElements[i])
		}
	}
	return tree
}

func create(value int) *treenode {
	return &treenode{value: value, left: nil, right: nil}
}

func PrintDFS(tree *treenode, nilString string) string {
	if tree != nil {

		nilString = PrintDFS(tree.left, nilString)
		nilString += fmt.Sprintf("%d ", tree.value)
		nilString = PrintDFS(tree.right, nilString)
	}
	return nilString
}

func Find(tree *treenode, value int) bool {
	if tree == nil {
		return false
	}
	if tree.value == value {
		return true
	}
	if tree.value < value {
		return Find(tree.right, value)
	}
	if tree.value > value {
		return Find(tree.left, value)
	}
	return false
}

// Delete method removes an element in this way: finds the element to be removed,
// and then replaces it with the smallest element from the right subtree of the element being searched for.
func Delete(tree *treenode, value int) (*treenode, error) {
	if tree == nil {
		return tree, errors.New("tree is empty")
	}

	if value < tree.value {
		tree.left, _ = Delete(tree.left, value)
	} else if value > tree.value {
		tree.right, _ = Delete(tree.right, value)
	} else {
		if tree.left == nil {
			temp := tree.right
			tree = nil
			return temp, nil
		} else if tree.right == nil {
			temp := tree.left
			tree = nil
			return temp, nil
		}
		temp, err := minValue(tree.right)
		if err != nil {
			return nil, err
		}

		tree.value = temp.value
		tree.right, _ = Delete(tree.right, temp.value)
	}
	return tree, nil
}

func minValue(tree *treenode) (*treenode, error) {
	if tree == nil {
		return nil, errors.New("node is nil")
	}
	current := tree
	for current.left != nil {
		current = current.left
	}
	return current, nil
}

func IsEmpty(tree *treenode) bool {
	if tree == nil {
		return true
	}
	return false
}

func PrintBFS(tree *treenode) string {
	result := ""
	h := height(tree)
	for i := 0; i < h; i++ {
		result += sprintLevel(tree, i)
	}
	return result
}

func sprintLevel(node *treenode, levelNumber int) string {
	level := ""
	if node == nil {
		return level
	}

	if levelNumber == 0 {
		level += fmt.Sprint(node.value, " ")
	} else if levelNumber > 0 {
		level += sprintLevel(node.left, levelNumber-1)
		level += sprintLevel(node.right, levelNumber-1)
	}
	return level
}

func height(tree *treenode) int {
	if tree == nil {
		return 0
	}
	lHeight := height(tree.left)
	rHeight := height(tree.right)

	return int(math.Max(float64(lHeight), float64(rHeight))) + 1
}
