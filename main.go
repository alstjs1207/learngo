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
	//Gorutines
	c := make(chan string) //main과 Goroutine 사이의 통신을 위한 channel 만들기
	people := [5]string{"ms", "hana", "star", "nico", "sun"}

	for _, person := range people {
		go learn.GoroutineTest(person, c) //go를 작성하여 Goroutines 만들기
	}
	for index, _ := range people {
		fmt.Println("waiting for ", index)
		fmt.Println("name is ", <-c) //생성된 message 만큼 만들기, 값을 전달 받을 때까지 main 함수는 정지
	}
}
