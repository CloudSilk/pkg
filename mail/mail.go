package mail

import (
	"log"
	"time"

	"github.com/go-gomail/gomail"
)

func NewMailEngine(host string, port int, username, password string) *MailEngine {
	e := &MailEngine{
		messages: make(chan *gomail.Message, 100),
		d:        gomail.NewDialer(host, port, username, password),
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
	return e
}

func NewMessage(from string, to []string, ccAddress, ccName string, subject string, body string, attachFileName string) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	if ccAddress != "" {
		m.SetAddressHeader("Cc", ccAddress, ccName)
	}
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	if attachFileName != "" {
		m.Attach(attachFileName)
	}

	return m
}

type MailEngine struct {
	messages chan *gomail.Message
	d        *gomail.Dialer
	host     string
	port     int
	username string
	password string
}

func (e *MailEngine) AddSendMessage(m *gomail.Message) {
	e.messages <- m
}

func (e *MailEngine) Run() {
	var s gomail.SendCloser
	var err error
	open := false
	for {
		select {
		case m, ok := <-e.messages:
			if !ok {
				return
			}
			if !open {
				if s, err = e.d.Dial(); err != nil {
					panic(err)
				}
				open = true
			}
			if err := gomail.Send(s, m); err != nil {
				log.Print(err)
			}
		// Close the connection to the SMTP server if no email was sent in
		// the last 30 seconds.
		case <-time.After(30 * time.Second):
			if open {
				if err := s.Close(); err != nil {
					panic(err)
				}
				open = false
			}
		}
	}
}
