package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "chAriZard typhlOSion   Eevee    ",
			expected: []string{"charizard", "typhlosion", "eevee"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("got %d words, want %d words", len(actual), len(c.expected))
			// You might also want to show what the slices contain
			t.Errorf("got %v, want %v", actual, c.expected)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("For input %q, got word %q at position %d, want %q",
					c.input, word, i, expectedWord)
			}
		}
	}
}
