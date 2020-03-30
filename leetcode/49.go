package leetcode

import "fmt"

func GroupAnagrams(strs []string) [][]string {
	el := 26
	var (
		ret   = make([][]string, len(strs))
		label = make(map[string]int)
		index = 0
	)
	for _, str := range strs {
		temp := make([]int, el)
		for j := range str {
			temp[str[j]-'a']++
		}
		esStr := ""
		for _, e := range temp {
			esStr += fmt.Sprintf("#%d", e)
		}

		i, ok := label[esStr]
		if !ok {
			label[esStr] = index
			i = index
			index++
		}
		ret[i] = append(ret[i], str)
	}

	return ret[0:index]
}
