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

func Run(program []int, userInput []int) []int {
	var output []int
	ip := 0
	inpPtr := 0

	mem := make([]int, len(program), len(program))
	copy(mem, program)

	for {
		instruction := mem[ip] % 100
		modes := mem[ip] / 100
		switch instruction {
		case OP_ADD:
			r1, r2, r3 := mem[ip+1], mem[ip+2], mem[ip+3]
			m1, m2 := Mode(modes, 0), Mode(modes, 1)
			v1, v2 := Read(mem, r1, m1), Read(mem, r2, m2)
			mem[r3] = v1 + v2
			ip += 4

		case OP_MUL:
			r1, r2, r3 := mem[ip+1], mem[ip+2], mem[ip+3]
			m1, m2 := Mode(modes, 0), Mode(modes, 1)
			v1, v2 := Read(mem, r1, m1), Read(mem, r2, m2)
			mem[r3] = v1 * v2
			ip += 4

		case OP_MOV:
			r1 := mem[ip+1]
			mem[r1] = userInput[inpPtr]
			inpPtr++
			ip += 2

		case OP_ECH:
			r1 := mem[ip+1]
			m1 := Mode(modes, 0)
			v1 := Read(mem, r1, m1)
			output = append(output, v1)
			ip += 2

		case OP_JNZ:
			r1, r2 := mem[ip+1], mem[ip+2]
			m1, m2 := Mode(modes, 0), Mode(modes, 1)
			v1, v2 := Read(mem, r1, m1), Read(mem, r2, m2)
			if v1 != 0 {
				ip = v2
				continue
			}
			ip += 3

		case OP_JZ:
			r1, r2 := mem[ip+1], mem[ip+2]
			m1, m2 := Mode(modes, 0), Mode(modes, 1)
			v1, v2 := Read(mem, r1, m1), Read(mem, r2, m2)
			if v1 == 0 {
				ip = v2
				continue
			}
			ip += 3

		case OP_SLT:
			r1, r2, r3 := mem[ip+1], mem[ip+2], mem[ip+3]
			m1, m2 := Mode(modes, 0), Mode(modes, 1)
			v1, v2 := Read(mem, r1, m1), Read(mem, r2, m2)
			if v1 < v2 {
				mem[r3] = 1
			} else {
				mem[r3] = 0
			}
			ip += 4

		case OP_SEQ:
			r1, r2, r3 := mem[ip+1], mem[ip+2], mem[ip+3]
			m1, m2 := Mode(modes, 0), Mode(modes, 1)
			v1, v2 := Read(mem, r1, m1), Read(mem, r2, m2)
			if v1 == v2 {
				mem[r3] = 1
			} else {
				mem[r3] = 0
			}
			ip += 4

		case OP_HLT:
			return output

		default:
			panic(fmt.Errorf("unknown opcode: %d", instruction))
		}
	}
}
