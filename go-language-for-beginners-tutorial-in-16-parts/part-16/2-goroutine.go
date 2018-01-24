package main

import (
  "fmt"
  "time"
)

func callWebService() {
  go func() {
    fmt.Println("Calling webservice")
    time.Sleep(5 * time.Second)
    fmt.Println("Webservice Finished")
  }()
}

func showToUser() {
  fmt.Println("Showing info to User..")
}

func main() {
  callWebService()

  showToUser()

  fmt.Println("Execution finished with the result:")
  time.Sleep(8 * time.Second)
}
