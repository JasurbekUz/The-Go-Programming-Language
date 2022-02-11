package main

import (
	"fmt"
	"reflect"
)

func main() {

	var idf bool = true
	fmt.Println(idf, reflect.TypeOf(idf).Name())

	var foo = (12 == 12)
	fmt.Println(foo, reflect.TypeOf(foo).Name())

	var goo = (12 == 24)

	fmt.Println(goo, reflect.TypeOf(goo).Name())

	var juba = (12 > 11)
	fmt.Println(juba)

	fmt.Println(reflect.TypeOf(reflect.TypeOf(juba)))
}
