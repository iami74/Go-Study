package main

import (
	"fmt"
	"strconv"
)

type Any interface{} //인터페이스에 메서드를 지정하지 않음

func f1(arg Any) { //모든 타입을 저장할 수 있음
}

type Person struct {
	name string
	age  int
}

/*
작성해보시오
func formatString(arg interface{}) string{
	switch

	case //arg가 int형이라면
	//int형으로 값을 가져옴
	//strconv.Itoa 함수를 사용하여 i를 문자열로 변환
	//arg가 float32형이라면
	//float32형으로 값을 가져옴
	//strconv.FormatFloat 함수를  사용하여 f를 문자열로 변
}
*/

func f2(arg interface{}) { //모든 타입을 저장할 수 있음
}

func formatString(arg interface{}) string {
	switch arg.(type) {
	//빈 인터페이스 변수는 그대로 사용할 수 없으므로 타입을 지정해서 가져온다 == Type assertion : arg.(int)
	case int:
		i := arg.(int)         //int형으로 값을 가져옴
		return strconv.Itoa(i) //strconv.Itoa 함수를 사용하여 i를 문자열로 변환

	case float32:
		f := arg.(float32)                                  //float32형으로 값을 가져옴
		return strconv.FormatFloat(float64(f), 'f', -1, 32) //strconv.FormatFloat 함수를  사용하여 f를 문자열로 변환

	case float64:
		f := arg.(float64)
		return strconv.FormatFloat(f, 'f', -1, 64)

	case string:
		s := arg.(string)
		return s

	case Person:
		p := arg.(Person)
		return p.name + " " + strconv.Itoa(p.age)

	case *Person:
		p := arg.(*Person)
		return p.name + " " + strconv.Itoa(p.age)

	default:
		return "Error"
	}
}

func main() {
	fmt.Println(formatString(1))
	fmt.Println(formatString(2.5))
	fmt.Println(formatString("hello"))

	//구조체 인스터스 및 포인터도 가능
	p := Person{"홍길동", 10}
	fmt.Println(formatString(p))
	fmt.Println(formatString(&p))

	//구조체 new 함수로 메모리 생성
	var p2 = new(Person)
	p2.name = "아무개"
	p2.age = 20

	fmt.Println(formatString(p2))

	fmt.Println("==인터페이스 저장된 타입이 특정 타입인지 검사 방법 - 인터페이스.(타입)==")
	//인터페이스 선언 t
	//인터페이스에 구조체 저장 : 구조체명 : Person , 값 : "Maria" ,20
	//if문 작성

	var t interface{}
	t = Person{"Maria", 20}

	if v, ok := t.(Person); ok {
		fmt.Println(v, ok)
	}

	fmt.Println("==빈 인터페이스는 모든 자료형을 저장할 수 있다. 값을 꺼낼때는 - Type assert를 사용한다==")
	//1) ifce1, ifce2 빈인터페이스 선언, 값 각각 저장(1, "hello")
	//2) 변수 num1, s1 에 값을 꺼내 저장한다
	//3) 출력한다
	//작성하시오

}
