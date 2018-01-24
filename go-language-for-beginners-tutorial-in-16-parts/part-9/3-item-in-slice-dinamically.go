package main

import "fmt"

func main() {
  languages := make([]string, 3)
  languages[0] = "Go"
  languages[1] = "Ruby"
  languages[2] = "Pony"

  fmt.Println("Slice size:", len(languages))

  languages = append(languages, "Erlang")
  languages = append(languages, "Java")

  fmt.Println("New slice size:", len(languages))
}
