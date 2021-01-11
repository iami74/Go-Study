package main

import (
	"fmt"
	"runtime"
)

//채널 버퍼링
//make(chan 자료형, 버퍼개수)

/*
	//CPU 개수를 한개로 설정
	//버퍼가 2개인 비동기 채널 생성, 채널명 : done, 채널 자료형 : bool
	//반복 횟수 변수 저장, 변수명 : count

	//고루틴 익명함수 생성

	//for문 작성 (고루틴에 true를 보냄. 버퍼가 가득차면 대기
	//반복문의 변수 출력

	//for문 작성
	//값을 꺼냄, 버퍼에 값이 없으면 대기
	//반복문의 변수 출력
*/

func main() {
	runtime.GOMAXPROCS(1)

	done := make(chan bool, 2)
	count := 4

	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("고루틴:", i)
		}
	}()

	for i := 0; i < count; i++ {
		<-done
		fmt.Println("메인함수:", i)
	}
}
