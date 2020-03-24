package model

import (
	"fmt"
)

type User struct {
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Email       string  `json:"email"`
	MobilePhone float64 `json:"mobilePhone"`
}

type Balances struct {
	Total     float64 `json:"total"`
	Available float64 `json:"available"`
	Blocked   float64 `json:"blocked"`
}

type OrderStatistics struct {
	TotalPaidAmount float64 `json:"totalPaidAmount"`
}

type UserInformation struct {
	Customer   User            `json:"customer"`
	Balance    []Balances      `json:"balances"`
	Statistics OrderStatistics `json:"retailOrderStatistics"`
}

func (u UserInformation) FirstName() string {
	return u.Customer.FirstName
}

func (u UserInformation) LastName() string {
	return u.Customer.LastName
}

func (u UserInformation) Email() string {
	return u.Customer.Email
}

func (u UserInformation) Phone() string {
	return fmt.Sprintf("%.0f", u.Customer.MobilePhone)
}

func (u UserInformation) TotalPaidAmount() float64 {
	return u.Statistics.TotalPaidAmount
}

func (u UserInformation) BonusesTotal() int32 {
	if len(u.Balance) == 0 {
		return 0
	}
	return int32(u.Balance[0].Total)
}

func (u UserInformation) BonusesAvailable() int32 {
	if len(u.Balance) == 0 {
		return 0
	}
	return int32(u.Balance[0].Available)
}

func (u UserInformation) BonusesBlocked() int32 {
	if len(u.Balance) == 0 {
		return 0
	}
	return int32(u.Balance[0].Blocked)
}

type OrdersHistory struct {
	TotalCount int
	Orders     []Order
}

type Order struct {
	ID                    string
	CreatedDate           string
	PaymentType           string
	DiscountedTotalPrice  float64
	PaymentAmount         float64
	AppliedDiscount       float64
	AcquiredBalanceChange float64
}

type OrdersXML struct {
	TotalCount int `xml:"totalCount"`
	Orders     struct {
		Order []struct {
			Ids struct {
				StoreezCom string `xml:"storeez.com"`
			} `xml:"ids"`
			CreatedDateTimeUtc   string  `xml:"createdDateTimeUtc"`
			DiscountedTotalPrice float64 `xml:"discountedTotalPrice"`
			Payments             struct {
				Payment struct {
					Type   string  `xml:"type"`
					Amount float64 `xml:"amount"`
				} `xml:"payment"`
			} `xml:"payments"`
			Lines struct {
				Line []struct {
					AppliedDiscounts struct {
						AppliedDiscount struct {
							Amount float64 `xml:"amount"`
						} `xml:"appliedDiscount"`
					} `xml:"appliedDiscounts"`
				} `xml:"line"`
			} `xml:"lines"`
			TotalAcquiredBalaneChange float64 `xml:"totalAcquiredBalaneChange"`
		} `xml:"order"`
	} `xml:"orders"`
}
