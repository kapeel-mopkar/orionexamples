// Code generated by protoc-gen-orion. DO NOT EDIT.
// source: productservice.proto

package product_proto

import (
	orion "github.com/carousell/Orion/orion"
)

// If you see error please update your orion-protoc-gen by running 'go get -u github.com/carousell/Orion/protoc-gen-orion'
var _ = orion.ProtoGenVersion1_0

// Encoders

// Handlers

// Decoders

//Streams

// RegisterProductServiceOrionServer registers ProductService to Orion server
// Services need to pass either ServiceFactory or ServiceFactoryV2 implementation
func RegisterProductServiceOrionServer(sf interface{}, orionServer orion.Server) error {
	err := orionServer.RegisterService(&_ProductService_serviceDesc, sf)
	if err != nil {
		return err
	}

	return nil
}

// DefaultEncoder
func RegisterProductServiceDefaultEncoder(svr orion.Server, encoder orion.Encoder) {
	orion.RegisterDefaultEncoder(svr, "ProductService", encoder)
}

// DefaultDecoder
func RegisterProductServiceDefaultDecoder(svr orion.Server, decoder orion.Decoder) {
	orion.RegisterDefaultDecoder(svr, "ProductService", decoder)
}

