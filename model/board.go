package model

import "fmt"

type Board struct {
	Tiles map[*Tile]Position
}

type Position struct {
	x int
	y int
}

type MoveAction struct  {
	deltaX int
	deltaY int
	name string
}

func (move MoveAction) Do(player *Player) {
	player.move(move.deltaX, move.deltaY)
}

func (move MoveAction) GetName() string {
	return move.name
}

var MoveLeft = MoveAction{-1, 0, "Go left"}
var MoveRight = MoveAction{1, 0, "Go right"}
var MoveUp = MoveAction{0, -1, "Go up"}
var MoveDown = MoveAction{0, 1, "Go down"}

type displaySkills struct {}
var DisplaySkills = displaySkills{}
type displayInv struct {}
var DisplayInv = displayInv{}

func (ds displaySkills) Do(player *Player) {
	fmt.Printf("%v", player.SkillsXp)
}
func (ds displaySkills) GetName() string { 
	return "Check Skills Xp"
}

func (di displayInv) Do(player *Player) { 
	fmt.Printf("%v", player.Items.Items)
}
func (di displayInv) GetName() string { 
	return "Show Inventory"
}

func NewPosition(x int, y int) Position {
	return Position{x, y}
}

func NewGameBoard() Board {
	return Board{make(map[*Tile]Position)}
}
