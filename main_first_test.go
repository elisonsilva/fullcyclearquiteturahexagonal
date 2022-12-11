package main

import (
	"database/sql"

	db2 "github.com/codeedu/go-hexagonal/adapters/db"
	"github.com/codeedu/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	ProductService := application.NewProductService(productDbAdapter)
	product, _ := ProductService.Create("Product Exemple", 30)
	ProductService.Enable(product)

}
