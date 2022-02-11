package main

import "fmt"

func bye(channel chan bool) {
	fmt.Println("goodbye, world!")
	channel <- true
}
func main() {
	channel := make(chan bool)

	go bye(channel)

	chanValue :=
		<-channel
	fmt.Println("hello, world")
	fmt.Println(chanValue)
}
