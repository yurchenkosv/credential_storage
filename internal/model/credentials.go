package model

import (
	"time"
)

type Metadata struct {
	Data string `json:"key"`
}

type BankingCard struct {
	Number         int       `json:"number"`
	ValidUntil     time.Time `json:"valid_until"`
	CardholderName string    `json:"cardholder_name"`
	CVV            int       `json:"cvv"`
}

type TextData struct {
	Data string
}

type BinaryData struct {
	Data []byte
}

type Credentials struct {
	Name     string       `json:"name"`
	Login    string       `json:"login,omitempty"`
	Password string       `json:"password,omitempty"`
	Card     *BankingCard `json:"card,omitempty"`
	Metadata []Metadata   `json:"metadata,omitempty"`
	Text     *TextData
	Binary   *BinaryData
}
