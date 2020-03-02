package main

import (
	"fmt"
	"github.com/R0L/golang-test/pool/pool"
	"io"
	"time"
)

type Test struct {
	a int
}

func (t *Test) Close() error {
	t.a = 0
	return nil
}

func main() {

	p, err := pool.New(func() (closer io.Closer, e error) {
		t := &Test{}
		t.a = time.Now().Second()
		return t, nil
	}, 10)

	fmt.Println(p, err)

	fmt.Println(p.Get())

}
