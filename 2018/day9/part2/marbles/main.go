package marbles

import (
	"container/ring"
	"fmt"
)

func PrintRing(r *ring.Ring) {
	n := r.Len()
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", r.Value)
		r = r.Next()
	}
	fmt.Println()
}

func Play(nPlayers int, lastMarble int) int {
	scores := make([]int, nPlayers, nPlayers)
	highScore := 0

	// Keeping a reference to the 0 marble (which will never get removed),
	// so that the printing is "consistent" for each turn.
	originalRing := ring.New(1)

	r := originalRing
	r.Value = 0

	lastMarble = (lastMarble / 23) * 23
	for marbleValue := 1; marbleValue <= lastMarble; marbleValue++ {
		if marbleValue%23 == 0 {
			player := (marbleValue - 1) % nPlayers
			scores[player] += marbleValue

			r = r.Move(-8)
			removedMarble := r.Unlink(1)
			r = r.Next()

			scores[player] += removedMarble.Value.(int)

			if scores[player] > highScore {
				highScore = scores[player]
			}

			continue
		}

		r = r.Next()

		newMarble := ring.New(1)
		newMarble.Value = marbleValue

		r.Link(newMarble)
		r = newMarble
	}

	return highScore
}
