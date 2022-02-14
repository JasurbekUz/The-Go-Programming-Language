package main

import "fmt"

func GetNumber() {
	var number int
	go func() { // ma'lumot yozilyapti
		number = 7
	}()
	fmt.Println(number) // malumot o'qilyapti
}
func main() {
	GetNumber()
}
