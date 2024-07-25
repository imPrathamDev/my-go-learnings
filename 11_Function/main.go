package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	//you can't create function inside a function
	result := addTwoNum(69,31)
	fmt.Println(result)

	//test for me
	myNums := []int{69783, 3887, 3467678, 36746}

	newResult, err := addManyNum(myNums...)
	fmt.Println(newResult, err)
}

func addTwoNum(numOne int, numTwo int)int{
	return numOne + numTwo
}

func addManyNum(nums...int)(int, string){
	total := 0
	for i := 0; i < len(nums); i++{
		total += nums[i]
	}
	// return total
	//you also return two values
	return total, "Oops Error"
}