package spinlock

type Spinlock struct {
	step int
}

func NewSpinlock(step int) *Spinlock {
	return &Spinlock{step}
}

func (s *Spinlock) Run(nIterations int) int {
	i := 0
	nextToZero := -1

	for it := 1; it <= nIterations; it++ {
		i = (i + s.step) % it

		if i == 0 {
			nextToZero = it
		}

		i++
	}

	return nextToZero
}
