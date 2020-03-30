package leetcode

func numMovesStones(a int, b int, c int) []int {

	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
		if a > b {
			a, b = b, a
		}
	}

	left := b - a - 1
	right := c - b - 1
	mix := 2
	if left == 0 || right == 0 || left == 1 || right == 1 {
		mix = 1
		if left+right == 0 {
			mix = 0
		}
	}
	return []int{mix, left + right}
}
