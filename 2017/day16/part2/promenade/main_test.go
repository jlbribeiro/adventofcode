package promenade_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2017/day16/part2/promenade"
)

var dancersDanceTT = []struct {
	name     string
	nDancers int
	steps    []string
	expected string
}{
	{
		"example",
		5,
		[]string{
			"s1",
			"x3/4",
			"pe/b",
		},
		"baedc",
	},
}

func TestDancersDance(t *testing.T) {
	for _, tc := range dancersDanceTT {
		t.Run(tc.name, func(t *testing.T) {
			dancers := promenade.NewDancers(tc.nDancers)
			dancers.Dance(tc.steps)
			actual := string(dancers.Alignment())

			if actual != tc.expected {
				t.Errorf("Dancers.Dance(%s): expected %s, got %s", tc.name, tc.expected, actual)
			}
		})
	}
}
