package main

import (
	"fmt"
	"sync"
)

func fn1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("fn1")
}
func fn2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("fn2")
}
func fn3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("fn3")
}
func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go fn1(&wg)
	go fn2(&wg)
	go fn3(&wg)

	wg.Wait()
}
