package learn

import (
	"fmt"
	"strings"
)

//Go 특징 #1. 하나의 입력값에 여러개의 값 return
func LenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

//Go 특징 #2. return할 variable를 굳이 명시하지 않아도 사용 가능
func LenAndUpper2(name string) (nameLen int, nameUpper string) {
	//Go 특징 #3. defer - function이 실행하고 return후에 실행됨
	defer fmt.Println("fuction done!!")
	nameLen = len(name)
	nameUpper = strings.ToUpper(name)
	return
}

//Go 특징 #4. for와 range를 이용한 loop와 ignore
//Go 특징 #5. ...args를 이용한 여러개의 variable를 받을 수 있다.
func SuperAdd(numbers ...int) int {
	fmt.Println(numbers)
	total := 0
	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}

	for _, number := range numbers {
		fmt.Println(number)
	}

	return total
}

//Go 특징 #6. if문 안에서 variable를 선언할 수 있다.
//			  elase는 생략할 수 있다.
func CanIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true

}

//Go 특징 #7. if ~ elseif 처럼 switch를 사용할 수 있다.
func WhatSwitch(dowork int) string {
	switch {
	case dowork < 12 || dowork > 9:
		return "working"
	case dowork > 12 || dowork < 13:
		return "break time"
	case dowork < 19 || dowork > 13:
		return "working"
	}
	return "Bye"
}

//Low Level Programming
//Go 특징 #8. 변수에 &를 붙이면 메모리주소를 볼수 있다.
//			  변수에 *를 붙이면 메모리주소를 살펴볼 수 있다.
func Pointers() {
	a := 2
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	// *b의 포인터를 변경하여 a를 변경할 수 있다.
	*b = 2021
	fmt.Println(a)
}

//Go 특징 #9. 길이를 선언 하면 array 길이를 선언하지 않으면 slice
func WhatArray() {
	names := [5]string{"yu", "jong", "joo"}
	names[3] = "min"
	names[4] = "ri"

	fmt.Println(names)
	//길이를 선언하지 않는건 slice 다.
	newNames := []string{"yu", "jong", "joo"}
	//append는 새로운 slice를 return하기 때문에 다시 담아줘야한다.
	newNames = append(newNames, "min")
	fmt.Println(newNames)
}

//Go 특징 #10. maps
func WhatMaps() {
	user := map[string]string{"ou": "departmentname", "oucode": "123", "displayname": "lotts"}
	fmt.Println(user)
	for key, value := range user {
		fmt.Println(key, value)
	}
}

//Go 특징 #11. Structs
type person struct {
	name    string
	age     int
	favFood []string
}

func WhatStructs() {
	favFood := []string{"candy", "rice"}
	//person에 어떤값을 넣어야하는지 알수 없다.
	//yulotts := person{"lotts", 10, favFood}
	yulotts := person{name: "lotts", age: 10, favFood: favFood}
	fmt.Println(yulotts.name, " / ", yulotts)

}
