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
    // b, err := json.Marshal(thing)
    // err := json.Unmarshal(thing, &foo)
    // er := os.WriteFile("./thing.json", b, 0644)
    // var foo Foo
    // thing, er := os.ReadFile("./.config")

}
