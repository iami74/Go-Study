package main

import "fmt"

var b1 bool = true
var b2 bool = true
var b3 bool = false
var b4 bool = false

var num1 int = 3
var num2 int = 10

func main() {

	fmt.Println(b1 && b2) //true
	fmt.Println(b1 && b3) //false
	fmt.Println(b3 && b4) //false(*)
	fmt.Println(b1 || b2) //true
	fmt.Println(b1 || b3) //true
	fmt.Println(b3 || b4) //false(*)
	fmt.Println(!b1)
	fmt.Println(!b3)

	fmt.Println("==비교연산자==")
	fmt.Println(num1 > num2)
	fmt.Println(num1 != num2)

}
