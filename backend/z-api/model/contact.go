package model

// Contact representa a estrutura de um contato individual
type GetContacts struct {
	Name   string `json:"name"`   // Nome completo do contato
	Short  string `json:"short"`  // Nome curto/apelido do contato
	Notify string `json:"notify"` // Nome exibido no WhatsApp
	VName  string `json:"vname"`  // Nome no vCard
	Phone  string `json:"phone"`  // Número do telefone
}
