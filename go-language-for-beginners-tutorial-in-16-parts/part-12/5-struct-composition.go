package main

import "fmt"

type Minute struct {
  Value int
}

type VideoCourse struct {
  Name string
  Language string
  Time Minute
}

func main() {
  minutes := Minute{Value: 60}
  course := VideoCourse{Name: "Go Language Overview", Language: "Go", Time: minutes}

  fmt.Println(course)
}
