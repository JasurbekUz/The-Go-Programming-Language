package main

import (
	"fmt"
	"reflect"
)

func for_uint_ints() {

	var (
		age, color         uint8
		year               uint
		integer1, integer2 int
	)

	age = 46
	color = 254
	year = 2021
	integer1 = 123988789
	integer2 = -23123

	fmt.Println(age, color, reflect.TypeOf(age).Name())
	fmt.Println(year, reflect.TypeOf(year).Name())
	fmt.Println(integer1, integer2, reflect.TypeOf(integer1).Name())
}

func for_floats() {

	var (
		percent float32
		area    float64
	)

	percent = 75.8
	area = 43455.3434

	fmt.Println(percent, reflect.TypeOf(percent).Name())
	fmt.Println(area, reflect.TypeOf(area).Name())
}

func main() {

	for_uint_ints()
	for_floats()
}
