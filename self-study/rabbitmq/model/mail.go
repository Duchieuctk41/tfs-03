package model

type EmailContent struct {
	ID               int64    `json:"id"`
	Subject          string   `json:"subject"`
	PlainTextContent string   `json:"plantext_content"`
	HtmlContent      string   `json:"html_content"`
	FromUser         ShopUser `json:"from"`
	ToUser           User     `json:"to"`
}
