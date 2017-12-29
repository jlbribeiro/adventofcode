package generators

type Generator struct {
	mem    int64
	factor int64
	modulo int64
}

func NewGenerator(mem int64, factor int64) *Generator {
	modulo := int64(2147483647)

	return &Generator{
		mem:    mem,
		factor: factor,
		modulo: modulo,
	}
}

func (g *Generator) Next() int64 {
	g.mem *= g.factor
	g.mem %= g.modulo
	return g.mem
}

func Match(a *Generator, b *Generator) int64 {
	mask := int64(1<<16 - 1)

	sum := int64(0)
	for i := 0; i < 40000000; i++ {
		ai := a.Next()
		bi := b.Next()

		if ai&mask == bi&mask {
			sum++
		}
	}

	return sum
}
