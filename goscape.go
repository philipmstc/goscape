package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"philipmstc/goscape/db"
	"philipmstc/goscape/feature"
	"philipmstc/goscape/model"
	"strconv"
	"strings"
	"time"
)


func ShowAnimation(action model.Action) {
	fmt.Print("[     ]")
	time.Sleep(1 * time.Second)
	fmt.Print("\r[#    ]")
	time.Sleep(1 * time.Second)
	fmt.Print("\r[##   ]")
	time.Sleep(1 * time.Second)
	fmt.Print("\r[###  ]")
	time.Sleep(1 * time.Second)
	fmt.Print("\r[#### ]")
	time.Sleep(1 * time.Second)
	fmt.Print("\r[#####]")
	fmt.Println()
}

func Load() (model.Player, model.Board, map[string]*model.Skill) {
	if feature.IsNewGame() {
		db.InitDb()
	}
	player := db.GetPlayer()
	board := db.GetBoard()
	skills := db.GetAllSkills()	
	return player, board, skills
}

func main() {
	player, board, skills := Load()	
	i := 1
	for {
		fmt.Println("Please choose an action: ")
		actions := model.GetAvailableActions(player, board, skills)

		for i, a := range actions {
			fmt.Printf("%v: %v", i+1, a.GetName())
			fmt.Println()
		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		input := strings.TrimSpace(scanner.Text())
		if choice, err := strconv.Atoi(input); err == nil && choice <= len(actions) && choice > -1 {
			if (feature.IsAnimated()) { 
				ShowAnimation(actions[choice-1])	
				if rand.Float64() > (1.0/math.Pow(2, float64(len(skills)))) { 
					skillsArray := []*model.Skill{}
					for _, skill := range(skills) {
						skillsArray = append(skillsArray, skill)
					}
					newSkill := model.GenerateNewSkill(skillsArray)
					db.PersistSkill(newSkill,"new-" + string(len(skills)))
					skills["new-" + string(len(skills))] = newSkill
				}
			}
			actions[choice-1].Do(&player)
		} else if err == nil {
			fmt.Println("Choice out of bounds")
		} else {
			switch input { 
			case "w": 
				model.MoveUp.Do(&player)
			case "W": 
				model.MoveUp.Do(&player)
			case "a": 
				model.MoveLeft.Do(&player)
			case "A": 
				model.MoveLeft.Do(&player)
			case "d": 
				model.MoveRight.Do(&player)
			case "D": 
				model.MoveRight.Do(&player)
			case "s": 
				model.MoveDown.Do(&player)
			case "S": 
				model.MoveDown.Do(&player)
			case "q":
				model.DisplayInv.Do(&player)
			case "Q": 
				model.DisplayInv.Do(&player)
			case "e": 
				model.DisplaySkills.Do(&player)
			case "E": 
				model.DisplaySkills.Do(&player)
			case "R": 
				fmt.Println(skills["secondary-skill-2"])
			case "f":
				skillsArray := []*model.Skill{}
				for _, skill := range(skills) { 
					skillsArray = append(skillsArray, skill)
				}
				p, s, t := model.TagSkills(skillsArray)
				fmt.Println(p)
				fmt.Println(s)
				fmt.Println(t)

				newSkill := model.GenerateProductLineNM("next",skillsArray , 1, 5);
				found := false
				for name, skill := range(skills)  {
					if skill == newSkill { 
						found = true	
						fmt.Println("Added recipe to existing skill ", name)
					}
				}
				if !found { 
					fmt.Println("Generated new skill entirely")
					newName := fmt.Sprintf("next#%v", i)
					i++
					skills[newName] = newSkill
					player.Skills[newName] = newSkill
					err := db.PersistSkill(newSkill, newName)
					if err != nil {
					    fmt.Println(err)
					}
				}
			default: 	
				fmt.Printf("Invalid Selection %v", input)
			}
		}
		fmt.Println()

	}
}
