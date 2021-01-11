package main

import "fmt"

func main() {

	var a [5]int //0으로 초기화
	var b [5]int = [5]int{32, 29, 78, 16, 81}

	var c = [5]int{32, 29, 78, 16, 81}
	d := [5]int{32, 29, 78, 16, 81}

	e := [...]int{1, 2, 3}

	f := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	a[2] = 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e[2])
	fmt.Println(f[1][1])

	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}

	fmt.Println("==배열의 길이를 구하지 않고 간단하게 배열의 요소를 가져오려면 - range==")
	for j, value := range b {
		fmt.Println(j, value)
	}

	fmt.Println("==range 키워드 (인덱스 생략하고, 값 변수만 출력) _ 사용==")
	for _, value := range b {
		fmt.Println(value)
	}

	fmt.Println("==배열을 다른 변수에 대입하면 첫번째 요소가 아닌 전체 복사==")
	a = b
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

}
