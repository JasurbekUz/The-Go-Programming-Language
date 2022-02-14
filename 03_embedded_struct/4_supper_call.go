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
