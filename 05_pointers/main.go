package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in go")

	//declearing/creating an int type pointer and it's initial value will be <nil>
	var myPointer *int
	fmt.Println("This is myPointer's value", myPointer)

	//
	myNum := 69
	var newPointer = &myNum
	//pointer value will be the memory address of referring variable
	//we use & for referring to an variable so, & === referance
	fmt.Println("newPointer value", newPointer)
	myNum = 96
	//for getting the memory location just simpily use the pointer name
	//for getting the value of refferecing variable use * 
	fmt.Println("newNum value using pointer", *newPointer)

	//for assiiaging new value do like below
	*newPointer = *newPointer + 4
	fmt.Println("value of myNum", myNum)

}