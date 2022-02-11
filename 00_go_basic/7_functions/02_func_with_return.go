package main

import (
	"fmt"
)

func isActive() bool {

	return true
}

func five() int {

	return 5
}

func sayHi() string {

	return "Hi"
}

func main() {

	fmt.Println(isActive())
	fmt.Println(five())
	fmt.Println(sayHi())
}
