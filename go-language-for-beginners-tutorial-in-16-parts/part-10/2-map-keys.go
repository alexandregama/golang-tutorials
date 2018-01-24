package main

import "fmt"

func main() {
  languages := map[string]int{}
  languages["java"] = 5
  languages["ruby"] = 4
  languages["go"] = 2

  fmt.Println("First value:", languages["java"])
  fmt.Println("Second value:", languages["ruby"])
  fmt.Println("Third value:", languages["go"])
}
