package model

import (
	"fmt"
	"math/rand"
)

type Inventory struct {
	Items map[string]int
}

type Skill struct {
	ProductLines [][]Recipe
}

func PrimaryResource(Name string, Tiers int) []Recipe {
	Recipes := []Recipe{}
	for i := range Tiers {
		dProduct := Product{Name, i + 1}
		Recipes = append(Recipes, Recipe{nil, dProduct})
	}
	return Recipes
}

func SelfSingleProcessResource(resource string, Name string, Tiers int) []Recipe {
	Recipes := []Recipe{}
	for i := range Tiers {
		dProduct := Product{Name, i + 1}
		Components := make(map[Product]int)
		cProd := Product{resource, i + 1}
		Components[cProd] = 1
		Recipes = append(Recipes, Recipe{Components, dProduct})
	}
	return Recipes
}

type Product struct {
	Name string
	Tier int
}

func (this Product) String() string {
	return fmt.Sprintf("%v-%v", this.Name, this.Tier)
}

type Recipe struct {
	Components map[Product]int
	Product    Product
}

func (this Recipe) String() string {
	if this.Components != nil {
		var comps string = ""
		for k, v := range this.Components {
			comps = comps + fmt.Sprintf("%v*%v + ", v, k)
		}
		if len(comps) > 2 {
			comps = comps[:len(comps)-2]
		}
		return fmt.Sprintf("Make %v from (%v)", this.Product, comps)
	} else {
		return fmt.Sprintf("(Free) %v", this.Product)
	}
}

type Tile struct {
	Actions []Action
}

type Action interface {
	Do(player *Player)
	GetName() string
}

type MakeProduct struct {
	Recipe Recipe
	XpGain uint64
	Skill *Skill
}

func (mk MakeProduct) Do(player *Player) {
	if (player.Items.CanCreate(mk.Recipe.Product, []Recipe{mk.Recipe})) {
		player.Items.Create(mk.Recipe.Product, mk.Recipe)
		player.SkillsXp[mk.Skill] += mk.XpGain;
		fmt.Printf("xp: %v Recipe: %v", mk.XpGain, mk.Recipe)
	} else { 
		fmt.Printf("Cannot create, not enoguh components in inventory");
	}
}

func (mk MakeProduct) GetName() string {
	return fmt.Sprintf("Make [%v] for %v xp", mk.Recipe.Product.Name, mk.XpGain)
}

// only works if theres only one Recipe per Product
func (this Inventory) CanCreate(Product Product, RecipeBook []Recipe) bool {
	for _, r := range RecipeBook {
		if r.Product.Name == Product.Name {
			for k, v := range r.Components {
				if this.Items[k.Name] < v {
					return false
				}
			}
			return true
		}
	}
	return false
}

func (this Inventory) Create(Product Product, Recipe Recipe) {
	for k, v := range Recipe.Components {
		this.Items[k.Name] -= v
	}
	this.Items[Recipe.Product.Name] += 1
}

func GenerateProductLineNM(Name string, skills []Skill, minSkills int, maxSkills int) []Recipe {
	const NEW_SKILL_RARITY_FACTOR = 2
	t := rand.Intn(len(skills)*NEW_SKILL_RARITY_FACTOR + 1)
	var targetSkill Skill
	if t == (len(skills) * NEW_SKILL_RARITY_FACTOR) {
		targetSkill = Skill{[][]Recipe{}}
	} else {
		targetSkill = skills[t%len(skills)]
	}
	comps := make(map[Product]int)
	var actualSkills int
	if minSkills >= maxSkills {
		actualSkills = minSkills
	} else {
		actualSkills = minSkills + rand.Intn(maxSkills-minSkills)
	}
	fmt.Printf("Making Product line from %v 'skills'\n", actualSkills)
	for i := 0; i <= actualSkills; i++ {
		var s int
		for {
			s = rand.Intn(len(skills))
			if s != ((t) % len(skills)) {
				break
			}
		}
		sourceSkill := skills[s]
		sourceRecipeIndex := rand.Intn(len(sourceSkill.ProductLines))
		sourceRecipe := sourceSkill.ProductLines[sourceRecipeIndex][0]
		_, ok := comps[sourceRecipe.Product]
		if ok {
			comps[sourceRecipe.Product] += 1
		} else {
			comps[sourceRecipe.Product] = 1
		}
	}
	fmt.Println("Target Skill = ", targetSkill)
	return []Recipe{{comps, Product{Name, 1}}}
}
