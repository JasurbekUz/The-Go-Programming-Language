package main

import "fmt"

func main() {
	channel :=
		make(chan int)

	fmt.Println(channel)
}
