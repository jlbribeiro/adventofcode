package fractal_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/day21/part1/fractal"
)

func TestMatrixRotations(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		left      string
		right     string
		one80     string
		transpose string
		fliph     string
		flipv     string
	}{
		{
			"made_up",
			"##./.../..#",
			"..#/#../#..",
			"..#/..#/#..",
			"#../.../.##",
			"#../#../..#",
			"..#/.../##.",
			".##/.../#..",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := fractal.NewBoolMatrix(tt.input)

			expected := fractal.MatrixHash(tt.left)
			actual := matrix.RotateL().Hash()
			if actual != expected {
				t.Errorf("matrix.RotateL(): expected %s, got %s", expected, actual)
			}

			expected = fractal.MatrixHash(tt.right)
			actual = matrix.RotateR().Hash()
			if actual != expected {
				t.Errorf("matrix.RotateR(): expected %s, got %s", expected, actual)
			}

			expected = fractal.MatrixHash(tt.one80)
			actual = matrix.Rotate180().Hash()
			if actual != expected {
				t.Errorf("matrix.Rotate180(): expected %s, got %s", expected, actual)
			}

			expected = fractal.MatrixHash(tt.transpose)
			actual = matrix.Transpose().Hash()
			if actual != expected {
				t.Errorf("matrix.Transpose(): expected %s, got %s", expected, actual)
			}

			expected = fractal.MatrixHash(tt.fliph)
			actual = matrix.FlipH().Hash()
			if actual != expected {
				t.Errorf("matrix.FlipH(): expected %s, got %s", expected, actual)
			}

			expected = fractal.MatrixHash(tt.flipv)
			actual = matrix.FlipV().Hash()
			if actual != expected {
				t.Errorf("matrix.FlipV(): expected %s, got %s", expected, actual)
			}
		})
	}
}

func TestFractalMatrixPixelsOn(t *testing.T) {
	tests := []struct {
		name        string
		rules       []string
		nIterations int
		expected    int
	}{
		{
			"example",
			[]string{
				"../.# => ##./#../...",
				".#./..#/### => #..#/..../..../#..#",
			},
			2,
			12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fractalGenerator := fractal.NewGeneratorFromInput(tt.rules)
			fractalGenerator.Run(tt.nIterations)
			actual := fractalGenerator.MatrixOnPixels()
			if actual != tt.expected {
				t.Errorf("fractalGenerator.MatrixOnPixels(%s): expected %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}
