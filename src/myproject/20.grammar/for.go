package main

import "fmt"

func main() {

	k := 0
	for k < 5 {
		fmt.Println(k)
		k++ //변화식이 없으면 무한루프에 빠진다
	}

	fmt.Println("==break문==")
	m := 0
	for m < 5 {
		if m > 2 {
			break
		}
		fmt.Println(m)
		m++
	}

	fmt.Println("==break문에 레이블 사용==")

This:
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j == 3 {
				break This
			}
			fmt.Println(i, j)
		}
	}

	fmt.Println("==반복문에서 특정부분 이하는 실행하지 않고 넘어가려면 - continue==")
	for a := 1; a < 5; a++ {

		if a == 3 {
			continue
		}
		fmt.Println(a)

	}

	fmt.Println("==continue문에 레이블 사용==")

This2:
	for i2 := 0; i2 < 4; i2++ {
		for j2 := 0; j2 < 4; j2++ {
			if j2 == 2 {
				continue This2
			}
			fmt.Println(i2, j2)
		}
	}

	fmt.Println("==반복문에서 변수 여러개 사용시==")
	for c, d := 0, 0; c < 10; c, d = c+1, d+1 {
		fmt.Println(c, d)
	}

}
