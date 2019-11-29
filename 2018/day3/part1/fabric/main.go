package fabric

import (
	"fmt"
	"strconv"
	"strings"
)

const FabricSide = 1000

type Claim struct {
	id int

	left   int
	top    int
	width  int
	height int
}

func (c *Claim) String() string {
	return fmt.Sprintf(
		"id: %d, left: %d, top: %d, width: %d, height: %d",
		c.id, c.left, c.top, c.width, c.height,
	)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func ClaimFromInput(claimInp string) *Claim {
	parts := strings.Split(claimInp, " @ ")

	id, _ := strconv.ParseInt(parts[0][1:], 10, 64)
	parts = strings.Split(parts[1], ": ")

	offsets := strings.Split(parts[0], ",")
	measurements := strings.Split(parts[1], "x")

	left, _ := strconv.ParseInt(offsets[0], 10, 64)
	top, _ := strconv.ParseInt(offsets[1], 10, 64)

	width, _ := strconv.ParseInt(measurements[0], 10, 64)
	height, _ := strconv.ParseInt(measurements[1], 10, 64)

	return &Claim{
		id:     int(id),
		left:   int(left),
		top:    int(top),
		width:  int(width),
		height: int(height),
	}
}

func OverlapsFromInput(claimInps []string) int {
	var claims []*Claim

	for _, claimInp := range claimInps {
		claim := ClaimFromInput(claimInp)
		claims = append(claims, claim)
	}

	overlaps := 0
	grid := make([][]int, FabricSide, FabricSide)
	for i := range grid {
		grid[i] = make([]int, FabricSide, FabricSide)
	}

	for _, claim := range claims {
		topL := min(claim.top+claim.height, FabricSide)
		leftL := min(claim.left+claim.width, FabricSide)

		for i := claim.top; i < topL; i++ {
			for j := claim.left; j < leftL; j++ {
				grid[i][j]++
				if grid[i][j] == 2 {
					overlaps++
				}
			}
		}
	}

	return overlaps
}
