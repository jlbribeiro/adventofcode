package marbles_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2018/day9/part2/marbles"
)

func TestPlay(t *testing.T) {
	var testCases = []struct {
		name       string
		nPlayers   int
		lastMarble int
		expected   int
	}{
		{
			name:       "example",
			nPlayers:   9,
			lastMarble: 25,
			expected:   32,
		},
		{
			name:       "example2",
			nPlayers:   10,
			lastMarble: 1618,
			expected:   8317,
		},
		{
			name:       "example3",
			nPlayers:   13,
			lastMarble: 7999,
			expected:   146373,
		},
		{
			name:       "example4",
			nPlayers:   17,
			lastMarble: 1104,
			expected:   2764,
		},
		{
			name:       "example5",
			nPlayers:   21,
			lastMarble: 6111,
			expected:   54718,
		},
		{
			name:       "example6",
			nPlayers:   30,
			lastMarble: 5807,
			expected:   37305,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := marbles.Play(tc.nPlayers, tc.lastMarble)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
