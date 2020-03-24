package bubble

import "fmt"

// 伪冒泡
func Sort(arr []int) []int {
	length := len(arr)
	var compareTimes, moveTimes int
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			compareTimes++
			fmt.Printf("bubble Sort compare %v times, %v -- %v\n", compareTimes, arr[j], arr[i])
			if arr[i] > arr[j] {
				moveTimes++
				fmt.Printf("bubble Sort move %v times, -- %v -- %v --> %v\n", moveTimes, arr, arr[j], arr[i])
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

// 冒泡
func SortOpt(arr []int) []int {
	length := len(arr)
	var flag = true
	var compareTimes, moveTimes int
	for i := 0; i < length && flag; i++ {
		for j := length - 1; j > i; j-- {
			flag = false
			compareTimes++
			fmt.Printf("bubble SortOpt compare %v times, %v -- %v\n", compareTimes, arr[j-1], arr[j])
			if arr[j-1] > arr[j] {
				flag = true
				moveTimes++
				fmt.Printf("bubble SortOpt move %v times, -- %v -- %v --> %v \n", moveTimes, arr, arr[j-1], arr[j])
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return arr
}
