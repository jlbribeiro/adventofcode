package thermal_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day11/part2/thermal"
)

func TestRun(t *testing.T) {
	var testCases = []struct {
		name     string
		program  []int64
		input    []int64
		expected []int64
	}{
		// comparisons
		{
			name:     "day5/example1_input_equals_8_position_mode",
			program:  []int64{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    []int64{8},
			expected: []int64{1},
		},
		{
			name:     "day5/example1_input_not_equals_8_position_mode",
			program:  []int64{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    []int64{7},
			expected: []int64{0},
		},
		{
			name:     "day5/example2_input_less_than_8_position_mode",
			program:  []int64{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    []int64{7},
			expected: []int64{1},
		},
		{
			name:     "day5/example2_input_not_less_than_8_position_mode",
			program:  []int64{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    []int64{9},
			expected: []int64{0},
		},
		{
			name:     "day5/example3_input_equals_8_immediate_mode",
			program:  []int64{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			input:    []int64{8},
			expected: []int64{1},
		},
		{
			name:     "day5/example3_input_not_equals_8_immediate_mode",
			program:  []int64{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			input:    []int64{14},
			expected: []int64{0},
		},
		{
			name:     "day5/example4_input_less_than_8_immediate_mode",
			program:  []int64{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			input:    []int64{4},
			expected: []int64{1},
		},
		{
			name:     "day5/example4_input_not_less_than_8_immediate_mode",
			program:  []int64{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			input:    []int64{10},
			expected: []int64{0},
		},

		// jumps
		{
			name:     "day5/example5_input_is_zero_jump_position_mode",
			program:  []int64{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			input:    []int64{0},
			expected: []int64{0},
		},
		{
			name:     "day5/example5_input_is_non_zero_jump_position_mode",
			program:  []int64{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			input:    []int64{123},
			expected: []int64{1},
		},
		{
			name:     "day5/example6_input_is_zero_jump_immediate_mode",
			program:  []int64{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			input:    []int64{0},
			expected: []int64{0},
		},
		{
			name:     "day5/example6_input_is_non_zero_jump_immediate_mode",
			program:  []int64{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			input:    []int64{123},
			expected: []int64{1},
		},

		// larger jump example
		{
			name:     "day5/example7_output_999_if_input_less_than_8_jumps",
			program:  []int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:    []int64{5},
			expected: []int64{999},
		},
		{
			name:     "day5/example7_output_1000_if_input_equals_8_jumps",
			program:  []int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:    []int64{8},
			expected: []int64{1000},
		},
		{
			name:     "day5/example7_output_1001_if_input_greater_than_8_jumps",
			program:  []int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:    []int64{12},
			expected: []int64{1001},
		},

		{
			name:     "day9/quine",
			program:  []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
			input:    []int64{},
			expected: []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			name:     "day9/16digit",
			program:  []int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
			input:    []int64{},
			expected: []int64{1219070632396864},
		},
		{
			name:     "day9/large_number",
			program:  []int64{104, 1125899906842624, 99},
			input:    []int64{},
			expected: []int64{1125899906842624},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cpu := thermal.NewCPU(tc.program)
			actual, waitingInput := cpu.Exec(tc.input)
			if waitingInput {
				t.Errorf("cpu should not be waiting input by the end of the test")
			}

			if len(actual) != len(tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}

			for i := range tc.expected {
				if actual[i] != tc.expected[i] {
					t.Errorf("expected %v, got %v", tc.expected, actual)
				}
			}
		})
	}
}
