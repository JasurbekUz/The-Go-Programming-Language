package main

import "fmt"

func message() {
	fmt.Println("goodbye world")
}
func main() {
	go message()
	fmt.Println("hello, world")
}
