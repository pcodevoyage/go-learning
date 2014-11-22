//to run : go run hello.go
package main

import (
  "fmt"
  "math/cmplx"
)

//func add( x int, y int) int {
func add(x,y int) int {
  return x + y
}

//return more than one variables
func swap (x,y string) (string,string) {
  return y,x
}


//named results
func split(sum int)(x,y int) {
  x = sum % 2
  y = sum / 2
  return
}

//variables outside the func has to have a var no short assignment
var i,j,k int

//declare multiple variables
var (
 t bool = false
 mi uint64 = 1 <<64 -1
 z complex128 = cmplx.Sqrt(-5 + 12i)
)

//constants: cannot use := to declare constants
const Pi = 3.14

func main() {
  fmt.Println("Hello, World\n")

  fmt.Println(add(333,2333))

  a,b := swap("First","Second")
  fmt.Println(a,b)

  fmt.Println(split(500))

  fmt.Println(i,j,k)

  //inside the function we can do short assignment
  m, name, isCool := 4, "Joe", true
  fmt.Println(m,name,isCool)

  //formatted output
  const f = "%T(%v)\n"
  fmt.Printf(f, t, t)
  fmt.Printf(f, mi, mi)
  fmt.Printf(f, z, z)

  //can do for ; sum < 1000 ;
  sum := 0
  for i:=0 ; i< 10 ; i++ {
    sum += i
  }
  fmt.Println("Sum is :",sum)

  //for is Go's while
  sum2 := 1
  for sum2 < 1000 {
    sum2 += sum2
  }
  fmt.Println("Generated by while",sum2)

  if sum2 == 1024 {
    fmt.Println("1024 Same as always")
  }


  if s := sum2%2; s==0 {
    fmt.Println("compute then check condition")
  }
}
