package knothash

import "encoding/hex"

var HARDCODED_LENGTHS = []int{17, 31, 73, 47, 23}

type Hash struct {
	n          int
	lengths    []int
	hashString []byte
	start      int
	skipSize   int
}

func NewHash(n int) *Hash {
	hashString := make([]byte, n)
	for i := range hashString {
		hashString[i] = byte(i)
	}

	return &Hash{
		n:          n,
		lengths:    nil,
		hashString: hashString,
		start:      0,
		skipSize:   0,
	}
}

func (h *Hash) calcLengthsFromInput(input []byte) {
	h.lengths = make([]int, len(input)+len(HARDCODED_LENGTHS))

	for i := range input {
		h.lengths[i] = int(input[i])
	}
	for i, length := range HARDCODED_LENGTHS {
		h.lengths[len(input)+i] = length
	}
}

func (h *Hash) reverse(length int) {
	half := length / 2

	for i := 0; i < half; i++ {
		a := (h.start + i) % h.n
		b := (h.start + length + h.n - i - 1) % h.n
		h.hashString[a], h.hashString[b] = h.hashString[b], h.hashString[a]
	}
}

func (h *Hash) hashRound() {
	for _, length := range h.lengths {
		h.reverse(length)
		h.start += length + h.skipSize
		h.skipSize++
	}
}

func (h *Hash) DenseHash() []byte {
	denseHash := make([]byte, 16)
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			if j == 0 {
				denseHash[i] = h.hashString[i*16]
				continue
			}

			denseHash[i] ^= h.hashString[i*16+j]
		}
	}

	return denseHash
}

func (h *Hash) Digest(input []byte) string {
	// Lengths from Input
	h.calcLengthsFromInput(input)

	// Sparse Hash
	for i := 0; i < 64; i++ {
		h.hashRound()
	}

	// Dense Hash
	denseHash := h.DenseHash()

	return hex.EncodeToString(denseHash)
}
