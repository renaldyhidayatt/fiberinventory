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

type CustomerHandlerGrpc struct {
	pb.UnimplementedCustomerServiceServer
	customer service.CustomerService
}

func NewCustomerHandlerGrpc(customer service.CustomerService) *CustomerHandlerGrpc {
	customerServer := CustomerHandlerGrpc{
		customer: customer,
	}

	return &customerServer
}

func (c *CustomerHandlerGrpc) CreateCustomer(ctx context.Context, req *pb.CreateCustomerInput) (*pb.CustomerResponse, error) {

	newCustomer := &domain.CreateCustomerRequest{
		Name:    req.Name,
		Alamat:  req.Alamat,
		Email:   req.Email,
		Telepon: req.Telepon,
	}

	if err := newCustomer.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	createdCustomer, err := c.customer.Create(newCustomer)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CustomerResponse{
		Customer: &pb.Customer{
			Id:      createdCustomer.ID,
			Name:    createdCustomer.Name,
			Alamat:  createdCustomer.Alamat,
			Email:   createdCustomer.Email,
			Telepon: createdCustomer.Telepon,
		},
	}

	return res, nil
}

func (c *CustomerHandlerGrpc) GetCustomers(ctx context.Context, req *pb.CustomersRequest) (*pb.CustomersResponse, error) {
	customers, err := c.customer.Results()

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var pbCustomers []*pb.Customer

	for _, customer := range *customers {
		pbCustomer := &pb.Customer{
			Id:        customer.ID,
			Name:      customer.Name,
			Alamat:    customer.Alamat,
			Email:     customer.Email,
			Telepon:   customer.Telepon,
			CreatedAt: timestamppb.New(customer.CreatedAt),
			UpdatedAt: timestamppb.New(customer.UpdatedAt),
		}
		pbCustomers = append(pbCustomers, pbCustomer)
	}

	return &pb.CustomersResponse{
		Customers: pbCustomers,
	}, nil
}

func (c *CustomerHandlerGrpc) GetCustomer(ctx context.Context, req *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	customerId := req.GetId()

	customer, err := c.customer.Result(customerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CustomerResponse{
		Customer: &pb.Customer{
			Id:      customer.ID,
			Name:    customer.Name,
			Alamat:  customer.Alamat,
			Email:   customer.Email,
			Telepon: customer.Telepon,
		},
	}

	return res, nil
}

func (c *CustomerHandlerGrpc) UpdateCustomer(ctx context.Context, req *pb.UpdateCustomerInput) (*pb.CustomerResponse, error) {
	customerId := req.GetId()

	customerToUpdate := &domain.UpdateCustomerRequest{
		ID:      customerId,
		Name:    req.Name,
		Alamat:  req.Alamat,
		Email:   req.Email,
		Telepon: req.Telepon,
	}

	if err := customerToUpdate.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	updatedCustomer, err := c.customer.Update(customerToUpdate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CustomerResponse{
		Customer: &pb.Customer{
			Id:      updatedCustomer.ID,
			Name:    updatedCustomer.Name,
			Alamat:  updatedCustomer.Alamat,
			Email:   updatedCustomer.Email,
			Telepon: updatedCustomer.Telepon,
		},
	}

	return res, nil
}

func (c *CustomerHandlerGrpc) DeleteCustomer(ctx context.Context, req *pb.CustomerRequest) (*pb.DeleteCustomerResponse, error) {
	customerId := req.GetId()

	_, err := c.customer.Delete(customerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteCustomerResponse{
		Success: true,
	}

	return res, nil
}
