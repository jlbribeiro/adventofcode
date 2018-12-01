package freq

func Analyse(deltas []int) int {
	curFreq := 0
	for _, delta := range deltas {
		curFreq += delta
	}

	return curFreq
}
