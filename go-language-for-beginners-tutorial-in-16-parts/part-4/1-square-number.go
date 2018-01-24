package main

import (
  "fmt"
  "math"
)

func main() {
  result := square(16)

  fmt.Println("Square of 16 is", result)
}

func square(value float64) float64 {
  return math.Sqrt(value)
}
