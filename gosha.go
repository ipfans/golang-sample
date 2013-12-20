package main

import (
    "crypto/sha1"
    "crypto/sha256"
    "io"
    "fmt"
)

func main() {
    h := sha1.New()
    io.WriteString(h, "Hello world!")
    fmt.Printf("%X\n", h.Sum(nil))
    h = sha256.New()
    io.WriteString(h, "Hello world!")
    fmt.Printf("%X\n", h.Sum(nil)) 
}