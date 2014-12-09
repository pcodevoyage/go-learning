package main

import (
  "fmt"
  "os"
  "log"
  "bufio"
)


func check(e error){
  if e!=nil{
    panic(e)
  }
}
func main(){

  args := os.Args[1:]

  file :=args[0]


  f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0600)
  if err != nil {
    log.Fatal(err)
  }

  defer f.Close()

  w := bufio.NewWriter(f)
  i,err :=w.WriteString("\n3 newtask written")
  if err != nil {
    log.Fatal(err)
  }
  w.Flush()

  fmt.Println(i)
}
