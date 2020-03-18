package utils

import (
	"github.com/alexcesaro/mail/mailer"
	"log"
	"net/mail"
	"strings"
)

func SendMail(subject string, content string, receiver []string) {
	msg := &mail.Message{
		mail.Header{
			"From":         {"1343785360@qq.com"},
			"To":           receiver,
			"Subject":      {subject},
			"Content-Type": {"text/plain"},
		},
		strings.NewReader(content),
	}

	m := mailer.NewMailer("smtp.qq.com", "1343785360@qq.com", "odngbvmihvmuhcjc", 25)
	err := m.Send(msg)
	if err != nil { // This will send the email to Bob and Cora
		log.Println(err)
	}
}
