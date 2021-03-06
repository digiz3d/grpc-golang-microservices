// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AddCommentInput struct {
	Text   string `json:"text"`
	ShowID string `json:"showId"`
	UserID string `json:"userId"`
}

type AddCommentPayload struct {
	Comment *Comment `json:"comment"`
}

type Comment struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	ShowID string `json:"showId"`
	Text   string `json:"text"`
}

type CreateShowInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	UserID      string  `json:"userId"`
}

type CreateShowPayload struct {
	Show *Show `json:"show"`
}

type CreateUserInput struct {
	Username string `json:"username"`
}

type CreateUserPayload struct {
	User *User `json:"user"`
}

type OnCommentAddedInput struct {
	ShowID string `json:"showId"`
}

type OnCommentAddedPayload struct {
	Comment *Comment `json:"comment"`
}
