package usecase

import (
	"encoding/json"
	"fmt"
	"modz/model"
	"modz/repository"
)

func NewAuthenticationService(rep repository.MobileRepo) Mobile {
	return &Authentication{rep: rep}
}

type Authentication struct {
	rep repository.MobileRepo
}

func (u Authentication) Authorization(deviceUUID string) (model.Answer, error) {
	result, err := u.rep.Authorization(deviceUUID)
	if err != nil {
		return model.Answer{}, err
	}

	var auth model.Answer
	err = json.Unmarshal(result, &auth)
	if err != nil {
		return model.Answer{}, fmt.Errorf("u.Authorization: json.Unmarshal: %v", err)
	}

	return auth, nil
}

func (u Authentication) Registration(deviceUUID string, user map[string]string) (model.Answer, error) {
	result, err := u.rep.Registration(deviceUUID, user)
	if err != nil {
		return model.Answer{}, err
	}

	var reg model.Answer
	err = json.Unmarshal(result, &reg)
	if err != nil {
		return model.Answer{}, fmt.Errorf("u.Registration: json.Unmarshal: %v", err)
	}

	return reg, nil
}

func (u Authentication) Code(numberPhone string) (model.Answer, error) {
	answer, err := u.rep.Code(numberPhone)
	if err != nil {
		return model.Answer{}, err
	}

	var result model.Answer

	err = json.Unmarshal(answer, &result)
	if err != nil {
		return model.Answer{}, fmt.Errorf("u.Code: json.Unmarshal: %v", err)
	}

	return result, nil
}

func (u Authentication) CheckCode(numberPhone string, code string) (model.Answer, error) {
	answer, err := u.rep.CheckCode(numberPhone, code)
	if err != nil {
		return model.Answer{}, err
	}

	var res map[string]interface{}

	err = json.Unmarshal(answer, &res)
	if err != nil {
		return model.Answer{}, fmt.Errorf("u.CheckCode: json.Unmarshal: %v", err)
	}

	if status, ok := res["status"].(string); ok {

		switch status {
		case "Success":
			var result model.Answer
			err = json.Unmarshal(answer, &result)
			if err != nil {
				return result, fmt.Errorf("u.CheckCode: switch.json.Unmarshal: %v", err)
			}
			return result, nil

		case "ValidationError":
			if len(res["validationMessages"].([]interface{})) > 0 {
				if msg, ok := res["validationMessages"].([]interface{})[0].(map[string]interface{})["message"]; ok {
					return model.Answer{
						Status: "",
						Customer: model.Customer{
							ProcessingStatus: msg.(string),
							Ids: model.Ids{
								MindboxId:        0,
								StoreezWebsiteId: "0",
							},
						},
					}, nil
				}
			}
			return model.Answer{}, nil

		default:
			return model.Answer{}, fmt.Errorf("u.CheckCode: %s", string(answer))
		}
	}

	return model.Answer{}, nil
}

func (u Authentication) Creation(deviceUUID string, id string, vendor string, md string) (model.Answer, error) {
	answer, err := u.rep.Creation(deviceUUID, id, vendor, md)
	if err != nil {
		return model.Answer{}, err
	}
	var result model.Answer

	err = json.Unmarshal(answer, &result)
	if err != nil {
		return model.Answer{}, fmt.Errorf("u.Creation: json.Unmarshal: %v", err)
	}

	return result, nil
}

func (u Authentication) EditUser(deviceUUID string, userID string, numberPhone string) (model.Answer, error) {
	answer, err := u.rep.EditUser(deviceUUID, userID, numberPhone)
	if err != nil {
		return model.Answer{}, err
	}
	var result model.Answer

	err = json.Unmarshal(answer, &result)
	if err != nil {
		return model.Answer{}, fmt.Errorf("u.EditUser: json.Unmarshal: %v", err)
	}

	return result, nil
}
