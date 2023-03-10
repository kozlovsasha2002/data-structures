package stack

import (
	"errors"
	"testing"
)

func TestStack_IsEmpty(t *testing.T) {
	st := New()

	t.Run("empty stack", func(t *testing.T) {
		actualResult := st.IsEmpty()
		expectedResult := true

		if actualResult != expectedResult {
			t.Errorf("stack must be empty. Actual result = %t and expected result = %t", actualResult, expectedResult)
		}
	})

	t.Run("stack is not empty", func(t *testing.T) {
		_, err := st.Push(5)
		if err != nil {
			t.Fatal("error in Push method")
		}

		actualResult := st.IsEmpty()
		expectedResult := false

		if actualResult != expectedResult {
			t.Errorf("stack must be empty. Actual result = %t and expected result = %t", actualResult, expectedResult)
		}
	})
}

func TestStack_Peek(t *testing.T) {
	t.Run("stack is empty", func(t *testing.T) {
		st := New()

		actualResult, err := st.Peek()

		if err == nil || actualResult != nil {
			t.Errorf("%s", err.Error())
		}
	})

	t.Run("stack contains one element", func(t *testing.T) {
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

	t.Run("stack contains more than one element", func(t *testing.T) {
		st := New()
		err := st.PushAll("cat", "elephant", "dog")
		if err != nil {
			t.Fatal("error in PushAll method")
		}

		actualResult, _ := st.Peek()
		actualLength := st.Size()
		expectedResult := "dog"
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
	t.Run("stack is empty", func(t *testing.T) {
		st := New()

		actualResult, err := st.Pop()
		expectedError := errors.New("stack is empty")

		if actualResult != nil || err == nil {
			t.Errorf("actual result = %v and expected result = %v", actualResult, expectedError)
		}
	})

	t.Run("stack contains one element", func(t *testing.T) {
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

	t.Run("stack contains more than one element", func(t *testing.T) {
		st := New()
		err := st.PushAll(8, 9, 11)
		if err != nil {
			t.Fatal("error in PushAll method")
		}

		actualResult, _ := st.Pop()
		actualLength := st.Size()
		expectedResult := 11
		expectedLength := 2

		if actualResult != expectedResult || actualLength != expectedLength {
			t.Errorf("actual result = %v, length = %v and expected result = %v, length = %v",
				actualResult, actualLength,
				expectedResult, expectedLength,
			)
		}
	})
}

func TestStack_Push(t *testing.T) {
	t.Run("stack is empty", func(t *testing.T) {
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
}

func TestStack_Size(t *testing.T) {
	t.Run("stack is empty", func(t *testing.T) {
		st := New()

		actual := st.Size()
		expected := 0

		if actual != expected {
			t.Errorf("actual result = %v and expected result = %v", actual, expected)
		}
	})

	t.Run("stack contains two elements", func(t *testing.T) {
		st := New()
		st.PushAll("dog", "cat")

		actual := st.Size()
		expected := 2

		if actual != expected {
			t.Errorf("actual result = %v and expected result = %v", actual, expected)
		}
	})

	t.Run("stack contains three element after removal", func(t *testing.T) {
		st := New()
		st.PushAll("dog", "cat", "elephant", "tiger")
		st.Pop()

		actual := st.Size()
		expected := 3

		if actual != expected {
			t.Errorf("actual result = %v and expected result = %v", actual, expected)
		}
	})
}
