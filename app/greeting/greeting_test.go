package greeting

import (
	"testing"
)

func TestShouldSayHello(t *testing.T) {
	greeting := SayHello("Ivo")

	if "Hello, Ivo!" != greeting {
		t.Error("Should say hello to Ivo", greeting)
	}
}
