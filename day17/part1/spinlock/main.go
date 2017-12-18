package spinlock

type Spinlock struct {
	step int
}

func NewSpinlock(step int) *Spinlock {
	return &Spinlock{step}
}

func (s *Spinlock) Run(nIterations int) int {
	l := make([]int, 1)
	l[0] = 0
	i := 0

	for it := 1; it <= nIterations; it++ {
		i = (i + s.step) % len(l)

		// insert
		l = append(l, 0)
		copy(l[i+1:], l[i:])
		l[i+1] = it

		i++
	}

	return l[i+1]
}
