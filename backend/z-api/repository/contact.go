package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"zapi/model"
)

func GetContacts(userInstance, userToken, accountToken string) ([]model.GetContacts, error) {
	url := fmt.Sprintf("https://api.z-api.io/instances/%s/token/%s/contacts", userInstance, userToken)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []model.GetContacts{}, err
	}

	req.Header.Add("client-token", accountToken)
	req.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return []model.GetContacts{}, err
	}

	defer response.Body.Close()

	var contacts []model.GetContacts

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []model.GetContacts{}, err
	}

	err = json.Unmarshal(body, &contacts)
	if err != nil {
		return []model.GetContacts{}, err
	}

	return contacts, nil
}

func CreateContact(userInstance, userToken, accountToken string, contact model.GetContacts) error {
	url := fmt.Sprintf("https://api.z-api.io/instances/%s/token/%s/contacts/add", userInstance, userToken)

	payload := model.CreateContact{
		FirstName: contact.Name,
		LastName:  contact.Short,
		Phone:     contact.Phone,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	payloadReader := strings.NewReader(string(jsonData))

	req, err := http.NewRequest("POST", url, payloadReader)
	if err != nil {
		return err
	}

	req.Header.Add("client-token", accountToken)
	req.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	return nil
}

func DeleteContact(contactPhone, userInstance, userToken, accountToken string) error {
	url := fmt.Sprintf("https://api.z-api.io/instances/%s/token/%s/contacts/remove", userInstance, userToken)

	payload := strings.NewReader(fmt.Sprintf("[\"%s\"]", contactPhone))

	req, err := http.NewRequest("DELETE", url, payload)
	if err != nil {
		return err
	}

	req.Header.Add("client-token", accountToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}

func GetContactPhoto(contactPhone, userInstance, userToken, accountToken string) (string, error) {
	url := fmt.Sprintf("https://api.z-api.io/instances/%s/token/%s/profile-picture?phone=%s&Client-Token=%s", userInstance, userToken, contactPhone, accountToken)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("client-token", accountToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
