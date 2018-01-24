package main

import "fmt"

func main() {
  var price float64 = 15

  fmt.Println("Price before:", price)

  var number int = 10

  price = float64(number)

  fmt.Println("Price after: ", price)
}
