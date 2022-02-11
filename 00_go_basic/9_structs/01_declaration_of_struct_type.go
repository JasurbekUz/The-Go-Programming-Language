// Dasturchi: Jasurbek Shamsitdinov
// Maqsad: struct type elon qilish va uning fieldlariga ma'lumot yozishni organish
// Sana: 2021 10 21

package main

import (
	"fmt"
)

type rectangle struct {
	length  float64
	breadth float64
	color   string
}

func main() {

	fmt.Println(rectangle{12.2, 13, "green"})

	// struct fieldlariga ma'lumot yozish:

	/*var rect = rectangle{12, 23, "yellow"}
	fmt.Println(rect)*/

	/*var rect1 rectangle = rectangle{12, 34, "grey"}
	fmt.Println(rect1)*/

	/*var react2 = rectangle{length: 12, breadth: 34, color: "white"}
	fmt.Println(react2)*/

	/*rect3 := rectangle{13, 34, "blue"}
	fmt.Println(rect3)*/

	/*var rect4 rectangle

	rect4.length = 12
	rect4.breadth = 13
	rect4.color = "black"

	fmt.Println(rect4)*/

}
