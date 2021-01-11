package main

import "fmt"

func hello(a int, b int) int {
	return a + b
}

//리턴명 지정시
func world(a int, b int) (r int) {
	r = a + b
	return //return r 로 하지 않는다
}

//리턴값 여러개 지정 - 리턴값 자료형만
func SumAndDiff1(a int, b int) (int, int) {
	return a + b, a - b
}

//리턴값 여러개 지정 - 리턴명 까지 지정
func SumAndDiff2(a int, b int) (s int, d int) {
	s = a + b
	d = a - b
	return
}

func returnmulti() (int, int, int, int, int) {
	return 1, 2, 3, 4, 5
}

//가변인자 - (매개변수명 ...자료형)리턴값 자료형
func sum(e ...int) int {
	total := 0
	for _, value := range e {
		total += value
	}

	return total
}

//팩토리얼 재귀함수
func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

//요소 두개를 합하고 빼는 함수 작성
func sum2(a int, b int) int {
	return a + b
}
func diff(a int, b int) int {
	return a - b
}

func main() {

	r := hello(1, 2)

	fmt.Println(r)

	fmt.Println("==리턴값 변수 지정시 - return 뒤에 변수를 지정하지 않음==")
	fmt.Println(world(5, 2))

	fmt.Println("==리턴값 변수 여러개 지정시 - 자료형만==")
	s, d := SumAndDiff1(4, 3)
	fmt.Println(s, d)

	fmt.Println("==리턴값 변수 여러개 지정시 - (자료형 리턴명) 지정 - return 뒤에 변수를 지정하지 않음==")
	s2, d2 := SumAndDiff2(4, 3)
	fmt.Println(s2, d2)

	fmt.Println("==리턴값 생략하고 싶을때 _==")
	a, b, _, c, _ := returnmulti()
	fmt.Println(a, b, c)

	fmt.Println("==가변인자 사용==")
	total := sum(1, 2, 3, 4, 5)
	fmt.Println(total)

	fmt.Println(("==가변인자는 슬라이스 타입이므로 ...을 사용하여 슬라이스로 바로 넘겨줄 수도 있다=="))
	n := []int{1, 2, 3, 4, 5} //슬라이스 선언
	total2 := sum(n...)
	fmt.Println(total2)

	fmt.Println("==재귀호출 - 팩토리얼 재귀함수 구현 예제==")
	fmt.Println(factorial(5))

	fmt.Println("==함수 변수에 저장 - var 변수명 func(매개변수명 자료형)리턴값_자료형 = 함수명 ==")
	var g func(e int, f int) int = hello //함수를 저장하는 변수를 선언하고 함수 대입
	f := hello                           //변수 선언과 동시에 함수 대입
	fmt.Println(g(1, 2))
	fmt.Println(f(1, 2))

	//원래 슬라이스 구조 - a:= []int{1,2,3,4,5}
	fmt.Println("==함수 슬라이스 배열에 저장 - 슬라이스 := []func(매개변수명의 자료형1, 매개변수명의 자료형2)리턴값_자료형{함수명1, 함수명2}==")

	//배열의 첫번째 요소로 함수 호출
	slice1 := []func(int, int) int{sum2, diff}
	fmt.Println(slice1[0](1, 2))
	fmt.Println(slice1[1](1, 2))

	//원래 맵의 구조 - a := map[string]int{"Hello":10, "World":20}
	fmt.Println("==함수 맵에 저장 - 맵 := map[키 자료형]func(매개변수명 자료형)리턴값_자료형{'키':함수명}==")

	//맵에 sum2키를 지정하여 함수 호출
	map1 := map[string]func(int, int) int{"Hello": sum2}
	fmt.Println(map1["Hello"](1, 2))
	//맵에 diff키를 지정하여 함수 호출
	map2 := map[string]func(int, int) int{"World": diff}
	fmt.Println(map2["World"](1, 2))

	map3 := map[string]func(int, int) int{
		"sum":  sum2,
		"diff": diff,
	}
	fmt.Println(map3["sum"](1, 2))
	fmt.Println(map3["diff"](1, 2))

	fmt.Println("==익명함수 - func(매개변수명 자료형) 리턴값 자료형{}() ==")
	func() { //함수에 이름이 없음 - 익명함수
		fmt.Println("Hello, world!")
	}()

	func(s string) {
		fmt.Println(s)
	}("익명함수를 정의한 뒤 바로 호출")

	noname := func(a int, b int) int { //변수 r에 저장
		return a + b
	}(1, 2) //바로 호출
	fmt.Println(noname)

}
