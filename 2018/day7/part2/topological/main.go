package topological

import (
	"bufio"
	"errors"
	"io"
	"math"
	"regexp"
	"unicode/utf8"
)

const NotWorking = -1

type WorkingCell struct {
	step     int
	timeLeft int
}

var InstructionRegex = regexp.MustCompile(`Step (.*) must be finished before step (.*) can begin.`)

func Order(reader io.Reader, nWorkers int, baseCost int) (int, error) {
	scanner := bufio.NewScanner(reader)

	requisites := make(map[int]map[int]struct{})
	steps := make(map[int]struct{})
	for scanner.Scan() {
		instruction := scanner.Text()
		parts := InstructionRegex.FindStringSubmatch(instruction)

		a, _ := utf8.DecodeLastRuneInString(parts[1])
		b, _ := utf8.DecodeLastRuneInString(parts[2])

		requisite := int(a - 'A')
		dependent := int(b - 'A')

		if _, ok := requisites[dependent]; !ok {
			requisites[dependent] = make(map[int]struct{})
		}

		requisites[dependent][requisite] = struct{}{}
		steps[requisite] = struct{}{}
		steps[dependent] = struct{}{}
	}

	for step := range steps {
		if _, ok := requisites[step]; !ok {
			requisites[step] = make(map[int]struct{})
		}
	}

	nSteps := len(steps)

	timeSpent := 0

	beingWorkedOn := make(map[int]int)

	for len(steps) > 0 {
		for len(beingWorkedOn) < nWorkers {
			nextStepRequisites := nSteps
			availableStep := nSteps

			for step := range steps {
				if _, ok := beingWorkedOn[step]; ok {
					continue
				}

				nRequisites := len(requisites[step])

				if nRequisites > 0 {
					continue
				}

				// Lexicographical order
				if step > availableStep {
					continue
				}

				nextStepRequisites = nRequisites
				availableStep = step
			}

			if nextStepRequisites > 0 {
				break
			}

			beingWorkedOn[availableStep] = baseCost + (availableStep + 1)
		}

		if len(beingWorkedOn) == 0 {
			return 0, errors.New("expected workers to be processing (bad input?)")
		}

		doneStep := -1
		minTimeLeft := math.MaxInt64
		for step, timeLeft := range beingWorkedOn {
			if timeLeft < minTimeLeft {
				doneStep = step
				minTimeLeft = timeLeft
			}
		}

		delete(beingWorkedOn, doneStep)
		delete(steps, doneStep)

		for step := range steps {
			delete(requisites[step], doneStep)
		}

		for step := range beingWorkedOn {
			beingWorkedOn[step] -= minTimeLeft
		}

		timeSpent += minTimeLeft
	}

	return timeSpent, nil
}
