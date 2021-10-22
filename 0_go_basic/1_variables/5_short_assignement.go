// sana: 2021 10 19
// dasturchi: Jasurbek Shamsitdinov
// maqsad: go da ozgaruvchilarni e'lon qilish va ma'lumot yozishni o'rganish
// ':=' - short-hand operatori

package main

import (
	"fmt"
)

func main() {

	// short-hand operatorda type lar faqat inplesid ishlaydi
	// barilgan qiymatga qarab typni bilib oladi
	// n := 20 <==> var n = 20

	number := 20
	age, color := 56, 254

	fmt.Println(number, age, color)
}
