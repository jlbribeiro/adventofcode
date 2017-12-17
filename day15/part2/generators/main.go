package generators

type Generator struct {
	mem            int64
	factor         int64
	moduloCriteria int64
	modulo         int64
}

func NewGenerator(mem int64, factor int64, moduloCriteria int64) *Generator {
	modulo := int64(2147483647)

	return &Generator{
		mem:            mem,
		factor:         factor,
		moduloCriteria: moduloCriteria,
		modulo:         modulo,
	}
}

func (g *Generator) Next() int64 {
	for {
		g.mem *= g.factor
		g.mem %= g.modulo

		if g.mem%g.moduloCriteria == 0 {
			return g.mem
		}
	}
}

func Match(a *Generator, b *Generator) int64 {
	mask := int64(1<<16 - 1)

	sum := int64(0)
	for i := 0; i < 5000000; i++ {
		ai := a.Next()
		bi := b.Next()

		if ai&mask == bi&mask {
			sum++
		}
	}

	return sum
}
