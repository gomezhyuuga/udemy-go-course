package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}
type contactInfo struct {
	email   string
	zipCode string
}

func main() {
	fer := person{firstName: "Fernando", lastName: "Gomez", contactInfo: contactInfo{
		email:   "fer@example.com",
		zipCode: "012345",
	}}
	var p person
	// Alternative way:
	// fer := person{"Fernando", "Gomez"}

	fmt.Println(fer)
	fer.print()
	p.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
