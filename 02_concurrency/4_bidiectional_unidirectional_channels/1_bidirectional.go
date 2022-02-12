package main

import "fmt"

func function(ch chan int) {
	fmt.Println(<-ch) //main funcdan kelgan ma'lumot o'qildi
	ch <- 225         // ma'lumotyozildi!
}
func main() {
	ch := make(chan int)
	go function(ch)
	ch <- 5           //channelga ma'lumot yozildi
	fmt.Println(<-ch) //functiondan kelgan ma'lumot o'qildi!
}
