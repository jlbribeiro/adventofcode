package particles

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"

	"github.com/jlbribeiro/adventofcode/day20/part2/particles/xmath"
)

const N_DIMS = 3

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

func (v *Vector) Array() [N_DIMS]int {
	return [N_DIMS]int{v.x, v.y, v.z}
}

func (v *Vector) Equals(v2 *Vector) bool {
	return v.x == v2.x && v.y == v2.y && v.z == v2.z
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
	destroyed bool
	acc       *Vector
	vel       *Vector
	pos       *Vector
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
		destroyed: false,
		pos:       pos,
		vel:       vel,
		acc:       acc,
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

func (p *Particle) AtInstant(n int) *Vector {
	// p0 + (v0 + a / 2) * n + (a / 2) * n^2
	nSquared := n * n
	// FIXME

	pos0 := p.pos.Array()
	vel0 := p.vel.Array()
	acc := p.acc.Array()

	pos := [N_DIMS]int{}
	for dim := 0; dim < N_DIMS; dim++ {
		halfAcc := float64(acc[dim]) / 2
		pos[dim] = int(float64(pos0[dim]) + (float64(vel0[dim])+halfAcc)*float64(n) + halfAcc*float64(nSquared))
	}

	return NewVector(pos[0], pos[1], pos[2])
}

func countIfSoundSolution(candidates []float64, counter map[int]int) {
	for _, candidate := range candidates {
		if candidate == math.Trunc(candidate) {
			solution := int(candidate)
			if solution < 0 {
				continue
			}

			counter[solution]++
		}
	}
}

func solutionsFromCounters(counter map[int]int, wildcard int) []int {
	solutions := []int(nil)

	for solution, count := range counter {
		if count+wildcard >= N_DIMS {
			if count+wildcard > N_DIMS {
				panic(fmt.Errorf("count + wildcard > N_DIMS: wtf"))
			}

			solutions = append(solutions, solution)
		}
	}

	return solutions
}

// p[n] = p0 + v0 * n + a * (n * (n + 1)) / 2
//      =               (a / 2) * n^2 + (a / 2) * n
//      = p0 + (v0 + a / 2) * n + (a / 2) * n^2
//
// p0_1 + (v0_1 + a_1 / 2) * n + (a_1 / 2) * n^2 = p0_2 + (v0_2 + a_2 / 2) * n + (a_2 / 2) * n^2
// p0_1 - p0_2 + ((v0_1 + a_1 / 2) - (v0_2 + a_2 / 2)) * n + ((a_1 - a_2) / 2) * n^2 = 0
func (p *Particle) CollidesWith(other *Particle) []int {
	counter := make(map[int]int, 0)
	wildcard := 0

	p1Acc, p2Acc := p.acc.Array(), other.acc.Array()
	p1Vel, p2Vel := p.vel.Array(), other.vel.Array()
	p1Pos, p2Pos := p.pos.Array(), other.pos.Array()

	for dim := 0; dim < N_DIMS; dim++ {
		a1, a2 := float64(p1Acc[dim]), float64(p2Acc[dim])
		v1, v2 := float64(p1Vel[dim]), float64(p2Vel[dim])
		p1, p2 := float64(p1Pos[dim]), float64(p2Pos[dim])

		// f(x) = ax^2 + bx + c
		//   a = (a_1 / 2) - (a_2 / 2) = (a_1 - a_2) / 2
		//   b = (v0_1 + a_1 / 2) - (v0_2 + a_2 / 2)
		//   c = p0_1 - p0_2
		a := (a1 - a2) / 2
		b := (v1 + a1/2) - (v2 + a2/2)
		c := p1 - p2

		if a == 0 {
			// 0x^2: not a quadratic function.

			if b == 0 {
				// 0x^2 + 0x + c = 0
				if c != 0 {
					// c = 0; c != 0 is an impossible equation: no solutions.
					return []int{}
				}

				// c = 0; c == 0 means infinite solutions.
				wildcard++
				continue
			}

			countIfSoundSolution([]float64{-c / b}, counter)

		} else if a != 0 && b != 0 && c != 0 {
			// Quadratic equation.
			possibleSolutions := xmath.SolveQuadraticNonComplexRoots(a, b, c)
			countIfSoundSolution(possibleSolutions, counter)
		}

	}

	return solutionsFromCounters(counter, wildcard)
}

type Collision struct {
	coords    Vector
	particles []int
}

func NewCollision(coords *Vector, particles ...int) *Collision {
	return &Collision{
		coords:    *coords,
		particles: particles,
	}
}

func (c *Collision) Add(particleIndex int) {
	for _, otherIndex := range c.particles {
		if particleIndex == otherIndex {
			return
		}
	}

	c.particles = append(c.particles, particleIndex)
}

type CollisionSet struct {
	collisions map[int][]*Collision
}

func NewCollisionSet() *CollisionSet {
	return &CollisionSet{
		collisions: make(map[int][]*Collision, 0),
	}
}

func (cs *CollisionSet) AddCollision(instant int, coords *Vector, particles ...int) {
	_, ok := cs.collisions[instant]
	if !ok {
		cs.collisions[instant] = []*Collision{NewCollision(coords, particles...)}
		return
	}

	for _, collision := range cs.collisions[instant] {
		if collision.coords.Equals(coords) {
			for _, particle := range particles {
				collision.Add(particle)
			}
			return
		}
	}

	cs.collisions[instant] = append(cs.collisions[instant], NewCollision(coords, particles...))
}

func (cs *CollisionSet) Instants() []int {
	instants := []int(nil)

	for instant := range cs.collisions {
		instants = append(instants, instant)
	}

	sort.Ints(instants)

	return instants
}

func (cs *CollisionSet) CollisionsAt(instant int) []*Collision {
	return cs.collisions[instant]
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

func (a *Analyser) buildCollisionSet() *CollisionSet {
	cs := NewCollisionSet()

	for i := 0; i < len(a.particles)-1; i++ {
		for j := i + 1; j < len(a.particles); j++ {
			p1 := a.particles[i]
			p2 := a.particles[j]

			collisions := p1.CollidesWith(p2)
			if len(collisions) > 0 {
				fmt.Printf("Particles %v and %v collide: %v\n", i, j, collisions)
			}

			for _, instant := range collisions {
				cs.AddCollision(instant, p1.AtInstant(instant), i, j)
			}
		}
	}

	return cs
}

func (a *Analyser) CollisionStillPossible(c *Collision) bool {
	count := 0
	for _, particleIndex := range c.particles {
		if !a.particles[particleIndex].destroyed {
			count++
		}
	}

	return count >= 2
}

func (a *Analyser) RemainingParticles() int {
	remaining := len(a.particles)

	cs := a.buildCollisionSet()
	instants := cs.Instants()

	for _, instant := range instants {
		collisions := cs.CollisionsAt(instant)
		for _, collision := range collisions {
			if !a.CollisionStillPossible(collision) {
				continue
			}

			for _, particleIndex := range collision.particles {
				particle := a.particles[particleIndex]
				if !particle.destroyed {
					particle.destroyed = true
					remaining--
				}
			}
		}
	}

	return remaining
}
