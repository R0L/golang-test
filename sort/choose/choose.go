package choose

import "fmt"

// 选择排序
func Sort(arr []int) []int {
	length := len(arr)
	var min int
	var compareTimes, moveTimes int
	for i := 0; i < length; i++ {
		min = i
		for j := i + 1; j < length; j++ {
			compareTimes++
			fmt.Printf("choose Sort compare %v times, %v -- %v\n", compareTimes, arr[min], arr[j])
			if arr[min] > arr[j] {
				min = j
			}
		}
		if min != i {
			moveTimes++
			fmt.Printf("choose Sort move %v times, -- %v -- %v --> %v\n", moveTimes, arr, arr[i], arr[min])
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}
