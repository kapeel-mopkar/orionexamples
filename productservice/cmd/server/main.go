package main

import (
	"github.com/carousell/Orion/orion"
	proto "github.com/carousell/orionexamples/productservice/pkg/api/v1"
	"github.com/carousell/orionexamples/productservice/pkg/service"
)

func main() {
	server := orion.GetDefaultServer("ProductService")                    // GetDefaultServer("StringService")
	proto.RegisterProductServiceOrionServer(service.GetFactory(), server) // RegisterStringServiceOrionServer(service.GetFactory(), server)
	server.Start()
	server.Wait()
}
