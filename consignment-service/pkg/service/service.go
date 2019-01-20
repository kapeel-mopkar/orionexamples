package service

import (
	"context"

	proto "github.com/carousell/orionexamples/consignment-service/pkg/api/v1"
	"github.com/jinzhu/gorm"
)

// svc implements proto.StringServiceServer
type svc struct {
	db *gorm.DB
}

type Consignment struct {
	gorm.Model
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Weight      int64  `json:"weight"`
}

func (s *svc) CreateConsignment(ctx context.Context, req proto.ConsignmentRequest) (proto.ConsignmentResponse, error) {
	consignmentResponse := new(proto.ConsignmentResponse)
	consign : = Consignment{
		
	}
	conn, err := s.db.Create()
}
