package db 
import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
	"philipmstc/goscape/model"
)

func GetAllSkills() map[string]*model.Skill {
	skills := make(map[string]*model.Skill)
	primarySkill := &model.Skill{}
	primarySkill.ProductLines = [][]model.Recipe{}
	skills["primary-skill-1"] = primarySkill 
	secondarySkill := &model.Skill{}
	skills["secondary-skill-2"] = secondarySkill
	return skills
}

// func GetSkills() ([]*model.Skill, error) { 
// 	db, err := sql.Open("sqlite", "skills.db")
// 	if err != nil { 
// 		log.Fatal(err) 
// 		return nil, err
// 	}
// 	query := `SELECT s.name as "Skill", p.name as "Product", r.tier as "RecipeTier", pc.name as "Component", rc.count as "Count", rc.tier as "ComponentTier"
// 			  FROM skills s 
// 			  JOIN recipe r ON r.skill_id = s.id 
// 			  JOIN recipe_components rc on rc.recipe_id = r.id 
// 			  JOIN product p ON p.id = r.product_id 
// 			  JOIN product pc ON pc.id = rc.component_id`
	
// }

func GetBoard() model.Board {
	board := model.NewGameBoard()
	tile1 := model.InitialTile(GetAllSkills())
	board.Tiles[&tile1] = model.NewPosition(0, 0)
	return board
}

func GetPlayer() model.Player { 
	return model.NewPlayer()
}

func PersistSkills(skills map[string]*model.Skill) error { 
	for name, skill := range(skills) {
		if err := PersistSkill(skill, name); err != nil {
			return err
		}
	}
	return nil
}

func PersistSkill(skill *model.Skill, name string) error {
	db, err := sql.Open("sqlite", "skills.db")
	if err != nil {
		log.Fatal(err)
		return err
	}
	query := fmt.Sprintf("INSERT INTO skills (name) VALUES ('%v');", name)
	statement, err := db.Prepare(query)
	if (err != nil) {
		log.Println("Error adding new skill")
		return err
	} else { 
		log.Println("New skill added to database")
	}
	statement.Exec()

	for _, productLine := range(skill.ProductLines) { 
		if err := PersistProductLine(name, productLine); err != nil { 
			return err
		}
	}

	return nil	
}

func PersistProductLine(skill string, productLine []model.Recipe) error { 
	db, err := sql.Open("sqlite", "skills.db")
	if err != nil {
		log.Fatal(err)
		return err 
	}
	query := fmt.Sprintf("INSERT INTO product (name, tiers) VALUES ('%v', %v);", productLine[0].Product.Name, len(productLine))
	statement, err := db.Prepare(query)
	if err != nil {
		return err 
	}
	statement.Exec()
	for t, recipe := range(productLine) {
		if err := PersistRecipe(skill, t, recipe); err != nil { 
			return err
		}
	}
	return nil
}

func PersistRecipe(skill string, t int, recipe model.Recipe) error { 
	db, err := sql.Open("sqlite", "skills.db")
	if err != nil { 
		log.Fatal(err)
		return err
	}
	query := fmt.Sprintf(
		`INSERT INTO recipe (skill_id, product_id, tier) 
		 SELECT s.id, p.id, %v 
		 FROM skills s 
		 JOIN product p on p.name = '%v' 
		 WHERE s.name = '%v';`,recipe.Product.Tier, recipe.Product.Name, skill)
	
	fmt.Println("ADDING recipe by \n", query)
	fmt.Println()

	statement, err := db.Prepare(query)

	if (err != nil) {
		return err
	}

	statement.Exec()
	for product, count := range(recipe.Components) { 
		query := fmt.Sprintf(
			`INSERT INTO recipe_components (recipe_id, component_id, count, tier)
			SELECT r.id, p.id, %v, %v 
			FROM recipe r 
			JOIN product p ON p.id = r.product_id 
			join skills s on s.id = r.skill_id 
			WHERE p.name = '%v' 
			AND s.name = '%v'`,count, t, product.Name, skill)
		
		fmt.Println("ADDING component by \n", query)
		fmt.Println()
		if statement, err := db.Prepare(query); err != nil {
			statement.Exec()
		} else { 
			return err
		}
	}
	return nil
}

func InitDb() {
	db, err := sql.Open("sqlite", "skills.db")
	if err != nil {
		log.Fatal(err)
	}
	statement, err := db.Prepare(
		`CREATE TABLE IF NOT EXISTS skills (
			id INTEGER PRIMARY KEY, name TEXT
		);
		CREATE TABLE IF NOT EXISTS product (
			id INTEGER PRIMARY KEY, name TEXT, tiers INTEGER
		);
		CREATE TABLE IF NOT EXISTS player (
			id INTEGER PRIMARY KEY
		);
		CREATE TABLE IF NOT EXISTS experience (
			player_id INTEGER, skill_id INTEGER, xp BIGINT
		);
		CREATE TABLE IF NOT EXISTS recipe (
			id INTEGER PRIMARY KEY, skill_id INTEGER, product_id INTEGER, tier INTEGER
		);
		CREATE TABLE IF NOT EXISTS recipe_components (
			recipe_id INTEGER, component_id INTEGER, count INTEGER, tier INTEGER
		);
		`,
	)
	if err != nil {
		log.Println("Error creating tables")
	} else {
		log.Println("Created tables")
	}
	statement.Exec()
}