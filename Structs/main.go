package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// Hameed:=person{"SHAIK","ABDULHAMEED"} //less readability
	Hameed := person{firstName: "SHAIK", lastName: "ABDULHAMEED", contactInfo: contactInfo{email: "a@m.com", zipCode: 32456}}
	// var Hameed person
	// fmt.Println(Hameed)
	// Hameed.firstName = "SHAIK"
	// Hameed.lastName = "ABDULHAMEED"
	// fmt.Println(Hameed)
	// fmt.Println(Hameed.firstName)
	// fmt.Println(Hameed.lastName)
	Hameed.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)

}
