package main

import (
  "fmt"
  "math"
  "log"
)

func main() {
  result, err := square(-1)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Square of 16 is", result)
}

func square(value float64) (float64, error) {
  if value < 0 {
    return 0, fmt.Errorf("You can not use negative numbers!")
  }
  return math.Sqrt(value), nil
}
