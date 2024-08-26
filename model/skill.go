package model

import (
	"slices"
	"fmt"
	"math/rand"
	"philipmstc/goscape/feature"
)

type Skill struct {
	ProductLines [][]Recipe
}

type TaggedSkill struct { 
	Skill Skill 
	Type string
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

const targetPrimarySkillDensity float64 = 0.33
const targetSecondarySkillDensity float64 = 0.33
const targetTertiarySkillDensity float64 = 0.33


func GenerateNewSkill(allSkills []*Skill) *Skill {
	primary, secondary, tertiary := TagSkills(allSkills)
	dp := delta(len(primary), len(allSkills)) / targetPrimarySkillDensity / 3 
	ds := delta(len(secondary), len(allSkills)) / targetSecondarySkillDensity / 3
	dt := delta(len(tertiary), len(allSkills)) / targetTertiarySkillDensity / 3
	decision := rand.Float64() 

	// 0 - - - - - - dp - - - - - - dp+ds - - - - - - dp+ds+dt == 1
	//   primary       secondary         tertiary
	
	if (decision < dp) { 
		// return newPrimarySkill()
	} else if (decision < dp+ds) {
		// return newSecondarySkill(primary)
	} else if (decision < dp+ds+dt) {
		// return newTertiarySkill(allSkills)
	}
	return nil
}

func delta(actual int, total int) float64 {
	return float64(total) / float64(actual) 
}

func TagSkills(allSkills []*Skill) (primary []*Skill, secondary []*Skill, tertiary []*Skill) { 
	productNameToSkill := make(map[string]*Skill)
	for _, skill := range(allSkills) {
		for _, productLine := range(skill.ProductLines) {
			productNameToSkill[productLine[0].Product.Name] = skill
		}
	}

	primarySkills := []*Skill{}
	secondarySkills := []*Skill{}
	tertiarySkills := []*Skill{}
	for _, skill := range(allSkills) {
		if (skill.isPrimary()) {
			primarySkills = append(primarySkills, skill)
		}
	}

	for _, skill := range(allSkills) { 
		if (!slices.Contains(primarySkills, skill)) {
			var candidatePrimarySource *Skill = nil 
			isSecondary := true
			for _, productLine := range(skill.ProductLines) {
				if !isSecondary {
					break
				}
				for _, recipe := range(productLine) {
					for component, _ := range(recipe.Components) {
						sourceSkill := productNameToSkill[component.Name]
						if candidatePrimarySource == nil {
							candidatePrimarySource = sourceSkill
						} else if candidatePrimarySource != sourceSkill { 
							isSecondary = false	
						}
					}
				}
			}
			if isSecondary {
				secondarySkills = append(secondarySkills, skill)
			} else { 
				tertiarySkills = append(tertiarySkills, skill)
			}
		}
	}
	return primarySkills, secondarySkills, tertiarySkills
}

func (skill *Skill) isPrimary() bool {
	if len(skill.ProductLines) > 1 {
		return false
	}
	for _, productLine := range(skill.ProductLines) {
		for _, recipe := range(productLine) {
			if len(recipe.Components) > 0 {
				return false	
			}
		}		
	}
	return true
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
		fmt.Println("`actual Skill` %v", sourceSkill)
		sourcePLCount := len(sourceSkill.ProductLines)
		fmt.Printf("############ %v ########", sourcePLCount) 
		sourceRecipeIndex := 0
		if (sourcePLCount > 0) {
			sourceRecipeIndex = rand.Intn(sourcePLCount)
		}
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
