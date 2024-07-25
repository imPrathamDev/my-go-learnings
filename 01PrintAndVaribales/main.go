package main

import "fmt"

func main() {
	var myVar string = "Hi My Love!" // string
	var myNum int64 = 3576766867878
	var myFloat float32 = 788.6678787887878778787878787878
	var myAnother int
	var testVar string
	var testFloat float64
	welcome := "Hello World"

	// fmt.Println(welcome, myVar, myNum, myFloat)

	//checking value and types
	fmt.Printf("Check type %T \n", myFloat)
	// fmt.Printf("Get Values %s \n Get Another %d %T Hi again %f\n default num %d\n default string %s\n", myVar, myNum, myFloat, myFloat, myAnother, testVar)
	fmt.Printf("Get Value and Type %s, %T\n", myVar, myVar)
	fmt.Printf("Get Value and Type %d, %T\n", myNum, myNum)
	fmt.Printf("Get Value and Tpe %f, %T\n", myFloat, myFloat)
	fmt.Printf("Get value and Type %s, %T", welcome, welcome)

	// checking default value
	fmt.Printf("Get Default Value and Type %d, %T\n", myAnother, myAnother)
	fmt.Printf("Get Default Value and Type %s, %T\n", testVar, testVar)
	fmt.Printf("Get Default Value and Type %f, %T\n", testFloat, testFloat)

	//multiple aasigning as same type
	var num1, num2, num3 int = 232,4545,45646
	fmt.Println(num1, num2, num3)

	//multiple aasigning as same with no type
	check1, check2, check3 := 5656.57657, "Fuckk Off", 6877
	fmt.Println(check1, check2, check3)

check1 = 667.78
	fmt.Println(check1)

}