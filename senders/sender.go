package senders

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
)

type smtpTemplateData struct {
	From    string
	To      []string
	Subject string
	Message string
}

// SendEmail is the function to send email for a list
func SendEmail(emailList []string) error {
	from := "email"
	pass := "senha"
	to := emailList

	auth := smtp.PlainAuth(
		"",
		from,
		pass,
		"smtp.gmail.com",
	)

	context := &smtpTemplateData{
		From:    from,
		To:      to,
		Subject: "My First App - Email",
		Message: "Hello World",
	}

	var buffer bytes.Buffer

	template := template.Must(template.New("emailTemplate").Parse(emailScript()))
	template.Execute(&buffer, context)
	log.Println(buffer.String())
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, buffer.Bytes())
	if err != nil {
		log.Fatal("Send email error:", err.Error())
		return err
	}
	return nil
}

func emailScript() string {
	return `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}
MIME-version: 1.0
Content-Type: text/html; charset="UTF-8"
<html>
<body>
<h1 style="text-align: center">My First App - Email</h1>
<h2 style="text-align: center">{{.Message}}</h2>
</body>
</html>`
}
