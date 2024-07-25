package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//what learn => input taking, conversions, string operations
	fmt.Println("Welcome to simple calculator")
	fmt.Println("Please enter any number")
	reader := bufio.NewReader(os.Stdin)
	firstNum,_ := reader.ReadString('\n')
	secondNum,_ := reader.ReadString('\n')
	numFirst,err1 := strconv.ParseFloat(strings.TrimSpace((firstNum)), 64)
	numSecond, err2 := strconv.ParseFloat(strings.TrimSpace(secondNum), 64)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		if err2 != nil{
			println("Error 2", err2)
		} else{
			fmt.Println("Calculated Values =>", numFirst + numSecond)
		}
	}
}