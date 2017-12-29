package digitalplumber_test

import (
	"fmt"
	"testing"

	"github.com/jlbribeiro/adventofcode/2017/day12/part2/digitalplumber"
)

var nGroupsTest = []struct {
	input    string
	expected int
}{
	{
		`0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`,
		2,
	},
}

func TestNGroups(t *testing.T) {
	for i, tt := range nGroupsTest {
		t.Run(fmt.Sprintf("ProgramNetwork.NGroups(test=%d)", i), func(t *testing.T) {
			pn := digitalplumber.NewProgramNetworkFromInput(tt.input)
			pn.Flood()
			actual := pn.NGroups()
			if actual != tt.expected {
				t.Errorf("%d: expected %d, got %d", i, tt.expected, actual)
			}
		})
	}
}
