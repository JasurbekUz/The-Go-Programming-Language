package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func main() {

	x := uuid.NewV4()
	fmt.Println(x)
}
