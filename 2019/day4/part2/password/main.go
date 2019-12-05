package password

func pow(b, e int) int {
	n := 1
	for i := 0; i < e; i++ {
		n *= b
	}
	return n
}

func digitN(n int, nth int) int {
	d := pow(10, nth)
	return (n / d) % 10
}

func IsValidPassword(n int) bool {
	foundPair := false
	curMatchingStreak := 0
	for i := 1; i < 6; i++ {
		d1, d2 := digitN(n, i), digitN(n, i-1)
		if d1 > d2 {
			return false
		}

		if d1 == d2 {
			curMatchingStreak++
			continue

		}

		if curMatchingStreak == 1 {
			foundPair = true
		}
		curMatchingStreak = 0
	}

	// last 2 digits may have been a pair
	if curMatchingStreak == 1 {
		foundPair = true
	}

	return foundPair
}

func TotalInRange(start, end int) int {
	count := 0

	for n := start; n <= end; n++ {
		if IsValidPassword(n) {
			count++
		}
	}

	return count
}
