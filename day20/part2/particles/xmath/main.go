package xmath

import (
	"math"
)

// Loosely based on https://github.com/asimone0/quadratic/blob/b79de8af/quadratic.go
func SolveQuadraticNonComplexRoots(a float64, b float64, c float64) []float64 {
	negB := -b
	twoA := 2 * a
	bSquared := b * b
	fourAC := 4 * a * c
	discrim := bSquared - fourAC

	// Discard complex roots...
	if discrim < 0 {
		return []float64{}
	}

	sq := math.Sqrt(discrim)

	rootA := (negB - sq) / twoA
	rootB := (negB + sq) / twoA

	// Return a single root if they're the same (discriminant = 0).
	if rootA == rootB {
		return []float64{rootA}
	}

	return []float64{rootA, rootB}
}
