package main

import "fmt"

type S struct {
	A string
}

func main() {
	var m map[string]int
	fmt.Printf("map statement variable %p\n", m)
	fmt.Printf("map statement pointer %p\n", &m)

	var s *S
	fmt.Printf("struct statement variable %p\n", s)
	fmt.Printf("struct statement pointer %p\n", &s)

	m = make(map[string]int)
	fmt.Printf("map initialize variable %p\n", m)
	fmt.Printf("map initialize pointer %p\n", &m)

	s = new(S)
	fmt.Printf("struct initialize variable %p\n", s)
	fmt.Printf("struct initialize pointer %p\n", &s)

	testMap(m)
	testStruct(s)
}

func testMap(m map[string]int) {
	fmt.Printf("map function variable %p\n", m)
	fmt.Printf("map function pointer %p\n", &m)
}

func testStruct(s *S) {
	fmt.Printf("struct function variable %p\n", s)
	fmt.Printf("struct function pointer %p\n", &s)
}
