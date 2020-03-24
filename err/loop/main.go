package main

import "fmt"

func main() {
	values := []string{"a", "b", "c"}
	var funcs []func()
	for _, val := range values {
		funcs = append(funcs, func() {
			fmt.Println(val)
		})
	}
	for _, f := range funcs {
		f()
	}
}
