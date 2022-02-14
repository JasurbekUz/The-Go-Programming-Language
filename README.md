# The-Go-Programming-Language
 ### 0 go_basic
 #### operators
 			Golang Operators 

An operator is a symbol that tells the compiler to perform certain actions. 
The following lists describe the different operators used in Golang.


    1.Arithmetic Operators
    2.Assignment Operators
    3.Comparison Operators
    4.Logical Operators
    5.Bitwise Operators
```
1.Arithmetic Operators in Go Programming Language

Operator |	Description   |   Example |	Result
_________|________________|___________|_______________________________________
+ 		 |	Addition 	  |	  x + y   |	Sum of x and y
- 		 |	Subtraction   |	  x - y   |	Subtracts one value from another
* 		 |	Multiplication|   x * y   |	Multiplies two values
/ 		 |	Division 	  |	  x / y   |	Quotient of x and y
% 		 |	Modulus 	  |	  x % y   |	Remainder of x divided by y
++ 		 |	Increment 	  |	  x++ 	  | Increases the value of a variable by 1
-- 		 |	Decrement 	  |	  x-- 	  | Decreases the value of a variable by 1
```
***
### 01 modules, packages
### 02 interface
1. go da interface ni any typlar ni qabul qilish uchun ishlata olamiz:
```
package main

import "fmt"

func Input(anyType interface{}) {
	switch anyType.(type) { // anyType.(type) - type switch 
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("Unsupported type")
	}
}
func main() {
	var anyType interface{}
	anyType = 25
	Input(anyType)
}
/*
out:
integer*/
```
* **2.** 
### 03 embedded struct
- embedded struct - bu bitta struct boshqa bir struct ni o'z ichiga olishi (= OOP(Inheritance)):
```
package main

import "fmt"

type Component struct {
	Name  string
	Sugar int
}
type Fruit struct {
	Component
}
type Vagetable struct {
	Component
}

func main() {
	var fruit Fruit = Fruit{Component{"Apple", 2}}
	fruit.Name = "Cherry"
	vegetable := Vagetable{}
	fmt.Println(fruit)
	fmt.Println(vegetable)
}
```
- Method overriding - Super klasda e'lon qilingan metodni child classda qaytadan yozish.
```
package main

import "fmt"

type OS struct{}

// method of supperclass.
func (OS) Version() {
	fmt.Println("OS version")
}

type Linux struct {
	OS
}

type MacOs struct {
	OS
}

// method overriding.
func (MacOs) Version() {
	fmt.Println("MacOs version")
}
func main() {
	m := MacOs{}
	l := Linux{}
	m.Version() //calling method overriding
	l.Version() //calling method of supperclass
}

/*
out:
MacOs version
OS version
*/
```
- supper call
```
package main

import "fmt"

type OS struct {
	Name string
}

// method of supperclass.
func (o OS) GetName() string {
	return o.Name
}

type Linux struct {
	OS
}

type MacOs struct {
	OS
}

// method overriding.
func (l Linux) GetName() string {
	//supper call
	//l.Os.Name
	return "Linux - " + l.OS.Name
}
func main() {
	m := MacOs{OS{"Big Sur"}}
	l := Linux{OS{"Ubuntu"}}
	fmt.Println(m.GetName()) // calling method of supperclass
	fmt.Println(l.GetName()) // calling method overridding and calling method of supperclass
}

/*
out:
Big Sur
Linux - Ubuntu
*/
```
- bitta shablonga asoslangan multiple struclar bilan yagona funksiya yozishham interace orqali bo'ladi
```
package main

import "fmt"

type Os struct {
	Name string
}

type MacOs struct { // birinchi struct
	Os
}

type Linux struct { // ikkinchi struct
	Os
}

func (o Os) GetName() string { // supperclassning methodi
	return o.Name
}

type OSInterface interface { // interface
	GetName() string
}

func Run(os OSInterface) { // funksiya interfacening talablarini qondiruvchi typeqabul qiladi, bu talabni supperclass qondiradi, shuningdek uning vorislariham

	fmt.Println("Loading " + os.GetName())
}

type Windows struct {
	Os
}

func main() {

	Run(MacOs{Os{"Big Sur"}})
	Run(Linux{Os{"Ubuntu"}})
	Run(Windows{Os{"Windows 11"}})
}
```

### 04 concurrency
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
* buffered chanellarda ma'lumot tugagan bo'lsayu, biz ma'lumot olmaoqchi bo'lsak deadlock bo'ladi!
```
package main

import "fmt"

func main() {
	channel := make(chan int, 3)

	chanValue := <-channel //ma'llumot olmoqchibo'lyapmiz,lekin ma'lumot bermyapmiz
	fmt.Println(chanValue)
}
/*
out:
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
	/home/ubuntu/Desktop/JasurbekUz/The-Go-Programming-Language/01_concurrency/3_deadlock_with_buffered-channels/1_deadlock-1.go:7 +0x3c
exit status 2
*/
```
* buffered chanellar to'lgan bo'lsayu, biz yana ma'lumot yozmoqchi bo'lsak deadlock bo'ladi!
```
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
/*
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/home/ubuntu/Desktop/JasurbekUz/The-Go-Programming-Language/01_concurrency/3_deadlock_with_buffered-channels/2_deadlock-2.go:10 +0x7b
exit status 2
*/
```
* **len(), cap() funlsiyalai**
* len() funksiyasi orqali biz channelda nechta ma'lumto borligini bilamiz
* cap() funksiyasi orqali biz channelning sig'imini bilamiz
```
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
/*
out:
uzunligi: 3; sig'imi: 7
7
uzunligi: 2; sig'imi: 7
*/
```
* buffered channellarda yozilgan ma'lumotni concurrent qabul qilinishi shart emas, ular malumot olinishini kutadi!
```
package main

func main() {
	channel := make(chan int, 3)
	channel <- 4
	channel <- 5
	channel <- 6
	//channel to'ldi, lekin ma'lumot qabulq qilmayapmi
	//bu deadlock emas!!!
}
```
*yuqorida korgan barcha channellar Bidirectional channellar hisoblanadi!*
***
* **Bidirectional Channels,<br> Receive-only Channels,<br> Send-only Channels**
```
var bidirectionalChan chan string // can read from, write to and close()
var receiveOnlyChan <-chan string // can read from, but cannot write to or close()
var sendOnlyChan chan<- string    // cannot read from, but can write to and close
```
* Bidirectional Channels:
bu channeldan ma'lumot ma'lumot o'qibham, yoziham bo'ladi!
```
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
/*
out:
5
225
*/
```

* receive only channel:
bu channeldan faqatgina ma'lumot ololamiz(ma'lumot yozib bo'lmaydi!)
```
package main

import "fmt"

func fn(ch <-chan int) {
	fmt.Println(<-ch)

	ch <- 4 // receive only channelga ma'lumot yozmib bo'lmaydi!!!
}
func main() {
	ch := make(chan int)

	go fn(ch)
	ch <- 2
}
/*
out:
.go:8:5: invalid operation: ch <- 4 (send to receive-only type <-chan int)

*/
```
* send onl channel
bu channle orqali faqat ma'lumot jo'nata olamiz(ma'lumot o'qib bo'lmaydi):
```
package main

import "fmt"

func fn(ch chan<- int) {
	fmt.Println(<-ch) //ma'lumot o'qib bo'lmaydi!
	ch <- 7           // ma'lumotjo'natdik
}
func main() {
	ch := make(chan int)

	go fn(ch)
	fmt.Println(<-ch)
}
/*
out:
.go:6:14: invalid operation: <-ch (receive from send-only type chan<- int)

*/
```

#### 5 waitgroup

concurrent yuz berayotgan hodisalarni kutib turish maqsadida *sync* paketidan foydalanamiz
- bir vaqta ishlab turgan funksiyalrni kutib turish uchun waitGroup dan foydalanamiz
```
package main

import (
	"fmt"
	"sync"
)

func fn1(wg *sync.WaitGroup) {
	defer wg.Done()		// 3-1=2
	fmt.Println("fn1")
}
func fn2(wg *sync.WaitGroup) {
	defer wg.Done()		// 1-1=0
	fmt.Println("fn2")
}
func fn3(wg *sync.WaitGroup) {
	defer wg.Done()		// 2-1=1
	fmt.Println("fn3")
}
func main() {
	var wg sync.WaitGroup
	wg.Add(3)		//+3

	go fn1(&wg)
	go fn2(&wg)
	go fn3(&wg)

	wg.Wait()
}
```
#### 6 data_race

- *shared variable* - barcha funksiyalar ishlatishi mumkin bo'lgan o'zgaruchi(deb tushunsak bo'ladi)
- **data_race** - *shared variable* bilan ishlayotganda bir vaqitning o'zida, ham read, ham write operatsiyaning bo'lishi data race ga olib keladi

- data race - concurrent programmningnda ihlatiladigan eng og'riqli nuqta hisoblanadi
chunki concurrent programming da debug qilish qiyinroq
- data rece bo'layotgan yoki bolmayotganini aniqlash uchun go tool chain ning ichida instrument bor:
```
go run -race main.go
``` 

```
package main

import "fmt"

func GetNumber() {
	var number int
	func() {
		number = 7
	}()
	fmt.Println(number)
}
func main() {
	GetNumber()
}
/*
go run -race 1_data_race.go 
out:
7
*/
```
yuqoridagi misolda data race yo'q, chunki ammalar ketma ketlikda bajarilyapti.

data race:
```
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
out:
0
==================
WARNING: DATA RACE
Write at 0x00c0000200c8 by goroutine 7:
  main.GetNumber.func1()
      /home/ubuntu/Desktop/JasurbekUz/The-Go-Programming-Language/02_concurrency/6_sync.Mutex-data_race/1_data_race.go:8 +0x30

Previous read at 0x00c0000200c8 by main goroutine:
  main.GetNumber()
      /home/ubuntu/Desktop/JasurbekUz/The-Go-Programming-Language/02_concurrency/6_sync.Mutex-data_race/1_data_race.go:10 +0xba
  main.main()
      /home/ubuntu/Desktop/JasurbekUz/The-Go-Programming-Language/02_concurrency/6_sync.Mutex-data_race/1_data_race.go:13 +0x24

Goroutine 7 (running) created at:
  main.GetNumber()
      /home/ubuntu/Desktop/JasurbekUz/The-Go-Programming-Language/02_concurrency/6_sync.Mutex-data_race/1_data_race.go:7 +0xb0
  main.main()
      /home/ubuntu/Desktop/JasurbekUz/The-Go-Programming-Language/02_concurrency/6_sync.Mutex-data_race/1_data_race.go:13 +0x24
==================
Found 1 data race(s)
exit status 66
```
yuqorudagi dasturimizda **channel**dan foydalansak bo'lar edi, kanalham kutib turish vazifasini bajaradi(blocklash orqali) lekin kutib turiladigan narsada channel ishlatish noto'gri, channel ning vazifasi goroutinelar aro ma'lumot lamashish

bundan to'g'ririg'i **waitgroup** bo'ladi, waitgroup hamma goroutinelar tuggalanishini kutib turadigan narsa! (belgrilgan tasklarizni) 

lekin aynan datarace lar uchun "sync" packagi ichida alohida type bor. bu **mutex**lar deyiladi!

* **sync.Mutex**