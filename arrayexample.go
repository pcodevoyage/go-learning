package main

import "fmt"

func main(){
  var a[2]string
  a[0] = "Hello"
  a[1] = "World"

  fmt.Println(a[0])
  fmt.Println(a)

  b := [2]string{"hello","there"}
  fmt.Println(b)
}
