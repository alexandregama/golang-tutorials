package main

import "fmt"

func main() {
  languages := map[string]int{}
  languages["java"] = 5
  languages["ruby"] = 4
  languages["go"] = 2

  for language, number := range languages {
    fmt.Println("Key:", language, "- Value:", number)
  }

  for _, number := range languages {
    fmt.Println("Value:", number)
  }
}
