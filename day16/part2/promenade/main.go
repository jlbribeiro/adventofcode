package promenade

import (
	"fmt"
	"strconv"
	"strings"
)

type Dancers struct {
	programs []rune
}

func NewDancers(nDancers int) *Dancers {
	programs := make([]rune, nDancers)
	for i := range programs {
		programs[i] = rune(int('a') + i)
	}

	return &Dancers{
		programs: programs,
	}
}

func (d *Dancers) String() string {
	return string(d.Alignment())
}

func (d *Dancers) Dance(steps []string) {
	for _, step := range steps {
		switch step[0] {
		case 's':
			d.Spin(step[1:])
		case 'x':
			d.Exchange(step[1:])
		case 'p':
			d.Partner(step[1:])
		}
	}
}

func (d *Dancers) Alignment() []rune {
	return append([]rune(nil), d.programs...)
}

func (d *Dancers) AlignmentEqualTo(alignment []rune) bool {
	if len(d.programs) != len(alignment) {
		return false
	}

	for i := range alignment {
		if alignment[i] != d.programs[i] {
			return false
		}
	}

	return true
}

func (d *Dancers) Spin(step string) {
	n, err := strconv.Atoi(step)
	if err != nil {
		panic(err)
	}

	d.programs = append(d.programs[len(d.programs)-n:], d.programs[:len(d.programs)-n]...)
}

func (d *Dancers) Exchange(step string) {
	swappers := strings.Split(step, "/")
	a, errA := strconv.Atoi(swappers[0])
	b, errB := strconv.Atoi(swappers[1])

	if errA != nil || errB != nil {
		panic(fmt.Errorf("Invalid exchange: %s", step))
	}

	d.programs[a], d.programs[b] = d.programs[b], d.programs[a]
}

func (d *Dancers) Partner(step string) {
	swappers := strings.Split(step, "/")
	indA := 0
	indB := 0

	for ; indA < len(d.programs); indA++ {
		if string(d.programs[indA]) == swappers[0] {
			break
		}
	}

	for ; indB < len(d.programs); indB++ {
		if string(d.programs[indB]) == swappers[1] {
			break
		}
	}

	d.programs[indA], d.programs[indB] = d.programs[indB], d.programs[indA]
}
