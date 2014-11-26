package main

import (
  "fmt"
  "io"
  // "io/ioutil"
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


  f, err := os.Open(file)
  if err != nil {
    log.Fatal(err)
  }

  bf := bufio.NewReader(f)
  for{
    line, isPrefix,err :=bf.ReadLine()

    if err == io.EOF{
      break
    }

    if err !=nil {
      log.Fatal(err)
    }

    if isPrefix{
      log.Fatal("Error : Unexpected line ")
    }

    fmt.Println(string(line))
  }
}
