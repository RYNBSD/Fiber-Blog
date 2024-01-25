package lib

import (
	"net/smtp"
)

type IMail interface {
	Send()
}

type Mail struct {
	To      []string
	Subject string
	Html    string
}

func (m *Mail) Send() {
	const SMTP_HOST = "smtp.gmail.com"
	const SMTP_PORT = "587"
	const SMTP_SERVER = SMTP_HOST + ":" + SMTP_PORT

	message := []byte("Subject: " + m.Subject + "\r\n" + m.Html)
	auth := smtp.PlainAuth("", "", "", SMTP_HOST)

	err := smtp.SendMail(SMTP_SERVER, auth, "", m.To, message)
	if err != nil {
		panic(err)
	}
}
