package painter_test

import "testing"

import "github.com/jlbribeiro/adventofcode/2019/day11/part1/painter"

func TestPainterRobot(t *testing.T) {
	var testCases = []struct {
		name               string
		paintInstructions  [][2]int
		expectedScanResult []int
		expectedCount      int
	}{
		{
			name:               "example",
			expectedScanResult: []int{0, 0, 0, 0, 1, 0, 0},
			paintInstructions: [][2]int{
				[2]int{1, 0},
				[2]int{0, 0},
				[2]int{1, 0},
				[2]int{1, 0},
				[2]int{0, 1},
				[2]int{1, 0},
				[2]int{1, 0},
			},
			expectedCount: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			paint := painter.NewPainter()
			for i, expectedScan := range tc.expectedScanResult {
				actualScan := paint.Scan()
				if actualScan != expectedScan {
					t.Errorf("expected %d scan, got %d scan", expectedScan, actualScan)
				}

				instruction := tc.paintInstructions[i]
				paint.Paint(instruction[0], instruction[1])
			}

			actualCount := paint.CountPaintedPanels()
			if actualCount != tc.expectedCount {
				t.Errorf("expected %d painted panels, got %d painted panels", tc.expectedCount, actualCount)
			}
		})
	}
}
