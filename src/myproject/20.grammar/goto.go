package main

import "fmt"

func main() {

	var a int = 1

	if a == 1 {
		fmt.Println("Error,1")
	}

	if a == 2 {
		fmt.Println("Error,2")

	}

	if a == 3 {
		fmt.Println("Error,3")

	}

	fmt.Println("==goto 키워드와 레이블 사용==")
	if a == 1 {
		goto Error1
		fmt.Println("goto 바로 뒤는 코드 실행되지 않는다")
	}

	if a == 2 {
		goto Error2
	}

	if a == 3 {
		fmt.Println("Error,3")
		return
	}

Error1:
	fmt.Println("레이블 Error,1")
	return

Error2:
	fmt.Println("레이블 Error,2")
	return

}
