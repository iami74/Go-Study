package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("==case에서 break 키워드 생략==")
	fmt.Println("==if 조건문 안에서 break 키워드 사용==")
	s := "Hello"
	i := 10

	switch i {
	case 10:
		if s == "Hello" {
			fmt.Println("Hello 2")
			break
		}
		fmt.Println(i)
	case 20:
		fmt.Println(i)
	default:
		fmt.Println(-1)
	}

	fmt.Println("==case문장을 실행한 뒤 다음 case의 문장을 실행하고 싶을때 - fallthrough ==")
	j := 3
	switch j {
	case 3:
		fmt.Println(j)
		fallthrough
	case 4:
		fmt.Println("fallthrough 조건에 맞지 않지만 그냥 실행-대박")
	case 5:
		fmt.Println(j)
	default:
		fmt.Println(-1) //마지막 행은 fallthrough 사용할 수 없음. 당연한거겠지
	}

	fmt.Println("==여러 조건 함께 처리하기 - ,(콤마)로 값을 구분==")
	h := 3
	switch h {
	case 1, 3, 5:
		fmt.Println("홀수")
	case 2, 4, 6:
		fmt.Println("짝수")
	}

	fmt.Println("==switch문에 조건식도 사용할 수 있다==")
	k := 3
	switch { //변수를 지정하지 않고 조건식 사용
	case k >= 5 && k < 10:
		fmt.Println("5이상, 10미만")
	case k >= 0 && k < 5:
		fmt.Println("0이상, 5미만")
	}

	fmt.Println("==switch 분기문 안에서 함수실행하고 결과값 분기할 수 있다==")
	rand.Seed(time.Now().UnixNano())
	switch h := rand.Intn(10); { //함수호출 뒤 ; 붙여준다
	case h > 1 && h < 6:
		fmt.Println("1이상, 6미만")
	case h == 9:
		fmt.Println("9")
	default:
		fmt.Println(h)
	}

	fmt.Println(h)

}
