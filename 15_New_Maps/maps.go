package main

import (
	"fmt"
	"maps"
)

func main() {
	//create map using make()
	// m := make(map[string]string)
	// m["name"] = "Pratham"
	// m["class"] = "BCA Part 2"

	//delete key from map (pass map variable and key name which to be delete)
	// delete(m, "class");

	// clear whole map
	// clear(m)

	// create map normal way
	m := map[string]int{"name": 6767, "age": 565476}
	// use to len as use for array and slices
	fmt.Println(len(m))

	//check if some key value or key exist or not
	// this _ is the value
	_, ok := m["hi"]

	if ok {
		fmt.Println("Hi dear")
	} else {
		fmt.Println("No hi dear")
	}

	m2 := map[string]int{"age": 565476, "name": 6767}

	//checking if two maps are equal or not
	fmt.Println(maps.Equal(m, m2))
}