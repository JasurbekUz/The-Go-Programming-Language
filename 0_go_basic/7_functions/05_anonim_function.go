package main

import (
	"fmt"
)

func main() {

	// anonim function 1
	func() {

		fmt.Println("hello, world!")
	}() // hello world

	// anonim function 2
	func(x int) {

		fmt.Println(x * x)
	}(4) //16

	// anonim function 3
	fmt.Println(func(x, y int) int {

		return x + y

	}(2, 3)) // 5

	// anonim function 4
	var square = func(n int) int {

		return n * n
	}(3)

	fmt.Println(square) //9

	// anonim function 5
	var cube = func(foo float64) float64 {

		return foo * foo * foo
	}

	fmt.Println(cube(3)) //27
}
