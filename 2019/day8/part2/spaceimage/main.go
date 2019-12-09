package spaceimage

import (
	"io"
	"io/ioutil"
	"math"
)

func ChecksumFromInput(in io.Reader, width int, height int) int {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}

	ptr := 0
	layerSize := width * height
	nLayers := len(input) / layerSize
	minN0 := math.MaxInt32
	checksum := 0
	for i := 0; i < nLayers; i++ {
		n0, n1, n2 := 0, 0, 0
		for j := 0; j < layerSize; j++ {
			switch input[ptr] {
			case '0':
				n0++
			case '1':
				n1++
			case '2':
				n2++
			}

			ptr++
		}

		if n0 < minN0 {
			minN0 = n0
			checksum = n1 * n2
		}
	}

	return checksum
}

func RenderFromInput(in io.Reader, width int, height int) [][]int {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}

	img := make([][]int, height)
	for i := 0; i < height; i++ {
		img[i] = make([]int, width)
		for j := 0; j < width; j++ {
			img[i][j] = 2
		}
	}

	layerSize := width * height
	nLayers := len(input) / layerSize
	for layer := 0; layer < nLayers; layer++ {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if img[y][x] != 2 {
					continue
				}

				ch := input[layer*layerSize+y*width+x]
				n := int(ch - '0')
				img[y][x] = n
			}
		}
	}

	return img
}
