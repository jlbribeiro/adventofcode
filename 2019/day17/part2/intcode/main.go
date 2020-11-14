package intcode

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

const OP_ADD int64 = 1
const OP_MUL int64 = 2
const OP_MOV int64 = 3
const OP_ECH int64 = 4
const OP_JNZ int64 = 5
const OP_JZ int64 = 6
const OP_SLT int64 = 7
const OP_SEQ int64 = 8
const OP_RBO int64 = 9
const OP_HLT int64 = 99

func ProgramFromInput(intcodeProgram io.Reader) []int64 {
	inputS, err := ioutil.ReadAll(intcodeProgram)
	if err != nil {
		panic(err)
	}

	programS := strings.Split(strings.TrimSpace(string(inputS)), ",")
	program := make([]int64, len(programS))
	for i, nS := range programS {
		n, err := strconv.ParseInt(nS, 10, 64)
		if err != nil {
			panic(err)
		}

		program[i] = n
	}

	return program
}

func Mode(modemap int64, paramPos int) int64 {
	for i := 0; i < paramPos; i++ {
		modemap /= 10
	}
	return modemap % 10
}

type CPU struct {
	Mem           []int64
	IP            int64
	Input         []int64
	InputPtr      int64
	RelBaseOffset int64
}

func NewCPU(program []int64) *CPU {
	mem := make([]int64, len(program), len(program))
	copy(mem, program)

	return &CPU{
		Mem:           mem,
		IP:            0,
		Input:         make([]int64, 0),
		InputPtr:      0,
		RelBaseOffset: 0,
	}
}

func (cpu *CPU) Read(off int64, mode int64) int64 {
	addr := int64(0)

	switch mode {
	case 1: // immediate
		return off

	case 0: // position
		addr = off
	case 2: // relative
		addr = cpu.RelBaseOffset + off

	default:
		panic("invalid mode for read addr")
	}

	if addr < 0 {
		panic(fmt.Errorf("invalid write memory address (%d)", addr))
	}

	// out of bounds = 0
	if addr >= int64(len(cpu.Mem)) {
		return 0
	}

	return cpu.Mem[addr]
}

func (cpu *CPU) Write(off int64, mode int64, value int64) {
	addr := int64(0)

	switch mode {
	case 0: // position
		addr = off
	case 2: // relative
		addr = cpu.RelBaseOffset + off
	default:
		panic("invalid mode for write addr")
	}

	if addr < 0 {
		panic(fmt.Errorf("invalid write memory address (%d)", addr))
	}

	if addr >= int64(len(cpu.Mem)) {
		mem := make([]int64, addr+1)
		copy(mem, cpu.Mem)
		cpu.Mem = mem
	}

	cpu.Mem[addr] = value
}

func (cpu *CPU) Exec(input []int64) ([]int64, bool) {
	cpu.Input = append(cpu.Input, input...)

	var output []int64
	for {
		instruction := cpu.Mem[cpu.IP] % 100
		modes := cpu.Mem[cpu.IP] / 100

		switch instruction {
		case OP_ADD:
			r1, r2, r3 := cpu.Mem[cpu.IP+1], cpu.Mem[cpu.IP+2], cpu.Mem[cpu.IP+3]
			m1, m2, m3 := Mode(modes, 0), Mode(modes, 1), Mode(modes, 2)
			v1, v2 := cpu.Read(r1, m1), cpu.Read(r2, m2)
			cpu.Write(r3, m3, v1+v2)
			cpu.IP += 4

		case OP_MUL:
			r1, r2, r3 := cpu.Mem[cpu.IP+1], cpu.Mem[cpu.IP+2], cpu.Mem[cpu.IP+3]
			m1, m2, m3 := Mode(modes, 0), Mode(modes, 1), Mode(modes, 2)
			v1, v2 := cpu.Read(r1, m1), cpu.Read(r2, m2)
			cpu.Write(r3, m3, v1*v2)
			cpu.IP += 4

		case OP_MOV:
			if cpu.InputPtr >= int64(len(cpu.Input)) {
				return output, true
			}
			r1 := cpu.Mem[cpu.IP+1]
			m1 := Mode(modes, 0)
			cpu.Write(r1, m1, cpu.Input[cpu.InputPtr])
			cpu.InputPtr++
			cpu.IP += 2

		case OP_ECH:
			r1 := cpu.Mem[cpu.IP+1]
			m1 := Mode(modes, 0)
			v1 := cpu.Read(r1, m1)
			output = append(output, v1)
			cpu.IP += 2

		case OP_JNZ:
			r1, r2 := cpu.Mem[cpu.IP+1], cpu.Mem[cpu.IP+2]
			m1, m2 := Mode(modes, 0), Mode(modes, 1)
			v1, v2 := cpu.Read(r1, m1), cpu.Read(r2, m2)
			if v1 != 0 {
				cpu.IP = v2
				continue
			}
			cpu.IP += 3

		case OP_JZ:
			r1, r2 := cpu.Mem[cpu.IP+1], cpu.Mem[cpu.IP+2]
			m1, m2 := Mode(modes, 0), Mode(modes, 1)
			v1, v2 := cpu.Read(r1, m1), cpu.Read(r2, m2)
			if v1 == 0 {
				cpu.IP = v2
				continue
			}
			cpu.IP += 3

		case OP_SLT:
			r1, r2, r3 := cpu.Mem[cpu.IP+1], cpu.Mem[cpu.IP+2], cpu.Mem[cpu.IP+3]
			m1, m2, m3 := Mode(modes, 0), Mode(modes, 1), Mode(modes, 2)
			v1, v2 := cpu.Read(r1, m1), cpu.Read(r2, m2)
			if v1 < v2 {
				cpu.Write(r3, m3, 1)
			} else {
				cpu.Write(r3, m3, 0)
			}
			cpu.IP += 4

		case OP_SEQ:
			r1, r2, r3 := cpu.Mem[cpu.IP+1], cpu.Mem[cpu.IP+2], cpu.Mem[cpu.IP+3]
			m1, m2, m3 := Mode(modes, 0), Mode(modes, 1), Mode(modes, 2)
			v1, v2 := cpu.Read(r1, m1), cpu.Read(r2, m2)
			if v1 == v2 {
				cpu.Write(r3, m3, 1)
			} else {
				cpu.Write(r3, m3, 0)
			}
			cpu.IP += 4

		case OP_RBO:
			r1 := cpu.Mem[cpu.IP+1]
			m1 := Mode(modes, 0)
			v1 := cpu.Read(r1, m1)
			cpu.RelBaseOffset += v1
			cpu.IP += 2

		case OP_HLT:
			return output, false

		default:
			panic(fmt.Errorf("unknown opcode: %d", instruction))
		}
	}
}
