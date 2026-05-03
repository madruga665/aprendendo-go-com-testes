package integers

import "testing"

func testAdd(t *testing.T) {

	t.Run("sum numbers", func(t *testing.T) {
		result := add(2, 2)
		expected := 4

		if result != expected {
			t.Errorf("result %d, expected %d", result, expected)
		}
	})
}
