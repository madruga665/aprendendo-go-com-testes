package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat test", func(t *testing.T) {
		result := repeat("a", 6)
		expected := "aaaaaa"

		if result != expected {
			t.Errorf("result %s, expected %s", result, expected)

		}
	})
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("a", 6)
	}
}
