package gapi

import (
	"context"
	"fiberinventory/internal/domain"
	"fiberinventory/internal/pb"
	"fiberinventory/internal/service"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandlerGrpc struct {
	pb.UnimplementedAuthServiceServer
	auth service.UserService
}

func NewAuthHandlerGrpcHandler(auth service.UserService) *AuthHandlerGrpc {
	return &AuthHandlerGrpc{
		auth: auth,
	}
}

func (a *AuthHandlerGrpc) RegisterUser(ctx context.Context, req *pb.SignUpUserInput) (*pb.SignUpUserResponse, error) {
	user := domain.RegisterInput{
		FirstName: req.GetFirstname(),
		LastName:  req.GetLastname(),
		Email:     req.GetEmail(),
		Password:  req.GetPassword(),
		Role:      req.GetRole(),
	}

	if err := user.Validate(); err != nil {

		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	newUser, err := a.auth.Register(&user)

	if err != nil {
		if strings.Contains(err.Error(), "email already exists") {
			return nil, status.Errorf(codes.AlreadyExists, "Email already exists")
		}
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	res := &pb.SignUpUserResponse{
		User: &pb.User{
			Firstname: newUser.FirstName,
			Lastname:  newUser.LastName,
			Email:     newUser.Email,
			Role:      newUser.Role,
		},
	}

	return res, nil
}

func (a *AuthHandlerGrpc) LoginUser(ctx context.Context, req *pb.SignInUserInput) (*pb.SignInUserResponse, error) {
	user := &domain.LoginInput{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	if err := user.Validate(); err != nil {

		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	res, err := a.auth.Login(user)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	result := &pb.SignInUserResponse{
		Status: "Success",
		Token:  res.Jwt,
	}

	return result, nil
}
