package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 100; i++ {
			time.Sleep(1 * time.Microsecond)
			fmt.Println("A:", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 100; i++ {
			time.Sleep(1 * time.Microsecond)
			fmt.Println("B:", i)
		}
	}()
	wg.Wait()
}
