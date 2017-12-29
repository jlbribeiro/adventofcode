package duet_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2017/day18/part1/duet"
)

var probSomethingTT = []struct {
	name              string
	instructionsInput string
	expected          int
}{
	{"example", `set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2`, 4},
}

func TestProbSomething(t *testing.T) {
	for _, tc := range probSomethingTT {
		t.Run(tc.name, func(t *testing.T) {
			instructions := strings.Split(tc.instructionsInput, "\n")

			musicPlayer := duet.NewMusicPlayer()
			actual := musicPlayer.Play(instructions)
			if actual != tc.expected {
				t.Errorf("...: expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
