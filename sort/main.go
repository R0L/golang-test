package main

import (
	"fmt"
	"github.com/R0L/golang-test/sort/bubble"
	"github.com/R0L/golang-test/sort/choose"
	"github.com/R0L/golang-test/sort/fast"
	"github.com/R0L/golang-test/sort/heap"
	"github.com/R0L/golang-test/sort/hill"
	"github.com/R0L/golang-test/sort/insert"
	"github.com/R0L/golang-test/sort/merge"
)

func main() {

	var arr = []int{
		7, 3, 2, 4, 5, 22, 66, 77, 21, 14, 12, 15, 100, 1,
	}

	var bs = make([]int, len(arr))
	copy(bs, arr)
	fmt.Println(bubble.Sort(bs))

	var bso = make([]int, len(arr))
	copy(bso, arr)
	fmt.Println(bubble.SortOpt(bso))

	var cs = make([]int, len(arr))
	copy(cs, arr)
	fmt.Println(choose.Sort(cs))

	var is = make([]int, len(arr))
	copy(is, arr)
	fmt.Println(insert.Sort(is))

	var hs = make([]int, len(arr))
	copy(hs, arr)
	fmt.Println(hill.Sort(hs))

	var ms = make([]int, len(arr))
	copy(ms, arr)
	fmt.Println(merge.Sort(ms))

	var heaps = make([]int, len(arr))
	copy(heaps, arr)
	fmt.Println(heap.Sort(heaps))

	var fs = make([]int, len(arr))
	copy(fs, arr)
	fmt.Println(fast.Sort(fs))

}
