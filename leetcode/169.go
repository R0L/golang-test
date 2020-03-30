package leetcode

func MajorityElement(nums []int) int {

	index := 0
	target := 0
	for _, num := range nums {
		if index == 0 {
			target = num
			index = 1
		} else if num == target {
			index++
		} else {
			index--
		}
	}
	return target

}
