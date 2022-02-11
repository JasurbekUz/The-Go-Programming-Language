package main

import (
	"fmt"
	"reflect"
)

func main() {

	var color string

	color = "bleak"

	fmt.Println(color, reflect.TypeOf(color).Name()) //bleak string

	bw := color + " " + "white"

	fmt.Println(bw, reflect.TypeOf(bw)) //bleak white string

	fmt.Println(bw[4], reflect.TypeOf(bw[4]))                 //107 uint8
	fmt.Println(string(bw[4]), reflect.TypeOf(string(bw[4]))) //k string

	fmt.Println(bw[2:3]) //e
	fmt.Println(bw[2:])  //eak white
	fmt.Println(bw[2:9]) //eak whi
	fmt.Println(bw[:9])  //bleak whi

}
