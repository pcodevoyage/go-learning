package main

import "strconv"

func toInt(s string)(i int){
  i1,_ := strconv.ParseInt(s, 0, 64)
  i = int(i1)
  return
}
