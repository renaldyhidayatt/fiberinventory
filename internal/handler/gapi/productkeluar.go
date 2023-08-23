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

type ProductKeluarGrpc struct {
	pb.UnimplementedProductKeluarServiceServer
	productkeluar service.ProductKeluarService
}

func NewProductKeluarHandlerGrpcHandler(product service.ProductKeluarService) *ProductKeluarGrpc {
	return &ProductKeluarGrpc{
		productkeluar: product,
	}
}

func (p *ProductKeluarGrpc) CreateProductKeluar(ctx context.Context, req *pb.CreateProductKeluarInput) (*pb.ProductKeluarResponse, error) {
	newProductKeluar := &domain.CreateProductKeluarRequest{
		Qty:        req.Qty,
		ProductID:  req.ProductId,
		CategoryID: req.CategoryId,
	}

	if err := newProductKeluar.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	createdProductKeluar, err := p.productkeluar.Create(newProductKeluar)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbProductKeluar := &pb.ProductKeluar{
		Id:         createdProductKeluar.ID,
		Qty:        createdProductKeluar.Qty,
		ProductId:  createdProductKeluar.ProductID,
		CategoryId: createdProductKeluar.CategoryID,
		CreatedAt:  timestamppb.New(createdProductKeluar.CreatedAt),
		UpdatedAt:  timestamppb.New(createdProductKeluar.UpdatedAt),
	}

	return &pb.ProductKeluarResponse{
		ProductKeluar: pbProductKeluar,
	}, nil
}

func (p *ProductKeluarGrpc) GetProductKeluars(ctx context.Context, req *pb.ProductKeluarsRequest) (*pb.ProductKeluarsResponse, error) {
	productKeluar, err := p.productkeluar.Results()

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var pbProductKeluars []*pb.ProductKeluar

	for _, product := range *productKeluar {
		pbProductKeluar := &pb.ProductKeluar{
			Id:         product.ID,
			Qty:        product.Qty,
			ProductId:  product.ProductID,
			CategoryId: product.CategoryID,
			CreatedAt:  timestamppb.New(product.CreatedAt),
			UpdatedAt:  timestamppb.New(product.UpdatedAt),
		}
		pbProductKeluars = append(pbProductKeluars, pbProductKeluar)
	}

	return &pb.ProductKeluarsResponse{
		ProductKeluars: pbProductKeluars,
	}, nil
}

func (p *ProductKeluarGrpc) GetProductKeluar(ctx context.Context, req *pb.ProductKeluarRequest) (*pb.ProductKeluarResponse, error) {
	productKeluarID := req.GetId()

	productKeluar, err := p.productkeluar.Result(productKeluarID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbProductKeluar := &pb.ProductKeluar{
		Id:         productKeluar.ID,
		Qty:        productKeluar.Qty,
		ProductId:  productKeluar.ProductID,
		CategoryId: productKeluar.CategoryID,
		CreatedAt:  timestamppb.New(productKeluar.CreatedAt),
		UpdatedAt:  timestamppb.New(productKeluar.UpdatedAt),
	}

	return &pb.ProductKeluarResponse{
		ProductKeluar: pbProductKeluar,
	}, nil
}

func (p *ProductKeluarGrpc) UpdateProductKeluar(ctx context.Context, req *pb.UpdateProductKeluarInput) (*pb.ProductKeluarResponse, error) {
	productKeluarID := req.GetId()

	productKeluarToUpdate := &domain.UpdateProductKeluarRequest{
		ID:         productKeluarID,
		Qty:        req.Qty,
		ProductID:  req.ProductId,
		CategoryID: req.CategoryId,
	}

	if err := productKeluarToUpdate.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	updatedProductKeluar, err := p.productkeluar.Update(productKeluarToUpdate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbProductKeluar := &pb.ProductKeluar{
		Id:         updatedProductKeluar.ID,
		Qty:        updatedProductKeluar.Qty,
		ProductId:  updatedProductKeluar.ProductID,
		CategoryId: updatedProductKeluar.CategoryID,
		CreatedAt:  timestamppb.New(updatedProductKeluar.CreatedAt),
		UpdatedAt:  timestamppb.New(updatedProductKeluar.UpdatedAt),
	}

	return &pb.ProductKeluarResponse{
		ProductKeluar: pbProductKeluar,
	}, nil
}

func (p *ProductKeluarGrpc) DeleteProductKeluar(ctx context.Context, req *pb.ProductKeluarRequest) (*pb.DeleteProductKeluarResponse, error) {
	productKeluarID := req.GetId()

	_, err := p.productkeluar.Delete(productKeluarID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.DeleteProductKeluarResponse{
		Success: true,
	}, nil
}
