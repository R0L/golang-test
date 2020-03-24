package heap

/**
比较 i, left, right 他们的大小，移动最大的到i上面去，递归计算变动的结点
*/
func heapify(arr []int, i int) {
	length := len(arr)
	max := i
	left := 2*i + 1
	right := 2*i + 2
	if left >= length || right >= length {
		return
	}
	if arr[left] > arr[max] {
		max = left
	}
	if arr[right] > arr[max] {
		max = right
	}

	if max != i {
		arr[i], arr[max] = arr[max], arr[i]
		heapify(arr, max)
	}
}

/**
从下到上依次的heapify，每次去打顶堆或者小顶堆的根结点，然后把最后一个元素移动到根结点，再次执行heapify
*/
func SortHeapify(arr []int) []int {
	length := len(arr)
	for i := (length - 3) / 2; i >= 0; i-- {
		heapify(arr, i)
	}
	var temp []int
	for j := 0; j < length; j++ {
		temp = append(temp, arr[0])
		l := length - j - 1
		arr[0] = arr[l]
		if l > 0 {
			heapify(arr[0:l], 0)
		}
	}
	return temp
}

/**
堆排序
*/
func Sort(arr []int) []int {
	return SortHeapify(arr)
}
