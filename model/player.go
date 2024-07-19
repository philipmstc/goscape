package model

import "fmt"

type Player struct {
	Skills      map[string]*Skill
	SkillsXp    map[*Skill]uint64
	SkillsLevel map[*Skill]byte
	CurrentPos  *Position
	Items       map[string]int
}

func (p *Player) move(dx int, dy int) {
	p.CurrentPos = &Position{p.CurrentPos.x + dx, p.CurrentPos.y + dy}
}

func (player *Player) processXp(skill string, xp uint64) {
	if player.Skills[skill] == nil {
		newSkill := &Skill{}
		player.Skills[skill] = newSkill
		player.SkillsXp[newSkill] = 0
		player.SkillsLevel[newSkill] = 1
	}
	actualSkill := player.Skills[skill]
	currentLevel := player.SkillsLevel[actualSkill]
	player.SkillsXp[actualSkill] += xp;
	for player.SkillsXp[actualSkill] > XP_PER_LEVEL[currentLevel] {
		currentLevel++
		player.SkillsLevel[actualSkill] += 1	
		fmt.Printf("%v levelled up! You are now level %v\n", skill,currentLevel)
	}
	
}

// only works if theres only one Recipe per Product
func (player *Player) CanCreate(Product Product, RecipeBook []Recipe) bool {
	for _, r := range RecipeBook {
		if r.Product.Name == Product.Name {
			for k, v := range r.Components {
				if player.Items[k.Name] < v {
					return false
				}
			}
			return true
		}
	}
	return false
}

func (player *Player) Create(Product Product, Recipe Recipe) {
	for k, v := range Recipe.Components {
		player.Items[k.Name] -= v
	}
	player.Items[Recipe.Product.Name] += 1
}

type displaySkills struct {}
var DisplaySkills = displaySkills{}
type displayInv struct {}
var DisplayInv = displayInv{}

func (ds displaySkills) Do(player *Player) {
	for name, skill := range(player.Skills) {
		currentLevel := player.SkillsLevel[skill]
		fmt.Printf("%v Level: %v (%v/%v)\n", 
		name,
		currentLevel,
		player.SkillsXp[skill],
		XP_PER_LEVEL[currentLevel])
	}
}
func (ds displaySkills) GetName() string { 
	return "Check Skills Xp"
}

func (di displayInv) Do(player *Player) { 
	fmt.Printf("%v", player.Items)
}
func (di displayInv) GetName() string { 
	return "Show Inventory"
}

func NewPlayer() Player {
	return Player{
		Skills:      make(map[string]*Skill),
		SkillsXp:    make(map[*Skill]uint64),
		SkillsLevel: make(map[*Skill]byte),
		CurrentPos:  &(Position{x: 0, y: 0}),
		Items:       make(map[string]int),
	}
}
