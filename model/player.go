package model

import "fmt"

type Player struct {
	Skills      map[string]*Skill
	SkillsXp    map[*Skill]uint64
	SkillsLevel map[*Skill]byte
	CurrentPos  *Position
	Items       Inventory
}

func (p *Player) move(dx int, dy int) {
	p.CurrentPos = &Position{p.CurrentPos.x + dx, p.CurrentPos.y + dy}
}

func (player *Player) processXp(skill string, xp uint64) {
	if player.Skills[skill] == nil {
		newSkill := &Skill{}
		player.Skills[skill] = newSkill
		player.SkillsXp[newSkill] = 0
		player.SkillsLevel[newSkill] = 0
	}
	actualSkill := player.Skills[skill]
	currentLevel := player.SkillsLevel[actualSkill]
	player.SkillsXp[actualSkill] += xp;
	for player.SkillsXp[actualSkill] > XP_PER_LEVEL[currentLevel] {
		currentLevel++
		player.SkillsLevel[actualSkill] += 1	
		fmt.Printf("%v levelled up! You are now level %v", skill,currentLevel)
	}
	
}



func NewPlayer() Player {
	return Player{
		Skills:      make(map[string]*Skill),
		SkillsXp:    make(map[*Skill]uint64),
		SkillsLevel: make(map[*Skill]byte),
		CurrentPos:  &(Position{x: 0, y: 0}),
		Items:       Inventory{Items: make(map[string]int)},
	}
}

func (this Position) equals(other Position) bool {
	return this.x == other.x && this.y == other.y
}
