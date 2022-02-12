package main

import "fmt"

func Add(channel chan int, num1, num2 int) {
	channel <- num1 + num2
}

func Mult(channel chan int, num1, num2 int) {
	channel <- num1 * num2
}
func main() {
	channel1 := make(chan int)
	go Add(channel1, 3, 4)
	go Mult(channel1, 3, 4)

	fmt.Println(<-channel1, <-channel1)
}
