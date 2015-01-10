package main

import(
  "os"
  "bufio"
  "fmt"
  "strings"
  "io"
)

func getFilenameFromConfig(){

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

  //dbFile is a global variable.
  dbFile =string(line)

  //TODO : validate that dbFile is a valid file path.
}

func addTask(task string){
  fmt.Println("ADD TASKS: file is "+dbFile+". Task :"+task)
}

func showTask(){
  f, err := os.Open(dbFile)
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

func reOrderTask(oldNumber string, newNumber string){
  fmt.Println("REORDER TASKS : file : "+dbFile+". Old :"+ oldNumber + ".New :"+newNumber)
}

func removeTask(taskNum string){
  fmt.Println("REMOVE TASKS : file : "+dbFile+" taskNum:"+taskNum)
}

var dbFile = ""
func main(){

  getFilenameFromConfig()
  // fmt.Println(dbFile)

  args := os.Args[1:]

  if len(args) == 0 {
    fmt.Println("Arguments\n-a Add a task \n-s Show all tasks\n-r remove task")
    return
  }

  switch args[0]{
    case "-a": addTask(strings.Join(args[1:]," "))
    case "-s": showTask()
    case "-r": removeTask(args[1])
    case "-c": reOrderTask(args[1],args[2])
  }

}
