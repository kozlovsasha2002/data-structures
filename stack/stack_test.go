package stack

import (
	"errors"
	"log"
	"testing"
)

func TestStack_IsEmpty(t *testing.T) {
	st := New()

	t.Run("empty Stack", func(t *testing.T) {
		actualResult := st.IsEmpty()
		expectedResult := true

		if actualResult != expectedResult {
			t.Errorf("Stack must be empty. Actual result = %t and expected result = %t", actualResult, expectedResult)
		}
	})

	t.Run("Stack is not empty", func(t *testing.T) {
		st = newFilledIntStack(t)

		actualResult := st.IsEmpty()
		expectedResult := false

		if actualResult != expectedResult {
			t.Errorf("Stack must be empty. Actual result = %t and expected result = %t", actualResult, expectedResult)
		}
	})
}

func TestStack_Peek(t *testing.T) {
	t.Run("Stack is empty", func(t *testing.T) {
		st := New()

		actualResult, err := st.Peek()

		if err == nil || actualResult != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("Stack contains one element", func(t *testing.T) {
		st := New()
		_, err := st.Push("test")
		if err != nil {
			t.Fatal("error in Push method")
		}

		actualResult, _ := st.Peek()
		expectedResult := "test"

		if actualResult != expectedResult {
			t.Errorf("invalid returning value. Actual result = %s and expected result = %s", actualResult, expectedResult)
		}
	})

	t.Run("Stack contains more than one element", func(t *testing.T) {
		st := newFilledStringStack(t)

		actualResult, _ := st.Peek()
		actualLength := st.Size()
		expectedResult := "cat"
		expectedLength := 3

		if actualResult != expectedResult || actualLength != expectedLength {
			t.Errorf("peek method error. Actual result = \"%s\", length = %d and expected result = \"%s\", length = %d",
				actualResult, actualLength,
				expectedResult, expectedLength,
			)
		}
	})
}

func TestStack_Pop(t *testing.T) {
	t.Run("Stack is empty", func(t *testing.T) {
		st := New()

		actualResult, err := st.Pop()
		expectedError := errors.New("Stack is empty")

		if actualResult != nil || err == nil {
			t.Errorf("actual result = %v and expected result = %v", actualResult, expectedError)
		}
	})

	t.Run("Stack contains one element", func(t *testing.T) {
		st := New()
		_, err := st.Push("one")
		if err != nil {
			t.Fatal("error in Push method")
		}

		actualResult, _ := st.Pop()
		actualLength := st.Size()
		expectedResult := "one"
		expectedLength := 0

		if actualResult != expectedResult || actualLength != expectedLength {
			t.Errorf("actual result = %v, length = %v and expected result = %v, length = %v",
				actualResult, actualLength,
				expectedResult, expectedLength,
			)
		}
	})

	t.Run("Stack contains more than one element", func(t *testing.T) {
		st := newFilledIntStack(t)

		actualResult, _ := st.Pop()
		actualLength := st.Size()
		expectedResult := 5
		expectedLength := 1

		if actualResult != expectedResult || actualLength != expectedLength {
			t.Errorf("actual result = %v, length = %v and expected result = %v, length = %v",
				actualResult, actualLength,
				expectedResult, expectedLength,
			)
		}
	})
}

func TestStack_Push(t *testing.T) {
	t.Run("Stack is empty", func(t *testing.T) {
		st := New()

		actualResult, _ := st.Push("dog")
		expectedResult := "dog"

		if actualResult != expectedResult {
			t.Errorf("actual result = %s and expected result = %s", actualResult, expectedResult)
		}
	})

	t.Run("push method works twice", func(t *testing.T) {
		st := New()

		actualResult := make([]string, 0, 2)
		for i, _ := range actualResult {
			item, err := st.Push(i)
			if err != nil {
				t.Fatal(err.Error())
			}
			actualResult[i] = item.(string)
		}
		expectedResult := []string{"rest", "pain"}

		for i, _ := range actualResult {
			if actualResult[i] != expectedResult[i] {
				t.Errorf("actual result = %v and expected = %v", actualResult, expectedResult)
			}
		}
	})

	t.Run("adding elements with different data type", func(t *testing.T) {
		st := New()

		actuals := make([]interface{}, 0)
		actuals = append(actuals, 15, "dog")

		for _, item := range actuals {
			expected := item
			actual, err := st.Push(item)
			if err == nil && actual != expected {
				t.Errorf("actual = %v and expected = %v, error = %v", actual, expected, err.Error())
			}
		}
	})
}

func TestStack_Size(t *testing.T) {
	t.Run("Stack is empty", func(t *testing.T) {
		st := New()

		actual := st.Size()
		expected := 0

		if actual != expected {
			t.Errorf("actual result = %v and expected result = %v", actual, expected)
		}
	})

	t.Run("Stack contains two elements", func(t *testing.T) {
		st := newFilledStringStack(t)

		actual := st.Size()
		expected := 3

		if actual != expected {
			t.Errorf("actual result = %v and expected result = %v", actual, expected)
		}
	})

	t.Run("Stack contains three element after removal", func(t *testing.T) {
		st := newFilledIntStack(t)
		_, err := st.Push(5)
		if err != nil {
			log.Fatal("pushing error")
		}
		_, err = st.Pop()
		if err != nil {
			log.Fatal("pop error")
		}

		actual := st.Size()
		expected := 2

		if actual != expected {
			t.Errorf("actual result = %v and expected result = %v", actual, expected)
		}
	})
}

func newFilledIntStack(t *testing.T) *Stack {
	t.Helper()
	st := New()
	ints := make([]int, 0)
	ints = append(ints, 12, 5)

	for _, item := range ints {
		_, err := st.Push(item)
		if err != nil {
			log.Fatal("pushing error")
		}
	}

	return st
}

func newFilledStringStack(t *testing.T) *Stack {
	t.Helper()
	st := New()
	strs := make([]string, 0)
	strs = append(strs, "dog", "mouse", "cat")

	for _, item := range strs {
		_, err := st.Push(item)
		if err != nil {
			log.Fatal("pushing error")
		}
	}

	return st
}
