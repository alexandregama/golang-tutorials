package main

import "fmt"

func main() {
  number := 10
  switch {
  case number > 2 && number <= 8:
    fmt.Println("The number belongs to the first conditional")
  case number > 5 && number <= 10:
    fmt.Println("The number belongs to the second conditional")
  case number > 15 && number <= 20:
    fmt.Println("The number belongs to the third conditional")
  }
}
