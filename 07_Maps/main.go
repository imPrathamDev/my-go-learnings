package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Maps in go")

	//creating maps
	myMap := make(map[string]string)
	myMap["email"] = "pratham@gamil.com"
	myMap["name"] = "Pratham"
	fmt.Printf("Get type: %T", myMap)
	//Getting all data
	fmt.Println("Get values og map: ", myMap)
	//Getting signle data
	fmt.Println("Get email", myMap["email"])
	//deleting data
	delete(myMap, "name")
	fmt.Println("Get values og map: ", myMap)

}