package main

import(
  "os"
  "bufio"
  "fmt"
  "strings"
  "io"
)

func getFilenameFromConfig()(dbFile string){

  f,err := os.Open("config")
  if err!=nil{
    panic(err)
  }

  defer f.Close()

  bf := bufio.NewReader(f)
  line, _,err :=bf.ReadLine()

  if err!=nil{
    panic(err)
  }

  dbFile =string(line)

  //TODO : validate that dbFile is a valid file path.
  return
}

func addTask(file string, task string){
  fmt.Println("ADD TASKS: file is "+file+". Task :"+task)
}

func showTask(file string){
  f, err := os.Open(file)
  if err != nil {
    panic(err)
  }

  defer f.Close()

  bf := bufio.NewReader(f)
  for{
    line, isPrefix,err :=bf.ReadLine()

    if err == io.EOF{
      break
    }

    if err !=nil || isPrefix{
      panic(err)
    }

    fmt.Println(string(line))
  }
}

func reOrderTask(file string, oldNumber string, newNumber string){
  fmt.Println("REORDER TASKS : file : "+file+". Old :"+ oldNumber + ".New :"+newNumber)
}

func removeTask(file string,taskNum string){
  fmt.Println("REMOVE TASKS : file : "+file+" taskNum:"+taskNum)
}

func main(){

  dbFile := getFilenameFromConfig()
  // fmt.Println(dbFile)

  args := os.Args[1:]

  if len(args) == 0 {
    fmt.Println("Arguments\n-a Add a task \n-s Show all tasks\n-r remove task")
    return
  }

  switch args[0]{
    case "-a": addTask(dbFile,strings.Join(args[1:]," "))
    case "-s": showTask(dbFile)
    case "-r": removeTask(dbFile,args[1])
    case "-c": reOrderTask(dbFile,args[1],args[2])
  }

}
