package main

import "fmt"

func main() {
	var loopType int
	fmt.Println("Choose a kind of 'for' loop:\n 1 - 'while'\n 2 - 'for'\n 3 - 'doWhile'\n 4 - 'nested loops'\n 5 - 'for range'")
	fmt.Scanln(&loopType)
	i := 0
	switch loopType {
	case 1:
		fmt.Println("for loop as while")
		for i < 5 {
			fmt.Println("This value is", i)
			i++
		}
	case 2:
		fmt.Println("for loop as for")
		for i := 0; i < 5; i++ {
			fmt.Println("This value is", i)
		}
	case 3:
		fmt.Println("for loop as doWhile - printing the even values from 1 to 10")
		for {
			if i <= 10 {
				i++
				if i%2 != 0 {
					continue
				}
				fmt.Println("This value is", i)
			} else {
				break
			}
		}
	case 4:
		fmt.Println("nested for loops")
		// Draw a square with character '*' 10 x 10
		for i := 1; i <= 10; i++ {
			for j := 1; j <= 10; j++ {
				fmt.Print("* ")
			}
			fmt.Print("\n")
		}
	case 5:
		fmt.Println("for range - printing the total sum of the array items")
		anualResult := []int{23452, 72622, 5433, 9273, 62533, 62353, 15625, 25162, 41312, 62762, 62532, 52426}
		var total int
		for _, value := range anualResult {
			total = total + value
		}
		fmt.Println("Total is :", total)
	default:
		fmt.Println("This value is not accepted (1 - 3)")
	}

}
