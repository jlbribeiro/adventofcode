package hexed

import "fmt"

type Direction int

const (
	NORTH Direction = iota + 1
	SOUTH
	NORTHEAST
	SOUTHWEST
	NORTHWEST
	SOUTHEAST
	UNDEFINED Direction = 0
)

func DirectionFromString(dir string) Direction {
	switch dir {
	case "n":
		return NORTH
	case "s":
		return SOUTH
	case "ne":
		return NORTHEAST
	case "sw":
		return SOUTHWEST
	case "nw":
		return NORTHWEST
	case "se":
		return SOUTHEAST
	default:
		panic(fmt.Errorf("Unexpected direction: %s", dir))
	}
}

// http://keekerdc.com/2011/03/hexagon-grids-coordinate-systems-and-distance-calculations/
func (d Direction) Vector() *Vector {
	switch d {
	case NORTH:
		return &Vector{x: 0, y: +1, z: -1}
	case SOUTH:
		return &Vector{x: 0, y: -1, z: +1}
	case NORTHEAST:
		return &Vector{x: +1, y: 0, z: -1}
	case SOUTHWEST:
		return &Vector{x: -1, y: 0, z: +1}
	case NORTHWEST:
		return &Vector{x: -1, y: +1, z: 0}
	case SOUTHEAST:
		return &Vector{x: +1, y: -1, z: 0}
	default:
		panic(fmt.Errorf("Unexpected direction: %d", d))
	}
}
