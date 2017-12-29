package pass_test

import (
	"testing"

	"git.jlbribeiro.com/adventofcode/day4/part1/pass"
)

var validPassphraseTests = []struct {
	input    string
	expected bool
}{
	{"aa bb cc dd ee", true},
	{"aa bb cc dd aa", false},
	{"aa bb cc dd aaa", true},
}

func TestValidPassphrase(t *testing.T) {
	for _, tt := range validPassphraseTests {
		actual := pass.ValidPassphrase(tt.input)
		if actual != tt.expected {
			t.Errorf("ValidPassphrase(%s): expected %v, got %v", tt.input, tt.expected, actual)
		}
	}
}
