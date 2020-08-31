package Week_02

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 因为丑数的积还是丑数，所以这个数列可以看成是三个数列的按顺序合并
// aList= {1*2,2*2,3*2,4*2,5*2,6*2,8*2...}
// bList= {1*3,2*3,3*3,4*3,5*3,6*3,8*3...}
// cList= {1*5,2*5,3*5,4*5,5*5,6*5,8*5...}
// 下面使用动态规划的方式解题
func nthUglyNumber(n int) int {
	res := []int{1}
	a, b, c := 0, 0, 0
	for i := 0; i < n; i++ {
		uNum := getMin(getMin(res[a], res[b]), res[c])

		//使用if，而不是ifelse来跳过重复项如 2*3和3*2
		if uNum == res[a] {
			a++
		}
		if uNum == res[b] {
			b++
		}
		if uNum == res[c] {
			c++
		}
		res = append(res, uNum)
	}
	return res[len(res)-1]
}
