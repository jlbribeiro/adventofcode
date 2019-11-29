package polymers_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2018/day5/part1/polymers"
)

func TestReact(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "example",
			input:    "dabAcCaCBAcCcaDA",
			expected: "dabCBAcaDA",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := polymers.React(tc.input)
			if actual != tc.expected {
				t.Errorf("expected '%s', got '%s'", tc.expected, actual)
			}
		})
	}
}
