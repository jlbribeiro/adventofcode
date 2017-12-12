package knothash

func Hash(n int, lengths []int) int {
	hashString := make([]int, n)
	for i := range hashString {
		hashString[i] = i
	}

	start := 0
	skipSize := 0

	for _, length := range lengths {
		reverse(hashString, n, start, length)
		start += length + skipSize
		skipSize++
	}

	return hashString[0] * hashString[1]
}

func reverse(hashString []int, n int, start int, length int) {
	half := length / 2

	for i := 0; i < half; i++ {
		a := (start + i) % n
		b := (start + length + n - i - 1) % n
		hashString[a], hashString[b] = hashString[b], hashString[a]
	}
}
