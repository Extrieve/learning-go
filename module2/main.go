package main

import (
	"fmt"
)

func main() {
	var participantNames = [10]string{"Nick", "Mat", "Dan"} // Array definition

	fmt.Printf("%v\n", participantNames)
	participantNames[4] = "John"
	fmt.Printf("%v\n", participantNames)
	fmt.Printf("This is an array of size: %v\n", len(participantNames)) // always size N

	var sliceDefinition = []string{"Nick", "Mat", "Dan"} // Slice definition
	fmt.Printf("%v\n", sliceDefinition)
	fmt.Printf("This is a slice of size: %v\n", len(sliceDefinition)) // size 3 in this case
	sliceDefinition = append(sliceDefinition, "Kennedy")
	fmt.Printf("%v\n", sliceDefinition)
}
