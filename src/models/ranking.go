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

func (r Ranking) TopTenSnake() ([]Ranking , error) {
	results := make([]Ranking , 0)

	rows, err := r.DB.Query("SELECT DISTINCT nome, pontos, email FROM RANKING WHERE ORIGEM = 1 ORDER BY pontos DESC LIMIT 10;")

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&r.Nome, &r.Email, &r.Pontos)
		if err != nil {
			panic(err)
		}
		results = append(results, r)
	}
	return results, nil
}

func (r Ranking) TopTenTetris() ([]Ranking , error) {
	results := make([]Ranking , 0)

	rows, err := r.DB.Query("SELECT DISTINCT nome, pontos, email FROM RANKING WHERE ORIGEM = 0 ORDER BY pontos DESC LIMIT 10;")

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&r.Nome, &r.Email, &r.Pontos)
		if err != nil {
			panic(err)
		}
		results = append(results, r)
	}
	return results, nil
}