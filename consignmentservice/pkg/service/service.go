package service

import (
	"context"
	"log"
	"time"

	proto "github.com/carousell/orionexamples/consignmentservice/pkg/api/v1"
	"github.com/jinzhu/gorm"
)

// svc implements proto.ConsignmentServiceServer
type svc struct {
	db *gorm.DB
}

func (s *svc) CreateConsignment(ctx context.Context, req *proto.ConsignmentRequest) (*proto.ConsignmentResponse, error) {
	consign := Consignment{
		Description: req.Consignment.Description,
		Weight:      req.Consignment.Weight,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Containers:  []Container{},
	}
	for _, cont := range req.Consignment.Containers {
		container := Container{
			Origin:     cont.Origin,
			CustomerID: cont.CustomerId,
			UserID:     cont.UserId,
		}
		consign.Containers = append(consign.Containers, container)

	}
	//log.Println(consign)
	dberr := s.db.Create(&consign).Error //.Model(&Consignment{})
	if dberr != nil {
		log.Fatalln("DB Error : ", dberr)
		return nil, nil
	}
	req.Consignment.Id = consign.ID

	return &proto.ConsignmentResponse{
		Api:         req.Api,
		Consignment: req.Consignment,
		Created:     true,
	}, nil
}
