package main

import "fmt"

type VideoCourse struct {
  Name string
  Language string
  TimeMinutes int
}

func main() {
  course := VideoCourse{}
  course.Name = "Go Language Overview"
  course.Language = "Go"
  course.TimeMinutes = 60

  fmt.Println("Name:", course.Name)
  fmt.Println("Language:", course.Language)
  fmt.Println("TimeMinutes:", course.TimeMinutes)
}
