package hexed

import (
	"strings"
)

type HexWalker struct {
	coords *Vector
}

func NewHexWalker() *HexWalker {
	return &HexWalker{
		coords: &Vector{x: 0, y: 0, z: 0},
	}
}

func (h *HexWalker) Walk(dir Direction) {
	h.coords.Add(dir.Vector())
}

func (h *HexWalker) WalkFromInput(input string) {
	stepsStr := strings.Split(input, ",")

	for _, stepStr := range stepsStr {
		h.Walk(DirectionFromString(stepStr))
	}
}

func (h *HexWalker) MinStepsToStart() int {
	return h.coords.Length()
}
