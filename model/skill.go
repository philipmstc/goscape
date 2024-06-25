package model

import (
    "fmt"
    "math/rand"
)

type Inventory struct {
    items map[string]int
}

type Skill struct { 
    ProductLines []Recipe
}

func PrimaryResource(name string, tiers int) []Recipe {
    recipes := []Recipe{}
    for i := range(tiers) {
        product := Product{name, i+1}
        recipes = append(recipes, Recipe{nil, product})
    }
    return recipes
}

type Product struct { 
    name string
    tier int
}
func (this Product) String() string { 
    return fmt.Sprintf("%v-%v", this.name, this.tier)
}

type Recipe struct { 
    components map[Product]int
    product Product
}

func (this Recipe) String() string {
    if this.components != nil {
        var comps string = ""
        for k, v := range(this.components) {
            comps = comps + fmt.Sprintf("%v*%v + ", v, k)
        }
        comps = comps[:len(comps)-2]
        return fmt.Sprintf("Make %v from (%v)", this.product, comps)
    } else {
        return fmt.Sprintf("(Free) %v", this.product)
    }
}

type Tile struct { 
    actions []Action
    x int
    y int
}

type Action interface {
    Do()
    GetName() string
}

type MakeProduct struct {
    recipe Recipe
    xpGain int
}

func (mk MakeProduct) Do() {
    fmt.Printf("xp: %v recipe: %v"  , mk.xpGain, mk.recipe)
}

func (mk MakeProduct) GetName() string { 
    return fmt.Sprintf("Make [%v] for %v xp", mk.recipe.product.name, mk.xpGain) 
}

// only works if theres only one recipe per product
func (this Inventory) CanCreate(product Product, recipeBook []Recipe) bool {
    for _, r := range recipeBook {
        if (r.product.name == product.name) { 
            for k,v := range r.components {
                if (this.items[k.name] < v) { 
                    return false
                }
            }
            return true
        }
    }
    return false
}

func (this Inventory) Create(product Product, recipe Recipe) {
    for k, v := range recipe.components {
        this.items[k.name] -= v
    }
    this.items[recipe.product.name] += 1
}

func GenerateProductLineNM(name string, skills []Skill, minSkills int, maxSkills int) []Recipe {
    const NEW_SKILL_RARITY_FACTOR = 2
    t := rand.Intn(len(skills) * NEW_SKILL_RARITY_FACTOR + 1)
    var targetSkill Skill
    if t == (len(skills)*NEW_SKILL_RARITY_FACTOR) { 
        targetSkill = Skill{[]Recipe{}}
    } else {
        targetSkill = skills[t % len(skills)]
    }
    comps := make(map[Product]int)
    var actualSkills int
    if minSkills >= maxSkills {
        actualSkills = minSkills
    } else { 
        actualSkills = minSkills + rand.Intn(maxSkills-minSkills)
    }
    fmt.Printf("Making product line from %v 'skills'\n", actualSkills)
    for i := 0; i <= actualSkills; i++ {
        var s int
        for {
            s = rand.Intn(len(skills))
            if s != ((t)%len(skills)) { 
                break;
            }
        }
        sourceSkill := skills[s]
        sourceRecipeIndex := rand.Intn(len(sourceSkill.ProductLines))
        sourceRecipe := sourceSkill.ProductLines[sourceRecipeIndex]
        _, ok := comps[sourceRecipe.product]
        if ok {
            comps[sourceRecipe.product] += 1
        } else {
            comps[sourceRecipe.product] = 1
        }
    }
    fmt.Println("Target Skill = ", targetSkill)
    return []Recipe{Recipe{comps, Product{name, 1}}}
}




//func main() {
//    p1 := Product{"ore", 1}
//    p2 := Product{"bar", 1}
//
//    p1_2 := Product{"ore", 2}
//    p2_2 := Product{"bar", 2}
//
//    comps := make(map[Product]int)
//    comps[p1] = 1
//    smeltBar := Recipe{comps, p2}
//
//
//    comps2 := make(map[Product]int)
//    comps2[p1_2] = 1
//    smeltBar2 := Recipe{comps2, p2_2}
//
//    otherRec := Recipe{map[Product]int{Product{"other", 1} : 1}, Product{"result", 1}}
//    otherRec2 := Recipe{map[Product]int{Product{"autre", 1} : 1}, Product{"resultat", 1}}
//
//    _inv := make(map[string]int)
//    _inv["ore"] = 1
//    inv := Inventory{_inv}
//
//    fmt.Println("can make bar?")
//    fmt.Println(inv.canCreate(p2, []Recipe{smeltBar}))
//    inv.create(p2, smeltBar)
//    fmt.Println("can do it again?")
//    fmt.Println(inv.canCreate(p2, []Recipe{smeltBar}))
//    inv.create(p2, smeltBar)
//    mp := MakeProduct{smeltBar, 20}
//    mp.do()
//    fmt.Println(mp.getName())
//
//    s1 := []Skill{Skill{[]Recipe{smeltBar, otherRec}}, Skill{[]Recipe{smeltBar2, otherRec2}}}
//    fmt.Printf("New product line: %v", generateProductLineNM(s1, 1,3))
//
//}

