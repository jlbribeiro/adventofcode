package marbles

func Insert(slice []int, el int, index int) []int {
	slice = append(slice, 0)
	copy(slice[index+1:], slice[index:])

	slice[index] = el

	return slice
}

func Play(nPlayers int, lastMarble int) int {
	scores := make([]int, nPlayers, nPlayers)
	highScore := 0

	marbles := make([]int, lastMarble, lastMarble)
	marbles = marbles[:1]

	marbles[0] = 0
	curMarble := 0
	for marbleValue := 1; marbleValue <= lastMarble; marbleValue++ {
		if marbleValue%23 == 0 {
			player := (marbleValue - 1) % nPlayers
			scores[player] += marbleValue

			curMarble = curMarble - 7
			if curMarble < 0 {
				curMarble = curMarble + len(marbles)
			}

			scores[player] += marbles[curMarble]

			marbles = append(marbles[:curMarble], marbles[curMarble+1:]...)

			if scores[player] > highScore {
				highScore = scores[player]
			}

			continue
		}

		curMarble = (curMarble + 2) % len(marbles)
		if curMarble == 0 {
			curMarble = len(marbles)
		}

		marbles = Insert(marbles, marbleValue, curMarble)
	}

	return highScore
}
