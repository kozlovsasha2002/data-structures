package prefixtree

import (
	"testing"
)

func TestPrefixTree_AmountOfAddedElements(t *testing.T) {
	t.Run("tree is empty", func(t *testing.T) {
		root := New()

		actual := root.AmountOfAddedElements()
		expected := 0

		if actual != expected {
			t.Errorf("expected amount of added elements = %d. Real amount = %d", expected, actual)
		}
	})

	t.Run("tree contains multiple elements", func(t *testing.T) {
		root := New()
		inputData := make([]string, 0)
		inputData = append(inputData, "april", "arc", "boy", "book", "boss", "barrel", "bar", "base")
		err := root.AddAll(inputData)
		if err != nil {
			panic(err.Error())
		}

		actual := root.AmountOfAddedElements()
		expected := 8

		if actual != expected {
			t.Errorf("expected amount of added elements = %d. Real amount = %d", expected, actual)
		}
	})

	t.Run("after performing add operation", func(t *testing.T) {
		root := New()
		inputData := make([]string, 0)
		inputData = append(inputData, "april", "arc", "boy", "book", "boss", "barrel", "bar", "base")
		err := root.AddAll(inputData)
		if err != nil {
			panic(err.Error())
		}

		root.Add("case")
		actual := root.AmountOfAddedElements()
		expected := 9

		if actual != expected {
			t.Errorf("expected amount of added elements = %d. Real amount = %d", expected, actual)
		}
	})

	t.Run("after performing remove operation", func(t *testing.T) {
		root := New()
		inputData := make([]string, 0)
		inputData = append(inputData, "april", "arc", "boy", "book", "boss", "barrel", "bar", "base")
		err := root.AddAll(inputData)
		if err != nil {
			panic(err.Error())
		}

		root.Remove("bar")
		actual := root.AmountOfAddedElements()
		expected := 7

		if actual != expected {
			t.Errorf("expected amount of added elements = %d. Real amount = %d", expected, actual)
		}
	})
}

func TestPrefixTree_Add(t *testing.T) {
	t.Run("tree is empty. Adding a non-null string", func(t *testing.T) {
		root := New()

		root.Add("rest")
		actualResult := root.IsContains("rest")
		expectedResult := true

		if actualResult != expectedResult {
			t.Errorf("expected result of adding element: %t and real result: %t", expectedResult, actualResult)
		}
	})

	t.Run("tree is not empty", func(t *testing.T) {
		root := New()
		inputData := make([]string, 0)
		inputData = append(inputData, "chef", "cherry", "char", "chain")
		err := root.AddAll(inputData)
		if err != nil {
			panic(err.Error())
		}

		root.Add("test")
		actualResult := root.IsContains("test")
		expectedResult := true

		if actualResult != expectedResult {
			t.Errorf("expected result of adding element: %t and real result: %t", expectedResult, actualResult)
		}

	})
}

func TestPrefixTree_AddAll(t *testing.T) {
	t.Run("adding zero elements", func(t *testing.T) {
		root := New()
		inputData := make([]string, 0)

		actualError := root.AddAll(inputData)
		expectedError := "can't add an empty slice"

		if actualError.Error() != expectedError {
			t.Errorf("expected result: %s and real result: %s", expectedError, actualError.Error())
		}
	})

	t.Run("adding one element", func(t *testing.T) {
		root := New()
		inputData := make([]string, 0)
		inputData = append(inputData, "apple")
		err := root.AddAll(inputData)
		if err != nil {
			panic(err.Error())
		}

		actualResult := root.IsContains("apple")
		expectedResult := true

		if actualResult != expectedResult {
			t.Errorf("expected result: %t and real result: %t", expectedResult, actualResult)
		}
	})

	t.Run("adding multiple elements", func(t *testing.T) {
		root := New()
		inputData := make([]string, 0)
		inputData = append(inputData, "chef", "cherry", "char")
		err := root.AddAll(inputData)
		if err != nil {
			panic(err.Error())
		}

		actualResult := true
		expectedResult := true

		for _, item := range inputData {
			if !root.IsContains(item) {
				actualResult = false
			}
		}

		if actualResult != expectedResult {
			t.Errorf("expected result: %t and real result: %t", expectedResult, actualResult)
		}
	})

	t.Run("adding multiple elements with empty string", func(t *testing.T) {
		root := New()
		inputData := make([]string, 0)
		inputData = append(inputData, "chef", "cherry", "char", "")

		actualError := root.AddAll(inputData)
		expectedError := "can't insert empty string"

		if actualError.Error() != expectedError {
			t.Errorf("expected result: %s and real result: %s", expectedError, actualError.Error())
		}
	})
}

func TestPrefixTree_IsContains(t *testing.T) {
	t.Run("tree is empty. Desired element = \"\"", func(t *testing.T) {
		root := New()

		actualResult := root.IsContains("")
		expectedResult := false

		if actualResult != expectedResult {
			t.Errorf("expected result: %t and real result: %t", expectedResult, actualResult)
		}
	})

	t.Run("tree is empty", func(t *testing.T) {
		root := New()

		actualResult := root.IsContains("any")
		expectedResult := false

		if actualResult != expectedResult {
			t.Errorf("expected result: %t and real result: %t", expectedResult, actualResult)
		}
	})

	t.Run("tree is not empty. Tree contains the desired element", func(t *testing.T) {
		root := New()
		inputData := make([]string, 0)
		inputData = append(inputData, "chef", "cherry", "char", "chain", "apple", "april")
		err := root.AddAll(inputData)
		if err != nil {
			panic(err.Error())
		}

		actualResult := root.IsContains("april")
		expectedResult := true

		if actualResult != expectedResult {
			t.Errorf("expected result: %t and real result: %t", expectedResult, actualResult)
		}
	})

	t.Run("tree is not empty. Tree doesn't contain the desired element", func(t *testing.T) {
		root := New()
		inputData := make([]string, 0)
		inputData = append(inputData, "chef", "cherry", "char", "chain", "apple", "april")
		err := root.AddAll(inputData)
		if err != nil {
			panic(err.Error())
		}

		actualResult := root.IsContains("china")
		expectedResult := false

		if actualResult != expectedResult {
			t.Errorf("expected result: %t and real result: %t", expectedResult, actualResult)
		}
	})
}

func TestPrefixTree_Remove(t *testing.T) {
	t.Run("tree is empty. Removing empty string", func(t *testing.T) {
		root := New()

		result := root.Remove("")
		expectedResult := false

		if result != expectedResult {
			t.Errorf("expected result: %t and real result: %t", expectedResult, result)
		}
	})

	t.Run("tree is empty. Removing non-empty string", func(t *testing.T) {
		root := New()

		result := root.Remove("any")
		expectedResult := false

		if result != expectedResult {
			t.Errorf("expected result: %t and real result: %t", expectedResult, result)
		}
	})

	t.Run("tree is not empty. Removing existing element", func(t *testing.T) {
		root := New()
		inputData := make([]string, 0)
		inputData = append(inputData, "chef", "cherry", "char", "chain")
		err := root.AddAll(inputData)
		if err != nil {
			panic(err.Error())
		}

		result := root.Remove("cherry")
		expectedResult := true

		isDeleted := root.IsContains("cherry")

		if result != expectedResult {
			t.Errorf("expected result:%t, real result:%t", expectedResult, result)
		}
		if isDeleted {
			t.Errorf("element was not removed")
		}
	})

	t.Run("tree is empty. Removing non-existing element", func(t *testing.T) {
		root := New()
		inputData := make([]string, 0)
		inputData = append(inputData, "chef", "cherry", "char", "chain")
		err := root.AddAll(inputData)
		if err != nil {
			panic(err.Error())
		}

		result := root.Remove("cheese")
		expectedResult := false

		if result != expectedResult {
			t.Errorf("expected result:%t, real result:%t", expectedResult, result)
		}
	})
}

func TestPrefixTree_GetAllByPrefix(t *testing.T) {
	t.Run("tree is empty", func(t *testing.T) {
		root := New()

		_, err := root.GetAllByPrefix("be")
		expectedError := "tree is empty"

		if err.Error() != expectedError {
			t.Errorf("expected error:%s and real error:%s", expectedError, err.Error())
		}
	})

	t.Run("tree contains strings with entered prefix", func(t *testing.T) {
		root := New()
		inputData := make([]string, 0)
		inputData = append(inputData, "april", "arc", "boy", "book", "boss", "barrel", "bar", "base")
		err := root.AddAll(inputData)
		if err != nil {
			panic(err.Error())
		}
		expectedData := make([]string, 0)
		expectedData = append(expectedData, "bar", "barrel", "base")

		result, _ := root.GetAllByPrefix("ba")
		actualResult := equals(result, expectedData, t)
		expectedResult := true

		if expectedResult != actualResult {
			t.Errorf("expected result:%t and real result:%t", expectedResult, actualResult)
		}
	})

	t.Run("tree doesn't contain strings with entered prefix", func(t *testing.T) {
		root := New()
		testData := make([]string, 0)
		testData = append(testData, "april", "arc", "boy", "book", "boss", "barrel", "bar", "base")
		err := root.AddAll(testData)
		if err != nil {
			panic(err.Error())
		}

		_, err = root.GetAllByPrefix("be")
		expectedError := "no words found beginning with prefix = \"be\""

		if err.Error() != expectedError {
			t.Errorf("expected error:%s and real error:%s", expectedError, err.Error())
		}
	})
}

func equals(input []string, expectedData []string, t *testing.T) bool {
	t.Helper()
	if len(input) != len(expectedData) {
		return false
	}

	contains := 0
	for _, item := range expectedData {
		for _, item2 := range input {
			if item == item2 {
				contains++
			}
		}
	}

	if contains != len(expectedData) {
		return false
	} else {
		return true
	}
}
