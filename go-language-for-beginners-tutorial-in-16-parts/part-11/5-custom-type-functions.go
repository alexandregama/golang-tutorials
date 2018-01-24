package main

import "fmt"

type Language string

func main() {
  language := Language("Java")

  fmt.Println("Language:", language)

  print(string(language))
}

func print(value string) {
  fmt.Println("Value:", value)
}
