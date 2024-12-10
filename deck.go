package main

import "fmt"

type deck []string

func (d deck) Println() {
	for i, student := range d {
		fmt.Println(i, student)
	}
}
