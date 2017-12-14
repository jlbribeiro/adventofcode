package knothash_test

import (
	"fmt"
	"testing"

	"github.com/jlbribeiro/adventofcode/day14/part1/knothash"
)

var hashTests = []struct {
	n        int
	input    []byte
	expected string
}{
	{
		256,
		[]byte(""),
		"a2582a3a0e66e6e86e3812dcb672a272",
	},
	{
		256,
		[]byte("AoC 2017"),
		"33efeb34ea91902bb2f59c9920caa6cd",
	},
	{
		256,
		[]byte("1,2,3"),
		"3efbe78a8d82f29979031a4aa0b16a9d",
	},
	{
		256,
		[]byte("1,2,4"),
		"63960835bcdc130f0b66d7ff4f6a5a8e",
	},
}

func TestHash(t *testing.T) {
	for _, tt := range hashTests {
		t.Run(fmt.Sprintf("knothash.Hash(%d, %v)", tt.n, tt.input), func(t *testing.T) {
			hash := knothash.NewHash(tt.n)
			actual := hash.Digest(tt.input)
			if actual != tt.expected {
				t.Errorf("knothash.Hash(%d, %v): expected %s, got %s", tt.n, tt.input, tt.expected, actual)
			}
		})
	}
}
