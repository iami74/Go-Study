package main

import "fmt"

func main() {

	//맵(해시테이블,딕셔너리) map[키_자료형]값 자료형 형태, 레퍼런스 형태

	//var a map[string]int //키는 string 자료형,값은 int자료형

	fmt.Println("==맵은 make 함수를 사용하여 공간을 할당해야 값을 넣을 수 있습니다 - make==")

	var b map[string]int = make(map[string]int)
	c := make(map[string]int)

	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("==맵을 생성하면서 초기화==")
	d := map[int]string{1: "Hello", 2: "world"}

	e := map[string]string{
		"a": "a",
		"b": "b", //여러 줄로 나열할때는 마지막 요소에 콤마
	}
	fmt.Println(d[1])
	fmt.Println(d[2])
	fmt.Println(e["a"])
	fmt.Println(e["b"])

	fmt.Println("==맵에 데이터 저장==")
	f := make(map[string]float32)
	f["Monday"] = 10.1
	f["TuesDay"] = 20.1

	fmt.Println(f["Mondy"])

	fmt.Println("==맵에 데이터 저장 - 존재하지 않는 키를 조회시==")
	g := make(map[int]string)
	g[1] = "October"

	fmt.Println(g[2])           //빈 문자열("") 출력
	fmt.Println(f["WednesDay"]) //float32 자료형은 0 출력

	fmt.Println("==맵 데이터 조회 - value, ok := 맵[]==")
	value, ok := f["TuesDay"]

	fmt.Println(value, ok)

	fmt.Println("==if value, ok := 맵[] ; ok==")
	if value, ok := f["TuesDay"]; ok {
		fmt.Println(value)
	}

	fmt.Println("==for반복문 - for 키,값 :=range 맵 - range==")
	for key, value := range f {

		fmt.Println(key, value)
	}

	fmt.Println("==맵 데이터 삭제 - delete(맵, 삭제할 키)==")
	delete(f, "Monday")
	fmt.Println(f)

	fmt.Println("==다시 저장==")
	f["Monday"] = 50.1
	fmt.Println(f)

	fmt.Println("==맵 안에 맵 만들기==")

	h := map[string]map[string]int{
		"월": map[string]int{
			"January":  1,
			"February": 2,
			"March":    3,
			"April":    4,
			"May":      5,
		},
		"주": map[string]int{
			"Monday":    1,
			"TuesDay":   2,
			"WendesDay": 3,
			"ThursDay":  4,
			"FriDay":    5,
		},
	}

	fmt.Println(h["월"]["January"])

	i := map[string]map[string]int{
		"hello": map[string]int{
			"world": 10,
		},
	}
	fmt.Println(i["hello"]["world"])
}
