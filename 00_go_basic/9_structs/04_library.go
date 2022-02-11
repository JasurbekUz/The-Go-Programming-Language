// Dasturchi: Jasurbek Shamsitdinov
// Maqsad: Structlar va ularning metodlari ustida Praktika qilish
// Sana: 2021 10 21

package main

import (
	"fmt"
)

type Author struct {
	author_id  uint
	first_name string
	last_name  string
}

type Category struct {
	category_id uint
	title       string
}

type Book struct {
	book_id   uint
	book_name string
	category  Category
	author    Author
}

func (book Book) getDeatail() string {

	return book.category.title + ": " + book.author.first_name + " " + book.author.last_name + "ning " + book.book_name + " asari"

}

func main() {

	var book Book = Book{
		1,
		"Yulduzli Tunlar",
		Category{
			1,
			"Badiiy asarlar",
		},

		Author{
			1,
			"Prirmqul",
			"Qodirov",
		},
	}

	var book1 Book = Book{
		2,
		"Amir Temur",
		Category{
			1,
			"Badiiy asarlar",
		},

		Author{
			5,
			"Yunus",
			"Oguz",
		},
	}

	fmt.Println(book.getDeatail())
	fmt.Println(book1.getDeatail())
}
