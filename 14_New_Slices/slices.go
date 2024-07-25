package main

import "fmt"

func main() {
	var numbers = make([]int, 5);
	numbers[0] = 68
	numbers[3] = 69
	numbers = append(numbers, 4)
	fmt.Println("Slice => ", numbers[:3])
	fmt.Println("Length => ", len(numbers))
	fmt.Println("Cap => ", cap(numbers))
}