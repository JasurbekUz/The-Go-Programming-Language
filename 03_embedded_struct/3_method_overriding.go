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
