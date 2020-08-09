package main

import (
	"testing"
)

func TestShouldGreet(t *testing.T) {
	greeting := Greet("Ivo")

	if "Hello, Ivo!" != greeting {
		t.Error("Should say hello to Ivo", greeting)
	}
}
