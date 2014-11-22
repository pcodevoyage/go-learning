//Go does not have classes
//We can defind methods on struct

package main

import (
    "fmt"
    "math"
)

type Vertex struct {
  X, Y float64
}

//Method 1 : Pointer receiver
// a) avoid copying the value on each methid call
// b) method can modify the value that its reveiver points to
func (v *Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


type MyFloat float64

func (f MyFloat) Abs() float64 {
  if f < 0 {
      return float64(-f)
  }
  return float64(f)
}

func main() {
  v := &Vertex{3, 4}
  fmt.Println(v.Abs())

  f := MyFloat(-math.Sqrt2)
  fmt.Println(f.Abs())
}
