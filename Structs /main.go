package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

//every person has 1 contactInfo :)
type person struct {
	firstName string
	lastName  string
	//we are not limited to just types,
	//we can have custom types and define them in other types
	contactInfo
}

func main() {
	name := "Bill"
	fmt.Println(&name)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	//go automatically understands that jim is a person, and update name is pointing to a person
	//so they work together anyways.
	jim.updateName("Jimmy")
	jim.print()
}

//need a pointer because just because you reference the name in memory does not mean you changed it
//point to the person
func (pointerToPerson *person) updateName(newFirstName string) {
	//point to the person's first name value and update that name at jims address
	(*pointerToPerson).firstName = newFirstName
}

//you can use this function to any person (reciever)
//in the body of this function, 'p' represents a person

func (p person) print() {
	fmt.Printf("%+v", p)
}
