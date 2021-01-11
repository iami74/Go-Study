package main

import "fmt"

//오리 구조체 정의
type Duck struct {
}

//오리의 quack 메서드 정의
func (d Duck) quack() {
	fmt.Println("꽥!")
}

//오리의 feathers 메서드 정의
func (f Duck) feathers() {
	fmt.Println("오리는 흰새과 회색 털을 가지고 있습니다.")
}

//사람 Person 구조체 정의
type Person struct {
}

//사람의 quack 메서드 정의
func (p Person) quack() {
	fmt.Println("사람은 오리를 흉내냅니다.")
}

//사람의 feathers 메서드 정의
func (f Person) feathers() {
	fmt.Println("사람은 땅에서 깃털을 주워서 보여줍니")
}

//quack, feathers 메서드를 가지는 Quacker 인터페이스 정의
type Quaker interface {
	quack()
	feathers()
}

func fncQuacker(qi Quaker) {
	qi.quack()    //Quacker 인터페이스로 quack 메서드 호출
	qi.feathers() //Quacker 인터페이스로 feathers 메서드 호출
}

func main() {

	/*
		//인터페이스 선언
		var qi Quaker

		//구조체 인스턴스 생성
		var d Duck
		var p Person

		//인터페이스를 통하여 오리의 quack, feathers 호출
		qi = d
		qi.quack()
		qi.feathers()

		//인터페이스를 통하여 사람 quack, feathers 호출
		qi = p
		qi.quack()
		qi.feathers()
	*/

	//구조체 인스턴스 생성
	var d Duck
	var p Person

	//인터페이스를 통하여 오리의 quack, feathers 호출
	fncQuacker(d)

	//인터페이스를 통하여 사람 quack, feathers 호출
	fncQuacker(p)

	fmt.Println("==타입이 특정 인터페이스를 구현하는지 검사할때 - interface{}(인스턴스).(인터페이스)==")
	//인스턴스를 빈 인터페이스에 넣은 뒤 인터페이스와 같은지 확인
	/*
		if b, err = ioutil.ReadFile("d:/error.txt"); err == nil {
			fmt.Printf("%s", b)
		} else {
			fmt.Println(err)
		}

		//================================================

		fmt.Println("==맵 데이터 조회 - value, ok := 맵[]==")
		value, ok := f["TuesDay"]

		fmt.Println(value, ok)

		fmt.Println("==if value, ok := 맵[] ; ok==")
		if value, ok := f["TuesDay"]; ok {
			fmt.Println(value)
		}
	*/

	//첫번째 리턴값 : 검사했던 인스턴스
	//두번째 리턴값 : 인스턴스가 해당 인터페이스를 구현하고 있는지 여부
	//Duck 타입의 인스턴스 d를 빈 인터페이스에 넣은 뒤 Quaker 인터페이스와 같은지 확인
	if value, ok := interface{}(d).(Quaker); ok {
		fmt.Println(value, ok)
	}

}
