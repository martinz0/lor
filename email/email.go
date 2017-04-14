package email

import (
	"github.com/go-gomail/gomail"
	"github.com/martinz0/lor/etc"
)

func Send(to, subject, contentType, content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", etc.String("email.user"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody(contentType, content)

	var (
		smtpDomain = etc.String("email.smtp_domain")
		smtpPort   = etc.Int("email.smtp_port")
		user       = etc.String("email.user")
		pass       = etc.String("email.pass")
	)
	d := gomail.NewDialer(smtpDomain, smtpPort, user, pass)
	return d.DialAndSend(m)
}
