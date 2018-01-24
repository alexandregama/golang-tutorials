package main

import "fmt"

type VideoCourse struct {
  Minute int
  Language string
}

func (v VideoCourse) Description() string {
  return fmt.Sprintf("Video Course for %s with a duration of %d minutes", v.Language, v.Minute)
}

func main() {
  course := VideoCourse{Language: "Go", Minute: 60}
  fmt.Println(course)

  description := course.Description()
  fmt.Println(description)
}
