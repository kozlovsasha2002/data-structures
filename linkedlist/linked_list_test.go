package linkedlist

import (
	"testing"
)

func NewFillingList(list []interface{}, t *testing.T) (*LinkedList, error) {
	t.Helper()
	result := New()
	for _, item := range list {
		_, err := result.Add(item)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func TestNewFillingTest(t *testing.T) {
	t.Run("elements have equal data type", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24, 15)
		list, _ := NewFillingList(elems, t)

		realResult := list.Sprint()
		expectedResult := "[12, 24, 15]"

		if realResult != expectedResult {
			t.Errorf("real result: %v, expected result: %v", realResult, expectedResult)
		}
	})

	t.Run("elements have different data type", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24, "15")
		_, err := NewFillingList(elems, t)

		realError := err.Error()
		expectedError := "invalid data type for adding"

		if realError != expectedError {
			t.Errorf("real error: %v, expected error: %v", realError, expectedError)
		}
	})
}

func TestLinkedList_Add(t *testing.T) {
	t.Run("list is empty. Adding one element", func(t *testing.T) {
		list := New()

		actionValue, err := list.Add(5)
		expectedValue := 5

		if err != nil || actionValue != expectedValue {
			t.Errorf("expected value = %v and action value = %v, error: %v", expectedValue, actionValue, err.Error())
		}
	})

	t.Run("list is empty. Adding multi elements", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24, 15, 11, 2)
		list, err := NewFillingList(elems, t)

		action := list.Sprint()
		expected := "[12, 24, 15, 11, 2]"

		if action != expected || err != nil {
			t.Errorf("expected list = %s and action list = %s. Error: %v", action, expected, err.Error())
		}
	})

	t.Run("list is not empty. Adding elements with different data type", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, "dog", 15, 11, 2)

		_, actionErr := NewFillingList(elems, t)
		expectedErr := "invalid data type for adding"

		if actionErr.Error() != expectedErr {
			t.Errorf("expected error: %v. Real error: %v", expectedErr, actionErr)
		}
	})
}

func TestLinkedList_Sprint(t *testing.T) {
	t.Run("list is empty", func(t *testing.T) {
		list := New()

		realResult := list.Sprint()
		expectedResult := "[]"

		if realResult != expectedResult {
			t.Errorf("real result = %v and expected result = %v", realResult, expectedResult)
		}
	})

	t.Run("list is not empty", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, "dog", "cat", "mouse")
		list, _ := NewFillingList(elems, t)

		realResult := list.Sprint()
		expectedResult := "[dog, cat, mouse]"

		if realResult != expectedResult {
			t.Errorf("real result = %v and expected result = %v", realResult, expectedResult)
		}
	})
}

func TestLinkedList_Clear(t *testing.T) {
	t.Run("list is empty", func(t *testing.T) {
		list := New()

		err := list.Clear()
		if err != nil {
			t.Errorf(err.Error())
		}
		realResult := list.Sprint()
		expectedResult := "[]"

		if realResult != expectedResult {
			t.Errorf("real result: %v, expected result: %v", realResult, expectedResult)
		}
	})

	t.Run("list is not empty", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24, 15, 11, 2)
		list, _ := NewFillingList(elems, t)

		err := list.Clear()
		if err != nil {
			t.Errorf(err.Error())
		}
		realResult := list.Sprint()
		expectedResult := "[]"

		if realResult != expectedResult {
			t.Errorf("real result: %v, expected result: %v", realResult, expectedResult)
		}
	})
}

func TestLinkedList_Remove(t *testing.T) {
	t.Run("list is empty", func(t *testing.T) {
		list := New()

		actionError := list.Remove()
		expectedError := "the list is empty"

		if actionError.Error() != expectedError {
			t.Errorf("real error: %v, expected error: %v", actionError, expectedError)
		}
	})

	t.Run("list is not empty", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24, 15, 11)
		list, _ := NewFillingList(elems, t)

		err := list.Remove()
		if err != nil {
			t.Error(err.Error())
		}
		actionList := list.Sprint()
		expectedList := "[12, 24, 15]"

		if actionList != expectedList {
			t.Errorf("list: %v, expected list: %v", actionList, expectedList)
		}
	})

	t.Run("list is not empty. Deletion happens multiple times", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24, 15, 11)
		list, _ := NewFillingList(elems, t)

		list.Remove()
		list.Remove()
		list.Remove()

		actionList := list.Sprint()
		expectedList := "[12]"

		if actionList != expectedList {
			t.Errorf("list: %v, expected list: %v", actionList, expectedList)
		}
	})
}

func TestLinkedList_AddElements(t *testing.T) {
	t.Run("list is empty", func(t *testing.T) {
		list := New()
		list.AddElements("dog", "cat")

		realResult := list.Sprint()
		expectedResult := "[dog, cat]"

		if realResult != expectedResult {
			t.Errorf("real list: %v, expected list: %v", realResult, expectedResult)
		}
	})

	t.Run("list is not empty", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24, 15, 11)
		list, _ := NewFillingList(elems, t)

		list.AddElements(7, 8, 9)
		realResult := list.Sprint()
		expectedResult := "[12, 24, 15, 11, 7, 8, 9]"

		if realResult != expectedResult {
			t.Errorf("real list: %v, expected list: %v", realResult, expectedResult)
		}
	})

	t.Run("adding elements with different data type", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24, 11)
		list, _ := NewFillingList(elems, t)

		realError := list.AddElements(5.5)
		expectedError := "invalid data type"

		if realError.Error() != expectedError {
			t.Errorf("real error: %v, expected error: %v", realError, expectedError)
		}
	})

	t.Run("adding zero elements", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24)
		list, _ := NewFillingList(elems, t)

		realError := list.AddElements()
		expectedError := "input parameters were not passed to the function"

		if realError.Error() != expectedError {
			t.Errorf("real error: %v, expected error: %v", realError, expectedError)
		}
	})
}

func TestLinkedList_AddFront(t *testing.T) {
	t.Run("list is empty", func(t *testing.T) {
		list := New()

		list.AddFront(5)
		realResult := list.Sprint()
		expectedResult := "[5]"

		if realResult != expectedResult {
			t.Errorf("real list: %v, expected list: %v", realResult, expectedResult)
		}
	})

	t.Run("list is not empty", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24, 11, 27)
		list, _ := NewFillingList(elems, t)

		list.AddFront(5)
		realResult := list.Sprint()
		expectedResult := "[5, 12, 24, 11, 27]"

		if realResult != expectedResult {
			t.Errorf("real list: %v, expected list: %v", realResult, expectedResult)
		}
	})

	t.Run("adding element with invalid data type", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24, 11, 27)
		list, _ := NewFillingList(elems, t)

		realError := list.AddFront("string")
		expectedError := "invalid data type for adding"
		if realError.Error() != expectedError {
			t.Errorf("real error: %v, expected error: %v", realError, expectedError)
		}
	})
}

func TestLinkedList_Clone(t *testing.T) {
	t.Run("list is empty", func(t *testing.T) {
		list := New()

		cloneList := list.Clone()
		realResult := cloneList.Sprint()
		expectedResult := list.Sprint()

		if realResult != expectedResult {
			t.Errorf("real result: %v, expected result: %v", realResult, expectedResult)
		}
	})

	t.Run("list is not empty", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, "cat", "dog", "lion", "fox")
		list, _ := NewFillingList(elems, t)

		cloneList := list.Clone()
		realResult := cloneList.Sprint()
		expectedResult := list.Sprint()

		if realResult != expectedResult {
			t.Errorf("real result: %v, expected result: %v", realResult, expectedResult)
		}
	})
}

func TestLinkedList_Size(t *testing.T) {
	t.Run("list is empty", func(t *testing.T) {
		list := New()

		realResult := list.Size()
		expectedResult := 0

		if realResult != expectedResult {
			t.Errorf("real result: %v, expected result: %v", realResult, expectedResult)
		}
	})

	t.Run("list is not empty", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, "cat", "dog", "lion", "fox")
		list, _ := NewFillingList(elems, t)

		realResult := list.Size()
		expectedResult := 4

		if realResult != expectedResult {
			t.Errorf("real result: %v, expected result: %v", realResult, expectedResult)
		}
	})
}

func TestLinkedList_Contains(t *testing.T) {
	t.Run("list is empty", func(t *testing.T) {
		list := New()

		realResult, _ := list.Contains(5)
		expectedResult := false

		if realResult != expectedResult {
			t.Errorf("real result: %v, expected result: %v", realResult, expectedResult)
		}
	})

	t.Run("list contains different data type", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24, 15)
		list, _ := NewFillingList(elems, t)

		_, realError := list.Contains("cat")
		expectedError := "list contains data of a different type"

		if realError.Error() != expectedError {
			t.Errorf("real result: %v, expected result: %v", realError.Error(), expectedError)
		}
	})

	t.Run("list is not empty", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24, 15)
		list, _ := NewFillingList(elems, t)

		realResult, _ := list.Contains(24)
		expectedResult := true

		if realResult != expectedResult {
			t.Errorf("real result: %v, expected result: %v", realResult, expectedResult)
		}
	})
}

func TestLinkedList_AddList(t *testing.T) {
	t.Run("list is empty. Adding an empty list", func(t *testing.T) {
		list := New()
		listForAdding := New()

		_, err := list.AddList(listForAdding)
		expectedError := "passed list is empty"

		if err.Error() != expectedError {
			t.Errorf("expected error = %v and real error = %v", expectedError, err.Error())
		}
	})

	t.Run("list is empty. Adding a non-empty list", func(t *testing.T) {
		list := New()
		listForAdding := New()
		listForAdding.AddElements(12, 15, 20, 25)

		isSuccess, _ := list.AddList(listForAdding)
		resultList := list.Sprint()
		expectedList := "[12, 15, 20, 25]"

		if isSuccess != true || resultList != expectedList {
			t.Errorf("expected list = %v and result list = %v", resultList, expectedList)
		}
	})

	t.Run("list is not empty. Adding elements with different data type", func(t *testing.T) {
		list := New()
		list.AddElements(1, 2, 3)
		listForAdding := New()
		listForAdding.AddElements("tree", "graph")

		isSuccess, err := list.AddList(listForAdding)
		expectedResult := false
		expectedError := "lists contain different types of data"

		if isSuccess != expectedResult || err.Error() != expectedError {
			t.Errorf("adding elements faild. Real error: %v, expected error: %v", err.Error(), expectedError)
		}
	})

	t.Run("list is not empty. Adding a non-empty list", func(t *testing.T) {
		list := New()
		list.AddElements(1, 2, 3)
		listForAdding := New()
		listForAdding.AddElements(12, 15, 20, 25)

		isSuccess, _ := list.AddList(listForAdding)
		resultList := list.Sprint()
		expectedList := "[1, 2, 3, 12, 15, 20, 25]"

		if isSuccess != true || resultList != expectedList {
			t.Errorf("expected list = %v and result list = %v", resultList, expectedList)
		}
	})
}

func TestLinkedList_Insert(t *testing.T) {
	t.Run("list is not empty. Adding element with different data type", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, "cat", "dog", "lion", "mouse")
		list, _ := NewFillingList(elems, t)

		_, err := list.Insert(50, 3)
		expectedError := "invalid data type for adding"

		if err.Error() != expectedError {
			t.Errorf("expected error: %v, real error: %v", expectedError, err)
		}
	})

	t.Run("list is empty. Index > 0", func(t *testing.T) {
		list := New()

		_, err := list.Insert("element", 2)
		expectedError := "invalid index. Out of range"

		if err.Error() != expectedError {
			t.Errorf("expected error: %v, real error: %v", expectedError, err)
		}
	})

	t.Run("list is not empty. Index = 0", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24, 15, 11, 2)
		list, _ := NewFillingList(elems, t)

		isSuccess, _ := list.Insert(0, 0)
		resultList := list.Sprint()
		expectedList := "[0, 12, 24, 15, 11, 2]"
		resultSize := list.Size()
		expectedSize := 6

		if resultList != expectedList || isSuccess == false || resultSize != expectedSize {
			t.Errorf("expected list: %v, result list: %v. Result size = %d, expected size = %d. Is success = %t",
				expectedList, resultList, resultSize, expectedSize, isSuccess)
		}
	})

	t.Run("list is not empty. Index = size of list", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24, 15, 11, 2)
		list, _ := NewFillingList(elems, t)

		isSuccess, _ := list.Insert(0, 5)
		resultList := list.Sprint()
		expectedList := "[12, 24, 15, 11, 2, 0]"
		resultSize := list.Size()
		expectedSize := 6

		if resultList != expectedList || isSuccess == false || resultSize != expectedSize {
			t.Errorf("expected list: %v, result list: %v. Result size = %d, expected size = %d. Is success = %t",
				expectedList, resultList, resultSize, expectedSize, isSuccess)
		}
	})

	t.Run("list is not empty. Adding to the middle of the list", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24, 15, 11, 2)
		list, _ := NewFillingList(elems, t)

		isSuccess, _ := list.Insert(0, 3)
		resultList := list.Sprint()
		expectedList := "[12, 24, 15, 0, 11, 2]"
		resultSize := list.Size()
		expectedSize := 6

		if resultList != expectedList || isSuccess == false || resultSize != expectedSize {
			t.Errorf("expected list: %v, result list: %v. Result size = %d, expected size = %d. Is success = %t",
				expectedList, resultList, resultSize, expectedSize, isSuccess)
		}
	})
}

func TestLinkedList_First(t *testing.T) {
	t.Run("list is empty", func(t *testing.T) {
		list := New()

		realResult, err := list.First()
		expectedError := "list is empty"

		if err.Error() != expectedError || realResult != nil {
			t.Errorf("expected error = %v and real error = %v. Real result = %v, expected = %v",
				expectedError, err.Error(), realResult, nil)
		}
	})

	t.Run("list is not empty", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24, 15)
		list, _ := NewFillingList(elems, t)

		realResult, _ := list.First()
		expectedResult := 12

		if realResult != expectedResult {
			t.Errorf("real result = %v, expected result = %v", realResult, expectedResult)
		}
	})
}

func TestLinkedList_Last(t *testing.T) {
	t.Run("list is empty", func(t *testing.T) {
		list := New()

		realResult, err := list.Last()
		expectedError := "list is empty"

		if err.Error() != expectedError || realResult != nil {
			t.Errorf("expected error = %v and real error = %v. Real result = %v, expected = %v",
				expectedError, err.Error(), realResult, nil)
		}
	})

	t.Run("list is not empty", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24, 15)
		list, _ := NewFillingList(elems, t)

		realResult, _ := list.Last()
		expectedResult := 15

		if realResult != expectedResult {
			t.Errorf("real result = %v, expected result = %v", realResult, expectedResult)
		}
	})
}

func TestLinkedList_RemoveByIndex(t *testing.T) {
	t.Run("list is empty", func(t *testing.T) {
		list := New()

		err := list.RemoveByIndex(0)
		actualError := err.Error()
		expectedError := "list is empty"

		if actualError != expectedError {
			t.Errorf("actual error: %v, expected error: %v", actualError, expectedError)
		}
	})

	t.Run("list is not empty. Index = 0", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 17, 29, 7, 8, 9, 11)
		list, _ := NewFillingList(elems, t)

		list.RemoveByIndex(0)
		actualResult := list.Sprint()
		expectedResult := "[29, 7, 8, 9, 11]"

		if actualResult != expectedResult {
			t.Errorf("actual result: %v, expected result: %v", actualResult, expectedResult)
		}
	})

	t.Run("list is not empty. Index = size of list", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 17, 29, 7, 8, 9, 11)
		list, _ := NewFillingList(elems, t)

		list.RemoveByIndex(5)
		actualResult := list.Sprint()
		expectedResult := "[17, 29, 7, 8, 9]"

		if actualResult != expectedResult {
			t.Errorf("actual result: %v, expected result: %v", actualResult, expectedResult)
		}
	})

	t.Run("list is not empty. Removing from middle of the list", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 17, 29, 7, 8, 9, 11)
		list, _ := NewFillingList(elems, t)

		list.RemoveByIndex(3)
		actualResult := list.Sprint()
		expectedResult := "[17, 29, 7, 9, 11]"

		if actualResult != expectedResult {
			t.Errorf("actual result: %v, expected result: %v", actualResult, expectedResult)
		}
	})

	t.Run("list is not empty. Index out of range", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 17, 29, 7, 8, 9, 11)
		list, _ := NewFillingList(elems, t)

		err := list.RemoveByIndex(6)
		actualError := err.Error()
		expectedError := "invalid index"

		if actualError != expectedError {
			t.Errorf("actual error: %v, expected error: %v", actualError, expectedError)
		}
	})

	t.Run("list is not empty. Index out of range", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 17, 29, 7, 8, 9, 11)
		list, _ := NewFillingList(elems, t)

		err := list.RemoveByIndex(-1)
		actualError := err.Error()
		expectedError := "invalid index"

		if actualError != expectedError {
			t.Errorf("actual error: %v, expected error: %v", actualError, expectedError)
		}
	})
}

func TestLinkedList_SetByIndex(t *testing.T) {
	t.Run("list is not empty. Adding element with different data type", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, "cat", "dog", "lion", "mouse")
		list, _ := NewFillingList(elems, t)

		err := list.SetByIndex(50, 3)
		expectedError := "invalid data type for adding"

		if err.Error() != expectedError {
			t.Errorf("expected error: %v, real error: %v", expectedError, err)
		}
	})

	t.Run("list is empty. Index > 0", func(t *testing.T) {
		list := New()

		err := list.SetByIndex("element", 2)
		expectedError := "invalid index. Out of range"

		if err.Error() != expectedError {
			t.Errorf("expected error: %v, real error: %v", expectedError, err)
		}
	})

	t.Run("list is not empty. Index = 0", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24, 15, 11, 2)
		list, _ := NewFillingList(elems, t)

		list.SetByIndex(0, 0)
		resultList := list.Sprint()
		expectedList := "[0, 24, 15, 11, 2]"
		resultSize := list.Size()
		expectedSize := 5

		if resultList != expectedList || resultSize != expectedSize {
			t.Errorf("expected list: %v, result list: %v. Result size = %d, expected size = %d.",
				expectedList, resultList, resultSize, expectedSize)
		}
	})

	t.Run("list is not empty. Index = size of list", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, "cat", "dog", "lion", "mouse")
		list, _ := NewFillingList(elems, t)

		err := list.SetByIndex("error", 4)
		expectedError := "invalid index. Out of range"

		if err.Error() != expectedError {
			t.Errorf("expected error: %v, real error: %v", expectedError, err)
		}
	})

	t.Run("list is not empty. Adding to the middle of the list", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 12, 24, 15, 11, 2)
		list, _ := NewFillingList(elems, t)

		list.SetByIndex(0, 3)
		resultList := list.Sprint()
		expectedList := "[12, 24, 15, 0, 2]"
		resultSize := list.Size()
		expectedSize := 5

		if resultList != expectedList || resultSize != expectedSize {
			t.Errorf("expected list: %v, result list: %v. Result size = %d, expected size = %d.",
				expectedList, resultList, resultSize, expectedSize)
		}
	})
}

func TestLinkedList_RemoveFirst(t *testing.T) {
	t.Run("list is empty", func(t *testing.T) {
		list := New()

		err := list.RemoveFirst()
		actualError := err.Error()
		expectedError := "the list is empty"

		if actualError != expectedError {
			t.Errorf("actual error: %v, expected error: %v", actualError, expectedError)
		}
	})

	t.Run("list is not empty", func(t *testing.T) {
		elems := make([]interface{}, 0, 5)
		elems = append(elems, 17, 29, 7, 8, 9, 11)
		list, _ := NewFillingList(elems, t)

		list.RemoveFirst()
		actualResult := list.Sprint()
		expectedResult := "[29, 7, 8, 9, 11]"

		if actualResult != expectedResult {
			t.Errorf("actual result: %v, expected result: %v", actualResult, expectedResult)
		}
	})
}
