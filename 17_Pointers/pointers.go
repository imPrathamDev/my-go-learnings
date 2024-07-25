package main

import "fmt"

type User struct {
	Age int
	Name string
	role string
}

func main() {
	//Pointers in go

	// i := 69
	// p := &i

	// *p = 96

	// fmt.Println(i)

	// struct in go
	// A struct is a collection of fields.
	// u := Test{69, "Pratham"}
	// u.Y = "Mohan"
	// fmt.Println(u.Y)

	//Pointer in struct
	user := User{52, "Mohan", "admin"}
	puser := &user
	puser.role = "manger"
	fmt.Println(user)
	fmt.Println(*puser)

}