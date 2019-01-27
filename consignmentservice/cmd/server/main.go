package main

import (
	"github.com/carousell/Orion/orion"
	proto "github.com/carousell/orionexamples/consignmentservice/pkg/api/v1"
	"github.com/carousell/orionexamples/consignmentservice/pkg/service"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {
	server := orion.GetDefaultServer("ShippingService")                       // GetDefaultServer("StringService")
	proto.RegisterConsignmentServiceOrionServer(service.GetFactory(), server) // RegisterStringServiceOrionServer(service.GetFactory(), server)
	server.Start()
	server.Wait()
}
