//
//
//

package main

import (
	"fmt"
)

// call-back func 1
func builder(fn func()) {

	fn()
}

// call-back func 2
func run(fn func(int) int) int {

	return fn(3)
}

func square(num int) int {

	return num * num
}

// call-back func 3
func calc(num1, num2 int) (add, div, mult, sub func() int) {

	add = func() int {

		return (num1 + num2)
	}

	div = func() int {

		return (num1 - num2)
	}

	mult = func() int {

		return (num1 * num2)
	}

	sub = func() int {

		return (num1 / num2)
	}

	return add, div, mult, sub

}

func main() {

	// call func 1
	builder(func() {

		fmt.Println("I'm built")
	})

	// call func 2
	fmt.Println(run(square))

	// call func 3
	add, div, mult, sub := calc(5, 5)

	fmt.Println(add())
	fmt.Println(div())
	fmt.Println(mult())
	fmt.Println(sub())
}
