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
	skills["primary-skill-1"] = &model.Skill{}
	skills["secondary-skill-2"] = &model.Skill{}
	return skills
}

func GetBoard() model.Board {
	board := model.NewGameBoard()
	tile1 := model.InitialTile(GetAllSkills())
	board.Tiles[&tile1] = model.NewPosition(0, 0)
	return board
}

func GetPlayer() model.Player { 
	return model.NewPlayer()
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
		`INSERT INTO recipe (skill_id, product_id, tier) VALUES (
			(SELECT id FROM skills WHERE name = '%v'), 
			(SELECT id FROM product WHERE name = '%v'),
			%v);`,skill, recipe.Product.Name, recipe.Product.Tier)
	
	fmt.Println("ADDING recipe by \n", query)
	fmt.Println()

	statement, err := db.Prepare(query)

	if (err != nil) {
		return err
	}

	statement.Exec()
	for product, count := range(recipe.Components) { 
		query := fmt.Sprintf(
			`INSERT INTO recipe_components (recipe_id, component_id, count, tier) VALUES (%v, %v, %v, %v)`,
			fmt.Sprintf(`
			(SELECT id FROM recipe r 
				WHERE r.skill_id = (SELECT id FROM skill s WHERE s.name = '%v')
				AND r.product_id = (SELECT id FROM product p WHERE p.name = '%v')
				AND r.tier = %v
			)`, skill, recipe.Product.Name, t), 
			fmt.Sprintf("(SELECT id FROM product c WHERE c.name = '%v')", product.Name),
			count,
			t)
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
			id INTEGER, skill_id INTEGER, product_id INTEGER, tier INTEGER
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

	statement, err= db.Prepare(`
		INSERT INTO skills (id, name) values (1, 'primary-skill-1'), (2, 'secondary-skill-2');
		INSERT INTO product (id, skill_id, name) values (1,1, 'ps1-p1');
		INSERT INTO player (id) values (1);
		INSERT INTO experience (player_id, skill_id, xp) VALUES (1, 1, 0);

	`)
	if err == nil {
		statement.Exec()
		log.Println("Inserted initial data")
	} else {  
		log.Fatal(err)
	}
}