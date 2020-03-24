package service

import (
	"context"
	"github.com/micro/go-micro/errors"
	pb "github.com/zhs/esb-protobufs/go/mindbox"
	"modz/repository"
	"modz/usecase"
)

type UserService struct {
	UserInfo usecase.User
}

func NewUserService(userInfo usecase.User) *UserService {
	return &UserService{
		UserInfo: userInfo,
	}
}

func (s UserService) Info(ctx context.Context, req *pb.ParamsUser, res *pb.ResponseUser) error {
	userInfo, err := s.UserInfo.Info(req.GetClientId())
	if err != nil {
		switch err {
		case repository.ErrNotFound:
			return errors.NotFound("mindbox.User.Info", "%v", err.Error())
		default:
			return errors.InternalServerError("mindbox.User.Info", "%v", err.Error())
		}
	}

	res.FirstName = userInfo.FirstName()
	res.LastName = userInfo.LastName()
	res.Email = userInfo.Email()
	res.Phone = userInfo.Phone()
	res.BonusTotal = userInfo.BonusesTotal()
	res.BonusAvailable = userInfo.BonusesAvailable()
	res.BonusBlocked = userInfo.BonusesBlocked()
	res.TotalPaidAmount = userInfo.TotalPaidAmount()

	return nil
}

func (s UserService) Orders(ctx context.Context, req *pb.ParamsOrders, res *pb.ResponseOrders) error {
	userInfo, err := s.UserInfo.Orders(req.GetClientId())
	if err != nil {
		switch err {
		case repository.ErrNotFound:
			return errors.NotFound("mindbox.User.Orders:", "%v", err.Error())
		default:
			return errors.InternalServerError("mindbox.User.Orders:", "%v", err.Error())
		}
	}

	res.Result = make([]*pb.Order, len(userInfo.Orders))
	res.Total = int32(userInfo.TotalCount)
	for i, ord := range userInfo.Orders {
		res.Result[i] = &pb.Order{
			OrderId:               ord.ID,
			CreatedDate:           ord.CreatedDate,
			DiscountedTotalPrice:  ord.DiscountedTotalPrice,
			PaymentType:           ord.PaymentType,
			PaymentAmount:         ord.PaymentAmount,
			AppliedDiscount:       ord.AppliedDiscount,
			AcquiredBalanceChange: ord.AcquiredBalanceChange,
		}
	}

	return nil
}

func (s UserService) SendOSMICard(ctx context.Context, req *pb.ParamsOSMICard, res *pb.ResponseOSMICard) error {
	isOk, err := s.UserInfo.SendOSMICard(req.GetClientId())
	if err != nil {
		return errors.InternalServerError("mindbox.User.SendOSMICard:", "%v", err.Error())
	}
	res.Ok = isOk
	return nil
}
