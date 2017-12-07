package memory

// State associates a given banks' state with an iteration occurrence.
type State struct {
	Banks     []int
	Iteration int
}

// NewState returns an instance of a memory state.
func NewState(banks []int, iteration int) *State {
	return &State{
		Banks:     append([]int(nil), banks...),
		Iteration: iteration,
	}
}

// HistoryKeeper is an interface for something capable of
// storing memory states over time.
type HistoryKeeper interface {
	Add(*State)
	Find([]int) (*State, bool)
}

// StateHistory stores every seen state, so they can be later referenced.
type StateHistory struct {
	states []*State
}

// NewStateHistory returns a StateHistory instance.
func NewStateHistory() *StateHistory {
	states := make([]*State, 0, 0)
	return &StateHistory{states}
}

// Add stores a given state for future reference.
func (p *StateHistory) Add(state *State) {
	p.states = append(p.states, state)
}

// Find returns a reference to a state (and ok = true) if the memory banks'
// arrangement has previously occurred; otherwise returns (nil, false).
func (p *StateHistory) Find(banks []int) (*State, bool) {
	for _, prevState := range p.states {
		// This shouldn't happen, but...
		if len(banks) != len(prevState.Banks) {
			continue
		}

		var i int
		for i = 0; i < len(prevState.Banks); i++ {
			if prevState.Banks[i] != banks[i] {
				break
			}
		}

		if i == len(prevState.Banks) {
			return prevState, true
		}
	}

	return nil, false
}

// RebalanceRepeatLoop calculates how many iterations go by before a rebalance loop
// is detected.
func RebalanceRepeatLoop(stateHistory HistoryKeeper, originalBanks []int) int {
	banks := append([]int(nil), originalBanks...)

	nIterations := 0

	for {
		if prevState, ok := stateHistory.Find(banks); ok {
			return nIterations - prevState.Iteration
		}

		stateHistory.Add(NewState(banks, nIterations))

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
