package main

import (
  "fmt"
  "os"
)

func main() {
  name := os.Args[1]
  name2 := os.Args[2]
  fmt.Println("First argument is ... " + name)
  fmt.Println("Second argument is ... " + name2)
}

/*
at the terminal:
go run main.go Todd
*/

