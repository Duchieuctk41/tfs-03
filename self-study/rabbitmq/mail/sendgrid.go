package email

import (
	"buoi11/db"
	"buoi11/model"
	"encoding/json"

	"github.com/sendgrid/sendgrid-go"
	mailHelper "github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sirupsen/logrus"
)

func SendMail(d []byte) {
	var mail model.EmailContent
	err := json.Unmarshal(d, &mail)
	if err != nil {
		logrus.Error("unmarshal fail: ", err)
	}

	from := mailHelper.NewEmail(mail.FromUser.Name, mail.FromUser.Email)
	subject := mail.Subject
	to := mailHelper.NewEmail(mail.ToUser.Name, mail.ToUser.Email)
	plantText := mail.PlainTextContent
	htmlContent := mail.HtmlContent
	message := mailHelper.NewSingleEmail(from, subject, to, plantText, htmlContent)
	client := sendgrid.NewSendClient("")
	res, err := client.Send(message)

	if err != nil {
		logrus.Error(err)
	} else {
		_, err := db.Connect().Exec("UPDATE ecommerce.orders SET thank_you_email_sent = true WHERE id = ?", mail.ID)
		if err != nil {
			logrus.Error(err)
		}
		logrus.Info(res)
	}
}
