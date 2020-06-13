package modules

import (
	E "github.com/yuw-mvc/yuw/exceptions"
	"gopkg.in/gomail.v2"
)

type (
	MailPoT struct {
		Username string
		Password string
		Host string
		Port int
	}

	MailParamsPoT struct {
		From string
		To string
		Subject string
		Body string
	}

	Mail struct {
		d *gomail.Dialer
		cfg *MailPoT
	}
)

func NewMail(cfg *MailPoT) *Mail {
	return (&Mail{cfg:cfg}).initialized()
}

func (mail *Mail) initialized() *Mail {
	if mail.cfg != nil {
		mail.d = gomail.NewDialer(
			mail.cfg.Host,
			mail.cfg.Port,
			mail.cfg.Username,
			mail.cfg.Password,
		)
	}

	return mail
}

func (mail *Mail) Send(mpPoT []*MailParamsPoT) (err error) {
	messages, err := mail.messages(mpPoT)
	if err != nil {
		return
	}

	return mail.d.DialAndSend(messages ...)
}

func (mail *Mail) messages(mpPoT []*MailParamsPoT) (messages []*gomail.Message, err error) {
	if mail.d == nil {
		err = E.Err("yuw^m_email_a", E.ErrPosition())
		return
	}

	if len(mpPoT) == 0 {
		err = E.Err("yuw^m_email_b", E.ErrPosition())
		return
	}

	messages = []*gomail.Message{}

	for _, mParams := range mpPoT {
		msg := gomail.NewMessage()
		msg.SetHeader("From", mParams.From)
		msg.SetHeader("To", mParams.To)
		msg.SetHeader("Subject", mParams.Subject)
		msg.SetBody("text/html", mParams.Body)

		messages = append(messages, msg)
	}

	return
}


