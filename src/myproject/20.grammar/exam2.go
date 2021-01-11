package main

import "fmt"

func main() {

	for i := 99; i >= 0; i-- {

		bottle := "bottles"

		switch {
		case i == 1:
			bottle = "bottle"
			fmt.Println(i, bottle, " of beer on the walls,", i, bottle, " of beer.")
			fmt.Println("Take on down, pass it aroud,", "No more bottles of beer on the wall")
		case i == 0:
			fmt.Println("No more bottles of beer on the walls, No more bottles of beer.")
			fmt.Println("Go to the stor and buy some more,", i+99, "bottles of beer on the wall")
		default:
			fmt.Println(i, bottle, " of beer on the walls,", i, bottle, " of beer.")
			fmt.Println("Take on down, pass it aroud,", i-1, "bottles of beer on the wall")
		}

	}

	fmt.Println("==Printf 로 구현==")
	for j := 99; j >= 0; j-- {

		bot := "bottles"
		switch {
		case j == 1:
			bot = "bottle"
			fmt.Printf("%d %s of beer on the wall, %d %s of beer.\n", j, bot, j, bot)
			fmt.Println("Take on down, pass it aroud,", "No more bottles of beer on the wall")
		case j == 0:
			fmt.Println("No more bottles of beer on the wall, No more bottles of beer.")
			fmt.Println("Go to the stor and buy some more,", j+99, "bottles of beer on the wall")
		default:
			fmt.Printf("%d %s of beer on the wall,%d %s of beer.\n", j, bot, j, bot)
			fmt.Println("Take on down, pass it aroud,", j-1, "bottles of beer on the wall")
		}

	}

}
