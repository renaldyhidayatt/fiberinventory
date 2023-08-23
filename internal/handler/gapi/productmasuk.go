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

type ProductMasukGrpc struct {
	pb.UnimplementedProductMasukServiceServer
	productmasuk service.ProductMasukService
}

func NewProductMasukHandlerGrpc(productmasuk service.ProductMasukService) *ProductMasukGrpc {
	return &ProductMasukGrpc{
		productmasuk: productmasuk,
	}
}

func (p *ProductMasukGrpc) CreateProductMasuk(ctx context.Context, req *pb.CreateProductMasukInput) (*pb.ProductMasukResponse, error) {
	newProductMasuk := &domain.CreateProductMasukRequest{
		Name:       req.Name,
		Qty:        req.Qty,
		ProductID:  req.ProductId,
		SupplierID: req.SupplierId,
	}

	if err := newProductMasuk.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	createdProductMasuk, err := p.productmasuk.Create(newProductMasuk)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbProductMasuk := &pb.ProductMasuk{
		Id:         createdProductMasuk.ID,
		Name:       createdProductMasuk.Name,
		Qty:        createdProductMasuk.Qty,
		ProductId:  createdProductMasuk.ProductID,
		SupplierId: createdProductMasuk.SupplierID,
		CreatedAt:  timestamppb.New(createdProductMasuk.CreatedAt),
		UpdatedAt:  timestamppb.New(createdProductMasuk.UpdatedAt),
	}

	return &pb.ProductMasukResponse{
		ProductMasuk: pbProductMasuk,
	}, nil
}

func (p *ProductMasukGrpc) GetProductMasuks(ctx context.Context, req *pb.ProductMasuksRequest) (*pb.ProductMasuksResponse, error) {
	productMasuk, err := p.productmasuk.Results()

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var pbProductMasuks []*pb.ProductMasuk

	for _, product := range *productMasuk {
		pbProductMasuk := &pb.ProductMasuk{
			Id:         product.ID,
			Name:       product.Name,
			Qty:        product.Qty,
			ProductId:  product.ProductID,
			SupplierId: product.SupplierID,
			CreatedAt:  timestamppb.New(product.CreatedAt),
			UpdatedAt:  timestamppb.New(product.UpdatedAt),
		}
		pbProductMasuks = append(pbProductMasuks, pbProductMasuk)
	}

	return &pb.ProductMasuksResponse{
		ProductMasuks: pbProductMasuks,
	}, nil
}

func (p *ProductMasukGrpc) GetProductMasuk(ctx context.Context, req *pb.ProductMasukRequest) (*pb.ProductMasukResponse, error) {
	productMasukID := req.GetId()

	productMasuk, err := p.productmasuk.Result(productMasukID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbProductMasuk := &pb.ProductMasuk{
		Id:         productMasuk.ID,
		Name:       productMasuk.Name,
		Qty:        productMasuk.Qty,
		ProductId:  productMasuk.ProductID,
		SupplierId: productMasuk.SupplierID,
		CreatedAt:  timestamppb.New(productMasuk.CreatedAt),
		UpdatedAt:  timestamppb.New(productMasuk.UpdatedAt),
	}

	return &pb.ProductMasukResponse{
		ProductMasuk: pbProductMasuk,
	}, nil
}

func (p *ProductMasukGrpc) UpdateProductMasuk(ctx context.Context, req *pb.UpdateProductMasukInput) (*pb.ProductMasukResponse, error) {
	productMasukID := req.GetId()

	productMasukToUpdate := &domain.UpdateProductMasukRequest{
		ID:         productMasukID,
		Name:       req.Name,
		Qty:        req.Qty,
		ProductID:  req.ProductId,
		SupplierID: req.SupplierId,
	}
	if err := productMasukToUpdate.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	updatedProductMasuk, err := p.productmasuk.Update(productMasukToUpdate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbProductMasuk := &pb.ProductMasuk{
		Id:         updatedProductMasuk.ID,
		Name:       updatedProductMasuk.Name,
		Qty:        updatedProductMasuk.Qty,
		ProductId:  updatedProductMasuk.ProductID,
		SupplierId: updatedProductMasuk.SupplierID,
		CreatedAt:  timestamppb.New(updatedProductMasuk.CreatedAt),
		UpdatedAt:  timestamppb.New(updatedProductMasuk.UpdatedAt),
	}

	return &pb.ProductMasukResponse{
		ProductMasuk: pbProductMasuk,
	}, nil
}

func (p *ProductMasukGrpc) DeleteProductMasuk(ctx context.Context, req *pb.ProductMasukRequest) (*pb.DeleteProductMasukResponse, error) {
	productMasukID := req.GetId()

	_, err := p.productmasuk.Delete(productMasukID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.DeleteProductMasukResponse{
		Success: true,
	}, nil
}
