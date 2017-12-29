package maze

// Walk walks the maze according to the steps on the `maze` list, incrementing
// them when walked "through them". Return the number of necessary steps.
func Walk(maze []int) int {
	nSteps := 0

	for i := 0; i >= 0 && i < len(maze); nSteps++ {
		jump := maze[i]
		maze[i]++
		i += jump
	}

	return nSteps
}
