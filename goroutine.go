//groutine is a lightweight thread managed by Go runtime

package main

import (
    "fmt"
    "time"
)

func say(s string) {
  for i := 0; i < 5; i++ {
      time.Sleep(100 * time.Millisecond)
      fmt.Println(s)
  }
}


func sum(a []int, c chan int) {
   sum := 0
   for _, v := range a {
       sum += v
   }
   c <- sum // send sum to c
}

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}

func main() {
  go say("world")
  say("hello")

  a := []int{7, 2, 8, -9, 4, 0}
  
  //make a channel
  c := make(chan int)
 
  go sum(a[:len(a)/2], c)
  go sum(a[len(a)/2:], c)
  
  x, y := <-c, <-c // receive from c

  fmt.Println(x, y, x+y)

  //Sender can close the channel  
  // v , ok <- ch ( here ok is of boolean type- set false when it is closed ) 
  c1 := make(chan int, 10)
  go fibonacci(cap(c1), c1)
  for i := range c1 {
        fmt.Println(i)
  }
}
