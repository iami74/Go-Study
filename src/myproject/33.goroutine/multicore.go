package main

import (
	"fmt"
	"runtime"
)

//시스템의 모든 CPU코어를 사용하는 방버
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //GOMAXPROCS : 사용할 최대 CPU 개수 설정함수, NumCPU : 현재 시스템의 CPU 코어 갯수
	fmt.Println(runtime.GOMAXPROCS(0))   //GOMAXPROCS 에 0을 넣으면 설정값은 바뀌지 않으며 현재 설정 값만 리턴

	runtime.GOMAXPROCS(1) //CPU 하나만 사용

	s := "Hello, world"

	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println(s, n)
		}(i)
	}

	fmt.Println("==클로저를 고루틴으로 실행시 반복문에 의해 바뀌는 변수는 반드시 매개변수로 넘겨주어야 한다")
	//고루틴은 반복문이 완전히 긑난 다음에 생성된다
	//이렇게 쓰면 안됩니다
	for j := 0; j < 100; j++ {
		go func() {
			fmt.Println(s, j)
		}()
	}

	fmt.Scanln()
}
