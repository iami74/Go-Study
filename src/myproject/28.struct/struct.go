package main

import "fmt"

//type 구조체명 struct{}
type Rectangle struct {
	width, height int //같은 자료형일때 한 줄로 나열 가능
}

//구조체는 new함수로 구조체의 메모리를 할당하는 동시에 값을 초기화하는 방법은 없다
//구조체 생성자 패턴 - 구조체의 값을 초기화한 뒤 포인터를 리턴하는 방식
func NewRectangle(a int, b int) *Rectangle {
	return &Rectangle{a, b}
}

//함수에 구조체 포인터가 아닌 구조체 인스턴스로 넘겨주면 값이 모두 복사되므로 주의한다
func rectangleArea(rect *Rectangle) int {
	return rect.width * rect.height
}

func rectangleScalA(rectA *Rectangle, f int) {
	rectA.width = rectA.width * f
	rectA.height = rectA.height * f
}

func rectangleScalB(rectB Rectangle, f int) {
	rectB.width = rectB.width * f
}

//클래스가 없는 대신 구조체에 메서드를 연결할 수 있다
//func (리시버명 *구조체_타입) 함수명() 리턴값_자료형{}
func (rect *Rectangle) area() int {
	return rect.height * rect.width
}

//리시버 변수를 받는 방식1 - 포인터 방식
func (rect *Rectangle) scaleA(factor int) {
	rect.width = rect.width * factor
}

//리시버 변수를 받는 방식1 - 일반구조체 방식
func (rect Rectangle) scaleB(factor int) {
	rect.width = rect.width * factor
}

func (_ Rectangle) scaleX() {
	fmt.Println("_밑줄문자")
}

//구조체 임베딩 - 상속의 효과
type Person struct {
	name string
	age  string
}

type Student struct {
	p      Person //구조체가 해당 타입을 가지고 있는 관계
	school string
	grade  int
}

type Student2 struct {
	Person //필드명 없이 구조체만 지정하면 - 구조체가 해당 타입을 포함하는 관계가 된다
	phone  string
	email  string
}

func (ps *Person) greeting() {
	fmt.Println("person 안녕하세요")
}

//메서드 오버라이딩 - 자식 구조체가 부모 구조체의 메소드를 오버라이드 하게 된다
func (st *Student2) greeting() {
	fmt.Println("student 안녕하세요")
}

func main() {
	//구조체 인스턴스는 일반 변수를 선언하는 방법과 동일
	//var rect Rectangle	//Rectangle 타입으로 인스턴스 생성

	var rect Rectangle = Rectangle{10, 20} //구조체 생성하면서 초기화
	fmt.Println(rect.width)

	fmt.Println("==구조체 포인터==")
	var rect1 *Rectangle = new(Rectangle) //구조체 포인터 선언
	rect2 := new(Rectangle)

	rect1.height = 10

	fmt.Println(rect1)
	fmt.Println(rect2)

	fmt.Println("==구조체 인스턴스 {} ==")
	rect3 := Rectangle{width: 30} //필드명을 지정할때는 필드 개수를 채우지 않아도 된다
	rect4 := Rectangle{40, 50}

	fmt.Println(rect3.width)
	fmt.Println(rect4.width)

	var rect5 Rectangle
	rect5.height = 100
	fmt.Println(rect5.height)

	fmt.Println("==구조체 생성자 패턴 활용 - 구조체는 new함수로 구조체의 메모리를 할당하는 동시에 값을 초기화하는 방법은 없다==")
	rect6 := NewRectangle(10, 20)
	fmt.Println(rect6)

	rect7 := &Rectangle{10, 20} //rect6을 줄이면 - 구조체를 초기화한 뒤 메모리 주소를 대입
	fmt.Println(rect7)

	fmt.Println("==함수에서 구조체 매개변수 사용==")
	rect8 := Rectangle{30, 30}

	area := rectangleArea(&rect8) //구조체의 포인터를 넘김
	fmt.Println(area)

	fmt.Println("==※주의 : 함수의 매개변수에 구조체 포인터가 아닌 구조체 인스턴스로 넘겨주면 값이 복사된다==")
	rect9 := Rectangle{10, 10}
	rectangleScalA(&rect9, 20) //구조체의 포인터를 넘겨야 한다
	fmt.Println(rect9)

	rect10 := Rectangle{10, 10}
	rectangleScalB(rect10, 20) //구조체의 인스턴스를 넘기면 값이 복사된다
	fmt.Println(rect10)

	fmt.Println("==클래스가 없는 대신 구조체에 메서드를 연결할 수 있다==")
	rect11 := Rectangle{10, 20}

	fmt.Println(rect11.area())

	fmt.Println("==리시버 변수를 받는 방식1 - 포인터 방식==")
	rect11.scaleA(9)
	fmt.Println(rect11)

	fmt.Println("==리시버 변수를 받는 방식2 - 일반구조체 방식==")
	rect12 := Rectangle{10, 20}
	rect12.scaleB(9)
	fmt.Println(rect12)

	fmt.Println("==리시버 변수를 사용하지 않을때==")
	rect12.scaleX()

	fmt.Println("==구조체 임베딩 - 상속과 같은 효과==")
	var s Student
	s.p.greeting()

	fmt.Println("==필드명 없이 구조체만 지정하면 - 구조체가 해당 타입을 포함하는 관계가 된다==")
	var s2 Student2
	s2.greeting()
	s2.Person.greeting()

	fmt.Println("==메서드 오버라이딩 - 자식 구조체가 부모 구조체의 메소드를 오버라이드 하게 된다==")
	s2.greeting()
}
