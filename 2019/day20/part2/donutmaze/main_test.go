package donutmaze_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day20/part2/donutmaze"
)

var testCases = []struct {
	name     string
	input    string
	expected int
}{
	{
		name: "example1",
		input: `         A           
         A           
  #######.#########  
  #######.........#  
  #######.#######.#  
  #######.#######.#  
  #######.#######.#  
  #####  B    ###.#  
BC...##  C    ###.#  
  ##.##       ###.#  
  ##...DE  F  ###.#  
  #####    G  ###.#  
  #########.#####.#  
DE..#######...###.#  
  #.#########.###.#  
FG..#########.....#  
  ###########.#####  
             Z       
             Z       `,
		expected: 26,
	},
	// {
	// 	name: "example2",
	// 	input: `                   A
	// A
	// #################.#############
	// #.#...#...................#.#.#
	// #.#.#.###.###.###.#########.#.#
	// #.#.#.......#...#.....#.#.#...#
	// #.#########.###.#####.#.#.###.#
	// #.............#.#.....#.......#
	// ###.###########.###.#####.#.#.#
	// #.....#        A   C    #.#.#.#
	// #######        S   P    #####.#
	// #.#...#                 #......VT
	// #.#.#.#                 #.#####
	// #...#.#               YN....#.#
	// #.###.#                 #####.#
	// DI....#.#                 #.....#
	// #####.#                 #.###.#
	// ZZ......#               QG....#..AS
	// ###.###                 #######
	// JO..#.#.#                 #.....#
	// #.#.#.#                 ###.#.#
	// #...#..DI             BU....#..LF
	// #####.#                 #.#####
	// YN......#               VT..#....QG
	// #.###.#                 #.###.#
	// #.#...#                 #.....#
	// ###.###    J L     J    #.#.###
	// #.....#    O F     P    #.#...#
	// #.###.#####.#.#####.#####.###.#
	// #...#.#.#...#.....#.....#.#...#
	// #.#####.###.###.#.#.#########.#
	// #...#.#.....#...#.#.#.#.....#.#
	// #.###.#####.###.###.#.#.#######
	// #.#.........#...#.............#
	// #########.###.###.#############
	// B   J   C
	// U   P   P               `,
	// 	expected: -1,
	// },
	{
		name: "example3",
		input: `             Z L X W       C                 
             Z P Q B       K                 
  ###########.#.#.#.#######.###############  
  #...#.......#.#.......#.#.......#.#.#...#  
  ###.#.#.#.#.#.#.#.###.#.#.#######.#.#.###  
  #.#...#.#.#...#.#.#...#...#...#.#.......#  
  #.###.#######.###.###.#.###.###.#.#######  
  #...#.......#.#...#...#.............#...#  
  #.#########.#######.#.#######.#######.###  
  #...#.#    F       R I       Z    #.#.#.#  
  #.###.#    D       E C       H    #.#.#.#  
  #.#...#                           #...#.#  
  #.###.#                           #.###.#  
  #.#....OA                       WB..#.#..ZH
  #.###.#                           #.#.#.#  
CJ......#                           #.....#  
  #######                           #######  
  #.#....CK                         #......IC
  #.###.#                           #.###.#  
  #.....#                           #...#.#  
  ###.###                           #.#.#.#  
XF....#.#                         RF..#.#.#  
  #####.#                           #######  
  #......CJ                       NM..#...#  
  ###.#.#                           #.###.#  
RE....#.#                           #......RF
  ###.###        X   X       L      #.#.#.#  
  #.....#        F   Q       P      #.#.#.#  
  ###.###########.###.#######.#########.###  
  #.....#...#.....#.......#...#.....#.#...#  
  #####.#.###.#######.#######.###.###.#.#.#  
  #.......#.......#.#.#.#.#...#...#...#.#.#  
  #####.###.#####.#.#.#.#.###.###.#.###.###  
  #.......#.....#.#...#...............#...#  
  #############.#.#.###.###################  
               A O F   N                     
               A A D   M                     `,
		expected: 396,
	},
}

func TestMinStepsFromInput(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.input)
			actual := donutmaze.MinStepsFromInput(input)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
