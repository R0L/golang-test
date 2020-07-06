package main

import (
	"time"
)

func main() {

	t := time.NewTimer(time.Second)

	go func(t *time.Timer) {
		select {
		case <-t.C:
			println(1)
		}
	}(t)

	tt := time.NewTimer(10 * time.Second)

	select {
	case <-tt.C:
	}

}
