package main

import (
    "fmt"
    "philipmstc/goscape/model"
)

// testing generating new skills from a previously existing set
func main() { 
    s1 := model.Skill{model.PrimaryResource("logs", 3)}
    s2 := model.Skill{model.PrimaryResource("bars", 3)}
    s3 := model.Skill{model.PrimaryResource("feathers", 1)}
    skills := []model.Skill{s1, s2, s3}
    fmt.Println(model.GenerateProductLineNM("new_item", skills, 2, 2))
    fmt.Println("Main")
}
