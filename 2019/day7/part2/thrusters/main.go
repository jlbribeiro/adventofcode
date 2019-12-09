package thrusters

import "io"

import "github.com/jlbribeiro/adventofcode/2019/day7/part2/thermal"

import "math"

func AmplifyFromInput(input io.Reader, nAmplifiers int) int {
	program := thermal.ProgramFromInput(input)
	phaseLevels := make([]int, nAmplifiers, nAmplifiers)
	return Amplify(program, 0, phaseLevels)
}

func Amplify(program []int, ampLevel int, phaseLevels []int) int {
	if ampLevel == len(phaseLevels) {
		cpus := make([]*thermal.CPU, len(phaseLevels))
		input := make([][]int, len(phaseLevels))
		for i := 0; i < len(phaseLevels); i++ {
			cpus[i] = thermal.NewCPU(program)
			input[i] = make([]int, 1)
			input[i][0] = phaseLevels[i]
		}

		var outputB []int
		output := 0
		waitingInput := true

		for waitingInput {
			for i := 0; i < len(phaseLevels); i++ {
				input[i] = append(input[i], output)
				outputB, waitingInput = cpus[i].Exec(input[i])
				output = outputB[0]
			}
		}

		return output
	}

	maxOut := math.MinInt32

phases:
	for phase := 5; phase <= 9; phase++ {
		i := 0
		for ; i < ampLevel; i++ {
			if phase == phaseLevels[i] {
				continue phases
			}
		}
		phaseLevels[i] = phase

		ampOut := Amplify(program, ampLevel+1, phaseLevels)
		if ampOut > maxOut {
			maxOut = ampOut
		}
	}

	return maxOut
}
