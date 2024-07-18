package model

type Player struct {
	SkillsXp    map[*Skill]uint64
	SkillsLevel map[*Skill]byte
	CurrentPos  *Position
	Items       Inventory
}

func (p *Player) move(dx int, dy int) {
	p.CurrentPos = &Position{p.CurrentPos.x + dx, p.CurrentPos.y + dy}
}

func NewPlayer() Player {
	return Player{
		SkillsXp:    make(map[*Skill]uint64),
		SkillsLevel: make(map[*Skill]byte),
		CurrentPos:  &(Position{x: 0, y: 0}),
		Items:       Inventory{Items: make(map[string]int)},
	}
}

func (this Position) equals(other Position) bool {
	return this.x == other.x && this.y == other.y
}
