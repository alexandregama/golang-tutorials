package main

import (
  "fmt"
  "reflect"
)

func main() {
  fmt.Println(reflect.TypeOf(1))
  fmt.Println(reflect.TypeOf(9.5))
  fmt.Println(reflect.TypeOf("Just a String"))
  fmt.Println(reflect.TypeOf(true))
}
