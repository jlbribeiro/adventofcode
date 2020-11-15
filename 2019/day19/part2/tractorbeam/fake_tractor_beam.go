package tractorbeam

import (
	"strings"
)

var _ BeamAnalyser = &FakeTractorBeam{}

type FakeTractorBeam struct {
	tractor [][]rune
}

func NewFakeTractorBeam(tractorS string) *FakeTractorBeam {
	var tractorBeam FakeTractorBeam

	lines := strings.Split(tractorS, "\n")
	for _, line := range lines {
		cells := []rune(line)
		tractorBeam.tractor = append(tractorBeam.tractor, cells)
	}

	return &tractorBeam
}

func (tb *FakeTractorBeam) IsBeam(y int, x int) bool {
	switch {
	case y < 0,
		y >= len(tb.tractor),
		x < 0,
		x >= len(tb.tractor[y]):
		return false
	}

	return tb.tractor[y][x] == '#'
}
