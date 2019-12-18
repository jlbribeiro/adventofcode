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

func RealFFT(input []int, nPhases int, period int) []int {
	messageOffsetA := input[:7]
	messageOffset := 0
	for i := range messageOffsetA {
		messageOffset *= 10
		messageOffset += messageOffsetA[i]
	}

	vl := len(input) * period

	n := [2][]int{}
	for i := range n {
		n[i] = make([]int, vl-messageOffset)
	}

	wPtr := 0
	rPtr := messageOffset
	for wPtr < len(n[0]) {
		n[0][wPtr] = input[rPtr%len(input)]
		wPtr++
		rPtr++
	}

	for phase := 0; phase < nPhases; phase++ {
		cur := phase & 0x1
		next := (phase + 1) & 0x1

		sum := 0
		for i := range n[next] {
			n[next][i] = 0
			sum += n[cur][i]
		}

		for i := range n[cur] {
			n[next][i] = sum % 10
			if n[next][i] < 0 {
				n[next][i] = -n[next][i]
			}

			sum -= n[cur][i]
		}
	}

	return n[nPhases&1][:8]
}

func FFTFromInput(in io.Reader, nPhases int) string {
	b, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}

	return ArrayToString(FFT(ArrayFromString(strings.TrimSpace(string(b))), nPhases))
}

func RealFFTFromInput(in io.Reader, nPhases int, period int) string {
	b, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}

	return ArrayToString(RealFFT(ArrayFromString(strings.TrimSpace(string(b))), nPhases, period))
}
