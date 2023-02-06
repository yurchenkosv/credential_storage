package model

type Metadata struct {
	ID    int
	Value string `json:"value"`
}

func (m Metadata) String() string {
	return m.Value
}

type BankingCardData struct {
	ID             int
	Name           string     `json:"name"`
	Number         string     `json:"number"`
	ValidUntil     string     `json:"valid_until"`
	CardholderName string     `json:"cardholder_name"`
	CVV            string     `json:"cvv"`
	Metadata       []Metadata `json:"metadata"`
}

type TextData struct {
	ID       int
	Name     string     `json:"name"`
	Data     string     `json:"data"`
	Metadata []Metadata `json:"metadata"`
}

type BinaryData struct {
	ID       int
	Name     string     `json:"name"`
	Link     string     `json:"link"`
	Metadata []Metadata `json:"metadata"`
}

type CredentialsData struct {
	ID       int
	Name     string     `json:"name"`
	Login    string     `json:"login"`
	Password string     `json:"password"`
	Metadata []Metadata `json:"metadata"`
}

type Credentials struct {
	ID              int
	Name            string           `json:"name"`
	CredentialsData *CredentialsData `json:"credentialsData,omitempty"`
	BankingCardData *BankingCardData `json:"bankingCardData,omitempty"`
	TextData        *TextData        `json:"textData,omitempty"`
	BinaryData      *BinaryData      `json:"binaryData,omitempty"`
	Metadata        []Metadata       `json:"metadata"`
}
