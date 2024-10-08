package model

import "fmt"

func GetAvailableActions(player Player, board Board, skills map[string]*Skill) []Action {
	for tile, pos := range board.Tiles {
		if player.CurrentPos.equals(pos) {
			return tile.Actions
		}
	}
	action := GenerateAction(player, skills) 
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
	initialRecipe := Recipe{Components: comps, Product: Product{"ps1-p1", 1}}
	skills["primary-skill-1"].ProductLines = [][]Recipe{}
	skills["primary-skill-1"].ProductLines = append(skills["primary-skill-1"].ProductLines, []Recipe{initialRecipe})
	mockAction := MakeProduct{
		initialRecipe, 100,  "primary-skill-1"}
	return NewTile(mockAction)
}

func GenerateAction(player Player, skills map[string]*Skill) Action {
	comps := make(map[Product]int)
	comps[Product{"ps1-p1",1}] = 1
	newRecipe := Recipe{
				Components: comps, 
				Product: Product{
					fmt.Sprintf("ss1-p(%d,%d)", player.CurrentPos.x, player.CurrentPos.y),
					1,
				},
			}
	skills["secondary-skill-2"].ProductLines = append(skills["secondary-skill-2"].ProductLines, []Recipe{newRecipe})
	return MakeProduct{
		newRecipe,
		100, 
		"secondary-skill-2",
	}
}