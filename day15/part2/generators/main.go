package generators

import "math/big"

type Generator struct {
	mem            *big.Int
	factor         *big.Int
	moduloCriteria int
	modulo         *big.Int
}

func NewGenerator(mem int, factor int, moduloCriteria int) *Generator {
	modulo := &big.Int{}
	modulo.SetInt64(2147483647)

	memB := &big.Int{}
	memB.SetInt64(int64(mem))

	factorB := &big.Int{}
	factorB.SetInt64(int64(factor))

	return &Generator{
		mem:            memB,
		factor:         factorB,
		moduloCriteria: moduloCriteria,
		modulo:         modulo,
	}
}

func (g *Generator) Next() int {
	for {
		g.mem.Mul(g.mem, g.factor)
		g.mem.Mod(g.mem, g.modulo)

		c := int(g.mem.Int64())
		if c%g.moduloCriteria == 0 {
			return c
		}
	}
}

func Match(a *Generator, b *Generator) int {
	mask := 1<<16 - 1

	sum := 0
	for i := 0; i < 5000000; i++ {
		ai := a.Next()
		bi := b.Next()

		if ai&mask == bi&mask {
			sum++
		}
	}

	return sum
}
