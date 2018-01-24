package main

import "fmt"

func main() {
  value := sum(9, 7)

  fmt.Println("Total value:", value)
}

func sum(first, second int) int {
  totalValue := first + second

  return totalValue
}
