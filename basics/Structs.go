package main

import "fmt"

type item struct {
	name     string
	price    int
	quantity int
}

func main() {
	//e_struct1()
	//e_struct2()
	annonym_struct1()
}

func e_struct1() {
	item1 := item{
		name:     "Item1",
		price:    10,
		quantity: 100,
	}

	item2 := item{
		name:     "item2",
		price:    20,
		quantity: 200,
	}

	fmt.Println(item1)
	fmt.Println(item2)
}

func e_struct2() {
	type person struct {
		name string
		age  int
	}

	type engineer struct {
		person
		ableToCreate bool
	}

	engineer1 := engineer{
		person: person{
			name: "Sergi",
			age:  43,
		},
		ableToCreate: true,
	}

	fmt.Println(engineer1)
	fmt.Println(engineer1.name, engineer1.age)
}

func annonym_struct1() {
	p := struct {
		name, surname string
		age           int
	}{
		name:    "Maria",
		surname: "Perez",
		age:     25,
	}
	fmt.Println(p)
}
