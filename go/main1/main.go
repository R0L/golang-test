package main

type base int

const (
	A base = iota
	C
	T
	G
)

type Baser interface {
	Base() base
}

// every base must fulfill the Baser interface
func (b base) Base() base {
	return b
}

func (b base) OtherMethod() {
}

func main() {
	testBase(32)
}

func testBase(b base) {
}
