package main

import "fmt"

func main() {
  printNumbers()
}

func printNumbers(numbers ...int) {
  fmt.Println(numbers)
}
