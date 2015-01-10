package main

import (
  "fmt"
  "io"
  // "io/ioutil"
  "os"
  "log"
  "bufio"
  "strings"
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

  defer f.Close()

  bf := bufio.NewReader(f)
  var entries map[string]string
  entries = make(map[string]string)
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

    s := strings.Split(string(line),"|")
    fmt.Println(s[0])
    entries[s[0]]= s[1]

    // fmt.Println(string(line))
  }

  for key, value := range entries{
    fmt.Println("Key :"+key + " - Value:"+value)
  }


  tmp := entries["1"]
  entries["1"] = entries["2"]
  entries["2"] = tmp

  for key, value := range entries{
    fmt.Println("Key :"+key + " - Value:"+value)
  }
}
