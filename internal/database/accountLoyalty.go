package database

import (
	"database/sql"

	"github.com/gaspartv/GO-thygas-coins-back/internal/entity"
)

type AccLoyaltyDB struct {
	db *sql.DB
}

func NewAccLoyaltyDB(db *sql.DB) *AccLoyaltyDB {
	return &AccLoyaltyDB{
		db: db,
	}
}

func (database *AccLoyaltyDB) Create(accountLoyalty *entity.AccLoyalty) (string, error) {
	stmt, err := database.db.Prepare("INSERT INTO account_loyalty (id, percentage, price) VALUES (?,?,?)")
	if err != nil {
		return "", err
	}

	if _, err := stmt.Exec(accountLoyalty.ID, accountLoyalty.Percentage, accountLoyalty.Price); err != nil {
		return "", err
	}

	row := database.db.QueryRow("SELECT * FROM account_loyalty WHERE id = ?", accountLoyalty.ID)

	var accLoyalty entity.AccLoyalty
	if err := row.Scan(&accLoyalty.ID, &accLoyalty.Percentage, &accLoyalty.Price); err != nil {
		return "", err
	}

	return "created successful", nil
}

func (database *AccLoyaltyDB) Delete(id string) (string, error) {
	stmt, err := database.db.Prepare("DELETE FROM account_loyalty WHERE id =?")
	if err != nil {
		return "", err
	}

	if _, err := stmt.Exec(id); err != nil {
		return "", err
	}

	return "deleted successful", nil
}

func (database *AccLoyaltyDB) Update(id string, accountLoyalty *entity.AccLoyalty) (string, error) {
	stmt, err := database.db.Prepare("UPDATE account_loyalty SET percentage =?, price =? WHERE id =?")
	if err != nil {
		return "", err
	}

	if _, err := stmt.Exec(accountLoyalty.Percentage, accountLoyalty.Price, id); err != nil {
		return "", err
	}

	return "updated successful", nil
}

func (database *AccLoyaltyDB) Get(id string) (*entity.AccLoyalty, error) {
	stmt, err := database.db.Prepare("SELECT * FROM account_loyalty WHERE id =?")
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(id)

	var accLoyalty entity.AccLoyalty
	if err := row.Scan(&accLoyalty.ID, &accLoyalty.Percentage, &accLoyalty.Price); err != nil {
		return nil, err
	}

	return &accLoyalty, nil
}

func (database *AccLoyaltyDB) List() ([]entity.AccLoyalty, error) {
	stmt, err := database.db.Prepare("SELECT * FROM account_loyalty")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var accountLoyalties []entity.AccLoyalty
	for rows.Next() {
		var accLoyalty entity.AccLoyalty
		err := rows.Scan(&accLoyalty.ID, &accLoyalty.Percentage, &accLoyalty.Price)
		if err != nil {
			return nil, err
		}
		accountLoyalties = append(accountLoyalties, accLoyalty)
	}

	return accountLoyalties, nil
}
