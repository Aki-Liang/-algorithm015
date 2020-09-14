package Week_04

//依然使用贪心
func lemonadeChange(bills []int) bool {
	fiveCount, tenCount := 0, 0
	for _, v := range bills {
		if v == 5 {
			fiveCount++
		} else if v == 10 {
			fiveCount--
			tenCount++
		} else if v == 20 {
			if tenCount > 0 {
				tenCount--
				fiveCount--
			} else {
				fiveCount -= 3
			}
		}

		if fiveCount < 0 {
			return false
		}
	}

	return true
}
