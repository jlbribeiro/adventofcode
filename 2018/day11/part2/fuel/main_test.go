package fuel_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2018/day11/part2/fuel"
)

func TestCellPowerLevel(t *testing.T) {
	var testCases = []struct {
		name         string
		serialNumber int
		x            int
		y            int
		expected     int
	}{
		{
			name:         "example1",
			serialNumber: 8,
			x:            3,
			y:            5,
			expected:     4,
		},
		{
			name:         "example2",
			serialNumber: 57,
			x:            122,
			y:            79,
			expected:     -5,
		},
		{
			name:         "example3",
			serialNumber: 39,
			x:            217,
			y:            196,
			expected:     0,
		},
		{
			name:         "example4",
			serialNumber: 71,
			x:            101,
			y:            153,
			expected:     4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := fuel.CellPowerLevel(tc.serialNumber, tc.x, tc.y)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}

}

func TestMaxTotalPower(t *testing.T) {
	var testCases = []struct {
		name               string
		serialNumber       int
		expectedX          int
		expectedY          int
		expectedSquareSize int
		expectedTotalPower int
	}{
		{
			name:               "example1",
			serialNumber:       18,
			expectedX:          90,
			expectedY:          269,
			expectedSquareSize: 16,
			expectedTotalPower: 113,
		},
		{
			name:               "example2",
			serialNumber:       42,
			expectedX:          232,
			expectedY:          251,
			expectedSquareSize: 12,
			expectedTotalPower: 119,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualTotalPower, actualX, actualY, actualSquareSize := fuel.MaxTotalPower(tc.serialNumber)
			if actualTotalPower != tc.expectedTotalPower {
				t.Errorf("expected %d max total power, got %d max total power", tc.expectedTotalPower, actualTotalPower)
			}
			if actualSquareSize != tc.expectedSquareSize {
				t.Errorf("expected %d square size, got %d square size", tc.expectedSquareSize, actualSquareSize)
			}
			if actualX != tc.expectedX || actualY != tc.expectedY {
				t.Errorf("expected (x,y)=(%d,%d), got (x,y)=(%d,%d)", tc.expectedX, tc.expectedY, actualX, actualY)
			}
		})
	}
}
