// sana: 2021 10 19
// dasturchi: Jasurbek Shamsitdinov
// maqsad: assaignement va short-hand operatorlari farqi

package main

import (
	"fmt"
)

func difference1() {

	var x, y, z = 0, 0, 0

	//x, y, z := 0, 0, 0 //wrong
	//no new variables on left side of :=

	//x, y, z ning qiymatini o'zgartiradi va t ni e'lon qilib qiymat bberadi
	x, y, z, t := 1, 2, 3, 4

	//fmt.Println(x, y, z)
	fmt.Println(x, y, z, t)
}

func difference2() {

	var x, y, z = 0, 0, 0 //typed

	//x, y, z, t := "AB", "BC", "CD", 'd' //wrong
	//cannot use "AB" (type untyped string) as type int in assignment

	//x, y, z e'lon qilingan va turi aniqlangan
	x, y, z, t := 2, 3, 4, "salom"

	fmt.Println(x, y, z, t)
}

func main() {

	difference1()
	difference2()
}
