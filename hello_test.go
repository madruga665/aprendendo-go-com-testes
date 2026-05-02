package main

import "testing"

func TestHello(t *testing.T) {
	verifyMessage := func(t testing.TB, result string, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result %s, expected %s", result, expected)
		}
	}
	t.Run("Say hello for people", func(t *testing.T) {
		result := hello("madruga665")
		expected := "Hello, madruga665"

		verifyMessage(t, result, expected)
	})

	t.Run("Say hello World when name is empty string", func(t *testing.T) {
		result := hello("")
		expected := "Hello, world"

		verifyMessage(t, result, expected)
	})
}
