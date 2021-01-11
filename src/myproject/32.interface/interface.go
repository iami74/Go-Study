package main

import "fmt"

type hello interface {
}

//Print 메서드를 가지는 인터페이스 정의
type Printer interface {
	Print()
}

type MyInt int //int형을 MyInt로 정의

type Rectangle struct {
	width, height int
}

//MyInt에 Print 메서드 연결
func (s MyInt) Print() {
	fmt.Println(s)
}

func (r Rectangle) Print() {
	fmt.Println(r.width, r.height)
}

func main() {
	var h hello
	fmt.Println(h)

	//인터페이스 선언
	var p Printer

	//s를 인터페이스 p에 대입
	var s MyInt = 5
	p = s
	//인터페이스를 통하여 MyInt의 Print 메서드 호출
	p.Print()

	fmt.Println("==구조체 인터페이스 연결==")
	r := Rectangle{10, 20}

	p = r
	p.Print()

	fmt.Println("==인터페이스 초기화 - ()를 사용하여 변수나 인스턴스를 넣어준다==")
	//var p2 Printer = Printer(s)
	p2 := Printer(s)
	p2.Print()

	P3 := Printer(r)
	P3.Print()

	fmt.Println("==배열 다시 상기==")
	var a [5]int
	var b = [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3, 4, 5}

	fmt.Println(a[1])
	fmt.Println(b[1])

	fmt.Println("==배열의 길이를 구하지 않고 간단하게 배열의 요소를 가져오려면 - range==")
	for j, value := range c {
		fmt.Println(j, value)
	}

	fmt.Println("==슬라이스 다시 상기==")
	//슬라이스 선언 - 배열처럼 선언만으로 0으로 초기화 되지 않음 make 함수를 써야 함
	//var d []int

	//슬라이스 초기화
	f := make([]int, 4)
	e := []int{6, 7, 8, 9, 10}
	fmt.Println(f[1])
	fmt.Println(e[1])

	fmt.Println("==인터페이스를 슬라이스 형태로 초기화==")
	p4 := []Printer{s, r} //사용자정의 자료형(s), 구조체(r)

	for key, value := range p4 {
		fmt.Println(key, value)
	}

	for key, _ := range p4 {
		p4[key].Print()
	}

	for _, value := range p4 {
		value.Print()
	}
}
