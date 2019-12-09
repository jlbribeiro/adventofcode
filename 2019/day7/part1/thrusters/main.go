package thrusters

import "io"

import "github.com/jlbribeiro/adventofcode/2019/day7/part1/thermal"

import "math"

func AmplifyFromInput(input io.Reader, nAmplifiers int) int {
	program := thermal.ProgramFromInput(input)
	phaseLevels := make([]int, nAmplifiers, nAmplifiers)
	return Amplify(program, 0, phaseLevels)
}

func Amplify(program []int, ampLevel int, phaseLevels []int) int {
	maxOut := math.MinInt32
	if ampLevel == len(phaseLevels) {
		output := 0
		for i := 0; i < len(phaseLevels); i++ {
			outputB := thermal.Run(program, []int{phaseLevels[i], output})
			output = outputB[0]
		}

		return output
	}

phases:
	for phase := 0; phase < len(phaseLevels); phase++ {
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
