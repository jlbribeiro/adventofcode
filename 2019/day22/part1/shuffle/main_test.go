package shuffle_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day22/part1/shuffle"
)

func TestShuffleFromInput(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		nCards   int
		expected []int
	}{
		{
			name: "example1",
			input: `deal with increment 7
deal into new stack
deal into new stack`,
			nCards:   10,
			expected: []int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7},
		},
		{
			name: "example2",
			input: `cut 6
deal with increment 7
deal into new stack`,
			nCards:   10,
			expected: []int{3, 0, 7, 4, 1, 8, 5, 2, 9, 6},
		},
		{
			name: "example3",
			input: `deal with increment 7
deal with increment 9
cut -2`,
			nCards:   10,
			expected: []int{6, 3, 0, 7, 4, 1, 8, 5, 2, 9},
		},
		{
			name: "example4",
			input: `deal into new stack
cut -2
deal with increment 7
cut 8
cut -4
deal with increment 7
cut 3
deal with increment 9
deal with increment 3
cut -1`,
			nCards:   10,
			expected: []int{9, 2, 5, 8, 1, 4, 7, 0, 3, 6},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.input)
			deck := shuffle.NewDeck(tc.nCards)
			deck.Shuffle(input)

			actual := deck.Cards()
			if len(actual) != len(tc.expected) {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}

			for i := range tc.expected {
				if actual[i] != tc.expected[i] {
					t.Fatalf("expected %v, got %v (index %d, expected %d, got %d)", tc.expected, actual, i, tc.expected[i], actual[i])
				}
			}
		})
	}
}
