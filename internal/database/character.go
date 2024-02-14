package database

import (
	"database/sql"

	"github.com/gaspartv/GO-thygas-coins-back/internal/entity"
)

type CharacterDB struct {
	db *sql.DB
}

func NewCharacterDB(db *sql.DB) *CharacterDB {
	return &CharacterDB{
		db: db,
	}
}

func (database *CharacterDB) Create(character *entity.Character) (string, error) {
	stmt, err := database.db.Prepare("INSERT INTO characters (id, vocation, level, world, description) VALUES (?,?,?,?,?)")
	if err != nil {
		return "", err
	}

	if _, err := stmt.Exec(character.ID, character.Vocation, character.Level, character.World, character.Description); err != nil {
		return "", err
	}

	row := database.db.QueryRow("SELECT * FROM characters WHERE id =?", character.ID)

	var char entity.Character
	if err := row.Scan(&char.ID, &char.Vocation, &char.Level, &char.World, &char.Description); err != nil {
		return "", err
	}

	return "created successful", nil
}

func (database *CharacterDB) Delete(id string) (string, error) {
	stmt, err := database.db.Prepare("DELETE FROM characters WHERE id =?")
	if err != nil {
		return "", err
	}

	if _, err := stmt.Exec(id); err != nil {
		return "", err
	}

	return "deleted successful", nil
}

func (database *CharacterDB) Update(id string, character *entity.Character) (string, error) {
	stmt, err := database.db.Prepare("UPDATE characters SET vocation =?, level =?, world =?, description =? WHERE id =?")
	if err != nil {
		return "", err
	}

	if _, err := stmt.Exec(character.Vocation, character.Level, character.World, character.Description, id); err != nil {
		return "", err
	}

	return "updated successful", nil
}

func (database *CharacterDB) Get(id string) (*entity.Character, error) {
	stmt, err := database.db.Prepare("SELECT * FROM characters WHERE id =?")
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(id)

	var character entity.Character
	if err := row.Scan(&character.ID, &character.Vocation, &character.Level, &character.World, &character.Description); err != nil {
		return nil, err
	}

	return &character, nil
}

func (database *CharacterDB) List() ([]entity.Character, error) {
	stmt, err := database.db.Prepare("SELECT * FROM characters")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var characters []entity.Character
	for rows.Next() {
		var character entity.Character
		err := rows.Scan(&character.ID, &character.Vocation, &character.Level, &character.World, &character.Description)
		if err != nil {
			return nil, err
		}
		characters = append(characters, character)
	}

	return characters, nil
}
