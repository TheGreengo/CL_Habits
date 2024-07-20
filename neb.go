package main 

import (
    "encoding/json"
    "fmt"
    "os"
)
    
type Task struct {
    Desc string `json:"desc"`
    Comp bool `json:"comp"`
}

type TodoList struct {
    Todos []Task `json:"todos"`
}

func (t TodoList) del(ind int) {
    t.Todos = append(t.Todos[:ind], t.Todos[ind+1:]...)

    fmt.Println("del Called")
    fmt.Println(t.Todos)
}

func check(err error, mess string) {
    if err != nil {
        fmt.Println("Issues with:", mess)
        fmt.Println(err)
        os.Exit(2)
    }
}

func main() {
    fmt.Println("Welcome to the contained nebula!")

    pathy := os.ExpandEnv("$HOME/neb/stuff.json")

    file, err := os.OpenFile(pathy, os.O_RDWR | os.O_CREATE, 0644)
    check(err, "file open")
    defer file.Close()
    
    shtat, err := os.Stat(pathy)
    check(err, "shtat")

    if (shtat.Size() == 0) {
        fmt.Println("Initializing")
        file.WriteString("{\n\"todos\":[]\n}")
    }

    var todos TodoList
    
    raw, err := os.ReadFile(pathy)
    check(err, "getting raw and wriggling")


    if err := json.Unmarshal(raw, &todos); err != nil {
        check(err, "unmarshaling")
    }

    var word string
    var num int
    for {
        fmt.Println(todos.Todos)
        for _, i := range todos.Todos {
            if !i.Comp {
                fmt.Println(i.Desc)
            } else {
                fmt.Printf("\033[9m%s\033[0m\n", i.Desc)
            }
        }
        
        fmt.Scanf("%s %d", &word, &num)
        if word == "quit" {
            break
        } else if word == "comp" {
            todos.Todos[num - 1].Comp = true
        } else if word == "del" {
            todos.del(num - 1)
        }
    }
}
