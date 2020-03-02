package main

import (
	"fmt"
	"time"
)

func main() {
	var flag = make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		close(flag)
	}()
	f, ok := <-flag
	fmt.Println(f, ok)
}
