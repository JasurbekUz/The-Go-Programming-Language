package main

import "fmt"

func fn(ch chan<- int) {
	//fmt.Println(<-ch) //ma'lumot o'qib bo'lmaydi!
	ch <- 7 // ma'lumotjo'natdik
}
func main() {
	ch := make(chan int)

	go fn(ch)
	fmt.Println(<-ch)
}
