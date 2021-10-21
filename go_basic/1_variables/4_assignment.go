package main

import (
	"fmt"
)

func inplesid() {

	var age, color, width = 27, "black", 200

	fmt.Println(age, color, width)
}

func explesid() {

	var (
		age, color uint8
		year       uint
	)

	age, color = 22, 255
	year = 2021

	fmt.Println(age, color, year)
}

func main() {

	inplesid()
	explesid()
}
