package Week_08

import "sync"

func isAnagram(s string, t string) bool {
	//题目中说明只包含小写字母
	bitmapS := [26]int{}
	bitmapT := [26]int{}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(s); i++ {
			index := s[i] - 'a'
			bitmapS[index] += 1
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(t); i++ {
			index := t[i] - 'a'
			bitmapT[index] += 1
		}
	}()
	wg.Wait()

	return bitmapS == bitmapT
}
