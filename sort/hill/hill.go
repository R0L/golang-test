package hill

import "fmt"

// 希尔排序
/**
最好使用链表来实现，当下面提供一种数组实现的方案
*/
func Sort(arr []int) []int {
	length := len(arr)
	sep := length / 2
	var compareTimes, moveTimes int

loop:
	for i := 0; i < length-sep; i++ {
		compareTimes++
		fmt.Printf("hill Sort compare %v times, %v -- %v\n", compareTimes, arr[i], arr[i+sep])
		if arr[i] > arr[i+sep] {
			moveTimes++
			fmt.Printf("hill Sort move %v times, -- %v -- %v --> %v\n", moveTimes, arr, arr[i], arr[i+sep])
			arr[i], arr[i+sep] = arr[i+sep], arr[i]
			for j := i; j > sep-1; j-- {
				compareTimes++
				fmt.Printf("hill Sort compare %v times, %v -- %v\n", compareTimes, arr[j-sep], arr[j])
				if arr[j-sep] > arr[j] {
					moveTimes++
					fmt.Printf("hill Sort move %v times, -- %v -- %v --> %v\n", moveTimes, arr, arr[j-sep], arr[j])
					arr[j-sep], arr[j] = arr[j], arr[j-sep]
				}
			}
		}
	}
	if 1 != sep {
		sep = sep / 2
		goto loop
	}
	return arr
}
