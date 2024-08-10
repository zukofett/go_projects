package main

import (
	"bufio"
	"log"
	"os"
    "fmt"
)

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        readFromStdin()
        return
    }

    sysStdin := os.Stdin
    defer func() { os.Stdin = sysStdin }()


    for _, file := range args {
        f, err := os.Open(file)
        if err != nil {
            log.Fatal(err)
        }

        defer func() {
            if err := f.Close(); err != nil {
                log.Fatal(err)
            }
        }()

        os.Stdin = f
        readFromStdin()
    }
}

func readFromStdin() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatal("reading standard input:", err)
    } 
}


