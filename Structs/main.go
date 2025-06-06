//NOTE: we can pass by value with slices, maps, channels, functions, interfaces, and pointers
//Use pointers to int, float64, string, struct, array, slice, map, channel, function, interface

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
	//Wont work because GO uses pass by value
	// Hameed.firstNameChange("SK") // This will not change the first name in the original struct
	// To change the first name, we need to use a pointer receiver
	//when ever we pass a value by function go creates a copy of the value
	// So, if we want to change the value of the original struct, we need to pass a pointer to the struct
	//we put & and some variable name we assign it to a variable it will give address of the variable where it is pointing to
	// HameedPointer := &Hameed            // Create a pointer to the struct
	Hameed.firstNameChange("SK") // This will change the first name in the original struct
	Hameed.print()
}

// in case of *person when * is infront of type is type description - it means we are dealing with a pointer to a person struct

func (pointerToPerson *person) firstNameChange(newFirstName string) {
	// * gives the value of this memory adress pointing at
	//* is operater here it means we want to manipulate the value at the memory address
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)

}
