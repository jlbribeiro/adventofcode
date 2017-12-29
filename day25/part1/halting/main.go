package halting

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const RESIZE_INCREMENT = 10000

type State struct {
	ID        int
	Write     [2]int
	Move      [2]int
	NextState [2]int
}

func NewStateFromInput(input []string) *State {
	var id int
	write := [2]int{}
	move := [2]int{}
	nextState := [2]int{}

	r := regexp.MustCompile(`In state (\w).+`)

	matches := r.FindStringSubmatch(input[0])
	if len(matches) != 2 {
		panic(fmt.Errorf("%v did not match", input[0]))
	}

	id = int(rune(matches[1][0]) - 'A')

	for i := 0; i < 2; i++ {
		inputString := strings.Join(input[1+i*4:1+(i+1)*4], "\n")
		r = regexp.MustCompile(
			`(?s)Write the value (\d)` +
				`.+Move one slot to the (left|right)` +
				`.+Continue with state (\w)`)
		matches := r.FindStringSubmatch(inputString)

		if len(matches) != 4 {
			panic(fmt.Errorf("%v did not match", inputString))
		}

		var err error
		if write[i], err = strconv.Atoi(matches[1]); err != nil {
			panic(err)
		}

		if matches[2] == "left" {
			move[i] = -1
		} else {
			move[i] = +1
		}

		nextState[i] = int(rune(matches[3][0]) - 'A')
	}

	return &State{
		ID:        id,
		Write:     write,
		Move:      move,
		NextState: nextState,
	}
}

func (s *State) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("State ID: %c\n", rune('A'+s.ID)))
	for i := 0; i < 2; i++ {
		var direction string
		if s.Move[i] > 0 {
			direction = "right"
		} else if s.Move[i] < 0 {
			direction = "left"
		} else {
			panic(fmt.Errorf("unexpected value for movement: %v", s.Move[i]))
		}

		buf.WriteString(fmt.Sprintf(
			"If %v: "+
				"write the value %v, "+
				"move one slot to the %s and "+
				"continue with state %c.\n",
			i, s.Write[i], direction, rune('A'+s.NextState[i])))
	}

	return buf.String()
}

type CPU struct {
	tape    []int
	pointer int
	state   int
	states  []*State
	nSteps  int
}

func parseInitialState(input string) int {
	parts := strings.Split(input, " ")
	lastPart := rune(parts[len(parts)-1][0])
	return int(lastPart - 'A')
}

func parseNSteps(input string) int {
	parts := strings.Split(input, " ")
	stepsPart := parts[len(parts)-2]

	nSteps, err := strconv.Atoi(stepsPart)
	if err != nil {
		panic(err)
	}

	return nSteps
}

func parseStates(input []string) []*State {
	const LINES_PER_STATE = 10

	nStates := len(input) / LINES_PER_STATE

	states := []*State(nil)
	for i := 0; i < nStates; i++ {
		state := NewStateFromInput(input[i*LINES_PER_STATE+1 : (i+1)*LINES_PER_STATE])
		states = append(states, state)
	}

	return states
}

func NewCPUFromBlueprint(blueprint string) *CPU {
	lines := strings.Split(blueprint, "\n")
	initialState := parseInitialState(lines[0])
	nSteps := parseNSteps(lines[1])

	states := parseStates(lines[2:])

	return &CPU{
		tape:    make([]int, 5),
		pointer: 0,
		state:   initialState,
		states:  states,
		nSteps:  nSteps,
	}
}

func (cpu *CPU) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("Current state: %c\n", rune('A'+cpu.state)))
	buf.WriteString(fmt.Sprintf("Tape size: %v\n", len(cpu.tape)))
	buf.WriteString(fmt.Sprintf("Pointer at: %v\n", cpu.pointer))
	buf.WriteString(fmt.Sprintf("Checksum at %v steps.\n", cpu.nSteps))
	buf.WriteString("\n")
	for i, state := range cpu.states {
		buf.WriteString(state.String())

		if i < len(cpu.states)-1 {
			buf.WriteString("\n")
		}
	}

	return buf.String()
}

func (cpu *CPU) Run() {
	for i := 0; i < cpu.nSteps; i++ {
		cpu.step()
	}
}

func (cpu *CPU) step() {
	currentState := cpu.states[cpu.state]

	curVal := cpu.tape[cpu.pointer]

	cpu.tape[cpu.pointer] = currentState.Write[curVal]
	cpu.pointer += currentState.Move[curVal]
	cpu.state = currentState.NextState[curVal]

	cpu.resizeTapeIfNeeded()
}

func (cpu *CPU) resizeTapeIfNeeded() {
	if cpu.pointer >= 0 && cpu.pointer < len(cpu.tape) {
		return
	}

	newTape := make([]int, len(cpu.tape)+RESIZE_INCREMENT)
	if cpu.pointer < 0 {
		copy(newTape[RESIZE_INCREMENT:], cpu.tape)
		cpu.pointer += RESIZE_INCREMENT
	} else {
		copy(newTape, cpu.tape)
	}

	cpu.tape = newTape
}

func (cpu *CPU) Checksum() int {
	checksum := 0

	for _, v := range cpu.tape {
		checksum += v
	}

	return checksum
}
