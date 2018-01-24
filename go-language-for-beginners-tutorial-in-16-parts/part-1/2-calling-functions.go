package main

import "fmt"

func main() {
  fmt.Println("Calling the function: main")

  anotherFunction()
}

func anotherFunction() {
  fmt.Println("Calling the function: anotherFunction()")
}
