package virus_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2017/day22/part2/virus"
)

func TestAntiVirusInfectionsCount(t *testing.T) {
	infectionStatus := "..#\n#..\n..."
	tests := []struct {
		name            string
		infectionStatus string
		bursts          int
		expected        int
	}{
		{"example100", infectionStatus, 100, 26},
		{"example10000000", infectionStatus, 10000000, 2511944},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			av := virus.NewAVFromInput(tt.infectionStatus)
			av.Run(tt.bursts)
			actual := av.NInfections()
			if actual != tt.expected {
				t.Errorf("av.NInfections(): expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
