package main

import "fmt"

func main() {
  number := 10
  if number > 5 {
    fmt.Println("Will not be printed!")
  } else if number < 7 {
    fmt.Println("Will not be printed!")
  } else if number < 9 {
    fmt.Println("Will not be printed!")
  } else {
    fmt.Println("Message Default!")
  }
}
