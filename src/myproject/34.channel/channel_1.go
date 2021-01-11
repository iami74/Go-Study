package main

import (
	"fmt"
	"time"
)

/*
	//동기 채널 생성, 채널명 : done, 채널 자료형 : bool
	//반복 횟수 변수 저장, 변수명 : count

	//고루틴 익명함수 생성


	//for문 작성 count까지 도는(고루틴에 true를 보냄. 값을 꺼낼때까지 대기
	//반복문의 변수 출력
	//1초 대기

	//for문 작성
	//값을 꺼냄, 채널에 값이 들어올 때까지 대기
	//반복문의 변수 출력
*/

func main() {
	//동기 채널 생성 - 채널명 : done, 채널 자료형 : bool
	done := make(chan bool)
	//반복 횟수 변수 저장 - 변수명 : count
	count := 5

	//고루틴 익명함수 생성
	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("고루틴:", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		<-done
		fmt.Println("메인함수:", i)
	}
}
