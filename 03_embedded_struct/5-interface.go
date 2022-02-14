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
