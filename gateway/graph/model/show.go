package model

type Show struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	UserID      string  `json:"userId"`
}
