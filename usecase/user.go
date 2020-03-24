package usecase

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"modz/model"
	"modz/repository"
)

func NewUserUsecase(rep repository.UserRepo) User {
	return &UserUsecase{rep: rep}
}

type UserUsecase struct {
	rep repository.UserRepo
}

func (u UserUsecase) Info(clientID string) (*model.UserInformation, error) {
	url := "&operation=ClientSearch"
	body := `{
  "customer": {
    "ids": {
      "storeezWebsiteId": "` + clientID + `"
    }
  }
}`

	result, err := u.rep.Get(url, body)
	if err != nil {
		return nil, err
	}

	var userInfo model.UserInformation
	if err := json.Unmarshal(result, &userInfo); err != nil {
		return nil, fmt.Errorf("u.Info: unmarshal error %v", err)
	}

	if len(userInfo.Balance) == 0 {
		return nil, repository.ErrNotFound
	}

	return &userInfo, nil
}

func (u UserUsecase) Orders(clientID string) (*model.OrdersHistory, error) {
	result, err := u.rep.GetUserOrders(clientID)
	if err != nil {
		return nil, err
	}

	var rez model.OrdersXML
	err = xml.Unmarshal(result, &rez)
	if err != nil {
		return nil, fmt.Errorf("u.Orders: unmarshal error %v", err)
	}

	ordersHistory := ordersXMLToOrdersHistory(rez)

	return &ordersHistory, nil
}

func (u UserUsecase) SendOSMICard(clientID string) (bool, error) {
	err := u.rep.SendOSMICard(clientID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ordersXMLToOrdersHistory(ord model.OrdersXML) model.OrdersHistory {
	var orders model.OrdersHistory

	orders.Orders = make([]model.Order, len(ord.Orders.Order))
	orders.TotalCount = ord.TotalCount

	for i, order := range ord.Orders.Order {
		discount := .0
		for _, amount := range order.Lines.Line {
			discount += amount.AppliedDiscounts.AppliedDiscount.Amount
		}
		orders.Orders[i] = model.Order{
			ID:                    order.Ids.StoreezCom,
			CreatedDate:           order.CreatedDateTimeUtc,
			DiscountedTotalPrice:  order.DiscountedTotalPrice,
			PaymentType:           order.Payments.Payment.Type,
			PaymentAmount:         order.Payments.Payment.Amount,
			AppliedDiscount:       discount,
			AcquiredBalanceChange: order.TotalAcquiredBalaneChange,
		}
	}

	return orders
}
