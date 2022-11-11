package main

import (
	"fmt"
	"strconv"
)

func main() {
	/******************* Declarations and assignations ******************/
	var num1 int = 11
	fmt.Printf("Type of num1 %T, value = %+v\n", num1, num1)

	num2 := 12
	fmt.Printf("Type of num2 %T, value = %+v\n", num2, num2)

	var str1 string
	fmt.Printf("Type of str1 %T, value = %q\n", str1, str1)

	var float64_1 float64
	fmt.Printf("Type of float64_1 %T, value = %+v\n", float64_1, float64_1)

	var bool1 bool
	fmt.Printf("Type of bool1 %T, value = %+v\n", bool1, bool1)

	var arrayInt1 [3]int
	fmt.Printf("Type of arrayInt1 %T, value = %+v\n", arrayInt1, arrayInt1)

	arrayInt2 := []int{24, 62, 7, 17, 3}
	fmt.Printf("Type of arrayInt2 %T, value = %+v\n", arrayInt2, arrayInt2)

	/******************* Conversions **********************************/
	str2 := strconv.Itoa(num1)
	fmt.Printf("Type of str2 %T, value = %q\n", str2, str2)

	num3, _ := strconv.Atoi(str2)
	fmt.Printf("Type of num3 %T, value = %+v\n", num3, num3)

	/******************* Bitwise *************************************/
	fmt.Printf("%d\t%b", num1, num1)

}
