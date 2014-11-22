//Maps -- functions 

package main

import (
 "fmt"
 "time"
)

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

var m1 = map[string]Vertex{
    "Bell Labs": Vertex{
        40.68433, -74.39967,
    },
    "Google": Vertex{
        37.42202, -122.08408,
    },
}


func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}


func main() {
   m = make(map[string]Vertex)
   m["Bell Labs"] = Vertex{
       40.68433, -74.39967,
   }
   fmt.Println(m["Bell Labs"])

   fmt.Println(m1)

   //modify  maps 
   m3 := make(map[string]int)
   m3["this"] = 2
   m3["is"] = 3 
   fmt.Println("m3['this'] :",m3["this"])

   //functions as a variables. 
   pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )
    }  
   
   //using switch as if then else 
   t := time.Now()
   switch {
    case t.Hour() < 12:
       fmt.Println("Good morning!")
    case t.Hour() < 17:
       fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
   }   
}
