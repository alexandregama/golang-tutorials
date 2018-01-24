package main

import "fmt"

func main() {
  number := 10
  switch number {
    case 3, 8, 15:
      fmt.Println("The number is in the first range!")
    case 9, 10, 20:
      fmt.Println("The number is in the second range!")
    case 16, 21, 25:
      fmt.Println("The number is in the third range!")
  }

}
