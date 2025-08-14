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
			input:    "   hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "IM ALL    CAPS! !  ",
			expected: []string{"im", "all", "caps!", "!"},
		},
		{
			input:    "Nothing Special",
			expected: []string{"nothing", "special"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("index length of output (%v) doesnt match expected length (%v)", len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word in output (%v) doesnt match expected word (%v)", word, expectedWord)
			}
		}
	}
}
