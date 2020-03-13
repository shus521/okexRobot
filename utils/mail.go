package utils

import (
	"github.com/alexcesaro/mail/mailer"
	"log"
	"net/mail"
	"strings"
)

func Send(content string, receiver []string) {
	msg := &mail.Message{
		mail.Header{
			"From":         {"1343785360@qq.com"},
			"To":           receiver,
			"Subject":      {"变价提醒!"},
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
