package emailns

import (
	"github.com/scorredoira/email"
	"github.com/svarlamov/bintrad/config"
	"net/mail"
)

func SendEmail(toStr string, subject string, htmlBody string) error {
	m := email.NewHTMLMessage(subject, htmlBody)
	m.From = mail.Address{Name: config.Conf.SMTPCreds.DisplayName, Address: config.Conf.SMTPCreds.EmailAddress}
	m.To = []string{toStr}
	return email.Send(config.Conf.SMTPCreds.ConnectionString,
		config.Conf.SMTPAuth, m)
}
