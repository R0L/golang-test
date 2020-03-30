package leetcode

func MinEatingSpeed(piles []int, H int) int {

	if len(piles) == 0 {
		return 0
	}
	var left, right, total int
	for _, pile := range piles {
		if pile > right {
			right = pile
		}
		total += pile
	}
	if H == 0 {
		return right
	}
	left = (total + H - 1) / H

	for left < right {
		var middle, total int
		middle = (left + right) >> 1
		for _, pile := range piles {
			total += (pile + middle - 1) / middle
		}
		if total > H {
			left = middle + 1
		} else if total <= H {
			right = middle
		}
	}
	return left
}
