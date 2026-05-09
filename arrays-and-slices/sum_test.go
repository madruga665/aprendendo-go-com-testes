package arraysandslices

import (
	"reflect"
	"testing"
)

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

func TestSumAllRest(t *testing.T) {
	t.Run(" sum slices", func(t *testing.T) {
		result := sumAllRest([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result %v expected %v", result, expected)
		}
	})

	t.Run(" sum empty slices", func(t *testing.T) {
		result := sumAllRest([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result %v expected %v", result, expected)
		}
	})
}
