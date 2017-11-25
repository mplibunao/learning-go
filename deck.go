package main

import "fmt"

// Create a new type of `deck` which is a slice of strings
type deck []string

// Receiver
// d - instance of deck type using the function
// like this or self
// deck - type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
