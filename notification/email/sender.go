package email

import "gopkg.in/gomail.v2"

type SmtpSender struct {
	Host string
	Port int
}

func (s *SmtpSender) Send(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "no-reply@store.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(s.Host, s.Port, "", "")
	return d.DialAndSend(m)
}