//simplejson.go
package main

import (
    "fmt"
    "github.com/likexian/simplejson"
)

func main() {
    data := `{"result":{"list":[0,1,2,3,4],"online":true,"rate":0.8},"status":{"code":1,"message":"success"},"hello":"1"}`
    json, _:= simplejson.Loads(data)
    status, _ := json.Get("status").Get("message").String()
    hello, _ := json.Get("hello").String()
    fmt.Println(status)
    fmt.Println(hello)
}