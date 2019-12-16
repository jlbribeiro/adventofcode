package fft

import (
	"io"
	"io/ioutil"
	"strings"
)

func ArrayFromString(in string) []int {
	out := make([]int, len(in))
	for i, c := range in {
		out[i] = int(c - '0')
	}
	return out
}

func ArrayToString(arr []int) string {
	out := strings.Builder{}
	for _, n := range arr {
		out.WriteRune(rune('0' + n))
	}

	return out.String()
}

func FFT(input []int, nPhases int) []int {
	n := [2][]int{}
	for i := range n {
		n[i] = make([]int, len(input))
	}

	copy(n[0], input)

	for phase := 0; phase < nPhases; phase++ {
		cur := phase & 0x1
		next := (phase + 1) & 0x1

		for i := 0; i < len(n[next]); i++ {
			n[next][i] = 0
		}

		for d := range n[cur] {
			offset := d + 1

			start := d
			for ; start < len(n[cur]); start += 4 * offset {
				startPositive := start
				endPositive := startPositive + offset
				for i := startPositive; i < endPositive && i < len(n[cur]); i++ {
					n[next][d] += n[cur][i]
				}

				startNegative := start + 2*offset
				endNegative := startNegative + offset
				for i := startNegative; i < endNegative && i < len(n[cur]); i++ {
					n[next][d] -= n[cur][i]
				}
			}

			n[next][d] %= 10
			if n[next][d] < 0 {
				n[next][d] = -n[next][d]
			}
		}
	}

	return n[nPhases&1]
}

func FFTFromInput(in io.Reader, nPhases int) string {
	b, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}

	return ArrayToString(FFT(ArrayFromString(strings.TrimSpace(string(b))), nPhases))
}
