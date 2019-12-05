package password

func pow(b, e int) int {
	n := 1
	for i := 0; i < e; i++ {
		n *= b
	}
	return n
}

func dn(n int, nth int) int {
	d := pow(10, nth)
	return (n / d) % 10
}

func IsValidPassword(n int) bool {
	foundPair := false
	for i := 1; i < 6; i++ {
		d1, d2 := dn(n, i), dn(n, i-1)
		if d1 > d2 {
			return false
		}
		if d1 == d2 {
			foundPair = true
			continue
		}
	}

	return foundPair
}

func TotalInRange(start, end int) int {
	count := 0

	for n := start; n <= end; n++ {
		if IsValidPassword(n) {
			count++
		}
	}

	return count
}
