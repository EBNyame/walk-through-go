package main

import "fmt"

func main() {

	students := []string{
		"Kofi",
		"John",
		"Jane",
	}

	for i, student := range students {
		fmt.Println(i, student)
	}
	students = append(students, newCard())

	fmt.Println(students)

}

func newCard() string {
	return "this is a new card"
}
