package email

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

type Message struct {
	Header       string
	UsersName    string
	Introduction string
	Content      string
	URL          string
	Action       string
}

type Email struct {
	From    string
	To      string
	Subject string
	Body    string
}

func LoadEmail(header, username, introduction, content, url, action string) (bytes.Buffer, error) {
	// Load and parse the email template
	tmpl, err := template.ParseFiles("../static/index.html")
	if err != nil {
		_ = fmt.Errorf("error loading email template: %v", err)
	}

	data := Message{
		Header:       header,
		UsersName:    username,
		Introduction: introduction,
		Content:      content,
		URL:          url,
		Action:       action,
	}

	var Body bytes.Buffer

	err = tmpl.Execute(&Body, data)
	if err != nil {
		log.Fatal(err)
	}

	return Body, nil
}

func SendEmail(to, subject string, loadedMail bytes.Buffer) {
	envVars := []string{"SMTP_HOST", "SMTP_PORT", "SMTP_USER", "SMTP_PASS"}
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("environment variable %s not set", envVar)
		}
	}

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	smtpAuth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	msg := []byte(subject +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		loadedMail.String())

	err := smtp.SendMail(smtpHost+":"+smtpPort, smtpAuth, smtpUser, []string{to}, msg)
	if err != nil {
		_ = fmt.Errorf("smtp error while sending email: %s", err)
		return
	}
}
