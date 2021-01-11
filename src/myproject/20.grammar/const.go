package main

import "fmt"

//상수는 코드에서 고정된 값을 체계적으로 관리하고 싶을때
//상수는 반드시 선언과 동시에 초기화
//선언한 뒤에는 값을 변경할 수 없다

const age int = 10

/*==컴파일 에러==
const score int //반드시 선언과 동시에 초기화
const age
age = 20
name = "가리봉"
*/

const (
	name, score     = "가리봉", 50
	x, y        int = 20, 30
)

/*
const (
	Sunday =0
	Monday = 1
	TuesDay  = 2
	Wednesday =3
	Thursday = 4
	Friday = 5
	Saturday = 6
	numberOfDays = 7
	)

위와 같이 일일이 값을 대입하지 않고
순서대로 생성하려면 - iota

*/

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	_ //특정 iota 를 생각하려
	_
	numberOfdays
)

const (
	a = 1 << iota // 1 << 0
	b = 1 << iota // 1 << 1

	c = iota * 30 // 2 * 30
	d = iota * 60 // 3 * 60
)

func main() {

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(score)
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println("==iota==")
	fmt.Println(Sunday)
	fmt.Println(numberOfdays)

	fmt.Println("==iota 이용해 특정순서로 조==")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
