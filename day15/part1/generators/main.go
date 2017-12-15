package generators

import "math/big"

type Generator struct {
	mem    *big.Int
	factor *big.Int
	modulo *big.Int
}

func NewGenerator(mem int, factor int) *Generator {
	modulo := &big.Int{}
	modulo.SetInt64(2147483647)

	memB := &big.Int{}
	memB.SetInt64(int64(mem))

	factorB := &big.Int{}
	factorB.SetInt64(int64(factor))

	return &Generator{
		mem:    memB,
		factor: factorB,
		modulo: modulo,
	}
}

func (g *Generator) Next() int {
	g.mem.Mul(g.mem, g.factor)
	g.mem.Mod(g.mem, g.modulo)

	return int(g.mem.Int64())
}

func Match(a *Generator, b *Generator) int {
	mask := 1<<16 - 1

	sum := 0
	for i := 0; i < 40000000; i++ {
		ai := a.Next()
		bi := b.Next()

		if ai&mask == bi&mask {
			sum++
		}
	}

	return sum
}
