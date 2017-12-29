package particles_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2017/day20/part2/particles"
)

var remainingParticlesTests = []struct {
	name      string
	particles []string
	expected  int
}{
	{
		"example",
		[]string{
			"p=<-6,0,0>, v=<3,0,0>, a=<0,0,0>",
			"p=<-4,0,0>, v=<2,0,0>, a=<0,0,0>",
			"p=<-2,0,0>, v=<1,0,0>, a=<0,0,0>",
			"p=<3,0,0>, v=<-1,0,0>, a=<0,0,0>",
		},
		1,
	},
	{
		"example_y",
		[]string{
			"p=<0,-6,0>, v=<0,3,0>, a=<0,0,0>",
			"p=<0,-4,0>, v=<0,2,0>, a=<0,0,0>",
			"p=<0,-2,0>, v=<0,1,0>, a=<0,0,0>",
			"p=<0,3,0>, v=<0,1,0>, a=<0,0,0>",
		},
		1,
	},
}

func TestRemainingParticles(t *testing.T) {
	for _, tt := range remainingParticlesTests {
		t.Run(tt.name, func(t *testing.T) {
			analyser := particles.NewAnalyserFromInput(tt.particles)
			actual := analyser.RemainingParticles()
			if actual != tt.expected {
				t.Errorf("Analyser.RemainingParticles(%s): expected %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}
