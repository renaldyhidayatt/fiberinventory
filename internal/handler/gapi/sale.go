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

type SaleHandlerGrpc struct {
	pb.UnimplementedSaleServiceServer
	sale service.SaleService
}

func NewSaleHandlerGrpc(sale service.SaleService) *SaleHandlerGrpc {
	saleServer := SaleHandlerGrpc{
		sale: sale,
	}

	return &saleServer
}

func (s *SaleHandlerGrpc) CreateSale(ctx context.Context, req *pb.CreateSaleInput) (*pb.SaleResponse, error) {
	newSale := &domain.CreateSaleRequest{
		Name:    req.Name,
		Alamat:  req.Alamat,
		Email:   req.Email,
		Telepon: req.Telepon,
	}

	if err := newSale.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	createdSale, err := s.sale.Create(newSale)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.SaleResponse{
		Sale: &pb.Sale{
			Id:      createdSale.ID,
			Name:    createdSale.Name,
			Alamat:  createdSale.Alamat,
			Email:   createdSale.Email,
			Telepon: createdSale.Telepon,
		},
	}

	return res, nil
}

func (s *SaleHandlerGrpc) GetSales(ctx context.Context, req *pb.SalesRequest) (*pb.SalesResponse, error) {
	sales, err := s.sale.Results()

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var pbSales []*pb.Sale

	for _, sale := range *sales {
		pbSale := &pb.Sale{
			Id:        sale.ID,
			Name:      sale.Name,
			Alamat:    sale.Alamat,
			Email:     sale.Email,
			Telepon:   sale.Telepon,
			CreatedAt: timestamppb.New(sale.CreatedAt),
			UpdatedAt: timestamppb.New(sale.UpdatedAt),
		}
		pbSales = append(pbSales, pbSale)
	}

	return &pb.SalesResponse{
		Sales: pbSales,
	}, nil
}

func (s *SaleHandlerGrpc) GetSale(ctx context.Context, req *pb.SaleRequest) (*pb.SaleResponse, error) {
	saleId := req.GetId()

	sale, err := s.sale.Result(saleId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.SaleResponse{
		Sale: &pb.Sale{
			Id:      sale.ID,
			Name:    sale.Name,
			Alamat:  sale.Alamat,
			Email:   sale.Email,
			Telepon: sale.Telepon,
		},
	}

	return res, nil
}

func (s *SaleHandlerGrpc) UpdateSale(ctx context.Context, req *pb.UpdateSaleInput) (*pb.SaleResponse, error) {
	saleId := req.GetId()

	saleToUpdate := &domain.UpdateSaleRequest{
		ID:      saleId,
		Name:    req.Name,
		Alamat:  req.Alamat,
		Email:   req.Email,
		Telepon: req.Telepon,
	}

	if err := saleToUpdate.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	updatedSale, err := s.sale.Update(saleToUpdate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.SaleResponse{
		Sale: &pb.Sale{
			Id:      updatedSale.ID,
			Name:    updatedSale.Name,
			Alamat:  updatedSale.Alamat,
			Email:   updatedSale.Email,
			Telepon: updatedSale.Telepon,
		},
	}

	return res, nil
}

func (s *SaleHandlerGrpc) DeleteSale(ctx context.Context, req *pb.SaleRequest) (*pb.DeleteSaleResponse, error) {
	saleId := req.GetId()

	_, err := s.sale.Delete(saleId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteSaleResponse{
		Success: true,
	}

	return res, nil
}
