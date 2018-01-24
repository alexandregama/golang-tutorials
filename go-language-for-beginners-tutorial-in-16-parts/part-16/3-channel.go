package main

import (
  "fmt"
  "time"
)

var webservice = make(chan int)

func callWebService() {
  go func() {
    fmt.Println("Calling webservice")
    time.Sleep(5 * time.Second)
    fmt.Println("Webservice Finished")
    webservice <-10
  }()
}

func showToUser() {
  fmt.Println("Showing info to User..")
}

func main() {
  callWebService()

  showToUser()

  result := <-webservice

  fmt.Println("Execution finished with the result:", result)
  time.Sleep(8 * time.Second)
}
