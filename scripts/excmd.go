package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("hello", os.Args[1])
    os.Exit(42)
}
