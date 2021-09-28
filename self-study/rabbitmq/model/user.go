package model

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ShopUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
