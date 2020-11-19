package shuffle_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day22/part2/shuffle"
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
			deck := shuffle.NewDeck(int64(tc.nCards))
			deck.Shuffle(input, 1)

			for expectedInd := range tc.expected {
				actualInd := deck.FindCard(int64(tc.expected[expectedInd]))
				if actualInd != int64(expectedInd) {
					t.Fatalf("expected to find card %d in index %v, got %d", tc.expected[expectedInd], expectedInd, actualInd)
				}
			}
		})
	}
}

func TestShuffleAgainstReferenceImplementation(t *testing.T) {
	routine := `deal into new stack
cut -2
deal with increment 7
cut 8
cut -4
deal into new stack
deal with increment 7
cut 3
deal with increment 9
deal with increment 3
cut -1`

	var testCases = []struct {
		nCards int64
		nTimes int64
	}{
		{
			nCards: 23,
			nTimes: 30,
		},
		{
			nCards: 53,
			nTimes: 101,
		},
	}

	for _, tc := range testCases {
		tc := tc
		tName := fmt.Sprintf("cards_%d_nTimes_%d", tc.nCards, tc.nTimes)

		t.Run(tName, func(t *testing.T) {
			deck := shuffle.NewDeck(tc.nCards)
			deck.Shuffle(strings.NewReader(routine), tc.nTimes)
			cards := deck.Cards()

			refDeck := shuffle.NewDeckRefImpl(tc.nCards)
			refDeck.Shuffle(strings.NewReader(routine), tc.nTimes)
			refCards := refDeck.Cards()
			for i, card := range refCards {
				actual := cards[i]
				if actual != card {
					t.Logf("expected: %v", refCards)
					t.Logf("  actual: %v", cards)
					t.Fatalf("expected to find %d in %d, found %d", card, i, actual)
				}
			}
		})
	}
}
