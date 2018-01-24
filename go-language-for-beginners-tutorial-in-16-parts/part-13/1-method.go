package main

import (
  "fmt"
  "strings"
)

type Description string

func (d Description) Upper() string {
  return strings.ToUpper(string(d))
}

func main() {
  description := Description("My Go special description")
  upper := description.Upper()

  fmt.Println(upper)
}
