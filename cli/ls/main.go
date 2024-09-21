package main

import (
    "os"
    "fmt"
)

func main() {
    cliArgs := os.Args[1:]


    fmt.Println(cliArgs)
}
