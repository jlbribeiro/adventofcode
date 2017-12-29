package hexed

import "fmt"

type Vector struct {
	x, y, z int
}

func (v *Vector) Add(v2 *Vector) {
	v.x += v2.x
	v.y += v2.y
	v.z += v2.z
}

// http://keekerdc.com/2011/03/hexagon-grids-coordinate-systems-and-distance-calculations/
func (v *Vector) Length() int {
	return maxAbs(v.x, v.y, v.z)
}

func (v *Vector) String() string {
	return fmt.Sprintf("{x: %d, y: %d, z: %d}", v.x, v.y, v.z)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func maxAbs(n ...int) int {
	max := abs(n[0])

	for i := 1; i < len(n); i++ {
		absN := abs(n[i])
		if absN > max {
			max = absN
		}
	}

	return max
}
