package runescape_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/day9/part2/runescape"
)

var removedGarbageCountTest = []struct {
	input    string
	expected int
}{
	{`<>`, 0},
	{`<random characters>`, 17},
	{`<<<<>`, 3},
	{`<{!>}>`, 2},
	{`<!!>`, 0},
	{`<!!!>>`, 0},
	{`<{o"i!a,<{i<a>`, 10},
}

func Test(t *testing.T) {
	for _, tt := range removedGarbageCountTest {
		t.Run(fmt.Sprintf("runescape.NewStream(%s).RemovedGarbageCount()", tt.input), func(t *testing.T) {
			reader := strings.NewReader(tt.input)

			stream := runescape.NewStream(reader)
			stream.Process()

			actual := stream.RemovedGarbageCount()
			if actual != tt.expected {
				t.Errorf(fmt.Sprintf("runescape.NewStream(%s).RemovedGarbageCount(): expected %d, got %d", tt.input, tt.expected, actual))
			}
		})
	}
}
