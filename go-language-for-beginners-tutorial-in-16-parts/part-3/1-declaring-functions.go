package main

import "fmt"

func main() {
  fmt.Println("Executing inside the main() function")

  anotherFunction()
}

func anotherFunction() {
  fmt.Println("Executing inside the anotherFuntion()")
}
