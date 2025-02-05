package main

import (
  "os"
  "io"
  "fmt"
)

func main() {
  file, err := os.Open(".gitignore/books/frankenstein.txt")
  if err != nil {
    fmt.Print(err)
  }

  output, err := io.ReadAll(file)
  fmt.Print(string(output))
}
