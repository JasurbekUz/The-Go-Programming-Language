package main

import (
	"fmt"
)

const (
	A int = iota + 5
	B
	_
	D
	E
)

const (
	F int = iota * 2
	G
	H
	_
	J
	K
)

const (
	FF int = iota * 2
	GG
	HH
	II = iota
	JJ
	KK
)

/* ... */

func main() {

	fmt.Println(A, B, D, E)         // 5, 6, 8, 9
	fmt.Println(F, G, H, J, K)      // 0 2 4 8 10
	fmt.Println(FF, GG, HH, JJ, KK) // 0 2 4 4 5

}
