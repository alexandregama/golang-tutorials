package main

import "fmt"

type VideoCourse struct {
  Name string
  Language string
  TimeMinutes int
}

func main() {
  course := VideoCourse{Name: "Go Language Overview", Language: "Go", TimeMinutes: 60}

  fmt.Println("Course:", course)
}
