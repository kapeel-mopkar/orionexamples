package main

import (
	"context"
	"log"

	proto "github.com/carousell/orionexamples/stringsvc/stringproto"
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

	client := proto.NewStringServiceClient(conn)

	countResponse, countErr := client.Count(context.Background(), &proto.CountRequest{
		Msg: "Kapeel Mopkar",
	})

	log.Println(countErr)
	log.Println(countResponse)

	upperResponse, upperErr := client.Upper(context.Background(), &proto.UpperRequest{
		Msg: "Kapeel Mopkar, Persistent",
	})

	log.Println(upperErr)
	log.Println(upperResponse)
}
