package database

import (
	"database/sql"

	"github.com/gaspartv/GO-thygas-coins-back/internal/entity"
)

type PromotionDB struct {
	db *sql.DB
}

func NewPromotionDB(db *sql.DB) *PromotionDB {
	return &PromotionDB{
		db: db,
	}
}

func (database *PromotionDB) Create(promotion *entity.Promotion) (string, error) {
	stmt, err := database.db.Prepare("INSERT INTO promotions (id, description, min, max, price, stack) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(promotion.ID, promotion.Description, promotion.Min, promotion.Max, promotion.Price, promotion.Stack); err != nil {
		return "", err
	}

	row := database.db.QueryRow("SELECT * FROM promotions WHERE id =?", promotion.ID)

	var p entity.Promotion
	if err := row.Scan(&p.ID, &p.Description, &p.Min, &p.Max, &p.Price, &p.Stack); err != nil {
		return "", err
	}

	return "created successful", nil
}

func (database *PromotionDB) Delete(id string) (string, error) {
	stmt, err := database.db.Prepare("DELETE FROM promotions WHERE id =?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return "", err
	}

	return "deleted successful", nil
}

func (database *PromotionDB) Update(id string, promotion *entity.Promotion) (string, error) {
	stmt, err := database.db.Prepare("UPDATE promotions SET description =?, min =?, max =?, price =?, stack =? WHERE id =?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(promotion.Description, promotion.Min, promotion.Max, promotion.Price, promotion.Stack, id); err != nil {
		return "", err
	}

	return "updated successful", nil
}

func (database *PromotionDB) Get(id string) (*entity.Promotion, error) {
	stmt, err := database.db.Prepare("SELECT * FROM promotions WHERE id =?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var promotion entity.Promotion
	if err := row.Scan(&promotion.ID, &promotion.Description, &promotion.Min, &promotion.Max, &promotion.Price, promotion.Stack); err != nil {
		return nil, err
	}

	return &promotion, nil
}

func (database *PromotionDB) List() ([]entity.Promotion, error) {
	stmt, err := database.db.Prepare("SELECT * FROM promotions")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var promotions []entity.Promotion
	for rows.Next() {
		var promotion entity.Promotion
		err := rows.Scan(&promotion.ID, &promotion.Description, &promotion.Min, &promotion.Max, &promotion.Price, &promotion.Stack)
		if err != nil {
			return nil, err
		}
		promotions = append(promotions, promotion)
	}

	return promotions, nil
}
