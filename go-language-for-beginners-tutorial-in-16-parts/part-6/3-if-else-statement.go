package main

import "fmt"

func main() {
  number := 10
  if number < 5 {
    fmt.Println("Will not execute this")
  } else {
    fmt.Println("The code in the else will be executed!")
  }
}
