package main

import (
	"fmt"
	"strings"
)

func main() {
	words := "mongodb://omngff"
	prefix := "mongodb://"
	fmt.Println(strings.TrimLeft(words, prefix))
	words = "mongodb://xgoff"
	prefix = "mongodb://"
	fmt.Println(strings.TrimLeft(words, prefix))

	words = "mongodb://omngff"
	prefix = "mongodb://"
	fmt.Println(strings.TrimPrefix(words, prefix))
	words = "mongodb://xgoff"
	prefix = "mongodb://"
	fmt.Println(strings.TrimPrefix(words, prefix))

	words = "mongodb://xgoff"
	prefix = "mongodb://ff"
	fmt.Println(strings.Trim(words, prefix))
}
