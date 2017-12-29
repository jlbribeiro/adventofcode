package duet_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2017/day18/part2/duet"
)

var duetPlayTT = []struct {
	name              string
	instructionsInput string
	expected          int
}{
	{"example", `snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d`, 3},
}

func TestProbSomething(t *testing.T) {
	for _, tc := range duetPlayTT {
		t.Run(tc.name, func(t *testing.T) {
			instructions := strings.Split(tc.instructionsInput, "\n")

			duet := duet.NewDuet()
			actual := duet.Play(instructions)
			if actual != tc.expected {
				t.Errorf("...: expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
