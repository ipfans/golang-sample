package main
import (
    "github.com/DHowett/go-plist"
    "os"
    "fmt"
    "io/ioutil"
    "bytes"
)
func main() {
    // Can not run with https://github.com/DHowett/go-plist/issues/2
    if len(os.Args) != 2 {
        fmt.Println("decoder.exe [Info.plist]")
        return
    }
    filename := os.Args[1]
    filebuf,err := ioutil.ReadFile(filename)
    if err != nil{
        fmt.Println(err)
        return
    }

    var bval interface{}
    reader := bytes.NewReader(filebuf)
    decoder := plist.NewDecoder(reader)
    decoder.Decode(&bval)
    fmt.Println(bval)
}