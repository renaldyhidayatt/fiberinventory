package gapi

import (
	"context"
	"fiberinventory/internal/domain"
	"fiberinventory/internal/pb"
	"fiberinventory/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductHandlerGrpc struct {
	pb.UnimplementedProductServiceServer
	product service.ProductService
}

func NewProductHandlerGrpcHandler(product service.ProductService) *ProductHandlerGrpc {
	productServer := ProductHandlerGrpc{
		product: product,
	}
	return &productServer
}

func (p *ProductHandlerGrpc) CreateProduct(ctx context.Context, req *pb.CreateProductInput) (*pb.ProductResponse, error) {
	newProduct := &domain.CreateProductRequest{
		Name:       req.Name,
		Image:      req.Image,
		Qty:        req.Qty,
		CategoryID: req.CategoryId,
	}

	if err := newProduct.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	createdProduct, err := p.product.Create(newProduct)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ProductResponse{
		Product: &pb.Product{
			Id:         createdProduct.ID,
			Name:       createdProduct.Name,
			Image:      createdProduct.Image,
			Qty:        createdProduct.Qty,
			CategoryId: createdProduct.CategoryID,
		},
	}

	return res, nil
}

func (p *ProductHandlerGrpc) GetProducts(ctx context.Context, req *pb.ProductsRequest) (*pb.ProductsResponse, error) {
	products, err := p.product.Results()

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var pbProducts []*pb.Product

	for _, product := range *products {
		pbProduct := &pb.Product{
			Id:    product.ID,
			Name:  product.Name,
			Image: product.Image,
			Qty:   product.Qty,
			Category: &pb.Category{
				Name: product.Category.Name,
			},
		}
		pbProducts = append(pbProducts, pbProduct)
	}

	return &pb.ProductsResponse{
		Products: pbProducts,
	}, nil
}

func (p *ProductHandlerGrpc) GetProduct(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	productID := req.GetId()

	product, err := p.product.Result(productID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ProductResponse{
		Product: &pb.Product{
			Id:         product.ID,
			Name:       product.Name,
			Image:      product.Image,
			Qty:        product.Qty,
			CategoryId: product.CategoryID,
		},
	}

	return res, nil
}

func (p *ProductHandlerGrpc) UpdateProduct(ctx context.Context, req *pb.UpdateProductInput) (*pb.ProductResponse, error) {
	productID := req.GetId()

	productToUpdate := &domain.UpdateProductRequest{
		ID:         productID,
		Name:       req.Name,
		Image:      req.Image,
		Qty:        req.Qty,
		CategoryID: req.CategoryId,
	}

	if err := productToUpdate.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	updatedProduct, err := p.product.Update(productToUpdate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ProductResponse{
		Product: &pb.Product{
			Id:         updatedProduct.ID,
			Name:       updatedProduct.Name,
			Image:      updatedProduct.Image,
			Qty:        updatedProduct.Qty,
			CategoryId: updatedProduct.CategoryID,
		},
	}

	return res, nil
}

func (p *ProductHandlerGrpc) DeleteProduct(ctx context.Context, req *pb.ProductRequest) (*pb.DeleteProductResponse, error) {
	productID := req.GetId()

	_, err := p.product.Delete(productID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteProductResponse{
		Success: true,
	}

	return res, nil
}
