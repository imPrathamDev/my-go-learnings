package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs in go")
	//Structs are basically replacement of class in go yes, you heard right there is no class in go
	note := Notes{69, "First Note", "Here is content"}
	fmt.Println(note)
	//ok for getting values
	fmt.Printf("Notes data: %+v\n", note)
	//getting single data
	fmt.Printf("Note Title is '%v', Note Content is '%v'", note.Title, note.Content)
}

//create structs
type Notes struct{
	id int
	Title string
	Content string
}