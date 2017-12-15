package generators_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/day15/part1/generators"
)

var generatorNextTT = []struct {
	name             string
	memA, memB       int
	factorA, factorB int
	expected         int
}{
	{"provided example", 65, 8921, 16807, 48271, 588},
}

func TestGeneratorNext(t *testing.T) {
	for _, tc := range generatorNextTT {
		t.Run(tc.name, func(t *testing.T) {
			a := generators.NewGenerator(tc.memA, tc.factorA)
			b := generators.NewGenerator(tc.memB, tc.factorB)

			actual := generators.Match(a, b)
			if actual != tc.expected {
				t.Errorf("generator.Next() count: expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
