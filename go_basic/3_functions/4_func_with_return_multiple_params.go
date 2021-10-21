//

package main

import (
	"fmt"
)

func getColorDec() (uint8, uint8, uint8) {

	return 235, 254, 255
}

func getColor() (string, string, string) {

	return "RED", "GREEN", "BLUE"
}

func calc(number1, number2 float64) (res1, res2, res3, res4 float64) {

	res1 = number1 + number2

	res2 = number1 - number2

	res3 = number1 * number2

	res4 = number1 / number2

	return res1, res2, res3, res4
}

func direction(num1 uint, num2 uint8) (uint, uint8) {

	return 2 * num1, 2 * num2
}

func main() {

	//func getColorDec
	fmt.Println("function getColorDec")
	R, G, B := getColorDec()
	fmt.Println(R, G, B)

	//func getColor
	//funksiyadan qaytgan qiymatlarning keraksizini olmaslik uchun 'empty idintifier' dan foydalanamiz:
	fmt.Println("\n\nfunction getColor")
	RR, _, BB := getColor()
	fmt.Println(RR, BB)

	//func calc
	fmt.Println("\n\nfunction calc")
	result1, result2, result3, result4 := calc(15, 5)
	fmt.Println(result1, result2, result3, result4)

	//func direction
	fmt.Println("\n\nfunction direction")
	L, R := direction(10, 15)

	fmt.Println(L) //20
	fmt.Println(R) //30
}
