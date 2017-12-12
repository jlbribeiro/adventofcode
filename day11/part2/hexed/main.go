package hexed

import (
	"strings"
)

type HexWalker struct {
	coords          *Vector
	furthestToStart int
}

func NewHexWalker() *HexWalker {
	return &HexWalker{
		coords:          &Vector{x: 0, y: 0, z: 0},
		furthestToStart: 0,
	}
}

func (h *HexWalker) Walk(dir Direction) {
	h.coords.Add(dir.Vector())
	stepsToStart := h.MinStepsToStart()
	if stepsToStart > h.furthestToStart {
		h.furthestToStart = stepsToStart
	}
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

func (h *HexWalker) FurthestToStart() int {
	return h.furthestToStart
}
