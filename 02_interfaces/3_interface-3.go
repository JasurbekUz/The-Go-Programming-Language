package main

import "fmt"

type Calc interface {
	Add(a, b int) int
}
type A struct{}

func (a A) Add(x, y int) int {
	return x + y
}
func kalk(c Calc, x int) {
	fmt.Println(c.Add(x, x))
}
func main() {
	var x Calc = A{}
	kalk(x, 4)
}
