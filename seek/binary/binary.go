package binary

import "fmt"

func Search(arr []int, dest int) int {

	fmt.Printf("arr: %v, dest: %v\n", arr, dest)
	left, right := 0, len(arr)-1
	for left <= right {
		middle := (left + right) >> 1
		fmt.Printf("left: %v, right: %v, middle: %v, value: %v\n", middle, left, right, arr[middle])
		if arr[middle] == dest {
			return middle
		} else if arr[middle] > dest {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1

}

//
