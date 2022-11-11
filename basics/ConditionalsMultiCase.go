package main

import "fmt"

func main() {
	var examResult int
	fmt.Println("Introduce your exam result (0 - 10):")
	fmt.Scanln(&examResult)

	switch examResult {
	case 1, 2, 3, 4:
		fmt.Println("You have Not passed the exam")
	case 5, 6, 7, 8:
		fmt.Println("You have passed the exam")
	case 9, 10:
		fmt.Println("You have passed the exam with an excellent result")
	default:
		fmt.Println("This result is not accepted (0 - 10)")
	}

}
