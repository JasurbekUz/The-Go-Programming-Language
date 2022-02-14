package main

import "fmt"

func Input(anyType interface{}) {
	switch anyType.(type) {
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("Unsupported type")
	}
}
func main() {
	var anyType interface{}
	anyType = 25
	Input(anyType)
}
