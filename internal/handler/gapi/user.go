package gapi

import (
	"context"
	"fiberinventory/internal/domain"
	"fiberinventory/internal/pb"
	"fiberinventory/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserHandlerGrpc struct {
	pb.UnimplementedUserServiceServer
	user service.UserService
}

func NewUserHandlerGrpcHandler(user service.UserService) *UserHandlerGrpc {
	userServer := UserHandlerGrpc{
		user: user,
	}

	return &userServer
}

func (u *UserHandlerGrpc) CreateUser(ctx context.Context, req *pb.CreateUserInput) (*pb.UserResponse, error) {

	newUser := &domain.RegisterInput{
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Email:     req.Email,
		Password:  req.Password,
		Role:      req.Role,
	}

	if err := newUser.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())

	}

	createdUser, err := u.user.Register(newUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.UserResponse{
		User: &pb.User{
			Id:        createdUser.ID,
			Firstname: createdUser.FirstName,
			Lastname:  createdUser.LastName,
			Email:     createdUser.Email,
			Role:      createdUser.Role,
		},
	}

	return res, nil
}

func (u *UserHandlerGrpc) GetUsers(ctx context.Context, req *pb.UsersRequest) (*pb.UsersResponse, error) {
	users, err := u.user.Results()

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var pbUsers []*pb.User

	for _, user := range *users {
		pbUser := &pb.User{
			Id:        user.ID,
			Firstname: user.FirstName,
			Lastname:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		}
		pbUsers = append(pbUsers, pbUser)
	}

	return &pb.UsersResponse{
		Users: pbUsers,
	}, nil
}

func (u *UserHandlerGrpc) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	userId := req.GetId()

	user, err := u.user.Result(userId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.UserResponse{
		User: &pb.User{
			Id:        user.ID,
			Firstname: user.FirstName,
			Lastname:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
		},
	}

	return res, nil
}

func (u *UserHandlerGrpc) UpdateUser(ctx context.Context, req *pb.UpdateUserInput) (*pb.UserResponse, error) {
	userId := req.GetId()

	userToUpdate := &domain.UpdateUserRequest{
		ID:        userId,
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Email:     req.Email,
		Role:      req.Role,
	}

	if err := userToUpdate.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	updatedUser, err := u.user.Update(userToUpdate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.UserResponse{
		User: &pb.User{
			Id:        updatedUser.ID,
			Firstname: updatedUser.FirstName,
			Lastname:  updatedUser.LastName,
			Email:     updatedUser.Email,
			Role:      updatedUser.Role,
		},
	}

	return res, nil
}

func (u *UserHandlerGrpc) DeleteUser(ctx context.Context, req *pb.UserRequest) (*pb.DeleteUserResponse, error) {
	userId := req.GetId()

	_, err := u.user.Delete(userId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteUserResponse{
		Success: true,
	}

	return res, nil
}
