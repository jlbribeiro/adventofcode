package virus

import "fmt"

type Vector struct {
	y int
	x int
}

func (v *Vector) String() string {
	return fmt.Sprintf("{y: %d, x: %d}", v.y, v.x)
}

func (v *Vector) Add(v2 *Vector) *Vector {
	return &Vector{y: v.y + v2.y, x: v.x + v2.x}
}
