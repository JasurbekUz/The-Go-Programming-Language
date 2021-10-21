package main

import (
	"fmt"
)

// ikki sonlarni qo'shish
func add(number1, number2 int) {

	fmt.Println(number1 + number2)
}

// ikki sonlarni ayirish
func sub(number1, number2 int) float64 {

	return float64(number1 - number2)
}

// ikki sonlarni ko'paytirish
func mult(number1, number2 int) (result int) {

	result = number1 * number2

	return result
}

// ikki sonlarni bo'lish
func Div(number1 float32, number2 int) (result float64) {

	result = float64(number1) / float64(number2)

	return result
}

func main() {

	add(3, 5)
	fmt.Println(sub(14, 7))
	fmt.Println(mult(3, 2))
	fmt.Println(Div(100, 20))
}
