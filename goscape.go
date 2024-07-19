package main

import (
	"bufio"
	"fmt"
	"os"
	"philipmstc/goscape/model"
	"strconv"
	"strings"
	"philipmstc/goscape/feature"
)

// testing generating new skills from a previously existing set
func main() {
	var player model.Player
	var board model.Board
	var skills map[string]*model.Skill
	feature.Load()
	if feature.IsNewGame() {
		player = model.NewPlayer()
		board = model.NewGameBoard()
		skills = make(map[string]*model.Skill)
		skills["primary-skill-1"] = &model.Skill{}
		skills["secondary-skill-2"] = &model.Skill{}
		tile1 := model.InitialTile(skills)
		board.Tiles[&tile1] = model.NewPosition(0, 0)
	}
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
				feature.ShowAnimation(actions[choice-1])	
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
			default: 	
				fmt.Printf("Invalid Selection %v", input)
			}
		}
		fmt.Println()

	}
}
