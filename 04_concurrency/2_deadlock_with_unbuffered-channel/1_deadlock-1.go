package main

import "fmt"

func main() {
	channel := make(chan int)

	chanValue := <-channel

	fmt.Println(chanValue)
}
