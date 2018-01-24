package main

import "fmt"

func main() {
  var languages [5]string
  languages[0] = "Go"
  languages[1] = "Ruby"
  languages[2] = "Pony"
  languages[3] = "Erlang"
  languages[4] = "Java"

  fmt.Println("Array size:", len(languages))  
}
