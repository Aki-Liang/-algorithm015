package Week_08

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return n&(-n) == n
}

func isPowerOfTwoV2(n int) bool {
	if n == 0 {
		return false
	}
	return n&(n-1) == 0
}
