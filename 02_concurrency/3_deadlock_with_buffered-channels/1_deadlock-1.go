package main

import "fmt"

func main() {
	channel := make(chan int, 3)

	chanValue := <-channel //ma'llumot olmoqchibo'lyapmiz,lekin ma'lumot bermyapmiz
	fmt.Println(chanValue)
}
