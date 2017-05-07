package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "fmt"
)

type Handlers func(w http.ResponseWriter, req *http.Request)

func CreateRoutes() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	fmt.Printf("Criando rotas \n")
	m := map[string]map[string]Handlers{
		"GET": {
			"/":			principal, // mesmo não servindo arquivos
			"/tetris":		servTetris,
			"/snake":		servSnake,
		},
		// "POST": {
		// 	"/foo": 		foo,
		// 	"/bar":			bar,
		// }
	}

	for method, routes := range m {
		for route, function := range routes {
			router.HandleFunc(route, function).Methods(method)
		}
	}
	return router
}

func principal(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusOK) // evita listagem das pastas
}

func servTetris(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "static/tetris/tetris.html")
}

func servSnake(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "static/tetris/tetris.html")
}

func main() {
	router := CreateRoutes();
	
	fmt.Println("Iniciando Servidor")
	http.ListenAndServe(":3000", router)
}