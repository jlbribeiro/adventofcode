package manhattan

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

const Unclaimed = -2
const Disputed = -1

type Coord struct {
	id  int
	row int
	col int
}

func CoordsFromReader(reader io.Reader) ([]Coord, int, int) {
	scanner := bufio.NewScanner(reader)

	var coords []Coord
	maxRow := 0
	maxCol := 0
	for scanner.Scan() {
		rawCoords := scanner.Text()
		coordParts := strings.Split(rawCoords, ", ")

		row64, _ := strconv.ParseInt(coordParts[1], 10, 64)
		col64, _ := strconv.ParseInt(coordParts[0], 10, 64)

		row := int(row64)
		col := int(col64)

		if row > maxRow {
			maxRow = row
		}

		if col > maxCol {
			maxCol = col
		}

		coords = append(coords, Coord{
			id:  len(coords) + 1,
			row: row,
			col: col,
		})
	}

	rows := maxRow + 1
	cols := maxCol + 1

	return coords, rows, cols
}

func LargestArea(reader io.Reader) int {
	coords, rows, cols := CoordsFromReader(reader)

	areaByCoordID := make(map[int]int)
	infiniteArea := make(map[int]struct{})
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			minDistance := rows + cols
			closestCoordID := Unclaimed

			for coordID, coord := range coords {
				distance := ManhattanDistance(row, col, coord.row, coord.col)
				if distance < minDistance {
					minDistance = distance
					closestCoordID = coordID
				} else if distance == minDistance {
					closestCoordID = Disputed
				}
			}

			if (row == 0 || col == 0 || row == rows-1 || col == cols-1) && closestCoordID != Disputed {
				infiniteArea[closestCoordID] = struct{}{}
			}

			if closestCoordID != Disputed {
				areaByCoordID[closestCoordID]++
			}
		}
	}

	maxArea := 0
	for coordID, area := range areaByCoordID {
		if _, hasInfiniteArea := infiniteArea[coordID]; hasInfiniteArea {
			continue
		}

		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func ManhattanDistance(row1, col1, row2, col2 int) int {
	return IntAbs(row1-row2) + IntAbs(col1-col2)
}

func IntAbs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
