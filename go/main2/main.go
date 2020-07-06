package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type (
	Dao struct {
		d int
	}
	Model struct {
		*Dao
		d Dao
	}
)

func main() {
	t := Model{Dao: &Dao{}}
	//fmt.Printf("%+v", t)
	spew.Dump(t)
	str := spew.Sdump(t)
	fmt.Println(str)
}
