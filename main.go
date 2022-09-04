package main

import (
	"database/sql"
	db2 "github.com/joaosouzadev/go-hexagonal-arch/adapters/db"
	"github.com/joaosouzadev/go-hexagonal-arch/application"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const dbUrl = "root:admin@tcp(mysql_main:3306)/hexagonal?parseTime=true"

func main() {
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, err := productService.Create("test", 500)
	if err != nil {
		log.Fatal(err)
	}

	_, err = productService.Enable(product)
}
