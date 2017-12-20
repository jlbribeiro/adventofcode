package tubes

type Vector struct {
	y, x int
}

func (v *Vector) Add(v2 *Vector) *Vector {
	return &Vector{
		v.y + v2.y,
		v.x + v2.x,
	}
}
