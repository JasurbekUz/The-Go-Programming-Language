// Dasturchi: Jasurbek Shamsidionov
// Maqsad: struct ichida struct ishlatishni o'rganish
// Sana: 2021 10 21

package main

import (
	"fmt"
)

//	  *** 1 ***    //
type rectangle struct {
	length  float64
	breadth float64
	color   string

	geometry struct {
		area      float64
		perimeter float64
	}
}

// 	 ***  2 ***   //
type Color struct {
	red, green, blue uint8
}

type Car struct {
	name  string
	color Color //<===
}

func main() {

	//*2*
	rect := rectangle{length: 10, breadth: 20, color: "green"}
	rect.geometry.area = 1234
	rect.geometry.perimeter = 123

	fmt.Println(rect)

	//*1*
	car := Car{name: "malibu", color: Color{red: 123, green: 234, blue: 213}}

	fmt.Println(car)
}
