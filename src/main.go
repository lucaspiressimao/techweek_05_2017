package main

import (
	"database/sql"
	_ "github.com/lib/pq"
  	"net/http"
  	"router"
  	"fmt"
)

func setupDB() *sql.DB {
	db, err := sql.Open("postgres", "dbname=fatec user=fatec password=123 host=localhost port=5432")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	r := router.Router{DB: setupDB()} // criando conexão com o banco de dados

	router := r.CreateRoutes() // criando as rotas
	
	fmt.Println("Iniciando Servidor")
	http.ListenAndServe(":3000", router)
}