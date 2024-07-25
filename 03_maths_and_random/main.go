package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome Golang")
	// var myNum int = 45
	// var myFloat float64 = 6767.7878
	// fmt.Println("Sum of nums", float64(myNum) + myFloat)

	//random number using math/rand
	// randNum := rand.Intn(10)
	// fmt.Println("Random Value", randNum)

	//random number using crypto/rand
	randomNum, _ := rand.Int(rand.Reader, big.NewInt(78787))
	fmt.Println(randomNum)
}