package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server/grpc"
	"github.com/micro/go-micro/util/log"
	pbhealth "github.com/zhs/esb-protobufs/go/health"
	pb "github.com/zhs/esb-protobufs/go/mindbox"
	"modz/config"
	"modz/repository"
	"modz/service"
	"modz/usecase"
)

func main() {
	cfg, err := config.ReadFromFile("config.yaml")
	if err != nil {
		log.Fatalf("can't read config: %v", err)
	}

	repMB := repository.NewUserRepo()
	repMobile := repository.NewMobileRepo()
	userInfo := usecase.NewUserUsecase(repMB)
	mobile := usecase.NewAuthenticationService(repMobile)
	userService := service.NewUserService(userInfo)
	mobileService := service.NewAuthenticationMobileService(mobile)
	healthService := service.NewHealthService()

	// MICRO gRPC SERVER
	srv := micro.NewService(
		micro.Server(grpc.NewServer()),
		micro.Name("mindbox"),
		micro.Address(cfg.App.Port),
	)

	srv.Init()
	_ = pb.RegisterUserHandler(srv.Server(), userService)
	_ = pb.RegisterMobileHandler(srv.Server(), mobileService)
	_ = pbhealth.RegisterHealthHandler(srv.Server(), healthService)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
