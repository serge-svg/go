package main

import "fmt"

func main() {
	// e1 - funtion with string as parameter
	fmt.Println(greeting("Serge"), ", how is it going")

	// e2 - funtion with array of integers as parameter
	listOfNumbers := []int{5, 2, 1, 7}
	fmt.Printf("The total value of the items in listOfNumbers is : %+v\n", sum(listOfNumbers))

	// e3 - function of type method
	student1 := subscribed{
		student: student{
			name:    "Jonh",
			surname: "Snow",
		},
		subject: "Go",
	}
	fmt.Println(student1)
	student1.introduceStudent()

	// e4 - defer funciton, bye is executed after hi
	defer bye()
	hi("user")

}

func greeting(name string) string {
	return fmt.Sprint("Hi ", name)
}

func sum(values []int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func hi(who string) {
	fmt.Printf("Hi %+v\n", who)
}
func bye() {
	fmt.Println("Bye user")
}

type student struct {
	name, surname string
}

type subscribed struct {
	student
	subject string
}

func (s subscribed) introduceStudent() {
	fmt.Printf("My name is %+v %+v, and I am subscribed to %+v\n", s.name, s.surname, s.subject)
}
