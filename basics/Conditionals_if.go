package main

import "fmt"

func main() {
	fmt.Println("*** if/elsif/else contidional examples ***")

	// Control with the conditional 'if' upon the result of an exam, if it passed or not.
	var examResult int
	fmt.Println("Introduce your exam result (0 - 10):")
	fmt.Scanln(&examResult)
	if examResult < 0 || examResult > 10 {
		fmt.Println("This value is not accepted (0 - 10)")
	} else if examResult < 5 {
		fmt.Println("You have Not passed the exam")
	} else if examResult > 6 && examResult < 8 {
		fmt.Println("You have passed the exam")
	} else {
		fmt.Println("You have passed the exam with an excellent result")
	}

}
