package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to methods in golang")
	//fist create Struct
	user := User{"imPrathamDev", "pratham@gmail.com", 69}
	fmt.Println(user)
	fmt.Printf("All data with key and value: %+v\n", user)
	user.GetUsernameWithId()
}

type User struct{
	Username string
	Email string
	Age int
}

func (us User) GetUsernameWithId(){
	fmt.Printf("Check %v_%d",us.Username, time.Now().UnixNano())
}