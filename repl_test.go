package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		want  []string
	}{
		{"  hello  world  ", []string{"hello", "world"}},
		{"foo			bar", []string{"foo", "bar"}},
		{"HellOOOOO wOrLLLLld", []string{"hellooooo", "worllllld"}},
		{"", []string{}},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.want) {
			t.Errorf("cleanInput(%q) returned %d words, want %d words", c.input, len(actual), len(c.want))
			continue
		}

		for i, word := range actual {
			if word != c.want[i] {
				t.Errorf("cleanInput(%q) returned word %q at index %d, want %q", c.input, word, i, c.want[i])
			}
		}
	}
}
