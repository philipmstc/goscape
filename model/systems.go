package model

import "fmt"

func GetAvailableActions(player Player, board Board, skills map[string]*Skill) []Action {
	for tile, pos := range board.Tiles {
		if player.CurrentPos.equals(pos) {
			return tile.Actions
		}
	}
	comps := make(map[Product]int)
	comps[Product{"no-cost",1}] = 1
	action := MakeProduct{
		Recipe{
			Components: comps, 
			Product: Product{
				fmt.Sprintf("test:%d,%d", player.CurrentPos.x, player.CurrentPos.y),
				1,
			},
		}, 
		100, 
		"test",
	}
	
	newTile := NewTile(action)
	board.Tiles[&newTile] = NewPosition(player.CurrentPos.x, player.CurrentPos.y)
	return newTile.Actions
}

func NewTile( action Action) Tile {
	return Tile{Actions: []Action{
		action,
		MoveLeft,
		MoveRight,
		MoveUp,
		MoveDown,
		DisplaySkills,
		DisplayInv,
	}}
}

func InitialTile(skills map[string]*Skill) Tile {
	comps := make(map[Product]int)
	mockAction := MakeProduct{
		Recipe{Components: comps, Product: Product{"no-cost", 1}}, 100,  "init"}
	return NewTile( mockAction)
}
