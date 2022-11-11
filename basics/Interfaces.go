package main

import "fmt"

func main() {
	fmt.Println("*** interface examples ***")
	student1 := student{
		name:    "Tyrion",
		surname: "Lannister",
	}

	subscribed1 := subscribed{
		student: student{
			name:    "Jonh",
			surname: "Snow",
		},
		subject: "Go",
	}

	fmt.Println(student1)
	fmt.Println(subscribed1)
	classroom(student1)
	classroom(subscribed1)

}

type student struct {
	name, surname string
}

type subscribed struct {
	student
	subject string
}

type user interface {
	introduceStudent()
}

func (s subscribed) introduceStudent() {
	fmt.Printf("My name is %+v %+v, and I am subscribed to %+v\n", s.name, s.surname, s.subject)
}

func (s student) introduceStudent() {
	fmt.Printf("My name is %+v %+v\n", s.name, s.surname)
}

func classroom(u user) {
	fmt.Println("I am in the classroom", u)
}
