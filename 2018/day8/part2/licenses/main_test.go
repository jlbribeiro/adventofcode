package licenses_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2018/day8/part2/licenses"
)

func TestFromTreeInput(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "example",
			input:    `2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`,
			expected: 66,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			actual := licenses.FromTreeInput(reader)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
