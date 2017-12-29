package fractal

import (
	"bytes"
	"fmt"
	"strings"
)

type MatrixHash string

type Generator struct {
	matrix BoolMatrix
	lookup map[MatrixHash]*Rule
}

type BoolMatrix [][]bool

func NewBoolMatrix(pattern string) BoolMatrix {
	lines := strings.Split(pattern, "/")

	matrix := make(BoolMatrix, len(lines))
	for i, line := range lines {
		matrix[i] = make([]bool, len(line))
		for j, ch := range line {
			matrix[i][j] = ch == '#'
		}
	}

	return matrix
}

func (m BoolMatrix) Hash() MatrixHash {
	return MatrixHash(m.String())
}

func (m BoolMatrix) String() string {
	var buffer bytes.Buffer

	lastLine := len(m) - 1
	for i, line := range m {
		for _, el := range line {
			if el {
				buffer.WriteRune('#')
			} else {
				buffer.WriteRune('.')
			}
		}

		if i < lastLine {
			buffer.WriteRune('/')
		}
	}

	return buffer.String()
}

func (m BoolMatrix) MatrixString() string {
	return strings.Replace(m.String(), "/", "\n", -1) + "\n"
}

func (m BoolMatrix) getSubMatrix(i int, j int, subMatrixSize int) BoolMatrix {
	subMatrix := make(BoolMatrix, subMatrixSize)
	for k := 0; k < subMatrixSize; k++ {
		subMatrix[k] = make([]bool, subMatrixSize)
		for l := 0; l < subMatrixSize; l++ {
			subMatrix[k][l] = m[i*subMatrixSize+k][j*subMatrixSize+l]
		}
	}

	return subMatrix
}

func (m BoolMatrix) writeSubMatrix(i int, j int, subMatrix BoolMatrix) {
	subMatrixSize := len(subMatrix)
	for k := 0; k < subMatrixSize; k++ {
		for l := 0; l < subMatrixSize; l++ {
			m[i*subMatrixSize+k][j*subMatrixSize+l] = subMatrix[k][l]
		}
	}
}

func (m BoolMatrix) emptyCopy() BoolMatrix {
	size := len(m)
	matrix := make(BoolMatrix, size)
	for i := range matrix {
		matrix[i] = make([]bool, size)
	}

	return matrix
}

func (m BoolMatrix) RotateL() BoolMatrix {
	matrix := m.emptyCopy()
	size := len(matrix)

	for i := range m {
		for j := range m[i] {
			matrix[size-j-1][i] = m[i][j]
		}
	}

	return matrix
}

func (m BoolMatrix) RotateR() BoolMatrix {
	matrix := m.emptyCopy()
	size := len(matrix)

	for i := range m {
		for j := range m[i] {
			matrix[j][size-i-1] = m[i][j]
		}
	}

	return matrix
}

func (m BoolMatrix) Rotate180() BoolMatrix {
	return m.RotateL().RotateL()
}

func (m BoolMatrix) Transpose() BoolMatrix {
	return m.RotateR().FlipV()
}

func (m BoolMatrix) FlipH() BoolMatrix {
	matrix := m.emptyCopy()
	size := len(matrix)

	for i := range m {
		for j := range m[i] {
			matrix[size-i-1][j] = m[i][j]
		}
	}

	return matrix
}

func (m BoolMatrix) FlipV() BoolMatrix {
	matrix := m.emptyCopy()
	size := len(matrix)

	for i := range m {
		for j := range m[i] {
			matrix[i][size-j-1] = m[i][j]
		}
	}

	return matrix
}

type Rule struct {
	inputPattern  BoolMatrix
	outputPattern BoolMatrix
}

func (r *Rule) MatrixHashes() []MatrixHash {
	patterns := []MatrixHash(nil)

	patterns = append(patterns, r.inputPattern.Hash())
	patterns = append(patterns, r.inputPattern.RotateL().Hash())
	patterns = append(patterns, r.inputPattern.RotateR().Hash())
	patterns = append(patterns, r.inputPattern.Rotate180().Hash())
	patterns = append(patterns, r.inputPattern.Transpose().Hash())
	patterns = append(patterns, r.inputPattern.Transpose().Rotate180().Hash())
	patterns = append(patterns, r.inputPattern.FlipV().Hash())
	patterns = append(patterns, r.inputPattern.FlipH().Hash())

	return patterns
}

func (r *Rule) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(r.inputPattern.String())
	buffer.WriteString(" => ")
	buffer.WriteString(r.outputPattern.String())

	return buffer.String()
}

func NewRuleFromInput(input string) *Rule {
	inputOutput := strings.Split(input, " => ")

	rule := &Rule{
		inputPattern:  NewBoolMatrix(inputOutput[0]),
		outputPattern: NewBoolMatrix(inputOutput[1]),
	}

	return rule
}

func NewGeneratorFromInput(rulesInput []string) *Generator {
	rules := []*Rule(nil)

	for _, ruleInput := range rulesInput {
		rule := NewRuleFromInput(ruleInput)
		rules = append(rules, rule)
	}

	return NewGenerator(rules)
}

func NewGenerator(rules []*Rule) *Generator {
	lookup := make(map[MatrixHash]*Rule, 0)

	g := &Generator{
		matrix: NewBoolMatrix(".#./..#/###"),
		lookup: lookup,
	}

	for _, rule := range rules {
		g.registerRule(rule)
	}

	return g
}

func (g *Generator) registerRule(rule *Rule) {
	hashes := rule.MatrixHashes()

	for _, hash := range hashes {
		if r, ok := g.lookup[hash]; ok {
			if r != rule {
				panic(fmt.Errorf("about to overwrite rule %v with another rule (%v)", r, rule))
			}

			// Avoid printing the same hash twice
			continue
		}

		g.lookup[hash] = rule
	}
}

func (g *Generator) Run(nIterations int) {
	for i := 0; i < nIterations; i++ {
		g.step()
	}
}

func (g *Generator) step() {
	patternSize := len(g.matrix)

	inputPatternSize := 0
	if patternSize%2 == 0 {
		inputPatternSize = 2
	} else if patternSize%3 == 0 {
		inputPatternSize = 3
	}
	outputPatternSize := inputPatternSize + 1

	nSubMatricesPerDim := patternSize / inputPatternSize

	nextGeneration := make(BoolMatrix, nSubMatricesPerDim*outputPatternSize)
	for i := range nextGeneration {
		nextGeneration[i] = make([]bool, nSubMatricesPerDim*outputPatternSize)
	}

	for i := 0; i < nSubMatricesPerDim; i++ {
		for j := 0; j < nSubMatricesPerDim; j++ {
			subMatrix := g.matrix.getSubMatrix(i, j, inputPatternSize)
			hash := subMatrix.Hash()

			rule, ok := g.lookup[hash]
			if !ok {
				panic(fmt.Errorf("Unexpected input pattern: %s", hash))
			}

			nextGeneration.writeSubMatrix(i, j, rule.outputPattern)
		}
	}

	g.matrix = nextGeneration
}

func (g *Generator) MatrixOnPixels() int {
	count := 0

	for _, row := range g.matrix {
		for _, ch := range row {
			if ch {
				count++
			}
		}
	}

	return count
}
