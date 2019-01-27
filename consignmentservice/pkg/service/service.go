package service

import (
	"context"

	proto "github.com/carousell/orionexamples/consignmentservice/pkg/api/v1"
	"github.com/jinzhu/gorm"
)

// svc implements proto.ConsignmentServiceServer
type svc struct {
	//db *gorm.DB
}

type Consignment struct {
	gorm.Model
	ID          int64  `gorm:"column:id; PRIMARY_KEY" json:"id"`
	Description string `json:"description"`
	Weight      int64  `gorm:"column:weight" json:"weight"`
}

func (s *svc) CreateConsignment(ctx context.Context, req proto.ConsignmentRequest) (proto.ConsignmentResponse, error) {
	// consign := Consignment{
	// 	Description: req.Consignment.Description,
	// 	Weight:      req.Consignment.Weight,
	// }
	//s.db.Save(&consign)
	//req.Consignment.Id = consign.ID

	return proto.ConsignmentResponse{
		Api:         req.Api,
		Consignment: req.Consignment,
		Created:     true,
	}, nil
}
