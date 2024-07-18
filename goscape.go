package main

import (
	"bufio"
	"fmt"
	"os"
	"philipmstc/goscape/model"
	"strconv"
	"strings"
)

// testing generating new skills from a previously existing set
func main() {
	player := model.NewPlayer()
	board := model.NewGameBoard()
	skills := make(map[string]*model.Skill)
	skills["init"] = &model.Skill{}
	skills["test"] = &model.Skill{}
	tile1 := model.InitialTile(skills)
	board.Tiles[&tile1] = model.NewPosition(0, 0)
	for {
		fmt.Println("Please choose an action: ")
		actions := model.GetAvailableActions(player, board, skills)

		for i, a := range actions {
			fmt.Printf("%v: %v", i+1, a.GetName())
			fmt.Println()
		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if choice, err := strconv.Atoi(strings.TrimSpace(scanner.Text())); err == nil && choice <= len(actions) && choice > 0 {
			actions[choice-1].Do(&player)
		} else if err == nil {
			fmt.Println("Choice out of bounds")
		} else {
			fmt.Println("Error from string conversion")
		}
		fmt.Println()

	}
}
