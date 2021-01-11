package main

import (
	"fmt"
	"math/rand"
	"time"
)

//고루틴 : 함수를 동시에 실행시키는 기능 (스레드)
func hello(n int) {

	r := rand.Intn(100)          //1~100까지의 랜덤한 숫자 리턴
	time.Sleep(time.Duration(r)) //time.Sleep : 설정한 시간동안 대기
	//time.Duration : 나노초 단위 (1/1000000000)

	fmt.Println(n)
}

func main() {
	for i := 0; i < 100; i++ {
		go hello(i)
	}
	fmt.Scanln() //메인 함수가 종료되지 않도록 대기
}
