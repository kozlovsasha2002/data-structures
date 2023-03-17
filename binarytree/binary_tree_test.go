package binarytree

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("adding one element", func(t *testing.T) {
		var root *treenode
		root = Add(root, 5)

		realResult := PrintBFS(root)
		expectedResult := "5 "

		if realResult != expectedResult {
			t.Errorf("real tree: \"%s\", expected tree: \"%s\"", realResult, expectedResult)
		}
	})

	t.Run("adding multiple elements", func(t *testing.T) {
		var root *treenode
		sl := make([]int, 0, 10)
		sl = append(sl, 20, 18, 16, 14, 15, 26, 24, 22, 25, 28, 27, 30, 29)

		for _, item := range sl {
			root = Add(root, item)
		}

		realResult := PrintBFS(root)
		expectedResult := "20 18 26 16 24 28 14 22 25 27 30 15 29 "

		if realResult != expectedResult {
			t.Errorf("real tree: \"%s\", expected tree: \"%s\"", realResult, expectedResult)
		}
	})
}

func TestAddAll(t *testing.T) {
	t.Run("adding multiple elements", func(t *testing.T) {
		inputElements := make([]int, 0)
		inputElements = append(inputElements, 40, 36, 32, 28, 34, 33, 38, 60, 56, 64, 52, 54, 58, 59, 57, 62, 63, 66, 70)
		tree := AddAll(inputElements)

		actualTree := PrintBFS(tree)
		expectedTree := "40 36 60 32 38 56 64 28 34 52 58 62 66 33 54 57 59 63 70 "

		if actualTree != expectedTree {
			t.Errorf("expected tree:\"%s\" and actual tree:\"%s\"", expectedTree, actualTree)
		}
	})

	t.Run("adding one element", func(t *testing.T) {
		inputElements := make([]int, 0)
		inputElements = append(inputElements, 7)
		tree := AddAll(inputElements)

		actualTree := PrintBFS(tree)
		expectedTree := "7 "

		if actualTree != expectedTree {
			t.Errorf("expected tree:\"%s\" and actual tree:\"%s\"", expectedTree, actualTree)
		}
	})

	t.Run("adding slice with zero elements", func(t *testing.T) {
		inputElements := make([]int, 0)
		tree := AddAll(inputElements)

		actualTree := PrintBFS(tree)
		expectedTree := ""

		if actualTree != expectedTree {
			t.Errorf("expected tree:\"%s\" and actual tree:\"%s\"", expectedTree, actualTree)
		}
	})

	t.Run("adding empty slice", func(t *testing.T) {
		var inputElements []int
		tree := AddAll(inputElements)

		actualTree := PrintBFS(tree)
		expectedTree := ""

		if actualTree != expectedTree {
			t.Errorf("expected tree:\"%s\" and actual tree:\"%s\"", expectedTree, actualTree)
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("removing an element with no children", func(t *testing.T) {
		inputElements := make([]int, 0)
		inputElements = append(inputElements, 40, 36, 32, 28, 34, 33, 38, 60, 56, 64, 52, 54, 58, 59, 57, 62, 63, 66, 70)
		tree := AddAll(inputElements)

		tree, _ = Delete(tree, 57)
		actualTree := PrintBFS(tree)
		expectedTree := "40 36 60 32 38 56 64 28 34 52 58 62 66 33 54 59 63 70 "

		if actualTree != expectedTree {
			t.Errorf("expected tree:\"%s\" and actual tree:\"%s\"", expectedTree, actualTree)
		}
	})

	t.Run("removing an element that has only one child", func(t *testing.T) {
		inputElements := make([]int, 0)
		inputElements = append(inputElements, 40, 36, 32, 28, 34, 33, 38, 60, 56, 64, 52, 54, 58, 59, 57, 62, 63, 66, 70)
		tree := AddAll(inputElements)

		tree, _ = Delete(tree, 34)
		actualTree := PrintBFS(tree)
		expectedTree := "40 36 60 32 38 56 64 28 33 52 58 62 66 54 57 59 63 70 "

		if actualTree != expectedTree {
			t.Errorf("expected tree:\"%s\" and actual tree:\"%s\"", expectedTree, actualTree)
		}
	})

	t.Run("deleting an element that has both children", func(t *testing.T) {
		inputElements := make([]int, 0)
		inputElements = append(inputElements, 40, 36, 32, 28, 34, 33, 38, 60, 56, 64, 52, 54, 58, 59, 57, 62, 63, 66, 70)
		tree := AddAll(inputElements)

		tree, _ = Delete(tree, 60)
		actualTree := PrintBFS(tree)
		expectedTree := "40 36 62 32 38 56 64 28 34 52 58 63 66 33 54 57 59 70 "

		if actualTree != expectedTree {
			t.Errorf("expected tree:\"%s\" and actual tree:\"%s\"", expectedTree, actualTree)
		}
	})

	t.Run("removing element doesn't exist", func(t *testing.T) {
		inputElements := make([]int, 0)
		inputElements = append(inputElements, 40, 36, 32, 28, 34, 33, 38, 60, 56, 64, 52, 54, 58, 59, 57, 62, 63, 66, 70)
		tree := AddAll(inputElements)

		tree, _ = Delete(tree, 80)
		actualTree := PrintBFS(tree)
		expectedTree := "40 36 60 32 38 56 64 28 34 52 58 62 66 33 54 57 59 63 70 "

		if actualTree != expectedTree {
			t.Errorf("expected tree:\"%s\" and actual tree:\"%s\"", expectedTree, actualTree)
		}
	})

	t.Run("tree contains one element", func(t *testing.T) {
		tree := Add(nil, 10)
		tree, _ = Delete(tree, 10)

		actualResult := IsEmpty(tree)
		expectedResult := true

		if actualResult != expectedResult {
			t.Errorf("expected result:%t, actual result:%t", expectedResult, actualResult)
		}
	})

	t.Run("tree contains zero elements", func(t *testing.T) {
		tree := Add(nil, 10)
		tree, _ = Delete(tree, 10)
		tree, err := Delete(tree, 20)

		actualResult := err.Error()
		expectedResult := "tree is empty"

		if actualResult != expectedResult {
			t.Errorf("expected error:%v, actual error:%v", expectedResult, actualResult)
		}
	})
}

func TestFind(t *testing.T) {
	t.Run("tree is empty", func(t *testing.T) {
		tree := Add(nil, 10)
		tree, _ = Delete(tree, 10)

		actualTree := Find(tree, 100)
		expectedTree := false

		if actualTree != expectedTree {
			t.Errorf("expected tree:\"%t\" and actual tree:\"%t\"", expectedTree, actualTree)
		}
	})

	t.Run("tree doesn't contain search element", func(t *testing.T) {
		inputElements := make([]int, 0)
		inputElements = append(inputElements, 40, 36, 32, 28, 34, 33, 38, 60)
		tree := AddAll(inputElements)

		actualTree := Find(tree, 100)
		expectedTree := false

		if actualTree != expectedTree {
			t.Errorf("expected tree:\"%t\" and actual tree:\"%t\"", expectedTree, actualTree)
		}
	})

	t.Run("tree contains search element", func(t *testing.T) {
		inputElements := make([]int, 0)
		inputElements = append(inputElements, 40, 36, 32, 28, 34, 33, 38, 60)
		tree := AddAll(inputElements)

		actualTree := Find(tree, 28)
		expectedTree := true

		if actualTree != expectedTree {
			t.Errorf("expected tree:\"%t\" and actual tree:\"%t\"", expectedTree, actualTree)
		}
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("tree is empty", func(t *testing.T) {
		tree := Add(nil, 10)
		tree, _ = Delete(tree, 10)

		actualResult := IsEmpty(tree)
		expectedResult := true

		if actualResult != expectedResult {
			t.Errorf("expected result:%t, actual result:%t", expectedResult, actualResult)
		}
	})

	t.Run("tree is not empty", func(t *testing.T) {
		tree := Add(nil, 10)

		actualResult := IsEmpty(tree)
		expectedResult := false

		if actualResult != expectedResult {
			t.Errorf("expected result:%t, actual result:%t", expectedResult, actualResult)
		}
	})
}

func TestPrintBFS(t *testing.T) {
	t.Run("tree is empty", func(t *testing.T) {
		tree := Add(nil, 5)
		tree, _ = Delete(tree, 5)

		actualResult := PrintBFS(tree)
		expectedResult := ""

		if actualResult != expectedResult {
			t.Errorf("expected result:\"%s\" and actual result:\"%s\"", expectedResult, actualResult)
		}
	})

	t.Run("tree contains one element", func(t *testing.T) {
		tree := Add(nil, 5)

		actualResult := PrintBFS(tree)
		expectedResult := "5 "

		if actualResult != expectedResult {
			t.Errorf("expected result:\"%s\" and actual result:\"%s\"", expectedResult, actualResult)
		}
	})

	t.Run("tree contains multiple elements", func(t *testing.T) {
		inputElements := make([]int, 0)
		inputElements = append(inputElements, 40, 36, 32, 28, 34, 33, 38, 60)
		tree := AddAll(inputElements)

		actualResult := PrintBFS(tree)
		expectedResult := "40 36 60 32 38 28 34 33 "

		if actualResult != expectedResult {
			t.Errorf("expected result:\"%s\" and actual result:\"%s\"", expectedResult, actualResult)
		}
	})
}

func TestPrintDFS(t *testing.T) {
	t.Run("tree is empty", func(t *testing.T) {
		tree := Add(nil, 5)
		tree, _ = Delete(tree, 5)

		actualResult := PrintDFS(tree, "")
		expectedResult := ""

		if actualResult != expectedResult {
			t.Errorf("expected result:\"%s\" and actual result:\"%s\"", expectedResult, actualResult)
		}
	})

	t.Run("tree contains one element", func(t *testing.T) {
		tree := Add(nil, 5)

		actualResult := PrintDFS(tree, "")
		expectedResult := "5 "

		if actualResult != expectedResult {
			t.Errorf("expected result:\"%s\" and actual result:\"%s\"", expectedResult, actualResult)
		}
	})

	t.Run("tree contains multiple elements", func(t *testing.T) {
		inputElements := make([]int, 0)
		inputElements = append(inputElements, 40, 36, 32, 28, 34, 33, 38, 60)
		tree := AddAll(inputElements)

		actualResult := PrintDFS(tree, "")
		expectedResult := "28 32 33 34 36 38 40 60 "

		if actualResult != expectedResult {
			t.Errorf("expected result:\"%s\" and actual result:\"%s\"", expectedResult, actualResult)
		}
	})
}

func Test_minValue(t *testing.T) {
	t.Run("node has not children", func(t *testing.T) {
		tree := Add(nil, 5)
		tree, _ = Delete(tree, 5)

		_, err := minValue(tree)
		actualResult := err.Error()
		expectedResult := "node is nil"

		if actualResult != expectedResult {
			t.Errorf("expected result:%s and actual result:%s", expectedResult, actualResult)
		}
	})

	t.Run("node has children", func(t *testing.T) {
		inputElements := make([]int, 0)
		inputElements = append(inputElements, 40, 36, 32, 28, 34, 33, 38, 60)
		tree := AddAll(inputElements)

		tree, _ = minValue(tree)
		actualResult := tree.value
		expectedResult := 28

		if actualResult != expectedResult {
			t.Errorf("expected result:\"%d\" and actual result:\"%d\"", expectedResult, actualResult)
		}
	})
}

func Test_sprintLevel(t *testing.T) {
	t.Run("print zero level", func(t *testing.T) {
		inputElements := make([]int, 0)
		inputElements = append(inputElements, 40, 36, 32, 28, 34, 33, 38, 60)
		tree := AddAll(inputElements)

		actualResult := sprintLevel(tree, 0)
		expectedResult := "40 "

		if actualResult != expectedResult {
			t.Errorf("expected result:\"%s\" and actual result:\"%s\"", expectedResult, actualResult)
		}
	})

	t.Run("print third level", func(t *testing.T) {
		inputElements := make([]int, 0)
		inputElements = append(inputElements, 40, 36, 32, 28, 34, 33, 38, 60, 56, 64, 52, 54, 58, 59, 57, 62, 63, 66, 70)
		tree := AddAll(inputElements)

		actualResult := sprintLevel(tree, 3)
		expectedResult := "28 34 52 58 62 66 "

		if actualResult != expectedResult {
			t.Errorf("expected result:\"%s\" and actual result:\"%s\"", expectedResult, actualResult)
		}
	})
}

func Test_height(t *testing.T) {
	t.Run("get height = 1", func(t *testing.T) {
		inputElements := make([]int, 0)
		inputElements = append(inputElements, 40)
		tree := AddAll(inputElements)

		actualResult := height(tree)
		expectedResult := 1

		if actualResult != expectedResult {
			t.Errorf("expected result:%d and actual result:%d", expectedResult, actualResult)
		}
	})

	t.Run("get height = 3", func(t *testing.T) {
		inputElements := make([]int, 0)
		inputElements = append(inputElements, 40, 36, 32, 38, 60, 56)
		tree := AddAll(inputElements)

		actualResult := height(tree)
		expectedResult := 3

		if actualResult != expectedResult {
			t.Errorf("expected result:%d and actual result:%d", expectedResult, actualResult)
		}
	})

	t.Run("get height = 5", func(t *testing.T) {
		inputElements := make([]int, 0)
		inputElements = append(inputElements, 40, 36, 32, 28, 34, 33, 38, 60, 56, 64, 52, 54, 58, 59, 57, 62, 63, 66, 70)
		tree := AddAll(inputElements)

		actualResult := height(tree)
		expectedResult := 5

		if actualResult != expectedResult {
			t.Errorf("expected result:%d and actual result:%d", expectedResult, actualResult)
		}
	})
}

func Test_create(t *testing.T) {
	t.Run("create one tree", func(t *testing.T) {
		tree := create(5)

		expectedValue := 5

		if tree.value != expectedValue || tree.left != nil || tree.right != nil {
			t.Errorf("actual tree:%d, %v, %v and expected:%d, %v, %v", tree.value, tree.left, tree.right,
				expectedValue, nil, nil)
		}
	})
}
