package main

import (
  "fmt"
  "math"
)

func main() {
  result, _ := square(-1)

  fmt.Println("Square of 16 is", result)
}

func square(value float64) (float64, error) {
    if value < 0 {
      return 0, fmt.Errorf("You can not use negative numbers!")
    }
    return math.Sqrt(value), nil
}
