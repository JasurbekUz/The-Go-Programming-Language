package main

import (
	"fmt"
)

// func 1
func calc(number1, number2 float64) (add, div func() float64) {

	add = func() float64 {

		return number1 + number2
	}

	div = func() float64 {

		return number1 - number2
	}

	return add, div
}

// super func
func super(num1 int) func(num2 int) func(num3 int) int {

	return func(num2 int) func(num3 int) int {

		return func(num3 int) int {

			return num1 * num2 * num3
		}
	}

}

func main() {

	// task 1
	plus, minus := calc(30, 10)

	fmt.Println(plus())  // 40
	fmt.Println(minus()) // 20

	// task 2
	a := super(4)(5)(6)

	fmt.Println(a) // 120
}
