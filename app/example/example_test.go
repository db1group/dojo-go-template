package example

import (
	"testing"
)

func TestShouldReturnSomeExpected(t *testing.T) {
	actual := MultiParametrosEUmRetorno(1, 13)

	if "private example" != actual {
		t.Error("Test public failed", actual)
	}

}

func TestShouldReturnSomeExpectedTwo(t *testing.T) {
	actual := MultiParametrosEUmRetorno(1, 17)

	if "private example" != actual {
		t.Error("Test public failed", actual)
	}

}
