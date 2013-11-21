package main
import (
    "net"
    "fmt"
    "strings"
)

func main() {
    conn, err := net.Dial("tcp", "www.baidu.com:80")
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    defer conn.Close()
    fmt.Println(strings.Split(conn.LocalAddr().String(), ":")[0])
}
