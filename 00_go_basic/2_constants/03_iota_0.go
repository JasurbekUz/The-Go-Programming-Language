package main

import (
	"fmt"
)

const (
	ADD uint = iota
	SUB
	DIV
	MUL
)

func main() {

	fmt.Println(ADD) // 0
	fmt.Println(SUB) // 1
	fmt.Println(DIV) // 2
	fmt.Println(MUL) // 3
}
