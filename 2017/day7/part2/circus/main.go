package circus

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Program struct {
	Name          string
	Weight        int
	Disc          []string
	DiscWeight    int
	ParentProgram *Program
}

func NewProgram(name string, weight int, disc []string) *Program {
	if disc == nil {
		disc = make([]string, 0, 0)
	}

	return &Program{
		Name:          name,
		Weight:        weight,
		Disc:          disc,
		DiscWeight:    0,
		ParentProgram: nil,
	}
}

func NewProgramFromYell(yell string) *Program {
	regex, err := regexp.Compile(`(\w+) \((\d+)\)( -> (\w+(, \w+)*))?`)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	matches := regex.FindStringSubmatch(yell)

	name := matches[1]

	weight, err := strconv.ParseInt(matches[2], 10, 0)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	disc := make([]string, 0, 0)

	if matches[4] != "" {
		disc = strings.Split(matches[4], ", ")
	}

	return NewProgram(name, int(weight), disc)
}

func (p *Program) Merge(src *Program) {
	p.Weight = src.Weight
	p.Disc = src.Disc
}

type Tower struct {
	structure map[string]*Program
}

func NewTower() *Tower {
	structure := make(map[string]*Program, 0)
	tower := &Tower{
		structure: structure,
	}

	return tower
}

func (t *Tower) RegisterProgram(p *Program) {
	program, ok := t.structure[p.Name]
	if ok {
		program.Merge(p)
		p = program
	} else {
		t.structure[p.Name] = p
	}

	for _, programName := range p.Disc {
		program, ok := t.structure[programName]
		if !ok {
			program = NewProgram(programName, 0, nil)
			t.structure[programName] = program
		}

		program.ParentProgram = p
	}
}

func (t *Tower) FindBottomProgram() *Program {
	for programName := range t.structure {
		program := t.structure[programName]
		if program.ParentProgram == nil {
			return program
		}
	}

	return nil
}

func (t *Tower) findIdealWeightForProgram(p *Program) int {
	if len(p.Disc) == 0 {
		return -1
	}

	childrenWeights := make([]int, len(p.Disc), len(p.Disc))
	for i, programName := range p.Disc {
		childProgram := t.structure[programName]

		idealWeight := t.findIdealWeightForProgram(childProgram)
		if idealWeight != -1 {
			return idealWeight
		}

		childCombinedWeight := childProgram.Weight + childProgram.DiscWeight
		p.DiscWeight += childCombinedWeight

		childrenWeights[i] = childCombinedWeight
	}

	idealWeightDiff, ind := Difference(childrenWeights)
	if ind != -1 {
		return t.structure[p.Disc[ind]].Weight + idealWeightDiff
	}

	return -1
}

func (t *Tower) FindWrongWeightProgramIdealWeight() int {
	bottom := t.FindBottomProgram()
	return t.findIdealWeightForProgram(bottom)
}

func Difference(list []int) (int, int) {
	if len(list) < 3 {
		return 0, -1
	}

	if list[0] != list[1] {
		// Tie breaker: the third element.
		if list[0] != list[2] {
			return list[1] - list[0], 0
		} else {
			return list[0] - list[1], 1
		}
	}

	// list[0] and list[1] are equal;
	// check list[2:].
	for i := 2; i < len(list); i++ {
		if list[i] != list[i-1] {
			return list[i-1] - list[i], i
		}
	}

	return 0, -1
}
