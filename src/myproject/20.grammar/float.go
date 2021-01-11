// test.go
package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {

	var a float64 = 10.0

	for i := 0; i < 10; i++ {
		a = a - 0.1
	}

	const epsilon = 1e-14
	fmt.Println(math.Abs(a-9.0) <= epsilon)

	//복소수
	var num1 complex64 = 1 + 2i
	var num2 complex64 = 4.2342 + 2.3527i

	var r1 float32 = real(num1)
	var r2 float32 = imag(num1)

	var r3 float32 = real(num2)
	var r4 float32 = imag(num2)

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)

	//byte - 16진수, 문자값으로 저
	var num3 byte = 10   //10진수 저장
	var num4 byte = 0x32 //16진수 저장
	var num5 byte = 'a'

	//문자저장 (문자열이나 큰따옴표로 저장할 수 없다)
	//var num6 byte = 'ab'
	//var num7 byte = "a"

	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(num5)
	//fmt.Println(num6)
	//fmt.Println(num7)

	//오버플로우,언더플로우
	fmt.Println("==오버플로우,언더플로우 - MaxUint8==")
	var num8 uint8 = math.MaxUint8
	//var num8 uint8 = 0 - 1		//오버플로우 에러

	var num9 int8 = math.MaxInt8
	var num10 int8 = math.MinInt8

	fmt.Println(num8)
	fmt.Println(num9)
	fmt.Println(num10)

	//변수의 크기저장
	fmt.Println("==변수의 크기저장 - Sizeof==")
	var num11 int8 = 1
	var num12 int64 = 1

	fmt.Println(unsafe.Sizeof(num11))
	fmt.Println(unsafe.Sizeof(num12))
}
