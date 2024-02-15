package database

import (
	"database/sql"

	"github.com/gaspartv/GO-thygas-coins-back/internal/entity"
)

type StoreDB struct {
	db *sql.DB
}

func NewStoreDB(db *sql.DB) *StoreDB {
	return &StoreDB{
		db: db,
	}
}

func (database *StoreDB) Create(store *entity.Store) (string, error) {
	stmt, err := database.db.Prepare("INSERT INTO stores (id, name, qrcode, email, cellphone, password) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(store.ID, store.Name, store.QRCode, store.Email, store.Cellphone, store.Password); err != nil {
		return "", err
	}

	row := database.db.QueryRow("SELECT * FROM stores WHERE id =?", store.ID)

	var s entity.Store
	if err := row.Scan(&s.ID, &s.Name, &s.QRCode, &s.Email, &s.Cellphone, &s.Password); err != nil {
		return "", err
	}

	return "created successful", nil
}

func (database *StoreDB) Delete(id string) (string, error) {
	stmt, err := database.db.Prepare("DELETE FROM stores WHERE id =?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return "", err
	}

	return "deleted successful", nil
}

func (database *StoreDB) Update(id string, store *entity.Store) (string, error) {
	stmt, err := database.db.Prepare("UPDATE stores SET name =?, qrcode =?, email =?, cellphone =?, password =? WHERE id =?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(store.Name, store.QRCode, store.Email, store.Cellphone, store.Password, id); err != nil {
		return "", err
	}

	return "updated successful", nil
}

func (database *StoreDB) Get() (*entity.Store, error) {
	stmt, err := database.db.Prepare("SELECT * FROM stores")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow()

	var store entity.Store
	if err := row.Scan(&store.ID, &store.Name, &store.QRCode, &store.Email, &store.Cellphone, &store.Password); err != nil {
		return nil, err
	}

	return &store, nil
}

func (database *StoreDB) List() ([]entity.Store, error) {
	stmt, err := database.db.Prepare("SELECT * FROM stores")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var stores []entity.Store
	for rows.Next() {
		var store entity.Store
		err := rows.Scan(&store.ID, &store.Name, &store.QRCode, &store.Email, &store.Cellphone, &store.Password)
		if err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}

	return stores, nil
}
