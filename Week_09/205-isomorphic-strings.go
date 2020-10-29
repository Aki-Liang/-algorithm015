package Week_09

// 给定两个字符串 s 和 t，判断它们是否是同构的。

// 如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。

// 所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。

// 示例 1:

// 输入: s = "egg", t = "add"
// 输出: true
// 示例 2:

// 输入: s = "foo", t = "bar"
// 输出: false
// 示例 3:

// 输入: s = "paper", t = "title"
// 输出: true

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/isomorphic-strings
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func isIsomorphicV2(s string, t string) bool {
	mapS := make([]int, 128)
	mapT := make([]int, 128)
	for i := 0; i < len(s); i++ {
		if mapS[s[i]] != mapT[t[i]] {
			return false
		} else {
			if mapS[s[i]] == 0 {
				mapS[s[i]] = i + 1
				mapT[t[i]] = i + 1
			}
		}
	}

	return true
}

func isIsomorphic(s string, t string) bool {
	return isIsomorphicHelper(s, t) && isIsomorphicHelper(t, s)
}

func isIsomorphicHelper(s string, t string) bool {
	n := len(s)
	if n == 0 {
		return true
	}
	checkMap := map[byte]byte{}

	for i := 0; i < n; i++ {
		c1 := s[i]
		c2 := t[i]
		if tmp, ok := checkMap[c1]; ok {
			if c2 != tmp {
				return false
			}
		}
		checkMap[c1] = c2
	}

	return true
}
