package Week_02

import "sort"

func groupAnagrams(strs []string) [][]string {
	resMap := make(map[string][]string)
	for _, str := range strs {
		strBytes := []byte(str)
		sort.Slice(strBytes, func(i, j int) bool {
			return strBytes[i] < strBytes[j]
		})
		sortedStr := string(strBytes)
		resList, ok := resMap[sortedStr]
		if ok {
			resList = append(resList, str)
		} else {
			resList = []string{str}

		}
		resMap[sortedStr] = resList
	}

	res := [][]string{}
	for _, resList := range resMap {
		res = append(res, resList)
	}

	return res
}
