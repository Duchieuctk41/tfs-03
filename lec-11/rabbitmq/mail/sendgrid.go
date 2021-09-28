package mail

import (
	"errors"
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Mailer interface {
	Send(*EmailContent) error
}

type EmailUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type EmailContent struct {
	ID               int64      `json:"id"`
	Subject          string     `json:"subject"`
	FromUser         *EmailUser `json:"from"`
	ToUser           *EmailUser `json:"to"`
	PlainTextContent string     `json:"plaintext_content"`
	HtmlContent      string     `json:"html_content"`
}

// validate will check whether the email content is valid
func (em *EmailContent) Validate() error {
	if em == nil || em.FromUser == nil || em.ToUser == nil || em.PlainTextContent == "" {
		return errors.New("wrong content")
	}
	return nil
}

// NewSendGrid creates new Sendgrid client
func NewSendGrid(apiKey string) *Sendgrid {
	client := sendgrid.NewSendClient(apiKey)
	return &Sendgrid{
		ApiKey: apiKey,
		Client: client,
	}
}

type Sendgrid struct {
	ApiKey string `json:"api_key"`
	Client *sendgrid.Client
}

// Send will send email based on email content
func (m *Sendgrid) Send(em *EmailContent) error {
	if err := em.Validate(); err != nil {
		return err
	}

	from := mail.NewEmail(em.FromUser.Name, em.FromUser.Email)
	subject := em.Subject
	to := mail.NewEmail(em.ToUser.Name, em.ToUser.Email)
	plainTextContent := em.PlainTextContent
	htmlContent := em.HtmlContent
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	response, err := m.Client.Send(message)
	if err != nil {
		fmt.Println("Cannot send email due to error: ", err)
		return err
	}
	fmt.Printf("Email sent with response code: %v, response body: %v, response headers: %v\n", response.StatusCode, response.Body, response.Headers)
	return nil
}
