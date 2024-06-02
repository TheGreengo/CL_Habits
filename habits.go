package main

import (
    "fmt"
    "encoding/json"
    "os"
    "time"
)

/*
  So what do I want to happen? I want to run `goals` from the command 
  line, and have the thing prepare a list of ongoing goals for me. 

  I want to be focused on goals. The notion of them being done on a 
  weekly/daily/monthly basis should be up to me. 

  The first thing a goal needs is whether or not it's going to repeat. 
  
  If not, then it just needs a completion date, and whether it's a 
  number goal, or a checkbox.

  If so, then we need when it repeats until, and for now we'll make it
  log every day, with only a checkbox option available.
  
  All goals should have a name, a description, and a due date. 
*/

type RepeatGoal struct {
    Name string `json:"name"`
    Done Time   `json:"done"`
    WeekDays bool `json:"weekdays"`
    Interval int `json:"inter"`
    Days [7]bool `json:"days"`
    Complete bool `json:"complete"`
}

func (* RepeatGoal) String string {
    return "your mam"
}

/*
  The once goals should have a description, a target date, and then a 
  response.
*/
type OnceGoal struct {
    Name string `json:"name"`
    Desc string `json:"desc"`
    Due Time   `json:"due"`
    Success bool `json:"success"`
}

func (* OnceGoal) String string {
    return "your mem"
}
/*
  The goal list should be track finished and active goals for both 
  repeating and one time goals.
*/
type GoalList struct {
    OnceGoals []OnceGoal `json:"one_time"` 
    RepeatGoals []RepeatGoal `json:"repeat"` 
    DoneOnce []OnceGoal `json:"done_once"` 
    DoneRepeat []RepeatGoal `json:"done_repeat"` 
}

func (* GoalList) String string {
    return "your mom"
}

func main() {
    // we read in the data
    raw, er := os.ReadFile("./.config")

    if (er != nil) {
        fmt.Println(er)
    }

    var data GoalList
    err := json.Unmarshal(raw, &data)

    if (err != nil) {
        fmt.Println(er)
    }

    // say hello
    fmt.Println("Welcome to Braden's silly goal thingy")
    
    // then we need to print out our current goals

    // then you should have the option to select what goal they would
    // like to update

    // if they finish one, then it should be moved into the completed 
    // array

    // 
    // b, err := json.Marshal(thing)
    // er := os.WriteFile("./thing.json", b, 0644)
}
