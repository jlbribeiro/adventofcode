package particles_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2017/day20/part1/particles"
)

var closestToOriginTests = []struct {
	name      string
	particles []string
	expected  int
}{
	{
		"example",
		[]string{
			"p=<3,0,0>, v=<2,0,0>, a=<-1,0,0>",
			"p=<4,0,0>, v=<0,0,0>, a=<-2,0,0>",
		},
		0,
	},
	{
		"tricky",
		[]string{
			"p=<-3,0,0>, v=<2,0,0>, a=<1,0,0>",
			"p=<2,0,0>, v=<2,0,0>, a=<1,0,0>",
		},
		0,
	},
}

func TestClosestToOrigin(t *testing.T) {
	for _, tt := range closestToOriginTests {
		t.Run(tt.name, func(t *testing.T) {
			analyser := particles.NewAnalyserFromInput(tt.particles)
			actual := analyser.ClosestToOrigin()
			if actual != tt.expected {
				t.Errorf("Analyser.ClosestToOrigin(%s): expected %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}
