package main

import (
	"context"
	"log"

	proto "github.com/carousell/orionexamples/consignmentservice/pkg/api/v1"
	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:9281"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Didnot connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewConsignmentServiceClient(conn)

	consResponse, consErr := client.CreateConsignment(context.Background(), &proto.ConsignmentRequest{
		Consignment: &proto.Consignment{
			Weight:      1000,
			Description: "Consignment 1000",
		},
		Api: "v1",
	})

	log.Println(consErr)
	log.Println(consResponse)
}
