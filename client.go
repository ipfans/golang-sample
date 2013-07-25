//client.go
package main

import (
    "net"
    "fmt"
    "bufio"
    "log"
)

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:2000")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    fmt.Fprintf(conn, "hello!\n")
    // recv data, end with '\n'.
    status, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil{
        log.Fatal(err)
    }
    fmt.Printf(status)
}