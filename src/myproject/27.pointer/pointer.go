package main

import "fmt"

func hello(num2 *int) { //포인터 변수로 받아
	*num2 = 2 //num2를 역참조하여ㅕ 메모리에 2를 대입
}

func main() {
	var numPtr *int //포인터형 변수를 선언하면 nil(NULL)로 초기화
	//*를 자료형 앞에 붙인다
	fmt.Println(numPtr)

	//new함수로 메모리를 할당해야 사용할 수 있다
	//new함수는 지정한 자료형 크기에 맞는 메모리 공간을 할당
	//메모리를 관리해주는 가비지 컬렉션을 지원하므로 메모리 할당 뒤 해제하지 않아도 된다
	var numPtr2 *int = new(int)
	fmt.Println(numPtr2)

	*numPtr2 = 1
	fmt.Println(*numPtr2)

	//일반 변수에 &변수명 - 참조(레퍼런스)를 사용하면 포인터형 변수에 대입할 수 있다
	//&변수명 : 변수의 메모리 주소를 뜻
	fmt.Println("==일반변수 포인터 변수에 대입==")

	var num int = 1
	var numPtr3 *int = &num //참조로 num변수의 메모리 주소를 구하여 대입

	fmt.Println(&num)
	fmt.Println(numPtr3)

	fmt.Println("==함수에서 포인터 매개변수 사용==")
	var num2 int = 1

	hello(&num2) //num2의 메모리 주소를 hello함수에 넘
	fmt.Println(num2)

}
