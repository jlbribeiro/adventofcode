package freq

func Analyse(deltas []int) int {
	history := make(map[int]struct{})

	curFreq := 0
	for {
		for _, delta := range deltas {
			history[curFreq] = struct{}{}

			curFreq += delta
			if _, ok := history[curFreq]; ok {
				return curFreq
			}
		}
	}
}
