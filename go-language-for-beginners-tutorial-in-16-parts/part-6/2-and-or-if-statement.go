package main

import "fmt"

func main() {
  number := 10
  amount := 15

  if number >= 5 && amount <= 20 {
    fmt.Println("This will be printed!")
  }

  if number >= 15 || amount <= 20 {
    fmt.Println("This will be printed!")
  }
}
