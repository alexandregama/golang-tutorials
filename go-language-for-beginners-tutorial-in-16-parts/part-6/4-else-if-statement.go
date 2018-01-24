package main

import "fmt"

func main() {
  number := 10
  if number < 5 {
    fmt.Println("Will not be printed!")
  } else if number < 8 {
    fmt.Println("Will not be printed!")
  } else if number < 12 {
    fmt.Println("Will be printed, since number is less than 12!")
  } else if number < 15 {
    fmt.Println("Will no be printed, since the last statement was printed out!")
  } else {
    fmt.Println("Default")
  }
}
