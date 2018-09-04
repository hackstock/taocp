package algos

func gcd(m, n int) int {
	r := m % n
	if r == 0 {
		return n
	}
	m = n
	n = r
	return gcd(m, n)
}
