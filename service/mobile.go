package service

import (
	"context"
	"github.com/micro/go-micro/errors"
	pb "github.com/zhs/esb-protobufs/go/mindbox"
	"modz/usecase"
)

type MindBoxMobile struct {
	Mobile usecase.Mobile
}

func NewAuthenticationMobileService(auth usecase.Mobile) *MindBoxMobile {
	return &MindBoxMobile{Mobile: auth}
}

func (s MindBoxMobile) Authorization(ctx context.Context, req *pb.ParamsAuthorization, res *pb.ResponseAuthorization) error {
	result, err := s.Mobile.Authorization(req.GetDeviceUuid())
	if err != nil {
		return errors.InternalServerError("mindbox.Mobile.Authorization:", "%v", err.Error())
	}
	res.Ok = result.Status == "Success"

	return nil
}

func (s MindBoxMobile) Registration(ctx context.Context, req *pb.ParamsRegistration, res *pb.ResponseRegistration) error {
	user := map[string]string{"clientID": req.GetClientId(), "fullName": req.GetFullName(), "email": req.GetEmail(), "phone": req.GetMobilePhone()}
	result, err := s.Mobile.Registration(req.GetDeviceUuid(), user)
	if err != nil {
		return errors.InternalServerError("mindbox.Mobile.Registration:", "%v", err.Error())
	}

	res.Ok = result.Status == "Success"

	return nil
}

func (s MindBoxMobile) Code(ctx context.Context, req *pb.ParamsCode, res *pb.ResponseCode) error {
	result, err := s.Mobile.Code(req.GetMobilePhone())
	if err != nil {
		return errors.InternalServerError("mindbox.Mobile.Code:", "%v", err.Error())
	}

	res.Ok = result.Status == "Success"

	return nil
}

func (s MindBoxMobile) CheckCode(ctx context.Context, req *pb.ParamsCheckCode, res *pb.ResponseCheckCode) error {
	result, err := s.Mobile.CheckCode(req.GetMobilePhone(), req.GetCode())
	if err != nil {
		return errors.InternalServerError("mindbox.Mobile.CheckCode:", "%v", err.Error())
	}

	res.Ok = result.Status == "Success"
	res.Status = result.ProcessingStatus
	res.ClientId = result.StoreezWebsiteId
	res.MindboxId = result.MindboxId

	return nil
}

func (s MindBoxMobile) Creation(ctx context.Context, req *pb.ParamsCreation, res *pb.ResponseCreation) error {
	result, err := s.Mobile.Creation(req.GetDeviceUuid(), req.GetId(), req.GetVendor(), req.GetModel())
	if err != nil {
		return errors.InternalServerError("mindbox.Mobile.Creation:", "%v", err.Error())
	}

	res.Ok = result.Status == "Success"

	return nil
}

func (s MindBoxMobile) EditUser(ctx context.Context, req *pb.ParamsEditUser, res *pb.ResponseEditUser) error {
	result, err := s.Mobile.EditUser(req.GetDeviceUuid(), req.GetClientId(), req.GetMobilePhone())
	if err != nil {
		return errors.InternalServerError("mindbox.Mobile.EditUser:", "%v", err.Error())
	}

	res.Ok = result.Status == "Success"

	return nil
}
