package tractorbeam

import (
	"io"
	"math"

	"github.com/jlbribeiro/adventofcode/2019/day17/part2/intcode"
)

type BeamAnalyser interface {
	IsBeam(y int, x int) bool
}

type TractorBeam struct {
	program []int64
}

func NewTractorBeam(program []int64) *TractorBeam {
	return &TractorBeam{
		program: program,
	}
}

func (t *TractorBeam) IsBeam(y int, x int) bool {
	cpu := intcode.NewCPU(t.program)
	output, _ := cpu.Exec([]int64{int64(x), int64(y)})
	return int(output[0]) == 1
}

// skipNRows is, without a doubt, a hack; given the provided examples, one would
// assume the beam's traction would be present in all rows past the emitter;
// that is not the case at least for my personal input program.
// However, past a (small) number of rows, that assumption holds true,
// so I chose to keep the algorithm very naïve and just implement it with the
// hack.
func FirstSquareUnderTractorBeam(beam BeamAnalyser, squareSize int, skipNRows int) int {
	firstInsideX := 0
	lastInsideX := 0
	for y := skipNRows; ; y++ {
		inside := false
		lastRowSize := lastInsideX - firstInsideX + 1
		for x := firstInsideX; ; x++ {
			if beam.IsBeam(y, x) {
				// Wasn't inside the beam, this is the first time.
				if !inside {
					firstInsideX = x

					// We know a given row is either lastRowSize or (slightly)
					// wider (probably 1 cell wider), so let's skip most of the
					// coordinates between the beam's boundaries.
					// -1 because this iteration is already inside the beam
					// (so we would need to advance lastRowSize - 1, tops);
					// -1 (again) to correct the loop's own iteration increment.
					x += lastRowSize - 1 - 1
				}

				inside = true
				lastInsideX = x

			} else {
				// Has never been inside the beam during this row scan, continue.
				if !inside {
					continue
				}

				// Was inside the beam,
				// this is the first non-beam cell in this row, which means it
				// is the right-most border of the traction of the beam.

				if lastInsideX-firstInsideX+1 >= squareSize {
					// The traction of the beam is at least squareSize wide.
					minDistance := math.MaxInt64
					xClosest := 0
					for _x := firstInsideX; lastInsideX-_x+1 >= squareSize; _x++ {
						// The square is inside the traction beam as long as we
						// verify that the bottom-left and bottom-right corners
						// are inside, too.
						bottomLeftY, bottomLeftX := y+squareSize-1, _x
						bottomRightY, bottomRightX := y+squareSize-1, _x+squareSize-1

						if beam.IsBeam(bottomLeftY, bottomLeftX) && beam.IsBeam(bottomRightY, bottomRightX) {
							distToEmitter := _x + y
							if distToEmitter < minDistance {
								xClosest = _x
								minDistance = distToEmitter
							}
						}
					}
					if minDistance != math.MaxInt64 {
						// If we found a square that starts in this row (y),
						// this is the closest possible square.
						return xClosest*10000 + y
					}
				}

				// No square found; move to next row,
				// as there's no more beam to analyse.

				inside = false
				break
			}
		}
	}
}

func FirstSquareFromInput(input io.Reader, squareSize int, skipNRows int) int {
	program := intcode.ProgramFromInput(input)
	beam := NewTractorBeam(program)
	return FirstSquareUnderTractorBeam(beam, squareSize, skipNRows)
}
