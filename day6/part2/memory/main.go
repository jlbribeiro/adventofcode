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
	stateCopy := append([]int(nil), state...)
	p.states = append(p.states, stateCopy)
}

// Find returns the index of a given state, if it has been seen before;
// otherwise returns -1.
func (p *PreviousStates) Find(state []int) int {
	for index, prevState := range p.states {
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
			return index
		}
	}

	return -1
}

// Remove _removes_ a previously stored state in index.
func (p *PreviousStates) Remove(index int) {
	p.states = append(p.states[:index], p.states[index+1:]...)
}

// RebalanceRepeatLoop calculates how many iterations go by before a rebalance loop
// is detected.
func RebalanceRepeatLoop(originalBanks []int) int {
	banks := append([]int(nil), originalBanks...)

	prevStates := NewPreviousStates()

	nIterations := 0

	loopFirstSeenAt := -1
	loopIndex := -1
	for {
		if stateIndex := prevStates.Find(banks); stateIndex != -1 {
			if loopFirstSeenAt == -1 {
				loopFirstSeenAt = nIterations
				loopIndex = stateIndex
			} else if loopIndex == stateIndex {
				return nIterations - loopFirstSeenAt
			}
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
