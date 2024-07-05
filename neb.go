package main 

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Welcome to the contained nebula!")
    file, err := os.OpenFile(os.ExpandEnv("$HOME/neb/stuff"), os.O_RDWR | os.O_CREATE, 0644)

    if err != nil {
        fmt.Println("something went wrong with the file")
        fmt.Println(err)
        fmt.Println("also you suck lolz")
        return
    }

    file.Close()

    fmt.Println("the file opened")
}
