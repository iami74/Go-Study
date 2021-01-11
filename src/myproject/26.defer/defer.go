package main

import "fmt"

//지연 호출 : 특정 함수를 현재 함수가 끝나기 직전에 실행하는 기능(try finally)
func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}

func HelloWorld() {
	defer func() {
		fmt.Println("d1")
	}()

	func() {
		fmt.Println("d2")
	}()
}

func f() {
	defer func() {
		g := recover()

		fmt.Println(g)
	}()

	//패닉 함수로 에러 메세지 설정, 패닉 발생
	panic("ddd")
}

//패닉을 발생시키는 함수 작성
func p() {
	defer func() {
		s := recover() //recover 함수로 런타임 에러(패닉) 상황을 복구

		fmt.Println(s)
	}()

	//패닉 발생상황 만듦
	a := [...]int{1, 2}

	for i := 0; i < 3; i++ {
		fmt.Println(a[i])
	}

}

func main() {
	defer world() //defer를 사용한 함수는 main()함수가 끝나기 직전에 호출됨
	hello()

	HelloWorld()

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	fmt.Println("==recover함수 - 변수 := recover() - 패닉이 발생했을때 프로그램을 바로 종료하지 않고 예외처리 할 수 있다==")
	f() //panic 함수로 패닉발생 시킴
	p() //프로그램 상황으로 패닉발생 시킴

	fmt.Println("==panic함수 - panic(에러메세지) - 로직에 따라 에러로 처리하고 싶을때==")
	panic("Error")
	fmt.Println("panic 다음은 실행되지 않음")

}
