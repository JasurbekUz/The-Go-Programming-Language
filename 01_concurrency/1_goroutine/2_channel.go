package main

import "fmt"

func main() {
	var channel chan bool //bool toifadagi ma'lumotlarnitashuvchikanal

	fmt.Printf("zero value: %v, type: %T", channel, channel)
}
