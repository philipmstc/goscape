package model

import (
	"fmt"
	"math/rand"
	"philipmstc/goscape/feature"
)

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

func (product Product) String() string {
	return fmt.Sprintf("%v-%v", product.Name, product.Tier)
}

type Recipe struct {
	Components map[Product]int
	Product    Product
}

func (recipe Recipe) String() string {
	if feature.DetailedRecipeStrings() {
		if recipe.Components != nil {
			var comps string = ""
			for k, v := range recipe.Components {
				comps = comps + fmt.Sprintf("%v*%v + ", v, k)
			}
			if len(comps) > 2 {
				comps = comps[:len(comps)-2]
			}
			return fmt.Sprintf("Make %v from (%v)", recipe.Product, comps)
		} else {
			return fmt.Sprintf("(Free) %v", recipe.Product)
		}
	}
	return fmt.Sprintf("%v tier %v", recipe.Product.Name, recipe.Product.Tier);
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
	SkillsName string
}

func (mk MakeProduct) Do(player *Player) {
	if (player.CanCreate(mk.Recipe.Product, []Recipe{mk.Recipe})) {
		player.Create(mk.Recipe.Product, mk.Recipe)
		player.processXp(mk.SkillsName, mk.XpGain)
		fmt.Printf("xp: %v Recipe: %v", mk.XpGain, mk.Recipe)
	} else { 
		fmt.Printf("Cannot create, not enoguh components in inventory");
	}
}

func (mk MakeProduct) GetName() string {
	return fmt.Sprintf("Make [%v] (+%v %v xp)", mk.Recipe.Product.Name, mk.XpGain, mk.SkillsName)
}

func GenerateProductLineNM(Name string, skills []*Skill, minSkills int, maxSkills int) *Skill {
	const NEW_SKILL_RARITY_FACTOR = 2
	t := rand.Intn(len(skills)*NEW_SKILL_RARITY_FACTOR + 1)
	var targetSkill *Skill
	if t == (len(skills) * NEW_SKILL_RARITY_FACTOR) {
		fmt.Println("Making new skill")
		targetSkill = &Skill{[][]Recipe{}}
	} else {
		fmt.Println("using existing skill")
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
		fmt.Printf("############ %v ########", len(sourceSkill.ProductLines)) 
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
	newProductLine := []Recipe{{comps, Product{Name, 1}}}
	targetSkill.ProductLines = append(targetSkill.ProductLines, newProductLine)
	return targetSkill
}
