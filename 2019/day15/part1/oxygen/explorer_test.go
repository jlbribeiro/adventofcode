package oxygen_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day15/part1/oxygen"
)

func TestShortestPathToOxygen(t *testing.T) {
	var testCases = []struct {
		name     string
		fn       func(x, y int) oxygen.Status
		expected int
	}{
		{
			// #############
			// #           #
			// #  #        #
			// # 7O#       #
			// # 6#        #
			// # 54321     #
			// # 4321X     #
			// #           #
			// #           #
			// #           #
			// #           #
			// #           #
			// #############
			name: "simple",
			fn: func(x, y int) oxygen.Status {
				if y < -5 || y > 5 || x < -5 || x > 5 {
					return oxygen.HitWall
				}

				switch {
				case y == -4 && x == -3:
					fallthrough
				case y == -3 && x == -2:
					fallthrough
				case y == -2 && x == -3:
					return oxygen.HitWall
				}

				if y == -3 && x == -3 {
					return oxygen.FoundOxygen
				}

				return oxygen.Moved
			},
			expected: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fnWalker := oxygen.NewFunctionWalker(tc.fn)
			explorer := oxygen.NewExplorer(fnWalker)
			explorer.ExploreMaze()
			actual := explorer.ShortestPathToOxygen()
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
