package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loop in golang")

	names := []string{"Pratham", "Milo", "Mohan", "Ramesh"}

	//normal way for looping an slice or array
	// for i := 0; i < len(names); i++{
	// 	fmt.Println(names[i])
	// }

	//looping using range
	// for i:= range names{
	// 	fmt.Println(names[i])
	// }

	//forEach type looping
	// for index, name := range names{
	// 	fmt.Printf("index is %v and value is %v\n", index, name)
	// }

	//another way if you don't want index or value use _,ok syntex
	for _, name := range names{
		fmt.Printf("index is NA and value is %v\n", name)
	}

	//while type loop
	//we also have break, continue, goto
	var num int = 1
	for num < 5{
		if num ==4{
			goto test
		}
		fmt.Println(num)
		num++
	}

	test: fmt.Println("Line GoTo")
}