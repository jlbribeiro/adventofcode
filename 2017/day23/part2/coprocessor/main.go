package coprocessor

import (
	"fmt"
	"strconv"
	"strings"
)

type Coprocessor struct {
	registers []int
}

func NewCoprocessor() *Coprocessor {
	registers := make([]int, 8)

	return &Coprocessor{
		registers: registers,
	}
}

func (cp *Coprocessor) GetValue(register string) int {
	val, err := strconv.Atoi(register)
	if err == nil {
		return val
	}

	registerCh := register[0]
	index := registerCh - 'a'
	return cp.registers[index]
}

func (cp *Coprocessor) SetValue(register string, value int) {
	registerCh := register[0]
	index := registerCh - 'a'
	cp.registers[index] = value
}

func (cp *Coprocessor) Run(program []string) {
	for i := 0; i < len(program); i++ {
		instruction := program[i]

		parts := strings.Split(instruction, " ")

		if len(parts) < 3 {
			panic(fmt.Errorf("unexpected instruction: %v", instruction))
		}

		opCode := parts[0]
		x := parts[1]
		y := parts[2]

		switch opCode {
		case "set":
			cp.SetValue(x, cp.GetValue(y))
		case "sub":
			cp.SetValue(x, cp.GetValue(x)-cp.GetValue(y))
		case "mul":
			cp.SetValue(x, cp.GetValue(x)*cp.GetValue(y))
		case "mod":
			cp.SetValue(x, cp.GetValue(x)%cp.GetValue(y))
		case "jnz":
			if cp.GetValue(x) != 0 {
				i += cp.GetValue(y) - 1
			}
		}
	}
}
