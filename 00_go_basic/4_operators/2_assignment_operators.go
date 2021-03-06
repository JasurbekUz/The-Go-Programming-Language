package main

import "fmt"

func main() {

	x, y := 15, 25

	x = y
	fmt.Println("= ", x)

	x = 15
	x += y
	fmt.Println("+=", x)

	x = 50
	x -= y
	fmt.Println("-=", x)

	x = 2
	x *= y
	fmt.Println("*=", x)

	x = 100
	x /= y
	fmt.Println("/=", x)

	x = 40
	x %= y
	fmt.Println("%=", x)
}

/*
results:
=  25
+= 40
-= 25
*= 50
/= 4
%= 15
*/
