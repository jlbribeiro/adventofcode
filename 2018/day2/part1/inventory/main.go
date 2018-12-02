package inventory

func Checksum(boxIDs []string) int {
	twoCount := 0
	threeCount := 0

	for _, boxID := range boxIDs {
		two, three := BoxChecks(boxID)

		twoCount += two
		threeCount += three
	}

	return twoCount * threeCount
}

func BoxChecks(boxID string) (int, int) {
	occ := make(map[rune]int)
	two := 0
	three := 0

	for _, c := range boxID {
		occ[c]++
	}

	for _, count := range occ {
		if count == 2 {
			two = 1
			continue
		}

		if count == 3 {
			three = 1
			continue
		}
	}

	return two, three
}
