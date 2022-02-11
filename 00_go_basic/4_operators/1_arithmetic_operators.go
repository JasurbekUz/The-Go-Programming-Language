package main

import "fmt"

func main() {
	var (
		x, y float64 = 25, 5
	)

	fmt.Printf("x + y = %f\n", x+y)
	fmt.Printf("x - y = %f\n", x-y)
	fmt.Printf("x * y = %f\n", x*y)
	fmt.Printf("x / y = %f\n", x/y)
	fmt.Printf("x mod y = %d\n", int(x)%int(y))

	x++
	fmt.Printf("x++ = %f\n", x)

	y--
	fmt.Printf("y-- = %f\n", y)
}

/*
results:
x + y = 30.000000
x - y = 20.000000
x * y = 125.000000
x / y = 5.000000
x mod y = 0
x++ = 26.000000
y-- = 4.000000
*/
