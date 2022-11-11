package main

import "fmt"

func main() {
	fmt.Println("*** simple switch case contidional examples ***")
	// Control with the conditional 'if' upon the result of an exam, if it passed or not.
	var examResult int
	fmt.Println("Introduce your exam result (0 - 10):")
	fmt.Scanln(&examResult)

	switch {
	case examResult < 0 || examResult > 10:
		fmt.Println("This value is not accepted (0 - 10)")
	case examResult < 5:
		fmt.Println("You have Not passed the exam")
	case examResult > 6 && examResult < 8:
		fmt.Println("You have passed the exam")
	case examResult > 8 && examResult <= 10:
		fallthrough
	default:
		fmt.Println("You have passed the exam with an excellent result")
	}

}
