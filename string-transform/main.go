package main

import (
	"fmt"
	"strconv"
)

func main() {

	//transNo := []byte(strconv.Itoa(1))
	transNo := 49
	fmt.Println(transNo)
	fmt.Println(stringToBin(strconv.Itoa(transNo)))
	fmt.Println(string(transNo))
	fmt.Println(stringToBin(string(transNo)))
}

/**
11111111 � 1111111111111101 &#xFFFD;
11111111 11111111 110001110001110001110001110001110001110001110001 11111111
*/

func stringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s--%b--%x", s, c, c)
	}
	return
}
