package service

import (
	"context"

	proto "github.com/carousell/orionexamples/productservice/pkg/api/v1"
)

// svc implements proto.ConsignmentServiceServer
type svc struct{}

func (s *svc) CreateProduct(ctx context.Context, req *proto.CreateProductRequest) (*proto.CreateProductResponse, error) {

	return &proto.CreateProductResponse{
		Api:     req.Api,
		Product: req.Product,
		Created: true,
	}, nil
}
