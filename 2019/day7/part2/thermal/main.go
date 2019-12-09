package thermal

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

const OP_ADD int = 1
const OP_MUL int = 2
const OP_MOV int = 3
const OP_ECH int = 4
const OP_JNZ int = 5
const OP_JZ int = 6
const OP_SLT int = 7
const OP_SEQ int = 8
const OP_HLT int = 99

func ProgramFromInput(intcodeProgram io.Reader) []int {
	inputS, err := ioutil.ReadAll(intcodeProgram)
	if err != nil {
		panic(err)
	}

	programS := strings.Split(strings.TrimSpace(string(inputS)), ",")
	program := make([]int, len(programS))
	for i, nS := range programS {
		n, err := strconv.ParseInt(nS, 10, 32)
		if err != nil {
			panic(err)
		}

		program[i] = int(n)
	}

	return program
}

func Read(mem []int, param int, mode int) int {
	if mode == 0 {
		return mem[param]
	}

	return param
}

func Mode(modemap, paramPos int) int {
	for i := 0; i < paramPos; i++ {
		modemap /= 10
	}
	return modemap % 10
}

type CPU struct {
	Mem      []int
	IP       int
	InputPtr int
}

func NewCPU(program []int) *CPU {
	mem := make([]int, len(program), len(program))
	copy(mem, program)

	return &CPU{
		Mem:      mem,
		IP:       0,
		InputPtr: 0,
	}
}

func (cpu *CPU) Exec(input []int) ([]int, bool) {
	var output []int

	for {
		instruction := cpu.Mem[cpu.IP] % 100
		modes := cpu.Mem[cpu.IP] / 100
		switch instruction {
		case OP_ADD:
			r1, r2, r3 := cpu.Mem[cpu.IP+1], cpu.Mem[cpu.IP+2], cpu.Mem[cpu.IP+3]
			m1, m2 := Mode(modes, 0), Mode(modes, 1)
			v1, v2 := Read(cpu.Mem, r1, m1), Read(cpu.Mem, r2, m2)
			cpu.Mem[r3] = v1 + v2
			cpu.IP += 4

		case OP_MUL:
			r1, r2, r3 := cpu.Mem[cpu.IP+1], cpu.Mem[cpu.IP+2], cpu.Mem[cpu.IP+3]
			m1, m2 := Mode(modes, 0), Mode(modes, 1)
			v1, v2 := Read(cpu.Mem, r1, m1), Read(cpu.Mem, r2, m2)
			cpu.Mem[r3] = v1 * v2
			cpu.IP += 4

		case OP_MOV:
			if cpu.InputPtr >= len(input) {
				return output, true
			}
			r1 := cpu.Mem[cpu.IP+1]
			cpu.Mem[r1] = input[cpu.InputPtr]
			cpu.InputPtr++
			cpu.IP += 2

		case OP_ECH:
			r1 := cpu.Mem[cpu.IP+1]
			m1 := Mode(modes, 0)
			v1 := Read(cpu.Mem, r1, m1)
			output = append(output, v1)
			cpu.IP += 2

		case OP_JNZ:
			r1, r2 := cpu.Mem[cpu.IP+1], cpu.Mem[cpu.IP+2]
			m1, m2 := Mode(modes, 0), Mode(modes, 1)
			v1, v2 := Read(cpu.Mem, r1, m1), Read(cpu.Mem, r2, m2)
			if v1 != 0 {
				cpu.IP = v2
				continue
			}
			cpu.IP += 3

		case OP_JZ:
			r1, r2 := cpu.Mem[cpu.IP+1], cpu.Mem[cpu.IP+2]
			m1, m2 := Mode(modes, 0), Mode(modes, 1)
			v1, v2 := Read(cpu.Mem, r1, m1), Read(cpu.Mem, r2, m2)
			if v1 == 0 {
				cpu.IP = v2
				continue
			}
			cpu.IP += 3

		case OP_SLT:
			r1, r2, r3 := cpu.Mem[cpu.IP+1], cpu.Mem[cpu.IP+2], cpu.Mem[cpu.IP+3]
			m1, m2 := Mode(modes, 0), Mode(modes, 1)
			v1, v2 := Read(cpu.Mem, r1, m1), Read(cpu.Mem, r2, m2)
			if v1 < v2 {
				cpu.Mem[r3] = 1
			} else {
				cpu.Mem[r3] = 0
			}
			cpu.IP += 4

		case OP_SEQ:
			r1, r2, r3 := cpu.Mem[cpu.IP+1], cpu.Mem[cpu.IP+2], cpu.Mem[cpu.IP+3]
			m1, m2 := Mode(modes, 0), Mode(modes, 1)
			v1, v2 := Read(cpu.Mem, r1, m1), Read(cpu.Mem, r2, m2)
			if v1 == v2 {
				cpu.Mem[r3] = 1
			} else {
				cpu.Mem[r3] = 0
			}
			cpu.IP += 4

		case OP_HLT:
			return output, false

		default:
			panic(fmt.Errorf("unknown opcode: %d", instruction))
		}
	}
}
