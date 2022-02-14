package main

import "fmt"

func main() {
	channel := make(chan int, 3) //channel o'zida 3ta ma'lumot saqlaydi
	channel <- 4                 //1-ma'lumot yozildi
	channel <- 5                 //2-ma'lumot yozildi
	channel <- 7                 //3-ma'lumot yozildi channelto'ldi
	channel <- 8                 //4-ma'lumotni ma'lumotni yozmoqchi bo'lyapmiz
	fmt.Println(<-channel, <-channel, <-channel)
}
