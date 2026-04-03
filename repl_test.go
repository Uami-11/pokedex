package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	t.Parallel()
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello   world   bye   ",
			expected: []string{"hello", "world", "bye"},
		},
		{
			input:    "   hello   world   bye   ",
			expected: []string{"hello", "world", "bye"},
		},
		{
			input:    "   PIKAchu   Volcarona   aegislash   ",
			expected: []string{"pikachu", "volcarona", "aegislash"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Expected %d pokemon, got %d pokemon instead.", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("INPUT: %q\nExpected: %q\nActual:  %q", c.input, expectedWord, word)
			}
		}
	}
}
