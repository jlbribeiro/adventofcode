package virus_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2017/day22/part1/virus"
)

func TestAntiVirusInfectionsCount(t *testing.T) {
	infectionStatus := "..#\n#..\n..."
	tests := []struct {
		name            string
		infectionStatus string
		bursts          int
		expected        int
	}{
		{"example7", infectionStatus, 7, 5},
		{"example70", infectionStatus, 70, 41},
		{"example10000", infectionStatus, 10000, 5587},
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
