package main

import (
	"fmt"
	_ "time" //사용하지 않는 패키지로 인한 컴파일 에러 방지
)

func main() {

	var x, y int
	var name string

	//자료형을 생략할때는 반드시 초기값을 대입
	var z = 50

	//병렬 할당 - 여러개의 값을 콤마로 구분하여 대입하는 방법
	x, y, name = 30, 50, "Maria"

	/* 아래와 같이도 사용 가능
	   var (
	   	x, y int = 30,50
	   	age,name = 10, "Maria"
	   )
	*/

	//사용하지 않는 변수로 인한 컴파일 에러 방지
	_ = z
	_ = y
	_ = name

	var num1 int = 0723
	fmt.Println("Hello, World!")

	fmt.Println(x)
	fmt.Println(num1)

}
