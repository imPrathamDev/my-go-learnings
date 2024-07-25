package main

import "fmt"

func main() {
	var age int = 12
	//normal ifelse
	if age >= 18 {
		fmt.Println("Safe no pdf")
	} else if age < 18{
		fmt.Println("pdf")
	}

	if age%2 == 0 {
		fmt.Println("Age is even!")
	} else {
		fmt.Println("Age is odd!")
	}

	//intialioning variable
	if num := 69; num == 69{
		fmt.Println("Prev")
	}

	//switch case
	var name string = "moohan"

	switch name{
	case "pratham":
		fmt.Println("Fuck you pratham")
	case "mohan":
		fmt.Println("Fuck offff mohan")
	default:
		fmt.Println("Fuccck yyoou offf Default")
	}
}