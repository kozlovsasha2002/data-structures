package binarytree

import "testing"

func TestAdd(t *testing.T) {
	t.Run("adding one element", func(t *testing.T) {
		var root *Treenode
		root = Add(root, 5)

		realResult := PrintDFS(root, "")
		expectedResult := "5 "

		if realResult != expectedResult {
			t.Errorf("real tree: \"%s\", expected tree: \"%s\"", realResult, expectedResult)
		}
	})

	t.Run("adding multi elements", func(t *testing.T) {
		var root *Treenode
		sl := make([]int, 0, 10)
		sl = append(sl, 20, 18, 16, 24, 22, 25, 23, 14, 15, 17)

		for _, item := range sl {
			root = Add(root, item)
		}

		realResult := PrintDFS(root, "")
		expectedResult := "20 18 16 14 15 17 24 22 23 25 "

		if realResult != expectedResult {
			t.Errorf("real tree: \"%s\", expected tree: \"%s\"", realResult, expectedResult)
		}
	})
}
