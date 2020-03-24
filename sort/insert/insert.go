package insert

import "fmt"

// 插入排序
/**
最好使用链表来实现，当下面提供一种数组实现的方案
*/
func Sort(arr []int) []int {
	length := len(arr)
	var compareTimes, moveTimes int
	for i := 0; i < length-1; i++ {
		compareTimes++
		fmt.Printf("insert Sort compare %v times, %v -- %v\n", compareTimes, arr[i], arr[i+1])
		if arr[i] > arr[i+1] {
			moveTimes++
			fmt.Printf("insert Sort move %v times, -- %v -- %v --> %v\n", moveTimes, arr, arr[i], arr[i+1])
			arr[i], arr[i+1] = arr[i+1], arr[i]
			for j := i; j > 0; j-- {
				compareTimes++
				fmt.Printf("insert Sort compare %v times, %v -- %v\n", compareTimes, arr[j-1], arr[j])
				if arr[j-1] > arr[j] {
					moveTimes++
					fmt.Printf("insert Sort move %v times, -- %v -- %v --> %v\n", moveTimes, arr, arr[j-1], arr[j])
					arr[j-1], arr[j] = arr[j], arr[j-1]
				}
			}
		}
	}
	return arr
}
