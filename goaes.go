package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/md5"
    "crypto/rand"
    "os"
    "fmt"
    "io"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Printf("%s [key] [Original Text]", os.Args[0])
    }

    keyhash := md5.New()
    io.WriteString(keyhash, os.Args[1])
    keyStr := fmt.Sprintf("%x", keyhash.Sum(nil))
    fmt.Printf("key is %s\n", keyStr)
    keyByte := []byte(keyStr)
    block, err:= aes.NewCipher(keyByte)
    if err != nil {
        fmt.Println(err)
    }
    blocksize := block.BlockSize()
    textByte := []byte(os.Args[2])
    needtoAdd := len(textByte) % blocksize - 16
    if needtoAdd != - 16{
        fmt.Println("Need to Extend Original text with 0x00")
        newTextByte := make([]byte, len(textByte) - needtoAdd)
        copy(newTextByte, textByte)
        textByte = newTextByte
    }
    ciphertext := make([]byte, blocksize + len(textByte))
    iv := ciphertext[:blocksize]
    //Rock your iv!
    _, err = io.ReadFull(rand.Reader, iv)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("iv: ", iv)
    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(ciphertext[blocksize:], textByte)
    fmt.Printf("Encrypted Text: %x\n", ciphertext[blocksize:])
    mode = cipher.NewCBCDecrypter(block, iv)
    mode.CryptBlocks(ciphertext[blocksize:], ciphertext[blocksize:])
    fmt.Printf("Original Text: %s\n", ciphertext[blocksize:])
}