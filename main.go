package main

import (
  "os"
  "io"
  "fmt"
  "strings"
)

func main() {
  file, err := os.Open("books/frankenstein.txt")
  if err != nil {
    fmt.Print(err)
  }

  output, err := io.ReadAll(file)
  words := strings.Fields(string(output))

  fmt.Print("Words in the book: ", len(words))
}
