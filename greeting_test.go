package main

import (
	"testing"
)

// a sample test using test tables
func TestGreet(t *testing.T) {
	// create a test table with some cases
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "Ivo",
			expected: "Hello, Ivo!",
		},
		{
			input:    "Lucas",
			expected: "Hello, Lucas!",
		},
	}

	// iterate over the cases and run the test for each one of them
	for _, tt := range cases {
		// the first parameter is the test case name
		// the anonymous function executes the actual test
		t.Run(tt.input, func(t *testing.T) {
			actual := Greet(tt.input)

			if actual != tt.expected {
				t.Errorf("Expected message to be: [%s], but got: [%s]", tt.expected, actual)
			}
		})
	}
}
