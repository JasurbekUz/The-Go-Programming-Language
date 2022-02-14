package main

import (
	"fmt"
	"time"
)

type IHuman interface {
	getName() string
	getAge() int
}

type Pupil struct {
	name      string
	BirthYear int
}

type Student struct {
	name      string
	BirthYear int
}

func (p Pupil) getName() string {
	return p.name
}

func (p Pupil) getAge() int {
	return time.Now().Year() - p.BirthYear
}

func (s Student) getName() string {
	return s.name
}

func (s Student) getAge() int {
	return time.Now().Year() - s.BirthYear
}

func (p Pupil) getAll() (string, int) {
	return p.name, time.Now().Year() - p.BirthYear
}
func main() {
	var pupil IHuman = Pupil{"Jasur", 1999}
	var student Student = Student{"Jasur", 1999}
	fmt.Println(pupil.getAge())
	fmt.Println(student)
	//fmt.Println(pupil.getAll())
}
