package main

import "fmt"

//채널은 고루틴끼리 데이터를 주고받고, 실행 흐름을 제어하는 기능
//채널 자체는 값이 아닌 레퍼런스 타입
//make(chan 자료형)

func sum(a int, b int, c chan int) {
	//채널에 a와 b의 합을 보냄
	c <- a + b
}

func main() {
	/*
		var c chan int	//chan int형 변수 선언
		c = make(chan int)//할
	*/
	c := make(chan int) //int형 채널 생성

	//★==채널을 매겨변수로 받은 함수는 반드시 go 키워드를 사용, 고루틴으로 실행해야 한다==★
	go sum(1, 2, c) //sum을 고루틴으로 실행한 뒤 채널을 매개변수로 넘겨줌

	n := <-c //채널에서 값을 꺼낸 뒤 변수 n에 대입

	fmt.Println(n)

}
