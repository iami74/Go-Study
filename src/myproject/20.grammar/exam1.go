package main

import "fmt"

func main() {

	var txt string

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			txt = "FizzBuzz"
		} else if i%3 == 0 && i%5 != 0 {
			txt = "Fizz"
		} else if i%5 == 0 && i%3 != 0 {
			txt = "Buzz"
		} else {
			txt = ""
		}

		fmt.Println(i, txt)
	}

	fmt.Println("==switch문으로==")
	for j := 1; j <= 100; j++ {
		switch {
		case j%3 == 0 && j%5 == 0:
			fmt.Println("FizzBuzz")
		case j%3 == 0:
			fmt.Println("Fizz")
		case j%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(j)
		}
	}
}
