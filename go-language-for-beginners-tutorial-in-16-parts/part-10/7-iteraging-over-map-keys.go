package main

import "fmt"

func main() {
  languages := map[string]int{}
  languages["java"] = 5
  languages["ruby"] = 4
  languages["go"] = 2

  for key, _ := range languages {
    fmt.Println("Value:", key)
  }
}
