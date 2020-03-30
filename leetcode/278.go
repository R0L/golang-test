package leetcode

func isBadVersion(n int) bool {
	return false
}

func firstBadVersion(n int) int {
	if n == 0 {
		return 0
	}
	left, right := 1, n
	for left < right {
		middle := left + (right-left)>>1
		if isBadVersion(middle) {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}
