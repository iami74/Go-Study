package main

import "fmt"

func main() {

	//슬라이스는 길이가 동적으로 늘어난다
	//var a []int 	//슬라이스 선언 - 배열처럼 선언만으로 0으로 초기화 되지 않음 make 함수를 써야 함

	fmt.Println("==슬라이스는 make함수 사용하여 공간을 할당해야 값을 넣을 수 있다 - make==")
	var b = make([]int, 5) //값은 모두 0 으로 초기화
	c := make([]int, 7)    //var, [], 자료형 생략하고 슬라이스 선언 가능

	d := []int{1, 2, 3, 4, 5} //슬라이스 생성하면서 초기화

	fmt.Println(b[2])
	fmt.Println(c[2])
	fmt.Println(d[2])

	fmt.Println("==슬라이스 길이,용량 크기 구하기 - len,cap==")
	//슬라이스는 배열을 내장하고 있다. 배열이 늘어났을때 사용할 공간 미리 할당 가능
	//용량 생략시 용량은 길이와 동일하게 설정됨
	var e = make([]int, 5, 10) //길이 5, 용량 10
	fmt.Println(len(e))        //길이 구하기 - len
	fmt.Println(cap(e))        //용량 구하기 - cap
	fmt.Println(e[4])          //최대길이까지만 인덱스 접근 가능
	//fmt.Println(e[5])          //런타임 에러

	fmt.Println("==슬라이기 값 추가 - append")
	d = append(d, 6, 7, 8)
	fmt.Println(d)

	fmt.Println("==슬라이스에 슬라이스 붙이기 e... - append")
	d = append(d, e...)
	fmt.Println(d)

	fmt.Println("==배열의 경우 - 값이 복사되지만==")
	f := [3]int{1, 2, 3}
	var g [3]int

	g = f
	g[0] = 9
	fmt.Println((f))
	fmt.Println((g))

	fmt.Println("==슬라이스는 레퍼런스 타입이라 값이 복사되지 않고 h 의 값도 바뀜==")
	h := []int{1, 2, 3}
	var i []int
	i = h
	i[0] = 9

	fmt.Println(i)
	fmt.Println(h)

	fmt.Println("==슬라이스 복사 - copy==")
	j := []int{1, 2, 3, 4, 5}
	k := make([]int, 3) //복사될 슬라이스에 make 함수로 공간을 할당, 빈 슬라이스에는 요소를 복사할 수 없다

	copy(k, j)
	k[0] = 9
	fmt.Println(j)
	fmt.Println(k)

	fmt.Println("==슬라이스 용량 - 요소가 늘어나면 Go런타임은 정해진 알고리즘에 의해 용량을 늘림==")
	l := []int{1, 2, 3, 4, 5}
	fmt.Println(len(l), cap(l))

	l = append(l, 6, 7, 8)
	fmt.Println(len(l), cap(l)) //3개만 추가했는데 용량은 10으로

	fmt.Println("==부분슬라이스[시작인덱스:끝인덱스] - 요소는 복사되지 않는다. 참조 형태이다==")
	m := []int{1, 2, 3, 4, 5}
	n := m[0:5] //끝 인덱스는 실제보다 +1

	fmt.Println(m)
	fmt.Println(n)

	fmt.Println("==부분슬라이스 값을 바꾸면 - 요소는 복사되지 않으므로 기존 슬라이스 내용도 바뀜==")
	n[2] = 9

	fmt.Println(m)
	fmt.Println(n)

	fmt.Println("==부분슬라이스 - 시작,끝인덱스 생략 가능==")
	o := l[:]
	p := l[:len(n)]

	fmt.Println(o)
	fmt.Println(p)

	fmt.Println("==부분슬라이스 - 배열에도 사용가능, 참조형태이므로 요소를 바꾸면 배열의 요소도 바뀜==")
	//배열선언
	q := [5]int{1, 2, 3, 4, 5}
	r := q[:]

	r[2] = 9
	fmt.Println(q)
	fmt.Println(r)

	fmt.Println("==부분슬라이스 - 슬라이스[시작:끝:용량], 용량도 지정 가능, 단 기존 슬라이스 용량범위 내에서==")
	s := q[0:3:5]

	fmt.Println(s)
	fmt.Println(len(s), cap(s))
}
