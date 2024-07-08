package main 

import (
    "encoding/json"
    "fmt"
    "os"
)

type Task struct {
    Desc string `json:"desc"`
}

func main() {
    fmt.Println("Welcome to the contained nebula!")

    pathy := os.ExpandEnv("$HOME/neb/stuff.json")

    file, err := os.OpenFile(pathy, os.O_RDWR | os.O_CREATE, 0644)

    if err != nil {
        fmt.Println("something went wrong with the file")
        fmt.Println(err)
        return
    }
    
    shtat, err := os.Stat(pathy)
    
    if err != nil {
        fmt.Println("something went wrong with the shtat")
        fmt.Println(err)
        return
    }

    fmt.Println("Size:", shtat.Size())

    if (shtat.Size() == 0) {
        fmt.Println("Initializing")
        file.WriteString("{\n\"Todos\":[]\n}")
    }

    file.Close()

    fmt.Println("the file opened")
}
