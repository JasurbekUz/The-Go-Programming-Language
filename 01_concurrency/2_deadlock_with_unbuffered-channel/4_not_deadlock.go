package main

import "fmt"

func main() {
	channel := make(chan int)
	go func(cahnnel chan int) {
		channel <- 25
	}(channel)

	fmt.Println(<-channel)
}
