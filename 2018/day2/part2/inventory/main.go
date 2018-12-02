package inventory

func FindCorrectBoxes(boxIDs []string) string {
	iMax := len(boxIDs) - 1
	jMax := iMax + 1

	for i := 0; i < iMax; i++ {
		for j := i + 1; j < jMax; j++ {
			if common, ok := IsCorrectPair(boxIDs[i], boxIDs[j]); ok {
				return common
			}
		}
	}

	return ""
}

func IsCorrectPair(a string, b string) (string, bool) {
	if len(a) != len(b) {
		return "", false
	}

	missIndex := -1
	for i := range a {
		if a[i] == b[i] {
			continue
		}

		// More than one different character.
		if missIndex != -1 {
			return "", false
		}

		missIndex = i
	}

	// Equal IDs. They are supposed to have at least one different character.
	if missIndex == -1 {
		return "", false
	}

	return a[:missIndex] + a[missIndex+1:], true
}
