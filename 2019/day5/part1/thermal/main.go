package thermal

import (
	"fmt"
)

const OP_ADD int = 1
const OP_MUL int = 2
const OP_MOV int = 3
const OP_ECH int = 4
const OP_HLT int = 99

func Read(mem []int, param int, mode int) int {
	if mode == 0 {
		return mem[param]
	}

	return param
}

func Run(mem []int, input int) []int {
	var output []int
	ip := 0

	for {
		instruction := mem[ip] % 100
		modes := mem[ip] / 100
		switch instruction {
		case OP_ADD:
			r1, r2, r3 := mem[ip+1], mem[ip+2], mem[ip+3]
			mem[r3] = Read(mem, r1, modes%10) + Read(mem, r2, (modes/10)%10)
			ip += 4

		case OP_MUL:
			r1, r2, r3 := mem[ip+1], mem[ip+2], mem[ip+3]
			mem[r3] = Read(mem, r1, modes%10) * Read(mem, r2, (modes/10)%10)
			ip += 4

		case OP_MOV:
			in := input
			r1 := mem[ip+1]
			mem[r1] = in
			ip += 2

		case OP_ECH:
			r1 := mem[ip+1]
			output = append(output, Read(mem, r1, modes%10))
			ip += 2

		case OP_HLT:
			return output

		default:
			panic(fmt.Errorf("unknown opcode"))
		}
	}
}
