package main

import "fmt"

type VideoCourse struct {
  Name string
  Language string
  TimeMinutes int
}

func main() {
  course := VideoCourse{"Go Language Overview", "Go", 60}

  fmt.Println("Course:", course)
}
