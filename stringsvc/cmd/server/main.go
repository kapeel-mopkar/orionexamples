package main

import (
	proto "github.com/carousell/orionexamples/stringsvc/stringproto"

	"github.com/carousell/Orion/orion"
	"github.com/carousell/orionexamples/stringsvc/service"
)

func main() {
	server := orion.GetDefaultServer("StringService")
	proto.RegisterStringServiceOrionServer(service.GetFactory(), server)
	server.Start()
	server.Wait()
}
