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
  go callWebService(1)
  go callWebService(2)
  go callWebService(3)
  time.Sleep(10 * time.Second)
}
