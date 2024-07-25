package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in go")

	//let's create our first array
	var myArray [4]string

	myArray[0] = "Milo"
	myArray[2] = "Pratham"
	myArray[3] = "Test Name"

	fmt.Println("Checkout the array", myArray)
	//length of array
	//length of array will same which we said or decleared in begning
	//no matter how many items in the array it will the initialized length
	fmt.Println("Checkout the length ", len(myArray))

	//we initialed and assign values at the same time
	var newArray = [2]string{"Test", "India"}
	fmt.Printf("Checkout the type: %T\n", newArray)
	fmt.Println("Checkout the array", newArray)
	fmt.Println("Checkout the length ", len(newArray))

}