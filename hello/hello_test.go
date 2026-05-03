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
		result := hello("madruga665", "english")
		expected := "Hello, madruga665"

		verifyMessage(t, result, expected)
	})

	t.Run("Say hello World when name is empty string", func(t *testing.T) {
		result := hello("", "english")
		expected := "Hello, world"

		verifyMessage(t, result, expected)
	})

	t.Run("Say hello in portuguese when languange is portuguese", func(t *testing.T) {
		result := hello("Dante", "portuguese")
		expected := "Olá, Dante"

		verifyMessage(t, result, expected)
	})

	t.Run("Say hello in spanish when languange is spanish", func(t *testing.T) {
		result := hello("Dante", "spanish")
		expected := "Hola, Dante"

		verifyMessage(t, result, expected)
	})
}
