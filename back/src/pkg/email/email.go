package email

import (
	"github.com/jordan-wright/email"
	"net/smtp"
	"valtec/pkg/config"
)

var fromEmail, fromName, authorizeCode, host, addr string

func init() {
	fromEmail = config.GetStringSevere("email", "fromEmail")
	authorizeCode = config.GetStringSevere("email", "authorizeCode")
	host = config.GetStringSevere("email", "host")
	addr = config.GetStringSevere("email", "addr")
	fromName = config.GetStringSevere("email", "fromName")
}
func Send(to, subject, html string) error {
	em := email.NewEmail()
	em.From = fromName + "<" + fromEmail + ">"
	em.To = []string{to}
	em.Subject = subject
	em.HTML = []byte(html)
	err := em.Send(addr, smtp.PlainAuth("", fromEmail, authorizeCode, host))
	return err
}
