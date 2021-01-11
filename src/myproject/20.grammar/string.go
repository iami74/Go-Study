package main

import (
	"fmt"
	"unicode/utf8"
)

var s1 string = "Hello, world\n"
var s2 string = "안녕하세요\n"
var s3 string = "\ud55c\uae00"             //유니코드 문자코드로 저장( \u )
var s4 string = "\U0000d55c\U0000ae00"     //유니코드 문자코드로 저장( \U: 16진수 8자리로 맞춰줘야 한다)
var s5 string = "\xed\x95\x9c\xea\xb8\x80" //UTF-8 인코딩의 바이트 값으로 저장 ( \x )

//여러 줄로 된 문자열 ``(백쿼트)
var s7 string = `안녕하세요
Hello World!`

var s8 string = "한글"
var s9 string = "Hello"

func main() {
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)

	fmt.Println("==여러 줄로 된 문자열 ``(백쿼트)==")
	fmt.Println(s7)

	fmt.Println("==문자열 길이 구하기 len ==")
	fmt.Println(len(s8)) //UTF-8 인코딩의 바이트 길이 : 6
	fmt.Println(len(s9))

	fmt.Println("==UTF-8로 저장했을때 2바이트가 넘는 문자열의 실제 길이구함 - 'unicode/utf8'추가 == ")
	fmt.Println(utf8.RuneCountInString(s8))

	fmt.Println("==문자열 비교 (==) ==")
	fmt.Println(s4 == s8)

	fmt.Println(s1[0])        //72 : 아스키코드
	fmt.Printf("%c\n", s1[1]) //Printf

	s1 = "다시 대입은 가능"
	fmt.Println(s1)

	fmt.Println("==문자는 배열의 형태이므로 s[0] 이처럼 가져올 수는 있지만==")
	//s1[0] = 'z'
	fmt.Println("\"" + s1 + "\" 하나 수정은 불가합니다.")

}
