package usecase

import "modz/model"

type User interface {
	Info(clientID string) (*model.UserInformation, error)
	Orders(clientID string) (*model.OrdersHistory, error)
	SendOSMICard(clientID string) (bool, error)
}

type Mobile interface {
	Creation(deviceUUID string, id string, vendor string, md string) (model.Answer, error)
	Authorization(deviceUUID string) (model.Answer, error)
	Registration(deviceUUID string, user map[string]string) (model.Answer, error)
	Code(numberPhone string) (model.Answer, error)
	CheckCode(numberPhone string, code string) (model.Answer, error)
	EditUser(deviceUUID string, userID string, numberPhone string) (model.Answer, error)
}
