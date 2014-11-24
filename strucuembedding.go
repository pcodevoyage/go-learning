package main

import "fmt"

type A struct {}

func (A) Hello(){
  fmt.Println("Hello from A")
}


type B struct{
  A // Struct embedding B now has all functions of A without knowing about it
}
// func (b B) Hello() { b.A.Hello() } // (implicitly!)

// If we have following Hello of B will be called.
// func (B) Hello(){
//   fmt.Println("Hello from B")
// }

func main(){
  var b B
  b.Hello() // print
}
