package mail

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sendgrid/sendgrid-go"
)

type Mailer interface {
	Send(*EmailContent) error
}

type EmailUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (eu *EmailUser) String() string {
	b, _ := json.Marshal(eu)
	return string(b)
}

type EmailContent struct {
	Id               int64      `json:"id"`
	Subject          string     `json:"subject"`
	FromUser         *EmailUser `json:"from"`
	ToUser           *EmailUser `json:"to"`
	PlainTextContext string     `json:"plaintext_content"` // body content
	HtmlContent      string     `json:"html_content"`
}

func (ec *EmailContent) Validate() error {
	if ec == nil || ec.FromUser == nil || ec.ToUser == nil || ec.PlainTextContext == "" {
		return errors.New("wrong content email")
	}
	return nil
}

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

func (m *Sendgrid) Send(ec *EmailContent) error {
	if err := ec.Validate(); err != nil {
		return err
	}

	fmt.Println("Sending email, infor: ", ec)
	return nil
}
