package models

type RegisterResponse struct {
	Token  string `json:"token"`
	Status int    `json:"status"`
}
