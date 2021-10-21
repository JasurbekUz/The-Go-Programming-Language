// sana: 2021 10 19
// dasturchi: Jasurbek Shamsitdinov
// maqsad: go da ozgaruvchilarni e'lon qilishni o'rganish

package main

import (
	"fmt"
)

func onceVariable() {

	var year uint

	fmt.Println(year)

}

func variableList() {

	//variable list || comma separate
	var age, color uint8

	fmt.Println(age, color)
}

func manyTypeVars() {

	var (
		age, color uint8
		year       uint
	)

	fmt.Println(age, color, year)
}

func main() {

	onceVariable()
	variableList()
	manyTypeVars()

}
