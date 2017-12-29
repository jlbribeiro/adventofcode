package generators_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2017/day15/part2/generators"
)

var generatorNextTT = []struct {
	name                             string
	memA, memB                       int64
	factorA, factorB                 int64
	moduloCriteriaA, moduloCriteriaB int64
	expected                         int64
}{
	{"provided example", 65, 8921, 16807, 48271, 4, 8, 309},
}

func TestGeneratorNext(t *testing.T) {
	for _, tc := range generatorNextTT {
		t.Run(tc.name, func(t *testing.T) {
			a := generators.NewGenerator(tc.memA, tc.factorA, tc.moduloCriteriaA)
			b := generators.NewGenerator(tc.memB, tc.factorB, tc.moduloCriteriaB)

			actual := generators.Match(a, b)
			if actual != tc.expected {
				t.Errorf("generator.Next() count: expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
