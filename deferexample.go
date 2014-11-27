// Defer statements are pushed on a list. This list of saved
// calls is called only after all the surrounding calls have returned.

package main

import "fmt"

func main() {
  a()
  b()
  x := c()
  fmt.Println("From C:",x)
}

//Defer rules

//1. A deferred function's arguments are evaluated
//when the defer statement is evaluated.

func a() {
  i := 0
  defer fmt.Println(i)
  i++
  return
}

// 2. Deferred functions are executed in Las In First Out order
// after the surrounding functions have returned
func b(){
  fmt.Println("Starting from i=0")
  for i:=0 ; i<4 ; i++ {
    defer fmt.Println("From b : ", i)
  }
}

// 3.Deferred functions may read and assign to the returning
// function's named return values.

func c() (i int) {
  defer func() { i++ }()
  return 1
}
