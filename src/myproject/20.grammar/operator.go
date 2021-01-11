package main

import "fmt"

func main() {

	a := 3
	b := 2
	//비트연산자 &
	c := a & b

	fmt.Println(c)
	fmt.Printf("%08b", c) //안됨

	d := 3
	e := -2

	f := +d
	g := +e
	h := -d
	i := -e

	fmt.Println("==단항 양수부호==")
	fmt.Println(f)
	fmt.Println(g)

	fmt.Println("==단항 음수부호==")
	fmt.Println(h)
	fmt.Println(i)

	j := [3]int{1, 2, 3}
	k := [3]int{4, 5, 6}

	l := []int{4, 5, 6}
	m := map[string]int{"Hello": 1}

	fmt.Println("==배열은 요소가 모두 같은지 비교==")
	fmt.Println(j != k)
	fmt.Println("==슬라이스를 nil과 비교하여 메모리 할당되었는지 비교==")
	fmt.Println(l == nil)
	fmt.Println("==맵을 nil과 비교하여 메모리 할당되었는지 비교==")
	fmt.Println(m == nil)

	n := 10
	n++

	o := 20
	//p := o++  //++ 연산자를 사용한 뒤 값을 대입할 수 없다

	fmt.Println("==변수 뒤에서만 사용")
	fmt.Println(n)
	fmt.Println(o)

	q := 3*2 + 4*1
	r := 3 - 5 + 4 - 7

	fmt.Println("==연산자 우선순위로 자동 정렬")
	fmt.Println(q)
	fmt.Println(r)
}
