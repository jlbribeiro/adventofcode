package polymers

import (
	"strings"
	"unicode"
)

func UnitsReact(a, b rune) bool {
	if unicode.IsUpper(a) {
		a, b = b, a
	}

	return unicode.IsLower(a) && unicode.ToUpper(a) == b
}

func React(input string) string {
	var stack []rune

	for _, unit := range input {
		if len(stack) == 0 {
			stack = append(stack, unit)
			continue
		}

		lastUnit := stack[len(stack)-1]
		if UnitsReact(lastUnit, unit) {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, unit)
	}

	return string(stack)
}

func RogueUnitReact(input string) string {
	shortestPolymer := input

	for l := 'A'; l <= 'Z'; l++ {
		attempt := input
		attempt = strings.Replace(attempt, string(l), "", -1)
		attempt = strings.Replace(attempt, string(unicode.ToLower(l)), "", -1)

		resultPolymer := React(attempt)
		if len(resultPolymer) < len(shortestPolymer) {
			shortestPolymer = resultPolymer
		}
	}

	return shortestPolymer
}
