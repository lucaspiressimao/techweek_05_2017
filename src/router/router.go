package router

import (
	"database/sql"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Handlers func(w http.ResponseWriter, req *http.Request)

type Router struct {
	DB *sql.DB
}

func (r Router) CreateRoutes() *mux.Router {

	router := mux.NewRouter().StrictSlash(true) //StrictSlash -> /tetris, /tetris/
	fmt.Printf("Criando rotas \n")
	m := map[string]map[string]Handlers{
		"GET": {
			"/":			principal, // mesmo não servindo arquivos
			"/tetris":		servTetris,
			"/snake":		servSnake,
		},
		"POST": {
			"/tetrisSave": 		r.sevePontosTetris,
			// "/snake":		sevePontosSnake,
		},
	}

	for method, routes := range m {
		for route, function := range routes {
			router.HandleFunc(route, function).Methods(method)
		}
	}
	return router
}