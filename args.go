package main

import (
  "fmt"
  "os"
)

func main() {
  name := os.Args[1]
  name2 := os.Args[2]
  fmt.Println("最初の引数は ... " + name)
  fmt.Println("二つ目の引数は ... " + name2)
}

/*
at the terminal:
go run main.go Todd
*/

