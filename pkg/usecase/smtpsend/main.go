package smtpsend

import (
	"bytes"
	"github.com/thisdougb/magiclink/config"
	"github.com/thisdougb/magiclink/pkg/entity/mlmsg"
	"html/template"
	"log"
	"net/smtp"
)

func (s *Service) SendMagicLink(to string, magicLinkURL string) error {

	var cfg *config.Config // dynamic config settings

	var input = struct {
		ToAddress    string
		MagicLinkURL string
	}{
		ToAddress:    to,
		MagicLinkURL: magicLinkURL,
	}

	request := mlmsg.NewMagicLinkMsg()

	templateFileName := config.GetTemplatePath("sendmagiclink.gohtml")
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, input); err != nil {
		return err
	}

	request.Body = buf.String()

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: Magic Link!\n"
	msg := []byte(subject + mime + "\n" + request.Body)

	auth := smtp.PlainAuth("", cfg.ValueAsStr("SMTP_USER"), cfg.ValueAsStr("SMTP_PASSWORD"), cfg.ValueAsStr("SMTP_HOST"))

	if err := smtp.SendMail(cfg.ValueAsStr("SMTP_HOST")+":"+cfg.ValueAsStr("SMTP_PORT"), auth, to, []string{to}, msg); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
