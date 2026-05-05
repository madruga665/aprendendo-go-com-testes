package arraysandslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("array 6 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		result := sum(numbers)
		expected := 21

		if result != expected {
			t.Errorf("result %d, expected %d", result, expected)
		}
	})

	t.Run("array 4 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		result := sum(numbers)
		expected := 10

		if result != expected {
			t.Errorf("result %d, expected %d", result, expected)
		}
	})

}
