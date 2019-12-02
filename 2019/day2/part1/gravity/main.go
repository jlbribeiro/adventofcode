package gravity

import "fmt"

const OP_ADD int = 1
const OP_MUL int = 2
const OP_HLT int = 99

func Run(mem []int) int {
	ip := 0

	for {
		switch mem[ip] {
		case OP_ADD:
			r1, r2, r3 := mem[ip+1], mem[ip+2], mem[ip+3]
			mem[r3] = mem[r1] + mem[r2]
			ip += 4

		case OP_MUL:
			r1, r2, r3 := mem[ip+1], mem[ip+2], mem[ip+3]
			mem[r3] = mem[r1] * mem[r2]
			ip += 4

		case OP_HLT:
			return mem[0]

		default:
			panic(fmt.Errorf("unknown opcode"))
		}
	}
}
