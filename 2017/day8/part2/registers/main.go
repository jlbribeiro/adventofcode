package registers

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type CPU struct {
	Program          []*Instruction
	Registers        map[string]int
	MaxRegisterValue int
}

func NewCPUFromProgramInput(input string) *CPU {
	textProgram := strings.Split(input, "\n")

	return NewCPUFromTextInstructions(textProgram)
}

func NewCPUFromTextInstructions(textProgram []string) *CPU {
	program := make([]*Instruction, 0)
	registers := make(map[string]int, 0)

	cpu := &CPU{
		Program:          program,
		Registers:        registers,
		MaxRegisterValue: math.MinInt32,
	}

	for _, textInstruction := range textProgram {
		instruction := NewInstructionFromInput(textInstruction)
		cpu.AddInstructionToProgram(instruction)
	}

	return cpu
}

func (cpu *CPU) AddInstructionToProgram(instruction *Instruction) {
	cpu.Program = append(cpu.Program, instruction)
}

func (cpu *CPU) RunProgram() {
	for _, instruction := range cpu.Program {
		cpu.runInstruction(instruction)
		registerVal := cpu.Registers[instruction.RegisterID]
		if registerVal > cpu.MaxRegisterValue {
			cpu.MaxRegisterValue = registerVal
		}
	}
}

func (cpu *CPU) runInstruction(instruction *Instruction) {
	conditionRegisterValue := cpu.Registers[instruction.Condition.RegisterID]
	result := instruction.Condition.Compare(conditionRegisterValue)
	if !result {
		return
	}

	cpu.Registers[instruction.RegisterID] += instruction.Inc
}

func (cpu *CPU) GetLargestRegisterValue() int {
	largest := math.MinInt32

	for _, registerVal := range cpu.Registers {
		if registerVal > largest {
			largest = registerVal
		}
	}

	return largest
}

func (cpu *CPU) GetLargestRegisterValueEver() int {
	return cpu.MaxRegisterValue
}

type Instruction struct {
	RegisterID string
	Inc        int
	Condition  Condition
}

func NewInstructionFromInput(input string) *Instruction {
	regex, err := regexp.Compile(`(\w+) (\w+) (\-?\d+) if (\w+) ([<>=!]+) (\-?\d+)`)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	matches := regex.FindStringSubmatch(input)
	if len(matches) != 7 {
		panic(fmt.Errorf("An instruction with an unexpected format appearted: %s", input))
	}

	registerID := matches[1]
	inc64, err := strconv.ParseInt(matches[3], 10, 64)
	if err != nil {
		panic(err)
	}
	inc := int(inc64)

	switch matches[2] {
	case "inc":
		break
	case "dec":
		inc *= -1
	default:
		panic(fmt.Errorf("Unknown inc/dec instruction: %s", matches[2]))
	}

	conditionRegisterID := matches[4]
	threshold64, err := strconv.ParseInt(matches[6], 10, 64)
	if err != nil {
		panic(err)
	}
	threshold := int(threshold64)

	comparison := NewComparisonFromString(matches[5])
	condition := Condition{
		RegisterID: conditionRegisterID,
		Comparison: comparison,
		Threshold:  threshold,
	}

	instruction := &Instruction{
		RegisterID: registerID,
		Inc:        inc,
		Condition:  condition,
	}

	return instruction
}

type Condition struct {
	RegisterID string
	Comparison Comparison
	Threshold  int
}

func (c *Condition) Compare(registerValue int) bool {
	switch c.Comparison {
	case GT:
		return registerValue > c.Threshold
	case LT:
		return registerValue < c.Threshold
	case GTE:
		return registerValue >= c.Threshold
	case LTE:
		return registerValue <= c.Threshold
	case EQ:
		return registerValue == c.Threshold
	case NEQ:
		return registerValue != c.Threshold
	}

	return false
}

type Comparison int

const (
	GT Comparison = iota + 1
	LT
	GTE
	LTE
	EQ
	NEQ
	UNDEFINED = 0
)

func NewComparisonFromString(input string) Comparison {
	switch input {
	case ">":
		return GT
	case "<":
		return LT
	case ">=":
		return GTE
	case "<=":
		return LTE
	case "==":
		return EQ
	case "!=":
		return NEQ
	default:
		return UNDEFINED
	}
}
