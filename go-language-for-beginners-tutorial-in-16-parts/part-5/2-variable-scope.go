package main

import "fmt"

func main() {
  numberOutside := 10
  for i := 0; i < 5; i++ {
    numberInside := 5
    numberOutside++
  }
  fmt.Println("Number outside loop:", numberOutside)
  fmt.Println("Number inside loop:", numberInside)
}
