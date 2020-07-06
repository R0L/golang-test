package main

import (
	"fmt"
	"github.com/R0L/golang-test/seek/binary"
)

func main() {

	var arr = []int{
		//1, 3, 4, 7, 9, 10, 13, 17, 24, 32, 33, 34, 43, 53,
		1, 2,
	}

	locate := binary.Search(arr, 33)
	fmt.Println(locate)
	locate = binary.SearchOpt(arr, 33)
	fmt.Println(locate)

}
