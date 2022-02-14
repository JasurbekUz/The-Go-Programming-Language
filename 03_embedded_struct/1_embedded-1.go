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
