package pass_test

import (
	"fmt"
	"testing"

	"github.com/jlbribeiro/adventofcode/2017/day4/part2/pass"
)

var validPassphraseTests = []struct {
	input    string
	expected bool
}{
	{"abcde fghij", true},
	{"abcde xyz ecdab", false},
	{"a ab abc abd abf abj", true},
	{"iiii oiii ooii oooi oooo", true},
	{"oiii ioii iioi iiio", false},
}

func TestValidPassphrase(t *testing.T) {
	for _, tt := range validPassphraseTests {
		t.Run(fmt.Sprintf("Check if \"%s\" is a valid passphrase", tt.input), func(t *testing.T) {
			actual := pass.ValidPassphrase(tt.input)
			if actual != tt.expected {
				t.Errorf("ValidPassphrase(%s): expected %v, got %v", tt.input, tt.expected, actual)
			}
		})
	}
}
