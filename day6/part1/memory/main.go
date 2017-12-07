package memory

// PreviousStates stores every seen state, so they can be later referenced.
type PreviousStates struct {
	states [][]int
}

// NewPreviousStates returns a PreviousStates instance.
func NewPreviousStates() *PreviousStates {
	states := make([][]int, 0, 1)
	return &PreviousStates{states}
}

// Add stores a given state for future reference.
func (p *PreviousStates) Add(state []int) {
	stateCopy := make([]int, len(state), len(state))
	copy(stateCopy, state)
	p.states = append(p.states, stateCopy)
}

// Has checks whether a given state has been seen before.
func (p *PreviousStates) Has(state []int) bool {
	for _, prevState := range p.states {
		// This shouldn't happen, but...
		if len(state) != len(prevState) {
			continue
		}

		var i int
		for i = 0; i < len(prevState); i++ {
			if prevState[i] != state[i] {
				break
			}
		}

		if i == len(prevState) {
			return true
		}
	}

	return false
}

// RebalanceLoop calculates how many iterations go by before a rebalance loop
// is detected.
func RebalanceLoop(originalBanks []int) int {
	banks := make([]int, len(originalBanks), len(originalBanks))
	copy(banks, originalBanks)

	prevStates := NewPreviousStates()

	nIterations := 0

	for {
		if prevStates.Has(banks) {
			return nIterations
		}

		prevStates.Add(banks)

		maxIndex := 0
		max := banks[maxIndex]

		for i := 1; i < len(banks); i++ {
			if banks[i] > max {
				maxIndex = i
				max = banks[maxIndex]
			}
		}

		banks[maxIndex] = 0
		for max > 0 {
			banks[(maxIndex+1)%len(banks)]++
			maxIndex++
			max--
		}

		nIterations++
	}
}
