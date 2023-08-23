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

type SupplierHandlerGrpc struct {
	pb.UnimplementedSupplierServiceServer
	supplier service.SupplierService
}

func NewSupplierHandlerGrpc(supplier service.SupplierService) *SupplierHandlerGrpc {
	supplierServer := SupplierHandlerGrpc{
		supplier: supplier,
	}

	return &supplierServer
}

func (s *SupplierHandlerGrpc) CreateSupplier(ctx context.Context, req *pb.CreateSupplierInput) (*pb.SupplierResponse, error) {
	newSupplier := &domain.CreateSupplierRequest{
		Name:    req.Name,
		Alamat:  req.Alamat,
		Email:   req.Email,
		Telepon: req.Telepon,
	}

	if err := newSupplier.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	createdSupplier, err := s.supplier.Create(newSupplier)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.SupplierResponse{
		Supplier: &pb.Supplier{
			Id:      createdSupplier.ID,
			Name:    createdSupplier.Name,
			Alamat:  createdSupplier.Alamat,
			Email:   createdSupplier.Email,
			Telepon: createdSupplier.Telepon,
		},
	}

	return res, nil
}

func (s *SupplierHandlerGrpc) GetSuppliers(ctx context.Context, req *pb.SuppliersRequest) (*pb.SuppliersResponse, error) {
	suppliers, err := s.supplier.Results()

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var pbSuppliers []*pb.Supplier

	for _, supplier := range *suppliers {
		pbSupplier := &pb.Supplier{
			Id:        supplier.ID,
			Name:      supplier.Name,
			Alamat:    supplier.Alamat,
			Email:     supplier.Email,
			Telepon:   supplier.Telepon,
			CreatedAt: timestamppb.New(supplier.CreatedAt),
			UpdatedAt: timestamppb.New(supplier.UpdatedAt),
		}
		pbSuppliers = append(pbSuppliers, pbSupplier)
	}

	return &pb.SuppliersResponse{
		Suppliers: pbSuppliers,
	}, nil
}

func (s *SupplierHandlerGrpc) GetSupplier(ctx context.Context, req *pb.SupplierRequest) (*pb.SupplierResponse, error) {
	supplierId := req.GetId()

	supplier, err := s.supplier.Result(supplierId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.SupplierResponse{
		Supplier: &pb.Supplier{
			Id:      supplier.ID,
			Name:    supplier.Name,
			Alamat:  supplier.Alamat,
			Email:   supplier.Email,
			Telepon: supplier.Telepon,
		},
	}

	return res, nil
}

func (s *SupplierHandlerGrpc) UpdateSupplier(ctx context.Context, req *pb.UpdateSupplierInput) (*pb.SupplierResponse, error) {
	supplierId := req.GetId()

	supplierToUpdate := &domain.UpdateSupplierRequest{
		ID:      supplierId,
		Name:    req.Name,
		Alamat:  req.Alamat,
		Email:   req.Email,
		Telepon: req.Telepon,
	}

	if err := supplierToUpdate.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	updatedSupplier, err := s.supplier.Update(supplierToUpdate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.SupplierResponse{
		Supplier: &pb.Supplier{
			Id:      updatedSupplier.ID,
			Name:    updatedSupplier.Name,
			Alamat:  updatedSupplier.Alamat,
			Email:   updatedSupplier.Email,
			Telepon: updatedSupplier.Telepon,
		},
	}

	return res, nil
}

func (s *SupplierHandlerGrpc) DeleteSupplier(ctx context.Context, req *pb.SupplierRequest) (*pb.DeleteSupplierResponse, error) {
	supplierId := req.GetId()

	_, err := s.supplier.Delete(supplierId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteSupplierResponse{
		Success: true,
	}

	return res, nil
}
