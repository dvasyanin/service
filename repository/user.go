package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/resty.v1"
)

const (
	AUTHORIZATION = ``
	BASE_URL_V3   = ""
)

var (
	ErrNotFound = errors.New("client not found")
)

func NewUserRepo() UserRepo {
	return &userRepo{}
}

type userRepo struct {
}

// отправить POST запрос в mindbox
// возвращает результат Body() запроса
func (r userRepo) Get(url string, body string) ([]byte, error) {
	url = BASE_URL_V3 + url
	headers := make(map[string]string, 3)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	headers["Authorization"] = AUTHORIZATION

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headers).
		SetBody(body).
		Post(url)
	if err != nil {
		return nil, fmt.Errorf("r.Get: %v", err)
	}

	return resp.Body(), nil
}

func (r userRepo) GetUserOrders(clientID string) ([]byte, error) {
	url := "" +
		"" + clientID
	headers := make(map[string]string, 3)
	headers["Accept"] = "application/xml"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = ``

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headers).
		Get(url)
	if err != nil {
		return nil, fmt.Errorf("r.GetUserOrders: %v", err)
	}

	if resp.StatusCode() == 404 {
		return nil, ErrNotFound
	}
	defer resp.RawBody().Close()
	return resp.Body(), nil
}

func (r userRepo) SendOSMICard(clientID string) error {
	url := ""
	headers := map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json; charset=utf-8",
		"Authorization": ``,
	}

	body := fmt.Sprintf(`{
			  "customer": {
				"ids": {
				  "storeezWebsiteId": %s
				}
			  }
			}`, clientID)

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headers).
		SetBody(body).
		Post(url)
	if err != nil {
		return fmt.Errorf("r.SendOSMICard: %v", err)
	}
	defer resp.RawBody().Close()

	if resp.StatusCode() != 200 {
		return fmt.Errorf("r.SendOSMICard: %s", string(resp.Body()))
	}

	var statusResp map[string]string
	if err := json.Unmarshal(resp.Body(), &statusResp); err != nil {
		return fmt.Errorf("r.SendOSMICard: unmarshal body error: %v", err)
	}

	if status, ok := statusResp["status"]; ok {
		if status != "Success" {
			return fmt.Errorf("r.SendOSMICard: request status: %v", status)
		}
	}

	return nil
}
