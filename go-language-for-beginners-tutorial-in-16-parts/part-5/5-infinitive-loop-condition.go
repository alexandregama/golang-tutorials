package main

import "fmt"

func main() {
  for {
    importantValue := importantFunction()
    if importantValue == 10 {
      fmt.Println(importantValue)
      break
    }
  }
}

func importantFunction() int {
  return 10
}
