package leetcode

func MySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left, right := 1, x/2
	for left < right {
		middle := left + (right-left+1)>>1
		if middle*middle > x {
			right = middle - 1
		} else {
			left = middle
		}
	}

	return left
}
