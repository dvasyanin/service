package repository

import (
	"fmt"
	"gopkg.in/resty.v1"
)

var (
	HEADERS = map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": "",
	}
)

func NewMobileRepo() MobileRepo {
	return &mobileRepo{}
}

type mobileRepo struct {
}

func request(url string, headers map[string]string, body string) ([]byte, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headers).
		SetBody(body).
		Post(url)
	if err != nil {
		return nil, fmt.Errorf("r.request: %w", err)
	}
	defer resp.RawBody().Close()

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("r.request: %s", string(resp.Body()))
	}

	return resp.Body(), nil
}

func (r mobileRepo) Authorization(deviceUUID string) ([]byte, error) {
	url := fmt.Sprintf("url%s", deviceUUID)
	result, err := request(url, HEADERS, "")
	if err != nil {
		return nil, fmt.Errorf("r.Authorization: %v", err)
	}

	return result, nil
}

func (r mobileRepo) Registration(deviceUUID string, clientInfo map[string]string) ([]byte, error) {
	url := fmt.Sprintf("=%s", deviceUUID)
	if _, ok := clientInfo["clientID"]; !ok {
		return nil, fmt.Errorf("r.Registration: clientID not found")
	}

	body := `{
  "customer": {
	"ids": {
  	"storeezWebsiteId": "` + clientInfo["clientID"] + `"
	},
	"email": "` + clientInfo["email"] + `",
	"fullName": "` + clientInfo["fullName"] + `",
	"mobilePhone": "` + clientInfo["phone"] + `"
  }
}`
	result, err := request(url, HEADERS, body)
	if err != nil {
		return nil, fmt.Errorf("r.Registration: %v", err)
	}

	return result, nil
}

func (r mobileRepo) Code(numberPhone string) ([]byte, error) {
	url := ""
	body := `{
  "customer": {
	"mobilePhone": "` + numberPhone + `"
  }
}
`
	result, err := request(url, HEADERS, body)
	if err != nil {
		return nil, fmt.Errorf("r.Code: %v", err)
	}

	return result, nil
}

func (r mobileRepo) CheckCode(numberPhone string, code string) ([]byte, error) {
	url := ""
	body := `{
  "customer": {
	"mobilePhone": "` + numberPhone + `"
  },
  "authentificationCode": "` + code + `"
}
`
	result, err := request(url, HEADERS, body)
	if err != nil {
		return nil, fmt.Errorf("r.CheckCode: %v", err)
	}

	return result, nil
}

func (r mobileRepo) Creation(deviceUUID string, id string, vendor string, model string) ([]byte, error) {
	url := fmt.Sprintf("=%s", deviceUUID)
	body := `{
  "mobileApplicationInstallation": {
	"id": "` + id + `",
	"vendor": "` + vendor + `",
	"model": "` + model + `"
  }
}
`
	result, err := request(url, HEADERS, body)
	if err != nil {
		return nil, fmt.Errorf("r.Creation: %v", err)
	}

	return result, nil
}

func (r mobileRepo) EditUser(deviceUUID string, userID string, numberPhone string) ([]byte, error) {
	url := fmt.Sprintf("=%s", deviceUUID)
	body := `{
  "customer": {
	"ids": {
  	"storeezWebsiteId": "` + userID + `"
	},
	"mobilePhone": "` + numberPhone + `"
  }
}
`
	result, err := request(url, HEADERS, body)
	if err != nil {
		return nil, fmt.Errorf("r.EditUser: %v", err)
	}

	return result, nil
}
