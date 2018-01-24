package main

import (
  "fmt"
  "time"
)

func callWebService() int {
  fmt.Println("Calling webservice")
  time.Sleep(5 * time.Second)
  fmt.Println("Webservice Finished")
  return 10
}

func showToUser() {
  fmt.Println("Showing info to User..")
}

func main() {
  result := callWebService()

  showToUser()

  fmt.Println("Execution finished with the result:", result)
  time.Sleep(8 * time.Second)
}
