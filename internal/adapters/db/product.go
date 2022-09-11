package db

import (
	"database/sql"
	"github.com/joaosouzadev/go-hexagonal-arch/internal/application"
	"log"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (a *ProductDb) Get(uuid string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := a.db.Prepare("SELECT id, uuid, name, price, active, on_stock FROM product WHERE uuid = ?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(uuid).Scan(&product.ID, &product.Uuid, &product.Name, &product.Price, &product.Active, &product.OnStock)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (a *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var uuid string
	stmt := `SELECT uuid FROM product where uuid = ?`
	a.db.QueryRow(stmt, product.GetUuid()).Scan(&uuid)

	if len(uuid) == 0 {
		_, err := a.create(product)
		if err != nil {
			return nil, err
		}
		return product, nil
	}

	_, err := a.update(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (a *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt := `INSERT INTO product (uuid, name, price, active, on_stock) VALUES (?, ?, ?, ?, ?)`
	_, err := a.db.Exec(stmt, product.GetUuid(), product.GetName(), product.GetPrice(), product.IsActive(), product.GetOnStock())
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (a *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	stmt := `UPDATE product set name = ?, price = ?, active = ?, on_stock = ? where uuid = ?`
	_, err := a.db.Exec(stmt, product.GetName(), product.GetPrice(), product.IsActive(), product.GetOnStock(), product.GetUuid())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return product, nil
}
