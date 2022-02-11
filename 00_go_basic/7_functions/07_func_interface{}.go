package main

import (
	"fmt"
	"reflect"
)

func echo(data interface{}) {

	fmt.Println(reflect.TypeOf(data).Name() == "int")
}

func main() {

	echo(20)
}
