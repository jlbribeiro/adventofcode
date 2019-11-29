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

func RegionSizeOfMaxDistance(reader io.Reader, maxDistance int) int {
	coords, rows, cols := CoordsFromReader(reader)

	counter := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			distance := 0
			for _, coord := range coords {
				distance += ManhattanDistance(row, col, coord.row, coord.col)
			}

			if distance < maxDistance {
				counter++
			}
		}
	}

	return counter
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
