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

}

type OnceGoal struct {

}

type GoalList struct {
    OnceGoals []OnceGoal `json:"one_time"` 
    RepeatGoals []RepeatGoal `json:"repeating"` 
}

func main() {
    // b, err := json.Marshal(thing)
    // err := json.Unmarshal(thing, &foo)
    // er := os.WriteFile("./thing.json", b, 0644)
    // var foo Foo
    thing, er := os.ReadFile("./.config")

    
}
