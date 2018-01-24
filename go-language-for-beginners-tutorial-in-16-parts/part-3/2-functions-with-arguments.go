package main

import "fmt"

func main() {
  sumAndPrint(3, 6)
}

func sumAndPrint(first int, second int) {
  sum := first + second

  fmt.Println(sum)
}
