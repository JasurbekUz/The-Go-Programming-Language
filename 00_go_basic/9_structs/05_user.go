// Dasturchi: Jasurbek Shamsitdionov
// Maqsad: structlar ustida amallar
// Sana: 2021 10 21

package main

import (
	"fmt"
	"time"
)

type Course struct {
	name string
}

type Group struct {
	duration int8
	course   Course
}

type User struct {
	first_name string
	last_name  string
	birth_day  string
	birth_year uint
	group      Group
}

func (u User) fullName() string {

	return u.first_name + " " + u.last_name
}

func (u User) age() uint {

	return uint(time.Now().Year()) - u.birth_year
}

func main() {

	var user = User{
		"John", "Doe", "'1999 10 7'", 1999,
		Group{
			80,
			Course{
				"Go",
			},
		},
	}

	fmt.Println(user.first_name)
	fmt.Println(user.last_name)
	fmt.Println(user.fullName())
	fmt.Println(user.age())
	fmt.Println(user.birth_day)
	fmt.Println(user.group.course.name)
	fmt.Println(user.group.duration)
}
