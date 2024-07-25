package main

import "fmt"
 
func main() {

  var var1 complex64 = complex(3,-5)  
  var var2 complex64 = complex(2, 7.5)
  fmt.Println("var1:", var1)
  fmt.Printf("Type of var1: %T", var1)
  fmt.Println("\nvar2:", var2)
  fmt.Printf("Type of var2: %T", var2)  
  
}