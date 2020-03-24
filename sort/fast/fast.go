package fast

import "fmt"

/**
默认left为basic
*/
func fast(arr []int, l int, r int) {
	if l == r {
		return
	}
	basic := r
	left := l
	right := r - 1
loop:
	for arr[left] < arr[basic] {
		fmt.Printf("left: %v -- %v<%v\n", left, arr[left], arr[basic])
		if left == right {
			fast(arr, l, r-1)
			return
		}
		left++
	}
	for arr[right] >= arr[basic] {
		fmt.Printf("right: %v -- %v>%v\n", right, arr[right], arr[basic])
		if left == right {
			fmt.Printf("1 ---%v--\n", arr)
			arr[basic], arr[left] = arr[left], arr[basic]
			fmt.Printf("2 ---%v--\n", arr)
			fast(arr, l+1, r)
			return
		}
		right--
	}
	if left == right {
		if l < left-1 {
			fast(arr, l, left-1)
		}
		if right+1 < r {
			fast(arr, right+1, r)
		}
		return
	}
	fmt.Printf("3 ---%v--\n", arr)
	arr[left], arr[right] = arr[right], arr[left]
	fmt.Printf("4 ---%v--\n", arr)
	goto loop
}

// 快速排序
func Sort(arr []int) []int {
	fast(arr, 0, len(arr)-1)
	return arr
}
