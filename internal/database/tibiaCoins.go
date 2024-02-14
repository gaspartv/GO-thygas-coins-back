package database

import (
	"database/sql"

	"github.com/gaspartv/GO-thygas-coins-back/internal/entity"
)

type TibiaCoinsDB struct {
	db *sql.DB
}

func NewTibiaCoinsDB(db *sql.DB) *TibiaCoinsDB {
	return &TibiaCoinsDB{
		db: db,
	}
}

func (database *TibiaCoinsDB) Create(coins *entity.TibiaCoins) (string, error) {
	stmt, err := database.db.Prepare("INSERT INTO coins (id, code, name, price, amount, min, max, image, step) VALUES (?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return "", err
	}

	if _, err := stmt.Exec(coins.ID, coins.Code, coins.Name, coins.Price, coins.Amount, coins.Min, coins.Max, coins.Image, coins.Step); err != nil {
		return "", err
	}

	row := database.db.QueryRow("SELECT * FROM coins WHERE id =?", coins.ID)

	var c entity.TibiaCoins
	if err := row.Scan(&c.ID, &c.Code, &c.Name, &c.Price, &c.Amount, &c.Min, &c.Max, &c.Image, &c.Step); err != nil {
		return "", err
	}

	return "created successful", nil
}

func (database *TibiaCoinsDB) Delete(id string) (string, error) {
	stmt, err := database.db.Prepare("DELETE FROM coins WHERE id =?")
	if err != nil {
		return "", err
	}

	if _, err := stmt.Exec(id); err != nil {
		return "", err
	}

	return "deleted successful", nil
}

func (database *TibiaCoinsDB) Update(id string, coins *entity.TibiaCoins) (string, error) {
	stmt, err := database.db.Prepare("UPDATE coins SET code =?, name =?, price =?, amount =?, min =?, max =?, image =?, step =? WHERE id =?")
	if err != nil {
		return "", err
	}

	if _, err := stmt.Exec(coins.Code, coins.Name, coins.Price, coins.Amount, coins.Min, coins.Max, coins.Image, coins.Step, id); err != nil {
		return "", err
	}

	return "updated successful", nil
}

func (database *TibiaCoinsDB) Get(id string) (*entity.TibiaCoins, error) {
	stmt, err := database.db.Prepare("SELECT * FROM coins WHERE id =?")
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(id)

	var c entity.TibiaCoins
	if err := row.Scan(&c.ID, &c.Code, &c.Name, &c.Price, &c.Amount, &c.Min, &c.Max, &c.Image, &c.Step); err != nil {
		return nil, err
	}

	return &c, nil
}

func (database *TibiaCoinsDB) List() ([]entity.TibiaCoins, error) {
	stmt, err := database.db.Prepare("SELECT * FROM coins")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var coins []entity.TibiaCoins
	for rows.Next() {
		var c entity.TibiaCoins
		err := rows.Scan(&c.ID, &c.Code, &c.Name, &c.Price, &c.Amount, &c.Min, &c.Max, &c.Image, &c.Step)
		if err != nil {
			return nil, err
		}
		coins = append(coins, c)
	}

	return coins, nil
}
