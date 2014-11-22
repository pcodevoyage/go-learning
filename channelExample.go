// Channel can be synchronous - if there is no receiver it will block.
// It can have a buffer too. So the sender can just stuff in it and continue doing thing

package main

import "fmt"

func main(){
  ch:=make(chan int) //channel of int
  go fibs(ch)

  for i:=0;i<20;i++ {
    fmt.Println(<-ch)
  }
}

func fibs(ch chan int){
  i,j := 0,1
  for {
    ch <- j
    i,j=j,i+j
  }
}
