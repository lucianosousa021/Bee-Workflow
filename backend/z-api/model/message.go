package model

type WhatsAppMessage struct {
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

type WhatsAppMessageWithImage struct {
	Phone   string `json:"phone"`
	Image   string `json:"image"`
	Caption string `json:"caption,omitempty"`
}

type WhatsAppMessageWithAudio struct {
	Phone string `json:"phone"`
	Audio string `json:"audio"`
}

type WhatsAppMessageWithDocument struct {
	Phone    string `json:"phone"`
	Document string `json:"document"`
}

type WhatsAppMessageWithContact struct {
	Phone        string `json:"phone"`
	ContactName  string `json:"contactName"`
	ContactPhone string `json:"contactPhone	"`
}

type CreateContact struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
}

type SendPixButton struct {
	Phone  string `json:"phone"`
	PixKey string `json:"pixKey"`
	Type   string `json:"type"` // Tipo da chave pix (CPF, CNPJ, PHONE, EMAIL, EVP)
}

type ReadMessage struct {
	MessageId string `json:"messageId"`
	Phone     string `json:"phone"`
}
