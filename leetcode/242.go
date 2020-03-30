package leetcode

func IsAnagram(s string, t string) bool {
	flag := true
	el := 26
	es, et := make([]int, el), make([]int, el)
	for si := range s {
		es[s[si]-'a']++
	}
	for ti := range t {
		et[t[ti]-'a']++
	}
	for i := 0; i < el; i++ {
		if es[i] != et[i] {
			flag = false
			break
		}
	}

	return flag

}
