package main

import (
	"encoding/json"
	"fmt"
)

//  json: invalid number literal "我是"
func main() {
	data := struct {
		Key json.Number `json:"key"`
	}{
		"我是",
	}

	dataByte, err := json.Marshal(data)
	fmt.Println(dataByte, err)
}
