package main

import "fmt"

func main() {
  languages := make([]string, 5)
  languages[0] = "Ruby"
  languages[1] = "JavaScript"
  languages[2] = "Java"
  languages[3] = "Scala"
  languages[4] = "Go"

  dynamicLanguages := languages[0:2]
  fmt.Println(dynamicLanguages)

  staticTyped := languages[2:5]
  fmt.Println(staticTyped)
}
