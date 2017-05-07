package router

import(
	"net/http"
	"models"
	"fmt"
)

func (r Router) sevePontosTetris(w http.ResponseWriter, req *http.Request) {
	nome := req.FormValue("nome")
	email := req.FormValue("email")
	pontos := req.FormValue("pontos")

	rank := models.Ranking{Nome: nome, Email: email, Pontos: pontos, DB: r.DB}
	_, err := rank.Save()

	if err != nil {
		fmt.Println("Erro ao salvar os dados")
	}
		
	http.Redirect(w, req, "/tetris", http.StatusFound)
}


func principal(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusOK) // evita listagem das pastas
	// lista de status https://golang.org/pkg/net/http/
}

func servTetris(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "static/tetris/tetris.html")
}

func servSnake(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "static/tetris/tetris.html")
}