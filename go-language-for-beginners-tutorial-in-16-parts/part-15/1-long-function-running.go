package main

import (
  "fmt"
  "time"
)

func callWebService(value int) {
  fmt.Println("WebService started:", value)
  time.Sleep(3 * time.Second)
  fmt.Println("WebService finished:", value)
}

func main() {
  callWebService(1)
  callWebService(2)
  callWebService(3)
}
