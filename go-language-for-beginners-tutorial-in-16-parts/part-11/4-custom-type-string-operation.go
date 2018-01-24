package main

import (
  "fmt"
)

type Language string

func main() {
  language := Language("Java")

  fmt.Println(language)

  language = language + " - Ruby"

  fmt.Println(language)
}
