package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to Files handling in golang")

	//for creating text file we use os
	file,err := os.Create("./myTextFile.txt")
	CheckNilErr(err)
	length,err := io.WriteString(file, "Hello Word!\nTesty\nFuck You")
	CheckNilErr(err)
	fmt.Println(length)
	defer file.Close()
	readFileContent("./myTextFile.txt")
}

func CheckNilErr(err error){
	if err != nil {
		panic(err)
	}
}

func readFileContent(filePath string){
	dataInBytes,err := ioutil.ReadFile(filePath)
	CheckNilErr(err)
	fmt.Println(string(dataInBytes))
}