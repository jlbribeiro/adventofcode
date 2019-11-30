package skylights

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

type Light struct {
	X    int
	Y    int
	VelX int
	VelY int
}

type Lights []*Light

var inputLineRe = regexp.MustCompile(`position=<\s*(\-?\d+),\s*(\-?\d+)>\s*velocity=<\s*(\-?\d+),[ ]*(\-?\d+)>`)

func NewLightsFromInput(input io.Reader) Lights {
	var lights []*Light

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		matches := inputLineRe.FindStringSubmatch(line)
		if len(matches) != 5 {
			panic(fmt.Errorf("failed to parse input line: %s", line))
		}

		posVel := make([]int, 4, 4)
		for i := 0; i < 4; i++ {
			n, err := strconv.ParseInt(matches[i+1], 10, 32)
			if err != nil {
				panic(err)
			}

			posVel[i] = int(n)
		}

		light := &Light{
			X:    posVel[0],
			Y:    posVel[1],
			VelX: posVel[2],
			VelY: posVel[3],
		}

		lights = append(lights, light)
	}

	return lights
}

func (lights Lights) GetLimits() (int, int, int, int) {
	minX, minY := lights[0].X, lights[0].Y
	maxX, maxY := minX, minY

	for _, light := range lights {
		if light.X < minX {
			minX = light.X
		}
		if light.Y < minY {
			minY = light.Y
		}
		if light.X > maxX {
			maxX = light.X
		}
		if light.Y > maxY {
			maxY = light.Y
		}
	}

	return minX, maxX, minY, maxY
}

func (lights Lights) Print(w io.Writer) {
	minX, maxX, minY, maxY := lights.GetLimits()
	width := maxX - minX + 1
	height := maxY - minY + 1

	sky := make([][]bool, height, height)
	for i := range sky {
		sky[i] = make([]bool, width, width)
	}

	for _, light := range lights {
		renderX := light.X - minX
		renderY := light.Y - minY
		sky[renderY][renderX] = true
	}

	for _, row := range sky {
		for _, cell := range row {
			if cell {
				fmt.Fprint(w, "#")
			} else {
				fmt.Fprint(w, ".")
			}
		}
		fmt.Fprintln(w)
	}
}

func (lights Lights) Step(nSteps int) {
	for _, light := range lights {
		light.X += nSteps * light.VelX
		light.Y += nSteps * light.VelY
	}
}

func Run(input io.Reader, output io.Writer) int {
	lights := NewLightsFromInput(input)

	_, _, minY, maxY := lights.GetLimits()
	minHeight := maxY - minY + 1
	for i := 1; ; i++ {
		lights.Step(1)

		_, _, minY, maxY = lights.GetLimits()
		height := maxY - minY + 1
		if height > minHeight {
			lights.Step(-1)
			lights.Print(output)
			return i - 1
		}

		minHeight = height
	}
}
