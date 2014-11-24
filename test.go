package main

import (
  "fmt"
  // "io"
  "io/ioutil"
  "os"
)


func check(e error){
  if e!=nil{
    panic(e)
  }
}
func main(){

  args := os.Args[1:]
  fmt.Println(args[0])

  file :=args[0]

  dat, err := ioutil.ReadFile(file)
  check(err)
  fmt.Print(string(dat))

  f,err := os.Open(file)
  check(err)

  b1 := make([]byte, 5)
  n1, err := f.Read(b1)
  check(err)
  fmt.Printf("%d bytes: %s\n", n1, string(b1))
}
