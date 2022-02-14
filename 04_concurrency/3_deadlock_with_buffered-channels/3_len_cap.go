package main

import "fmt"

func main() {
	channel := make(chan int, 7)
	channel <- 7
	channel <- 5
	channel <- 3
	fmt.Printf("uzunligi: %d; sig'imi: %d\n", len(channel), cap(channel))
	fmt.Println(<-channel) //bitta ma'lumot olyapmiz, bunda birinchi yozgan ma'lumotimiz keladi!
	fmt.Printf("uzunligi: %d; sig'imi: %d", len(channel), cap(channel))
}
