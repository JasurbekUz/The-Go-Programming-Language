package main

import "fmt"

func fn(ch <-chan int) {
	fmt.Println(<-ch)
}
func main() {
	ch := make(chan int)

	go fn(ch)
	ch <- 2
}
