package fft_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day16/part2/fft"
)

func TestFFT(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		nPhases  int
		expected string
	}{
		{
			name:     "example1/1",
			input:    "12345678",
			nPhases:  1,
			expected: "48226158",
		},
		{
			name:     "example1/2",
			input:    "12345678",
			nPhases:  2,
			expected: "34040438",
		},
		{
			name:     "example1/3",
			input:    "12345678",
			nPhases:  3,
			expected: "03415518",
		},
		{
			name:     "example1/4",
			input:    "12345678",
			nPhases:  4,
			expected: "01029498",
		},
		{
			name:     "example2",
			input:    "80871224585914546619083218645595",
			nPhases:  100,
			expected: "24176176",
		},
		{
			name:     "example3",
			input:    "19617804207202209144916044189917",
			nPhases:  100,
			expected: "73745418",
		},
		{
			name:     "example4",
			input:    "69317163492948606335995924319873",
			nPhases:  100,
			expected: "52432133",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := fft.ArrayFromString(tc.input)
			actual := fft.FFT(input, tc.nPhases)
			actual = actual[:8]

			if len(actual) != len(tc.expected) {
				t.Fatalf("expected \"%s\", got \"%s\"", tc.expected, fft.ArrayToString(actual))
			}

			expected := fft.ArrayFromString(tc.expected)
			for i := range expected {
				if actual[i] != expected[i] {
					t.Fatalf("expected \"%s\", got \"%s\"", tc.expected, fft.ArrayToString(actual))
				}
			}
		})
	}
}

func TestRealFFT(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		nPhases  int
		period   int
		expected string
	}{
		{
			name:     "example1",
			input:    "03036732577212944063491565474664",
			nPhases:  100,
			period:   10000,
			expected: "84462026",
		},
		{
			name:     "example2",
			input:    "02935109699940807407585447034323",
			nPhases:  100,
			period:   10000,
			expected: "78725270",
		},
		{
			name:     "example3",
			input:    "03081770884921959731165446850517",
			nPhases:  100,
			period:   10000,
			expected: "53553731",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := fft.ArrayFromString(tc.input)
			actual := fft.RealFFT(input, tc.nPhases, tc.period)
			actual = actual[:8]

			if len(actual) != len(tc.expected) {
				t.Fatalf("expected \"%s\", got \"%s\"", tc.expected, fft.ArrayToString(actual))
			}

			expected := fft.ArrayFromString(tc.expected)
			for i := range expected {
				if actual[i] != expected[i] {
					t.Fatalf("expected \"%s\", got \"%s\"", tc.expected, fft.ArrayToString(actual))
				}
			}
		})
	}
}
