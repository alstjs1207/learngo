package main

import (
	"fmt"

	"github.com/learngo/learn"
)

func main() {
	fmt.Print("Hello!!")

	//#1.
	nameLen1, nameUpper1 := learn.LenAndUpper("yulotts")
	fmt.Println(nameLen1, nameUpper1)
	//#2~3.
	nameLen, nameUpper := learn.LenAndUpper2("yulotts")
	fmt.Println(nameLen, nameUpper)
	//#4~5.
	total := learn.SuperAdd(1, 2, 3, 4, 5, 6)
	fmt.Println("total cnt :", total)
	//#6.
	fmt.Println(learn.CanIDrink(16))
	//#7.
	fmt.Println(learn.WhatSwitch(11))
	//#8.
	learn.Pointers()
	//#9.
	learn.WhatArray()
	//#10.
	learn.WhatMaps()
	//#11.
	learn.WhatStructs()
}
