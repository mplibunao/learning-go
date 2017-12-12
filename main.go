package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Neutron",
		contact: contactInfo{
			email:   "jimmyneutron@gmail.com",
			zipCode: 1243,
		},
	}
	fmt.Printf("%+v", jim)
}
