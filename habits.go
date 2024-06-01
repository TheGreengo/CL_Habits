package main

import (
    "fmt"
    "encoding/json"
    "os"
)

type Foo struct {
    Name string `json:"nombre"` 
    Age int `json:"alder"` 
    Height int `json:"hoejde"`
}

func main() {
    thing := Foo{ Name:"Jason", Age:24, Height:70 }
    fmt.Println(thing)

    b, err := json.Marshal(thing)

    if (err != nil) {
	    fmt.Println("Oh no, our table")
    }

	fmt.Println(string(b))
    er := os.WriteFile("./thing.json", b, 0644)

    if (er != nil) {
	    fmt.Println(er)
    }
}
