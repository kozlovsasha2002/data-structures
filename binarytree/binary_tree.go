package binarytree

import (
	"fmt"
	"math"
)

type Treenode struct {
	value int
	left  *Treenode
	right *Treenode
}

func Add(tree *Treenode, value int) *Treenode {
	if tree == nil {
		tree = create(value)
	} else if value < tree.value {
		tree.left = Add(tree.left, value)
	} else {
		tree.right = Add(tree.right, value)
	}
	return tree
}

func create(value int) *Treenode {
	return &Treenode{value: value, left: nil, right: nil}
}

func PrintDFS(tree *Treenode, nilString string) string {
	if tree != nil {
		nilString += fmt.Sprintf("%d ", tree.value)
		nilString = PrintDFS(tree.left, nilString)
		nilString = PrintDFS(tree.right, nilString)
	}
	return nilString
}

func Find(tree *Treenode, value int) *int {
	if &tree.value == nil {
		return nil
	}
	if tree.value == value {
		return &tree.value
	}
	if tree.value < value {
		return Find(tree.right, value)
	}
	if tree.value > value {
		return Find(tree.left, value)
	}
	return nil
}

func Delete(tree *Treenode, value int) *Treenode {
	if tree == nil {
		return tree
	}

	if value < tree.value {
		tree.left = Delete(tree.left, value)
	} else if value > tree.value {
		tree.right = Delete(tree.right, value)
	} else {
		if tree.left == nil {
			temp := tree.right
			tree = nil
			return temp
		} else if tree.right == nil {
			temp := tree.left
			tree = nil
			return temp
		}
		temp := minValue(tree.right)

		tree.value = temp.value
		tree.right = Delete(tree.right, temp.value)
	}
	return tree
}

func minValue(tree *Treenode) *Treenode {
	current := tree
	for current.left != nil {
		current = current.left
	}
	return current
}

func IsEmpty(tree *Treenode) bool {
	if tree == nil {
		return true
	}
	return false
}

func PrintBFS(tree *Treenode) {
	h := height(tree)
	for i := 0; i < h; i++ {
		printLevel(tree, i)
	}
}

func printLevel(tree *Treenode, level int) {
	if tree == nil {
		return
	}
	if level == 0 {
		fmt.Print(tree.value, " ")
	} else if level > 0 {
		printLevel(tree.left, level-1)
		printLevel(tree.right, level-1)
	}
}

func height(tree *Treenode) int {
	if tree == nil {
		return 0
	}
	lHeight := height(tree.left)
	rHeight := height(tree.right)

	return int(math.Max(float64(lHeight), float64(rHeight))) + 1
}
