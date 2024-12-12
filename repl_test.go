package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " Hello World ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    " HOWDY partner",
			expected: []string{"howdy", "partner"},
		},
	}

	for _, c := range cases {
		actual := cleanInputString(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Actual and expected slices have different length. Expected: %v, Actual: %v", c.expected, actual)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words do not match. Expected: %v, Actual: %v", expectedWord, word)
			}
		}
	}
}
