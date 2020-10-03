package Week_06

import "math"

func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	length := math.MaxInt32
	ansLeft, ansRight := -1, -1
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}

		for check() && l <= r {
			tmp := r - l + 1
			if tmp < length {
				length = tmp
				ansLeft = l
				ansRight = l + tmp
			}

			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansLeft == -1 {
		return ""
	}
	return s[ansLeft:ansRight]
}
