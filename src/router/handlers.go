package router

import(
	"net/http"
	"models"
	"fmt"
	"text/template"
	"mail"
)

func (r Router) sevePontosTetris(w http.ResponseWriter, req *http.Request) {
	nome := req.FormValue("nome")
	email := req.FormValue("email")
	pontos := req.FormValue("pontos")

	rank := models.Ranking{Nome: nome, Email: email, Pontos: pontos, DB: r.DB}
	_, err := rank.Save(0)

	if err != nil {
		fmt.Println("Erro ao salvar os dados")
	}
		
	emailEnv := mail.Email{
		From: "Lucas Pires", 
		To: email,
		ToName: nome, 
		Subject: nome + " obrigado por jogar",
		Body: ("Obrigado por jogar Tetris. Sua pontuação foi de " + pontos),
		Email: "lucaspiressimao@gmail.com",
		Senha: "rqsvwmvnixowkbxu",
		Smtp: "smtp.gmail.com"}

	go emailEnv.Enviar()

	http.Redirect(w, req, "/tetris", http.StatusFound)
}

func (r Router) sevePontosSnake(w http.ResponseWriter, req *http.Request) {
	nome := req.FormValue("nome")
	email := req.FormValue("email")
	pontos := req.FormValue("pontos")

	rank := models.Ranking{Nome: nome, Email: email, Pontos: pontos, DB: r.DB}
	_, err := rank.Save(1)

	if err != nil {
		fmt.Println("Erro ao salvar os dados")
	}
		
	emailEnv := mail.Email{
		From: "Lucas Pires", 
		To: email,
		ToName: nome, 
		Subject: nome + " obrigado por jogar",
		Body: ("Obrigado por jogar Snake. Sua pontuação foi de " + pontos),
		Email: "lucaspiressimao@gmail.com",
		Senha: "rqsvwmvnixowkbxu",
		Smtp: "smtp.gmail.com"}

	go emailEnv.Enviar()

	http.Redirect(w, req, "/snake", http.StatusFound)
}


func (r Router) principal(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusOK) // evita listagem das pastas
	// lista de status https://golang.org/pkg/net/http/
}

func (r Router) servTetris(w http.ResponseWriter, req *http.Request) {
	ranking := models.Ranking{DB: r.DB}
	result, err := ranking.TopTenTetris()

	w.Header().Set("Content-Type", "text/html")
	
	tpl, err := template.ParseFiles("static/tetris/tetris.gohtml")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tpl.Execute(w, result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (r Router) servSnake(w http.ResponseWriter, req *http.Request) {
	ranking := models.Ranking{DB: r.DB}
	result, err := ranking.TopTenSnake()

	w.Header().Set("Content-Type", "text/html")
	
	tpl, err := template.ParseFiles("static/snake/snake.gohtml")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tpl.Execute(w, result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}