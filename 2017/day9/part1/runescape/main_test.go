package runescape_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2017/day9/part1/runescape"
)

var scoreTest = []struct {
	input    string
	expected int
}{
	{`{}`, 1},
	{`{{{}}}`, 6},
	{`{{},{}}`, 5},
	{`{{{},{},{{}}}}`, 16},
	{`{<a>,<a>,<a>,<a>}`, 1},
	{`{{<ab>},{<ab>},{<ab>},{<ab>}}`, 9},
	{`{{<!!>},{<!!>},{<!!>},{<!!>}}`, 9},
	{`{{<a!>},{<a!>},{<a!>},{<ab>}}`, 3},
}

func Test(t *testing.T) {
	for _, tt := range scoreTest {
		t.Run(fmt.Sprintf("runescape.NewStream(%s).Score()", tt.input), func(t *testing.T) {
			reader := strings.NewReader(tt.input)

			stream := runescape.NewStream(reader)
			stream.Process()

			actual := stream.Score()
			if actual != tt.expected {
				t.Errorf(fmt.Sprintf("runescape.NewStream(%s).Score(): expected %d, got %d", tt.input, tt.expected, actual))
			}
		})
	}
}
