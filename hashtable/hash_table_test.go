package hashtable

import "testing"

func TestHashtable_Size(t *testing.T) {
	t.Run("hashtable is empty", func(t *testing.T) {
		table := New()

		actualSize := table.Size()
		expectedSize := 0

		if actualSize != expectedSize {
			t.Errorf("expected size=%d, and actual size=%d", expectedSize, actualSize)
		}
	})

	t.Run("hashtable is not empty", func(t *testing.T) {
		table := New()
		table.Put("first", 21)
		table.Put("second", 32)

		actualSize := table.Size()
		expectedSize := 2

		if actualSize != expectedSize {
			t.Errorf("expected size=%d, and actual size=%d", expectedSize, actualSize)
		}
	})
}

func TestHashtable_RemoveAll(t *testing.T) {
	t.Run("hashtable is empty", func(t *testing.T) {
		table := New()
		table.RemoveAll()

		actualSize := table.Size()
		expectedSize := 0

		if actualSize != expectedSize {
			t.Errorf("expected size=%d, and actual size=%d", expectedSize, actualSize)
		}
	})

	t.Run("hashtable is not empty", func(t *testing.T) {
		table := New()
		table.Put("first", 21)
		table.Put("second", 32)
		table.RemoveAll()

		actualSize := table.Size()
		expectedSize := 0

		if actualSize != expectedSize {
			t.Errorf("expected size=%d, and actual size=%d", expectedSize, actualSize)
		}
	})

	t.Run("hashtable is not empty. After removing adding one element", func(t *testing.T) {
		table := New()
		table.Put("first", 21)
		table.Put("second", 32)
		table.RemoveAll()

		table.Put("third", 43)
		actualSize := table.Size()
		expectedSize := 1

		if actualSize != expectedSize {
			t.Errorf("expected size=%d, and actual size=%d", expectedSize, actualSize)
		}
	})
}

func TestHashtable_Put(t *testing.T) {
	t.Run("adding one element", func(t *testing.T) {
		table := New()
		key := "first"
		value := "first value"
		table.Put(key, value)

		actualSize := table.Size()
		expectedSize := 1

		if actualSize != expectedSize {
			t.Errorf("expected size=%d, and actual size=%d", expectedSize, actualSize)
		}
	})

	t.Run("adding two elements", func(t *testing.T) {
		table := New()
		table.Put("first", 21)
		table.Put("second", 32)

		actualSize := table.Size()
		expectedSize := 2

		if actualSize != expectedSize {
			t.Errorf("expected size=%d, and actual size=%d", expectedSize, actualSize)
		}
	})

	t.Run("adding element with nil value", func(t *testing.T) {
		table := New()
		_, err := table.Put("key", nil)

		actualError := err.Error()
		expectedError := "value is nil"

		if actualError != expectedError {
			t.Errorf("expected error=%s, and actual error=%s", expectedError, actualError)
		}
	})

	t.Run("resize table", func(t *testing.T) {
		table := New()
		table.Put("first", "first")
		table.Put("second", "second")
		table.Put("test", "test")
		table.Put("lion", "lion")
		table.Put("mouse", "mouse")
		table.Put("elephant", "elephant")
		table.Put("dog", "dog")
		table.Put("cat", "cat")

		actualSize := table.Size()
		expectedSize := 8
		actualCapacity := table.capacity
		expectedCapacity := 20

		if actualSize != expectedSize || actualCapacity != expectedCapacity {
			t.Errorf("expected size=%d, and actual size=%d. Expected capacity=%d, actual capacity=%d",
				expectedSize, actualSize, expectedCapacity, actualCapacity)
		}
	})
}

func TestHashtable_Find(t *testing.T) {
	t.Run("hashtable is empty", func(t *testing.T) {
		table := New()

		actualResult := table.Find("f")
		expectedResult := false

		if actualResult != expectedResult {
			t.Errorf("does element exist? expected result=%t, and actual result=%t", expectedResult, actualResult)
		}
	})

	t.Run("hashtable is not empty", func(t *testing.T) {
		table := New()
		table.Put("first", 21)
		table.Put("second", 32)

		actualResult := table.Find("second")
		expectedResult := true

		if actualResult != expectedResult {
			t.Errorf("does element exist? expected result=%t, and actual result=%t", expectedResult, actualResult)
		}
	})
}

func TestHashtable_Get(t *testing.T) {
	t.Run("hashtable is empty", func(t *testing.T) {
		table := New()

		actualResult := table.Get("f")

		if actualResult != nil {
			t.Errorf("expected result=%v, and actual result=%v",
				actualResult, nil)
		}
	})

	t.Run("hashtable is not empty. Element exists", func(t *testing.T) {
		table := New()
		table.Put("first", 21)
		table.Put("second", 32)

		actualResult := table.Get("second")
		expectedResult := 32

		if expectedResult != (*actualResult).(int) {
			t.Errorf("expected result=%v, and actual result=%v", expectedResult, actualResult)
		}
	})

	t.Run("hashtable is not empty. Element doesn't exist", func(t *testing.T) {
		table := New()
		table.Put("first", 21)
		table.Put("second", 32)

		actualResult := table.Get("non-exists")

		if actualResult != nil {
			t.Errorf("expected result=%v, and actual result=%v", nil, actualResult)
		}
	})
}

func TestHashtable_Remove(t *testing.T) {
	t.Run("hashtable is empty", func(t *testing.T) {
		table := New()

		actualError := table.Remove("element")
		expectedError := "element with given key does not exist"

		if actualError.Error() != expectedError {
			t.Errorf("expected error=%v, and actual error=%v",
				expectedError, actualError)
		}
	})

	t.Run("hashtable is not empty. Element exists", func(t *testing.T) {
		table := New()
		table.Put("first", 21)
		table.Put("second", 32)

		table.Remove("first")

		actualResult := table.Find("first")
		expectedResult := false

		if actualResult != expectedResult {
			t.Errorf("does element exist? expected result=%t, and actual result=%t", expectedResult, actualResult)
		}
	})

	t.Run("hashtable is not empty. Element doesn't exist", func(t *testing.T) {
		table := New()
		table.Put("first", 21)
		table.Put("second", 32)

		err := table.Remove("third")
		actualError := err.Error()
		expectedError := "element with given key does not exist"

		if expectedError != actualError {
			t.Errorf("does element exist? expected error=%v, and actual error=%v", expectedError, actualError)
		}
	})
}

func TestHashtable_Replace(t *testing.T) {
	t.Run("hashtable is empty", func(t *testing.T) {
		table := New()

		err := table.Replace("f", "new value")
		actualError := err.Error()
		expectedError := "value cannot be replaced because element with given key does not exist"

		if expectedError != actualError {
			t.Errorf("expected error=%v, and actual error=%v", expectedError, actualError)
		}
	})

	t.Run("hashtable is not empty. Element exists", func(t *testing.T) {
		table := New()
		table.Put("first", 21)
		table.Put("second", 32)

		table.Replace("second", 44)
		actualResult := table.Get("second")
		expectedResult := 44

		if expectedResult != (*actualResult).(int) {
			t.Errorf("expected result=%v, and actual result=%v", expectedResult, actualResult)
		}
	})

	t.Run("hashtable is not empty. Element doesn't exist", func(t *testing.T) {
		table := New()
		table.Put("first", 21)
		table.Put("second", 32)

		err := table.Replace("f", "new value")
		actualError := err.Error()
		expectedError := "value cannot be replaced because element with given key does not exist"

		if expectedError != actualError {
			t.Errorf("expected error=%v, and actual error=%v", expectedError, actualError)
		}
	})
}
