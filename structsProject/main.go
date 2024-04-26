package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// Struct declaration
type person struct {
	//person struc "properties"
	firstName string
	lastName  string

	//Property of custom type contactinfo
	contact contactInfo
}

func main() {

	//Create new record (pers1) of type person
	//syntax: type{property1Val,property2Val} not recommended
	// pers1 := person{"name", "surname"}
	//syntax type{propetyname: propertyVal,propertyName2: propertyVal2} BETTER
	pers1 := person{
		firstName: "name",
		lastName:  "surname",
		contact: contactInfo{
			email:   "em@mail.com",
			zipCode: 12345},
	}

	pers1.print()

	pers1.updateFirstName("passedByValueName")
	pers1.print()
	pers1.updateFirstNamePoint("passedByReference")
	pers1.print()
	//fmt.Println("Person Struct: ", pers1)
	//Create a var of type person wth properties initialized with its zero value
	// var varPerson person
	// varPerson.firstName = "name"
	// varPerson.lastName = "surname"
}

func (p person) print() {

	fmt.Printf("\n%+v", p)
}

// p passed by value, the original person is not updated
func (p person) updateFirstName(newFristName string) {
	p.firstName = newFristName
}

// using pointer to a person p is passed by reference
func (p *person) updateFirstNamePoint(newFristName string) {
	p.firstName = newFristName
}
