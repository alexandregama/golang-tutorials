package main

import "fmt"

func main() {
  total := sum(8, 12)

  fmt.Println("Total value:", total)
}

func sum(first, second int) (amount int) {
  amount = first + second

  return
}
