package merge

func merge(arr []int) []int {
	length := len(arr)
	var temp []int

	middle := length / 2
	var value int
	i := 0
	j := middle
	for i < middle && j < length {
		if arr[i] < arr[j] {
			value = arr[i]
			i++
		} else {
			value = arr[j]
			j++
		}
		temp = append(temp, value)
	}

	if middle > i {
		temp = append(temp, arr[i:middle]...)
	}
	if length > j {
		temp = append(temp, arr[j:length]...)
	}
	return temp
}

/**
归并排序
*/
func Sort(arr []int) []int {
	length := len(arr)
	if length == 1 {
		return arr
	}

	tempA := arr[0 : length/2]
	tempA = Sort(tempA)
	tempB := arr[length/2 : length]
	tempB = Sort(tempB)
	return merge(append(tempA, tempB...))
}
