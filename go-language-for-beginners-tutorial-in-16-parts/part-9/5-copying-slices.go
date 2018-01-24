package main

import "fmt"

func main() {
  languages := make([]string, 3)
  languages[0] = "Go"
  languages[1] = "Ruby"
  languages[2] = "Pony"

  learn := make([]string, len(languages))
  copy(learn, languages)

  fmt.Println(learn)
}
