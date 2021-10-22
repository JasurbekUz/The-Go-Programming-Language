// Dasturchi: Jasurbek Shamsitdinov
// Maqsad: structlarga metodlar e'lon qilishni o'rganish
// Sana: 2021 10 21

package main

import (
	"fmt"
)

type Color struct {
	R uint8
	G uint8
	B uint8
}

func (c Color) getColor() int { //<===

	return int(c.R) + int(c.G) + int(c.B)
}

type Car struct {
	name  string
	color Color
}

func main() {

	var white = Color{123, 123, 213}
	fmt.Println(white.getColor())

	var car Car = Car{"nexia I", Color{12, 233, 122}}
	fmt.Println(car)
	fmt.Println(car.color.getColor())

}
