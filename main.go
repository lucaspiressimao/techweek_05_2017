package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "fmt"
)

func servTetris(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "static/tetris/tetris.html")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	
	fmt.Println("Criando Rotas")
	router.HandleFunc("/tetris", servTetris) // GET por default
	//router.HandleFunc("/tetris", servTetris).Methods("GET") mas pode especificar o metodo

	fmt.Println("Iniciando Servidor")
	http.ListenAndServe(":3000", router)
}