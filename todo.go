package main

import(
  "os"
  "bufio"
  "fmt"
)

func getFilenameFromConfig()(dbFile string){

  f,err := os.Open("config")
  if err!=nil{
    panic(err)
  }

  defer f.Close()
  bf := bufio.NewReader(f)
  line, isPrefix,err :=bf.ReadLine()
  if err!=nil || isPrefix{
    panic(err)
  }

  dbFile =string(line)
  return
}


func main(){

  dbFile := getFilenameFromConfig()

  fmt.Println(dbFile)

}
