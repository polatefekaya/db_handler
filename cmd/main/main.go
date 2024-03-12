package main

import (
	"DatabaseHandler/pkg/data/entities"
	"DatabaseHandler/pkg/handlers"
	"DatabaseHandler/pkg/usecases"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	//var app config.AppConfig

	req := usecases.GetPlayerWithId(3)

	a := handlers.Convert(*req).(entities.PlayerEntity)
	a.FirstName = req.Responses[0].Player.FirstName
	log.Println(a.FirstName)

}

type Product struct {
	Name  string
	Price string
}

func dbCall() {

	var connStr string
	connStr = "postgres://postgres:root@localhost:5432/goptest?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createProductTable(db)
	product := Product{
		Name:  "Polat",
		Price: "55.98",
	}

	pk := insertProduct(db, product)
	fmt.Printf("ID = %d\n", pk)

	var name string
	var price string

	query := "SELECT name, price FROM product WHERE id = $1"
	err = db.QueryRow(query, pk).Scan(&name, &price)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s", name)
	fmt.Printf("price: %s", price)

}

// instead of this, use migrations
func createProductTable(db *sql.DB) {
	query := "CREATE TABLE IF NOT EXISTS product (id SERIAL PRIMARY KEY, name VARCHAR(100) NOT NULL, price NUMERIC(6,2) NOT NULL)"

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func insertProduct(db *sql.DB, product Product) int {
	query := "INSERT INTO product (name, price) VALUES ($1, $2) RETURNING id"

	var pk int
	err := db.QueryRow(query, product.Name, product.Price).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}
