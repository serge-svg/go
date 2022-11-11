package main

import (
	"fmt"
)

func main() {
	//e_arrays1()
	//e_arrays2()
	//e_arrays3()
	//e_arrays4()
	//e_arrays5()
	//e_arrays6()
	//e_arrays7()
	//e_array8()
	//e_maps1()
	//e_maps2()
}

// Change the value at position 1 by 'cogombre'
func e_arrays1() {
	array17 := []string{"pastanagues", "pebrots", "aigua", "formatge"}
	array17[1] = "cogombre"
	fmt.Println(array17)
}

// Print data indexes and values of the array
func e_arrays2() {
	array18 := []string{"pastanagues", "pebrots", "aigua", "formatge"}
	fmt.Println(array18)
	for index, value := range array18 {
		fmt.Println(index, " ", value)
	}
}

// Print result of the first quarter
func e_arrays3() {
	anualResult := []int{23452, 72622, 5433, 9273, 62533, 62353, 15625, 25162, 41312, 62762, 62532, 52426}
	firstQuarter := anualResult[0:3]
	fmt.Println("First quarter values are :", firstQuarter)
	var total int
	for _, value := range firstQuarter {
		total = total + value
	}
	fmt.Println("First quarter result is :", total)
}

// Add an element to the list
func e_arrays4() {
	shoppingList := []string{"pastanagues", "pebrots", "aigua", "formatge"}
	shoppingList = append(shoppingList, "water")
	fmt.Println("items in the list are :", shoppingList)
}

// Join two arrys in one
func e_arrays5() {
	workers := []string{"Josep", "Cristina", "Helena", "Robert"}
	newWorkers := []string{"Ivan", "Estel", "Aitana"}
	workers = append(workers, newWorkers...)
	fmt.Println(workers)
}

// Delete Robert and Ivan
func e_arrays6() {
	workers := []string{"Josep", "Cristina", "Helena", "Robert"}
	newWorkers := []string{"Ivan", "Estel", "Aitana"}
	workers = append(workers, newWorkers...)
	workers = append(workers[:3], workers[5:]...)
	fmt.Println(workers)
}

func e_arrays7() {
	x := make([]int, 9, 10)
	x[2] = 3
	x[8] = 7
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	//x[11] = 11 // error
	x = append(x, 11)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) // duplicate space memory
}

func e_array8() {
	list1 := []string{"Samarreta", "Lacoste", "Talla L", "Blava"}
	fmt.Println(list1)
	list2 := []string{"Polo", "Decathlon", "Talla XL", "Verd"}
	fmt.Println(list2)
	shoppingCart := [][]string{list1, list2}
	fmt.Println(shoppingCart)
}

func e_maps1() {
	m := map[string]int{"Josep": 6, "Joana": 8, "Enric": 4} // Last element also needs ','
	fmt.Println(m)
	fmt.Println(m["Josep"])
}

func e_maps2() {
	m := map[string]int{"Josep": 6, "Joana": 8, "Enric": 4}
	fmt.Println(m)
	fmt.Println(m["Josep"])

	v, ok := m["Emilia"]
	fmt.Println(v, ok) // 0 false
	if v, ok := m["Joana"]; ok {
		fmt.Println("Has been deleted the value", v, ok)
		delete(m, "Joana")
	} else {
		fmt.Println("The item 'Joana' doen't exist", v, ok)
	}
}
