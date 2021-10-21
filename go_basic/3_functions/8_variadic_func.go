package main

import (
	"fmt"
)

func numbers(numbers ...int) {

	fmt.Println(numbers)
}

func main() {

	numbers(1, 2, 3, 4, 5, 6)
}
