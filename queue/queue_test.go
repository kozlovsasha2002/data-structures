package queue

import (
	"testing"
)

func TestQueue_IsEmpty(t *testing.T) {
	t.Run("queue is empty", func(t *testing.T) {
		q := New()

		actual := q.IsEmpty()
		expected := true

		if actual != expected {
			t.Errorf("actual = %t and expected = %t", actual, expected)
		}
	})

	t.Run("queue contains multi elements", func(t *testing.T) {
		strings := make([]interface{}, 0)
		strings = append(strings, 5, 2)
		q := NewFilledQueue(strings, t)

		actual := q.IsEmpty()
		expected := false

		if actual != expected {
			t.Errorf("actual = %t and expected = %t", actual, expected)
		}
	})
}

func TestQueue_Push(t *testing.T) {
	t.Run("adding one element", func(t *testing.T) {
		q := New()

		actual, err := q.Push(15)
		expected := 15

		if err != nil || actual != expected {
			t.Errorf("actual = %v and expected = %v, error = %v", actual, expected, err.Error())
		}
	})

	t.Run("adding three elements", func(t *testing.T) {
		q := New()

		actuals := make([]int, 0)
		actuals = append(actuals, 15, 20, 25)

		for _, item := range actuals {
			expected := item
			actual, err := q.Push(item)
			if err != nil || actual != expected {
				t.Errorf("actual = %v and expected = %v, error = %v", actual, expected, err.Error())
			}
		}
	})

	t.Run("adding elements with different data type", func(t *testing.T) {
		q := New()

		actuals := make([]interface{}, 0)
		actuals = append(actuals, 15, "dog")

		for _, item := range actuals {
			expected := item
			actual, err := q.Push(item)
			if err == nil && actual != expected {
				t.Errorf("actual = %v and expected = %v, error = %v", actual, expected, err.Error())
			}
		}
	})
}

func TestQueue_Peek(t *testing.T) {
	t.Run("queue is empty", func(t *testing.T) {
		q := New()

		_, err := q.Peek()

		if err.Error() != "queue is empty" {
			t.Errorf("error in Peek method: %s", err.Error())
		}
	})

	t.Run("queue contains multi elements", func(t *testing.T) {
		strings := make([]interface{}, 0)
		strings = append(strings, 12, 14, 20)
		q := NewFilledQueue(strings, t)

		actual, err := q.Peek()
		actualSize := q.Size()
		expectedValue := 12
		expectedSize := 3

		if err != nil {
			panic(err.Error())
		}

		if expectedValue != actual || expectedSize != actualSize {
			t.Errorf("actual returning value = %v and expected = %v, expected queue size = %d, actual size = %d",
				actual, expectedValue, expectedSize, actualSize)
		}
	})
}

func TestQueue_Pop(t *testing.T) {
	t.Run("queue is empty", func(t *testing.T) {
		q := New()

		actual := q.Pop()

		if actual != nil {
			t.Errorf("expected returning nil, and actual = %v", actual)
		}
	})

	t.Run("queue contains two elements, method works one time", func(t *testing.T) {
		strings := make([]interface{}, 0)
		strings = append(strings, "dog", "cat", "mouse")
		q := NewFilledQueue(strings, t)

		actual := q.Pop()
		expected := "dog"

		if actual != expected {
			t.Errorf("actual result = %v and expected result = %v", actual, expected)
		}
	})
}

func TestQueue_Size(t *testing.T) {
	t.Run("queue is empty", func(t *testing.T) {
		q := New()

		actual := q.Size()
		expected := 0

		if actual != expected {
			t.Errorf("actual = %v and expected = %v", actual, expected)
		}
	})

	t.Run("queue contains multi elements", func(t *testing.T) {
		strings := make([]interface{}, 0)
		strings = append(strings, "dog", "cat", "mouse", "tiger")
		q := NewFilledQueue(strings, t)

		actual := q.Size()
		expected := 4

		if actual != expected {
			t.Errorf("actual = %v and expected = %v", actual, expected)
		}
	})

	t.Run("after deleting one element", func(t *testing.T) {
		strings := make([]interface{}, 0)
		strings = append(strings, "dog", "cat", "mouse", "lion")
		q := NewFilledQueue(strings, t)
		elem := q.Pop()

		actual := q.Size()
		expected := 3
		expectedElem := "dog"

		if actual != expected || elem != expectedElem {
			t.Errorf("actual = %v and expected = %v, returning elem = %v and expected elem = %v",
				actual, expected, elem, expectedElem)
		}
	})
}

func NewFilledQueue(arr []interface{}, t *testing.T) *Queue {
	q := New()
	t.Helper()

	for _, item := range arr {
		_, err := q.Push(item)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	return q
}
