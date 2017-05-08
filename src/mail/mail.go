package mail

import (
	"log"
	"net/mail"
	"encoding/base64"
	"net/smtp"
	"fmt"
)

type Email struct {
	From 		string
	To 			string
	ToName		string
	Subject		string
	Body 		string

	// Autenticação
	Email  		string
	Senha		string
	Smtp 		string
}

func (e Email) Enviar() (error){
	auth := smtp.PlainAuth(
		"",
		e.Email,
		e.Senha,
		e.Smtp)

	from := mail.Address{e.From, e.Email}
	to := mail.Address{e.ToName, e.To}

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = e.Subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	body := e.Body

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	err := smtp.SendMail(
		e.Smtp + ":587",
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(message),
	)
	if err != nil {
		log.Fatal(err)
	}
	return err
}