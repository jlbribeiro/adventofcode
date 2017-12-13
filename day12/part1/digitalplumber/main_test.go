package digitalplumber_test

import (
	"fmt"
	"testing"

	"github.com/jlbribeiro/adventofcode/day12/part1/digitalplumber"
)

var nConnectionsOfTests = []struct {
	input     string
	programID int
	expected  int
}{
	{
		`0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`,
		0,
		6,
	},
}

func TestNConnectionsOf(t *testing.T) {
	for _, tt := range nConnectionsOfTests {
		t.Run(fmt.Sprintf("ProgramNetwork.NConnectionsOf(%d)", tt.programID), func(t *testing.T) {
			pn := digitalplumber.NewProgramNetworkFromInput(tt.input)
			actual := pn.NConnectionsOf(tt.programID)
			if actual != tt.expected {
				t.Errorf("ProgramNetwork.NConnectionsOf(%d): expected %d, got %d", tt.programID, tt.expected, actual)
			}
		})
	}
}
