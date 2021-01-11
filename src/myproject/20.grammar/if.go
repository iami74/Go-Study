package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	i := 10

	if i >= 10 {
		fmt.Println("10 이상")
	} else if i >= 5 && i < 10 { // ==> } else { 는 같은 줄에
		fmt.Println("5 이상 10 미만")
	}

	var b []byte
	var err error

	//ioutil.ReadFile 함수로 파일을 열면 b에는 파일의 내용, err 에는 에러값이 들어간다
	b, err = ioutil.ReadFile("c:/error.txt")

	if err == nil {
		fmt.Printf("%s", b)
	}

	fmt.Println("\n==if 조건문 안에서 함수 사용 - 함수실행 ; 조건식==")
	//if 조건문 안에서 변수를 생성했을때 if 문 안에서만 사용할 수 있음
	if b, err = ioutil.ReadFile("c:/error.txt"); err == nil {
		fmt.Printf("%s", b)
	} else {
		fmt.Println(err)
	}

	//변수 b를 사용할 수 없음.
	//fmt.Println(b)
}
