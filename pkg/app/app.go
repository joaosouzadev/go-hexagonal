package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const dbUrl = "root:admin@tcp(mysql_main:3306)/hexagonal?parseTime=true"

type App struct {
	DBConn *sql.DB
}

func NewApp() *App {
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	return &App{
		DBConn: db,
	}
}
