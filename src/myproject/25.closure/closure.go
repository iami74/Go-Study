package main

import "fmt"

//클로저는 함수 안에서 함수를 선언 및 정의할 수 있다. 바깥쪽 함수에 선언된 변수에도 접근 할 수 있는 함수
//리턴값이 정수형 매개변수, 정수형 리턴값 형태의 익명함수를 리턴하는 calc2함수 작성
func calc2() func(x int) int {
	a, b := 3, 5

	return func(x int) int {
		return a*x + b
	}
}

func main() {

	a, b := 3, 5

	calc := func(x int) int {
		return a*x + b
	}(1) //바로 호출

	fmt.Println(calc)

	calc1 := func(x int) int {
		return a*x + b
	}

	r := calc1(1)
	fmt.Println(r)

	r2 := calc2()
	fmt.Println(r2(1))
}
