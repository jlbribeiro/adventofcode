package particles

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func signal(x int) int {
	if x == 0 {
		return 0
	}

	return x / abs(x)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

type Vector struct {
	x, y, z int
}

func NewVector(x, y, z int) *Vector {
	return &Vector{x, y, z}
}

func NewVectorFromInput(input []string) *Vector {
	xs, ys, zs := input[0], input[1], input[2]

	x, errx := strconv.Atoi(xs)
	y, erry := strconv.Atoi(ys)
	z, errz := strconv.Atoi(zs)

	if errx != nil || erry != nil || errz != nil {
		panic(fmt.Errorf("Unexpected input: %s", input))
	}

	return NewVector(x, y, z)
}

func (v *Vector) String() string {
	return fmt.Sprintf("<%d,%d,%d>", v.x, v.y, v.z)
}

func (v *Vector) Array() [3]int {
	return [3]int{v.x, v.y, v.z}
}

func (v *Vector) Add(v2 *Vector) *Vector {
	return NewVector(v.x+v2.x, v.y+v2.y, v.z+v2.z)
}

func (v *Vector) ScalarMult(factor int) *Vector {
	return NewVector(v.x*factor, v.y*factor, v.z*factor)
}

func (v *Vector) ManhattanDistance() int {
	return abs(v.x) + abs(v.y) + abs(v.z)
}

type Particle struct {
	acc *Vector
	vel *Vector
	pos *Vector
}

func NewParticleFromInput(input string) *Particle {
	regex, err := regexp.Compile(`p=<(\-?\d+),(\-?\d+),(\-?\d+)>, v=<(\-?\d+),(\-?\d+),(\-?\d+)>, a=<(\-?\d+),(\-?\d+),(\-?\d+)>`)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	matches := regex.FindStringSubmatch(input)
	if len(matches) < 10 {
		panic(fmt.Errorf("Unexpected input: %s", input))
	}

	pos := NewVectorFromInput(matches[1:4])
	vel := NewVectorFromInput(matches[4:7])
	acc := NewVectorFromInput(matches[7:10])

	return &Particle{
		pos: pos,
		vel: vel,
		acc: acc,
	}
}

func (p *Particle) String() string {
	return fmt.Sprintf("p=%s, v=%s, a=%s", p.pos, p.vel, p.acc)
}

func (p *Particle) Step(steps int) {
	for i := 0; i < steps; i++ {
		p.vel = p.vel.Add(p.acc)
		p.pos = p.pos.Add(p.vel)
	}
}

// Less returns whether `p` particle will be closer to (0,0,0) when approaching
// infinity than `p2`.
// Since the update function for each step is
//   "Increase the velocity by the acceleration."
//   "Increase the position by the velocity."
// we can deduce the position of the particle is given by the function
// p[n] = p0 + v0 * n + a * (n * (n + 1)) / 2
func (p *Particle) Less(p2 *Particle) bool {
	v1 := p.acc
	v2 := p2.acc

	diff := v1.ManhattanDistance() - v2.ManhattanDistance()
	if diff != 0 {
		return diff < 0
	}

	v1 = v1.Add(p.vel)
	v2 = v2.Add(p2.vel)

	diff = v1.ManhattanDistance() - v2.ManhattanDistance()
	if diff != 0 {
		return diff < 0
	}

	v1 = v1.Add(p.pos)
	v2 = v2.Add(p2.pos)

	diff = v1.ManhattanDistance() - v2.ManhattanDistance()
	if diff != 0 {
		return diff < 0
	}

	return true
}

type Analyser struct {
	particles []*Particle
}

func NewAnalyserFromInput(particlesInput []string) *Analyser {
	particles := []*Particle(nil)

	for _, particleStr := range particlesInput {
		particles = append(particles, NewParticleFromInput(particleStr))
	}

	return &Analyser{
		particles: particles,
	}
}

func (a *Analyser) ClosestToOrigin() int {
	minI := 0

	for i := 1; i < len(a.particles); i++ {
		if a.particles[i].Less(a.particles[minI]) {
			minI = i
		}
	}

	return minI
}
