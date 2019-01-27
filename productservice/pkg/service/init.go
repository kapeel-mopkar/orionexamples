package service

import (
	"log"

	"github.com/carousell/Orion/orion"
)

type svcFactory struct{}

func (s *svcFactory) NewService(orion.Server) interface{} {
	return &svc{}
}

func (s *svcFactory) DisposeService(svc interface{}) {
	log.Println("Disposing SVCFactory service")
}

func GetFactory() orion.ServiceFactory {
	return &svcFactory{}
}
