package main

func main() {
	channel := make(chan int, 3)
	channel <- 4
	channel <- 5
	channel <- 6
	//channel to'ldi, lekin ma'lumot qabulq qilmayapmi
	//bu deadlock emas!!!
}
