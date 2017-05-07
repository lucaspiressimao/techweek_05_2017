package models

import (
	"database/sql"
	// "errors"
	_ "github.com/lib/pq"
	// "fmt"
)

type Ranking struct {
	DB       *sql.DB `json:"-"`
	Id       string  `json:"id"`
	Nome     string  `json:"nome"`
	Email	 string  `json:"email"`
	Pontos   string  `json:"pontos"`
}

func (r Ranking) Save(i int) (Ranking , error) {

	err := r.DB.QueryRow("INSERT INTO Ranking (nome, email, pontos, origem) VALUES($1, $2, $3, $4) RETURNING id", r.Nome, r.Email, r.Pontos, i).Scan(&r.Id)
	if err != nil {
		return r, err
	}
	return r, nil
}