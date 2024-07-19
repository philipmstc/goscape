package model

import "fmt"

func GetAvailableActions(player Player, board Board, skills map[string]*Skill) []Action {
	for tile, pos := range board.Tiles {
		if player.CurrentPos.equals(pos) {
			return tile.Actions
		}
	}
	action := GenerateAction(player) 
	newTile := NewTile(action)
	board.Tiles[&newTile] = NewPosition(player.CurrentPos.x, player.CurrentPos.y)
	return newTile.Actions
}

func NewTile(action Action) Tile {
	return Tile{Actions: []Action{
		action,
	}}
}

func InitialTile(skills map[string]*Skill) Tile {
	comps := make(map[Product]int)
	mockAction := MakeProduct{
		Recipe{Components: comps, Product: Product{"ps1-p1", 1}}, 100,  "primary-skill-1"}
	return NewTile( mockAction)
}

func GenerateAction(player Player) Action {
	comps := make(map[Product]int)
	comps[Product{"ps1-p1",1}] = 1
	return MakeProduct{
			Recipe{
				Components: comps, 
				Product: Product{
					fmt.Sprintf("ss1-p(%d,%d)", player.CurrentPos.x, player.CurrentPos.y),
					1,
				},
			}, 
			100, 
			"secondary-skill-1",
		}
}