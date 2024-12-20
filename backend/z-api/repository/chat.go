package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"zapi/model"
)

func (r *Repository) GetChat(userInstance, userToken, accountToken string) ([]model.Chat, error) {
	url := fmt.Sprintf("https://api.z-api.io/instances/%s/token/%s/chats?page=1&pageSize=10", userInstance, userToken)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []model.Chat{}, err
	}

	req.Header.Add("client-token", accountToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []model.Chat{}, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []model.Chat{}, err
	}

	var chats []model.Chat
	json.Unmarshal(body, &chats)

	return chats, nil
}

func (r *Repository) GetChatMetadata(phone, userInstance, userToken, accountToken string) (model.ChatMetadata, error) {
	url := fmt.Sprintf("https://api.z-api.io/instances/%s/token/%s/chats/%s", userInstance, userToken, phone)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model.ChatMetadata{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("client-token", accountToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return model.ChatMetadata{}, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.ChatMetadata{}, err
	}

	var chatMetadata model.ChatMetadata
	json.Unmarshal(body, &chatMetadata)

	return chatMetadata, nil
}

func (r *Repository) ModifyChatStatus(phone, action, userInstance, userToken, accountToken string) error {
	url := fmt.Sprintf("https://api.z-api.io/instances/%s/token/%s/modify-chat", userInstance, userToken)

	payload := model.ModifyChatStatus{
		Phone:  phone,
		Action: action, // read | unread | archive | unarchive | pin | unpin | mute | unmute | spam | unspam | clear | delete
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
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
