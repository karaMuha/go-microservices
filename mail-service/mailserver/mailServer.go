package mailserver

import (
	"bytes"
	"html/template"
	"mailer/models"
	"os"
	"strconv"
	"time"

	"github.com/vanng822/go-premailer/premailer"
	simpleMail "github.com/xhit/go-simple-mail/v2"
)

type MailServer struct {
	Domain      string
	Host        string
	Port        int
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	FromName    string
}

func NewMailServer() *MailServer {
	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		panic(err)
	}

	return &MailServer{
		Domain:      os.Getenv("MAIL_DOMAIN"),
		Host:        os.Getenv("MAIL_HOST"),
		Port:        port,
		Username:    os.Getenv("MAIL_USERNAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		Encryption:  os.Getenv("MAIL_ENCRYPTION"),
		FromAddress: os.Getenv("MAIL_FROMADDRESS"),
		FromName:    os.Getenv("MAIL_FROMNAME"),
	}
}

func (m *MailServer) SendSMTPMail(mail *models.Mail) error {
	if mail.From == "" {
		mail.From = m.FromAddress
	}

	if mail.FromName == "" {
		mail.FromName = m.FromName
	}

	data := map[string]any{
		"message": mail.Message,
	}

	mail.DataMap = data

	formattedMessage, err := m.buildHTMLMail(mail)

	if err != nil {
		return err
	}

	mailServer := simpleMail.NewSMTPClient()
	mailServer.Host = m.Host
	mailServer.Port = m.Port
	mailServer.Username = m.Username
	mailServer.Password = m.Password
	mailServer.Encryption = m.getEncryption(m.Encryption)
	mailServer.KeepAlive = false
	mailServer.ConnectTimeout = 10 * time.Second
	mailServer.SendTimeout = 10 * time.Second

	smtpClient, err := mailServer.Connect()

	if err != nil {
		return err
	}

	email := simpleMail.NewMSG()
	email.SetFrom(mail.From).AddTo(mail.To).SetSubject(mail.Subject)
	email.SetBody(simpleMail.TextHTML, formattedMessage)

	err = email.Send(smtpClient)

	if err != nil {
		return err
	}

	return nil
}

func (m *MailServer) buildHTMLMail(mail *models.Mail) (string, error) {
	templateToRender := "./templates/mail.html.gohtml"

	t, err := template.New("email-html").ParseFiles(templateToRender)

	if err != nil {
		return "", err
	}

	var template bytes.Buffer
	err = t.ExecuteTemplate(&template, "body", mail.DataMap)

	if err != nil {
		return "", err
	}

	formattedMessage := template.String()
	formattedMessage, err = m.inlineCSS(formattedMessage)

	if err != nil {
		return "", err
	}

	return formattedMessage, nil
}

func (m *MailServer) inlineCSS(s string) (string, error) {
	options := premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}

	prem, err := premailer.NewPremailerFromString(s, &options)

	if err != nil {
		return "", err
	}

	html, err := prem.Transform()

	if err != nil {
		return "", err
	}

	return html, nil
}

func (m *MailServer) getEncryption(s string) simpleMail.Encryption {
	switch s {
	case "tls":
		return simpleMail.EncryptionSTARTTLS
	case "ssl":
		return simpleMail.EncryptionSSLTLS
	case "none", "":
		return simpleMail.EncryptionNone
	default:
		return simpleMail.EncryptionSTARTTLS
	}
}
