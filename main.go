package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./wire.db")
	if err != nil {
		panic(err)
	}

	usecase := NewProductUsecase(db)
	product, _ := usecase.Execute(2)
	fmt.Println(product.Name)
}
