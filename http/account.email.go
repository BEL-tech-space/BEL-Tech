package http

import (
	"beltech/frontend"
	"bytes"
	"fmt"
	"html/template"
	"io"

	gomail "gopkg.in/mail.v2"
)

func (a *App) sendGreetingEmail(email string) error {

	var (
		from = a.config.Mail.EmailAcc
		tpl  bytes.Buffer
	)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Beltech - Welcome to the platform")

	if err := prepareMail(&tpl); err != nil {
		return err
	}

	// fmt.Println(email, a.config.Mail.EmailSmtpServerHost, a.config.Mail.EmailSmtpServerPort, from, a.config.Mail.EmailPswd)
	data, err := readPNGAsBytes("public/assets/images/logo.png")
	if err != nil {
		return err
	}

	m.Attach("beltech.png", gomail.SetHeader(map[string][]string{
		"Content-ID": {"BeltechLogo"},
	}), gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(data)
		return err
	}))

	m.SetBody("text/html", tpl.String())

	d := gomail.NewDialer(a.config.Mail.EmailSmtpServerHost, a.config.Mail.EmailSmtpServerPort, from, a.config.Mail.EmailPswd)
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email, err at dial: %w", err)
	}

	return nil
}

func prepareMail(mailBuf *bytes.Buffer) error {

	t, err := template.ParseFS(frontend.EmailHTML, "pages/account/mail.html")
	if err != nil {
		return fmt.Errorf("failed to parse email template: %w", err)
	}

	if err := t.Execute(mailBuf, nil); err != nil {
		return fmt.Errorf("failed to execute email template: %w", err)
	}

	return nil
}