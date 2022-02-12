package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Horse struct {
	name string
}

func (Horse) getSpeed() time.Duration {
	rand.Seed(time.Now().UnixNano())

	speed := rand.Intn(10000)

	return time.Duration(speed) * time.Microsecond
}

func run(horse Horse, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100000 - horse.getSpeed())
	fmt.Println(horse.name + " is finished!")
}
func main() {
	var wg sync.WaitGroup
	var horses []Horse = []Horse{
		Horse{"Boychibor"},
		Horse{"Gulsari"},
		Horse{"Qorabayir"},
		Horse{"Jiyron Qush"},
	}
	for _, horse := range horses {
		wg.Add(1)
		go run(horse, &wg)
	}
	wg.Wait()
}
