package main

import "fmt"

type Framework string
type Language string

func main() {
  language := Language("Java")
  framework := Framework("Spring 5")

  fmt.Println("Language to learn:", language)
  fmt.Println("Framework to learn:", framework)
}
