package main

import "fmt"

func main() {
	var a, b = 1, 4

	//b, a = a+1, b

	tempA := a
	tempB := b
	b = tempA + 1
	a = tempB

	fmt.Println(a, b)
}
