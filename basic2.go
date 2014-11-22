//Struct -- Pointer -- Array -- Slices -- range
package main 

import "fmt"

type Coordinates struct { 
 latitude int 
 longitude int 
}


func main() { 
  fmt.Println(Coordinates{53,65})

  v := Coordinates{23,56}
  fmt.Println(v.latitude)

  //Pointers [but no pointer arithmetic]
  point := &v
  point.longitude = 100
  fmt.Println(v)

  //The expression new(T) allocates a zeroed T value and returns a pointer to it. 
  coor := new(Coordinates)
  fmt.Println(coor)
  coor.latitude, coor.longitude = 300,90
  fmt.Println(coor)

  //Array 
  var isp[4]string
  isp[0] = "I"
  isp[1] = "am"
  isp[2] = "learning"
  isp[3] = "Go"
  fmt.Println(isp)

  //slices : Like array but does not have specific length
  letters := []string{"a", "b", "c", "d"}  
  fmt.Println(isp[2:])
  
  //slices are created with make 
  a := make([]int,5)
  fmt.Println(a)
  //func make([]T, len, cap) []T
  b := make([]int, 0, 5) // len(b)=0, cap(b)=5

  b = b[:cap(b)] // len(b)=5, cap(b)=5
  b = b[1:]      // len(b)=4, cap(b)=4

  for i,v := range letters { 
    fmt.Println( i ," - ",v)
  }   
}
