package password_test

import "testing"

import "github.com/jlbribeiro/adventofcode/2019/day4/part1/password"

func TestTotalInRange(t *testing.T) {
	var testCases = []struct {
		name     string
		start    int
		end      int
		expected int
	}{
		{
			name:     "test1",
			start:    111111,
			end:      111112,
			expected: 2,
		},
		{
			name:     "test2",
			start:    123440,
			end:      123460,
			expected: 7,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := password.TotalInRange(tc.start, tc.end)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
