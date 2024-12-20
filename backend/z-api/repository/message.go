package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"zapi/model"
)

func ZAPISendMessage(message, phone, userInstance, userToken, accountToken string) error {
	url := fmt.Sprintf("https://api.z-api.io/instances/%s/token/%s/send-text", userInstance, userToken)

	payload := model.WhatsAppMessage{
		Phone:   phone,
		Message: message,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Erro ao criar JSON:", err)
		return err
	}

	payloadReader := strings.NewReader(string(jsonData))

	req, err := http.NewRequest("POST", url, payloadReader)
	if err != nil {
		return err
	}

	req.Header.Add("client-token", accountToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}

func ZAPISendMessageWithImage(caption, imageBase64, phone, userInstance, userToken, accountToken string) error {
	url := fmt.Sprintf("https://api.z-api.io/instances/%s/token/%s/send-image", userInstance, userToken)

	payload := model.WhatsAppMessageWithImage{
		Phone: phone,
		Image: imageBase64,
	}

	if caption != "" {
		payload.Caption = caption
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Erro ao criar JSON:", err)
		return err
	}

	payloadReader := strings.NewReader(string(jsonData))

	req, err := http.NewRequest("POST", url, payloadReader)
	if err != nil {
		return err
	}

	req.Header.Add("client-token", accountToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	return nil
}

func ZAPISendAudioMessage(audioBase64, phone, userInstance, userToken, accountToken string) error {
	url := fmt.Sprintf("https://api.z-api.io/instances/%s/token/%s/send-audio", userInstance, userToken)

	payload := model.WhatsAppMessageWithAudio{
		Phone: phone,
		Audio: audioBase64,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Erro ao criar JSON:", err)
		return err
	}

	payloadReader := strings.NewReader(string(jsonData))

	req, err := http.NewRequest("POST", url, payloadReader)
	if err != nil {
		return err
	}

	req.Header.Add("client-token", accountToken)
	req.Header.Add("Content-Type", "application/json")

	return nil
}

func ZAPISendDocumentMessage(documentBase64, phone, userInstance, userToken, accountToken string) error {
	url := fmt.Sprintf("https://api.z-api.io/instances/%s/token/%s/send-document", userInstance, userToken)

	payload := model.WhatsAppMessageWithDocument{
		Phone:    phone,
		Document: documentBase64,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Erro ao criar JSON:", err)
		return err
	}

	payloadReader := strings.NewReader(string(jsonData))

	req, err := http.NewRequest("POST", url, payloadReader)
	if err != nil {
		return err
	}

	req.Header.Add("client-token", accountToken)
	req.Header.Add("Content-Type", "application/json")

	return nil
}

func ZAPISendContactMessage(phone, contactName, contactPhone, userInstance, userToken, accountToken string) error {
	url := fmt.Sprintf("https://api.z-api.io/instances/%s/token/%s/send-contact", userInstance, userToken)

	payload := model.WhatsAppMessageWithContact{
		Phone:        phone,
		ContactName:  contactName,
		ContactPhone: contactPhone,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Erro ao criar JSON:", err)
		return err
	}

	payloadReader := strings.NewReader(string(jsonData))

	req, err := http.NewRequest("POST", url, payloadReader)
	if err != nil {
		return err
	}

	req.Header.Add("client-token", accountToken)
	req.Header.Add("Content-Type", "application/json")

	return nil
}

func ZAPISendPixButton(phone, pixKey, pixName, pixType, userInstance, userToken, accountToken string) error {
	url := fmt.Sprintf("https://api.z-api.io/instances/%s/token/%s/send-pix-button", userInstance, userToken)

	payload := model.SendPixButton{
		Phone:  phone,
		PixKey: pixKey,
		Type:   pixType,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Erro ao criar JSON:", err)
		return err
	}

	payloadReader := strings.NewReader(string(jsonData))

	req, err := http.NewRequest("POST", url, payloadReader)
	if err != nil {
		return err
	}

	req.Header.Add("client-token", accountToken)
	req.Header.Add("Content-Type", "application/json")

	return nil
}

func DeleteMessage(messageId, userInstance, userToken, accountToken string, owner bool) error {
	url := fmt.Sprintf("https://api.z-api.io/instances/%s/token/%s/messages?messageId=%s&phone=%s&owner=%t", userInstance, userToken, messageId, accountToken, owner)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("client-token", accountToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}

func ReadMessage(messageId, userInstance, userToken, accountToken string) error {
	url := fmt.Sprintf("https://api.z-api.io/instances/%s/token/%s/read-message", userInstance, userToken)

	payload := model.ReadMessage{
		MessageId: messageId,
		Phone:     accountToken,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Erro ao criar JSON:", err)
		return err
	}

	payloadReader := strings.NewReader(string(jsonData))

	req, err := http.NewRequest("POST", url, payloadReader)
	if err != nil {
		return err
	}

	req.Header.Add("client-token", accountToken)
	req.Header.Add("Content-Type", "application/json")

	return nil
}
