package Week_01

func plusOne(digits []int) []int {
	length := len(digits)

	for i := length - 1; i >= 0; i-- {
		if 9 == digits[i] {
			digits[i] = 0
		} else {
			digits[i] = digits[i] + 1
			return digits
		}
	}
	return append([]int{1}, digits...)
}
