package main

import "fmt"

type Animal struct {
	Name string
	mean bool
}

type Cat struct {
	Basics       Animal
	MeowStrength int
}

type Dog struct {
	Animal
	BarkStrength int
}

func main() {
	fmt.Printf("%+v", Cat{})
}
