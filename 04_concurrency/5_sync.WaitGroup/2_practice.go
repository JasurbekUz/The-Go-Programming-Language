package main

import (
	"fmt"
	"sync"
)

type Student struct {
	firsName string
	lastName string
}

func (s Student) FullName() string {
	return fmt.Sprintf("%s %s", s.firsName, s.lastName)
}
func study(student Student, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s o'qishni boshladi\n", student.FullName())
	fmt.Printf("%s o'qishni tamomladi\n", student.FullName())
}
func main() {
	var wg sync.WaitGroup

	var students []Student = []Student{
		Student{"Jasurbek", "Shamsitdinov"},
		Student{"Zohid", "Hamidov"},
		Student{"Nuraziz", "Eshmurodov"},
		Student{"Ibrohid", "Muhammadaliyev"},
	}
	for _, student := range students {
		wg.Add(1)
		go study(student, &wg)
	}
	wg.Wait()
}
