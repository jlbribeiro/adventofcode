package thermal_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day7/part1/thermal"
)

func TestRun(t *testing.T) {
	var testCases = []struct {
		name     string
		program  []int
		input    []int
		expected int
	}{
		// comparisons
		{
			name:     "example1_input_equals_8_position_mode",
			program:  []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    []int{8},
			expected: 1,
		},
		{
			name:     "example1_input_not_equals_8_position_mode",
			program:  []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    []int{7},
			expected: 0,
		},
		{
			name:     "example2_input_less_than_8_position_mode",
			program:  []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    []int{7},
			expected: 1,
		},
		{
			name:     "example2_input_not_less_than_8_position_mode",
			program:  []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    []int{9},
			expected: 0,
		},
		{
			name:     "example3_input_equals_8_immediate_mode",
			program:  []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			input:    []int{8},
			expected: 1,
		},
		{
			name:     "example3_input_not_equals_8_immediate_mode",
			program:  []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			input:    []int{14},
			expected: 0,
		},
		{
			name:     "example4_input_less_than_8_immediate_mode",
			program:  []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			input:    []int{4},
			expected: 1,
		},
		{
			name:     "example4_input_not_less_than_8_immediate_mode",
			program:  []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			input:    []int{10},
			expected: 0,
		},

		// jumps
		{
			name:     "example5_input_is_zero_jump_position_mode",
			program:  []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			input:    []int{0},
			expected: 0,
		},
		{
			name:     "example5_input_is_non_zero_jump_position_mode",
			program:  []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			input:    []int{123},
			expected: 1,
		},
		{
			name:     "example6_input_is_zero_jump_immediate_mode",
			program:  []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			input:    []int{0},
			expected: 0,
		},
		{
			name:     "example6_input_is_non_zero_jump_immediate_mode",
			program:  []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			input:    []int{123},
			expected: 1,
		},

		// larger jump example
		{
			name:     "example7_output_999_if_input_less_than_8_jumps",
			program:  []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:    []int{5},
			expected: 999,
		},
		{
			name:     "example7_output_1000_if_input_equals_8_jumps",
			program:  []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:    []int{8},
			expected: 1000,
		},
		{
			name:     "example7_output_1001_if_input_greater_than_8_jumps",
			program:  []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:    []int{12},
			expected: 1001,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := thermal.Run(tc.program, tc.input)
			if len(actual) != 1 {
				t.Errorf("expected len=1, got %v", actual)
			}

			if actual[0] != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual[0])
			}
		})
	}
}
