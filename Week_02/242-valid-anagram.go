package Week_02

import "sync"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	//bitmap 方式，假定只有26个小写字母
	sBitmap := [26]int{}
	tBitMap := [26]int{}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(s); i++ {
			j := s[i] - 'a'
			sBitmap[j]++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(t); i++ {
			j := t[i] - j
			tBitMap[j]++
		}
	}()

	wg.Wait()

	//go 的数组是可以比较的
	return sBitmap == tBitMap
}
