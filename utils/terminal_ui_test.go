package utils

import "testing"

func TestMakeBold(t *testing.T) {
	input := "Hello, %s!"
	expected := "\033[1mHello, World!\033[0m"
	result := MakeBold(input, "World")

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}
