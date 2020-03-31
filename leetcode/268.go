package leetcode

func MissingNumber(nums []int) int {
	length := len(nums)
	total := length * (length + 1) / 2
	for _, num := range nums {
		total -= num
	}
	return total
}
