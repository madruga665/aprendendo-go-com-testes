package main

import "testing"

func TestHello(t *testing.T) {
	result := hello("madruga665")
	expected := "Olá, madruga665"

	if result != expected {
		t.Errorf("result %s, expected %s", result, expected)
	}
}
