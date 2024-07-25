package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices in go")

	//creating slices
	var myFirstSlice = []string{"Teri","Ma", "Ki"}
	fmt.Printf("Checking values and types %s, %T\n", myFirstSlice, myFirstSlice)

	//adding values in existing slice using append which takes two params first is the slice in which we are adding data and second is the data
	myFirstSlice = append(myFirstSlice, "Chut")
	fmt.Printf("Checking values and types %s, %T\n", myFirstSlice, myFirstSlice)

	//now data range using append and also range are not inclusive
	// myFirstSlice = append(myFirstSlice[3:])
	myFirstSlice = append(myFirstSlice[1:3])
	fmt.Printf("Checking values and types %s, %T\n", myFirstSlice, myFirstSlice)

	//make

	newSlice := make([]int, 5)

	newSlice[0] = 69
	newSlice[1] = 78
	newSlice[2] = 8989
	newSlice[3] = 8778
	newSlice[4] = 7678
	fmt.Printf("Checking values and types %d, %T\n", newSlice, newSlice)
	newSlice = append(newSlice, 66996699)
	fmt.Println("If is sorted", sort.IntsAreSorted(newSlice))
	fmt.Printf("Checking values and types %d, %T\n", newSlice, newSlice)

	//simple sorting
	sort.Ints(newSlice)
	fmt.Println("If is sorted", sort.IntsAreSorted(newSlice))
	fmt.Printf("Checking values and types %d, %T\n", newSlice, newSlice)

	//removing data from slice
	var testSlice = []string{"Mohan", "Pratham", "Rajesh", "Lodu"}
	var index int = 2
	//now logic for removing data from slice
	testSlice = append(testSlice[:index], testSlice[index+1:]...)
	fmt.Println(testSlice)
}