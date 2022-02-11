# The-Go-Programming-Language
 ### 0 go_basic

 ### 01 concurrency
 - **concurrency** - ikki yoki undan ortiq **vazifalarni** bir vaqtning o'zida (parallel) bajarish tushunchasi. **vazifalar** metodlar (funksiyalari), dastur qismlari yoki boshqa dasturlarni o'z ichiga olishi mumkin.
- Golangdagi parallellik - bu funktsiyalarning bir-biridan mustaqil ishlashi qobiliyati. Goroutine - bu boshqa funktsiyalar bilan bir vaqtda ishlashga qodir bo'lgan funktsiya. Funktsiyani gorutin sifatida yaratganingizda, u rejalashtirilgan bo'lib, keyin mavjud mantiqiy protsessorda bajariladigan mustaqil ish birligi sifatida ko'rib chiqiladi.

 #### 1 goroutain
* **goroutina** - bu boshqa funktsiyalar bilan bir vaqtda ishlashga qodir bo'lgan funktsiya.

* Go tilida multi **trade** applarni qilish uchun goroutina dan foydalanamiz
 
 - **trade** - bitta protsesda multiple qilib ochiladigan protses
 - main funksiya ham goroutine hisoblanadi va dastur ishlaganda boshqa gorutinalarni ishlashini kutmaydi. Yani main funksiya bilan uning ichidagi goroutina lar bir vaqtda ishga tushadi
 ```
 package main

import "fmt"

func main() {
	fmt.Println("hello, world")
}
 ```
 boshqa goroutine lar **go** key wordi orqali yasaladi:
 ```
 package main

import "fmt"

func message() {
	fmt.Println("goodbye world")
}
func main() {
	go message() // goroutine.
	fmt.Println("hello, world")
}
 ```
#### 1 channel
- **channel** - **goroutine** lar bir birini kutib turmaydi yani bitta vaqtning o'zida ishga tushadi. Shuning uchun ular orasida comunicatsiya boishi kerak (gorotine lar aro ma'lumot almashish vostasi) - biz bunda **channel**dan foydalanamiz
* **channel**ni elon qilish:
```
package main

import "fmt"

func main() {
	var channel chan bool //bool toifadagi ma'lumotlarni tashuvchi kanal

	fmt.Printf("zero value: %v, type: %T", channel, channel)
}
```
* channellarham slice va map kabi buildIn taype va yasamasdan oldin ishlatib bolmaydi!
* channellardan ma'lumot yuborish va ma'lumot qabul qilish uchun foydalanamiz! <br><br>
* making channel:
```package main

import "fmt"

func main() {
	channel := make(chan int)

	fmt.Println(channel)
}

```
* channel dan foydalanish (reading, writing)
```
package main

import "fmt"

func bye(channel chan bool) {
	fmt.Println("goodbye, world!")
	channel <- true // kanalga ma'lumot yozish.
}
func main() {
	channel := make(chan bool)

	go bye(channel)

	chanValue := <- channel // kanaldan ma'lumot olish.
	fmt.Println("hello, world")
	fmt.Println(chanValue)
}

``` 
#### 1 Deadlock
* by default channellar *unbaffered channel* deyiladi: ***make(chan int)***
##### unbuffered channellarda: 
* Agar bizda boror bir kanal bo'lsayu, biz kanladan ma'lumot qabul qilmoqchi bo'lsagu kanaldan concurent hech kim ma'lumot jo'natmasa **Deadlock** bo'ladi
```
package main

import "fmt"

func main() {
	channel := make(chan int)

	chanValue := <-channel // deadlock

	fmt.Println(chanValue)
}
/*
out:
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
	/home/ubuntu/Desktop/JasurbekUz/The-Go-Programming-Language/01_concurrency/2_deadlock/1_deadlock-1.go:8 +0x39
exit status 2
*/
```
* agar biz kanalga ma'lumot yozsagu, lekin concurent tarzda uni qabul qiluvchi bo'lmasaham **deadlock** bo'ladi!
```
package main

func main() {
	channel := make(chan int)

	channel <- 7 // deadlock
}
/*
out:
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/home/ubuntu/Desktop/JasurbekUz/The-Go-Programming-Language/01_concurrency/2_deadlock/tempCodeRunnerFile.go:6 +0x31
exit status 2
*/
```
* deadlock - runtime paytida aniqlanadi!

ex3:
```
package main

import "fmt"

func main() {
	channel := make(chan int)
	channel <- 25
	fmt.Println(<-channel)
}
/*
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/home/ubuntu/Desktop/JasurbekUz/The-Go-Programming-Language/01_concurrency/2_deadlock/3_deadlock-3.go:7 +0x37
exit status 2
*/
```
* yuqoridagi misolda channelga malumot yozilyapti va u o'qilyapti lekin bu jarayon concurent tarzda ro'y bermayapti shuning uchun ushbu holat deadlock hisoblanadi

buni oldini olish uchun channelga main bilan bir vaqtda ishlaydigan funksiyada ya'ni goroutine da ma'lumot yozishimiz kerak:
```
package main

import "fmt"

func main() {
	channel := make(chan int)
	go func(cahnnel chan int) {
		channel <- 25
	}(channel)

	fmt.Println(<-channel)
}
```
### buffered channel
* 