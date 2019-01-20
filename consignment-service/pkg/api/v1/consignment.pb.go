// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consignment.proto

package consignment_proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ConsignmentRequest struct {
	Api                  string       `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Consignment          *Consignment `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ConsignmentRequest) Reset()         { *m = ConsignmentRequest{} }
func (m *ConsignmentRequest) String() string { return proto.CompactTextString(m) }
func (*ConsignmentRequest) ProtoMessage()    {}
func (*ConsignmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{0}
}

func (m *ConsignmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsignmentRequest.Unmarshal(m, b)
}
func (m *ConsignmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsignmentRequest.Marshal(b, m, deterministic)
}
func (m *ConsignmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsignmentRequest.Merge(m, src)
}
func (m *ConsignmentRequest) XXX_Size() int {
	return xxx_messageInfo_ConsignmentRequest.Size(m)
}
func (m *ConsignmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsignmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConsignmentRequest proto.InternalMessageInfo

func (m *ConsignmentRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ConsignmentRequest) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{1}
}

func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (m *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(m, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{2}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ConsignmentResponse struct {
	Api                  string       `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Created              bool         `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment `protobuf:"bytes,3,opt,name=consignment,proto3" json:"consignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ConsignmentResponse) Reset()         { *m = ConsignmentResponse{} }
func (m *ConsignmentResponse) String() string { return proto.CompactTextString(m) }
func (*ConsignmentResponse) ProtoMessage()    {}
func (*ConsignmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{3}
}

func (m *ConsignmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsignmentResponse.Unmarshal(m, b)
}
func (m *ConsignmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsignmentResponse.Marshal(b, m, deterministic)
}
func (m *ConsignmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsignmentResponse.Merge(m, src)
}
func (m *ConsignmentResponse) XXX_Size() int {
	return xxx_messageInfo_ConsignmentResponse.Size(m)
}
func (m *ConsignmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsignmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConsignmentResponse proto.InternalMessageInfo

func (m *ConsignmentResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ConsignmentResponse) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *ConsignmentResponse) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func init() {
	proto.RegisterType((*ConsignmentRequest)(nil), "consignment_proto.ConsignmentRequest")
	proto.RegisterType((*Consignment)(nil), "consignment_proto.Consignment")
	proto.RegisterType((*Container)(nil), "consignment_proto.Container")
	proto.RegisterType((*ConsignmentResponse)(nil), "consignment_proto.ConsignmentResponse")
}

func init() { proto.RegisterFile("consignment.proto", fileDescriptor_3804bf87090b51a9) }

var fileDescriptor_3804bf87090b51a9 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x36, 0x4d, 0xff, 0x32, 0x01, 0xb5, 0x2b, 0x68, 0x50, 0xd1, 0x10, 0x50, 0x72, 0xea, 0xa1,
	0x5e, 0x3d, 0x08, 0x3d, 0xf5, 0xba, 0x7d, 0x80, 0x52, 0xb3, 0x43, 0x3a, 0x60, 0x77, 0xe3, 0xee,
	0xb6, 0x7d, 0x00, 0xdf, 0xc8, 0x27, 0x94, 0x64, 0x5b, 0x5d, 0x6d, 0xa1, 0x78, 0xcb, 0x7c, 0x33,
	0xdf, 0x4f, 0xf8, 0x16, 0x06, 0x85, 0x92, 0x86, 0x4a, 0xb9, 0x44, 0x69, 0x87, 0x95, 0x56, 0x56,
	0x31, 0x1f, 0x9a, 0x35, 0x50, 0xb6, 0x00, 0x36, 0xfe, 0x01, 0x39, 0xbe, 0xaf, 0xd0, 0x58, 0x76,
	0x0e, 0xe1, 0xbc, 0xa2, 0x24, 0x48, 0x83, 0x3c, 0xe2, 0xf5, 0x27, 0x7b, 0x81, 0xd8, 0x23, 0x27,
	0xad, 0x34, 0xc8, 0xe3, 0xd1, 0xdd, 0x70, 0x4f, 0x70, 0xe8, 0xab, 0xf9, 0x94, 0xec, 0x33, 0x80,
	0xd8, 0x5b, 0xb2, 0x53, 0x68, 0x91, 0xd8, 0x5a, 0xb4, 0x48, 0xb0, 0x14, 0x62, 0x81, 0xa6, 0xd0,
	0x54, 0x59, 0x52, 0xb2, 0x71, 0x88, 0xb8, 0x0f, 0xb1, 0x4b, 0xe8, 0x6e, 0x90, 0xca, 0x85, 0x4d,
	0xc2, 0x34, 0xc8, 0x3b, 0x7c, 0x3b, 0xb1, 0x67, 0x80, 0x42, 0x49, 0x3b, 0x27, 0x89, 0xda, 0x24,
	0xed, 0x34, 0xcc, 0xe3, 0xd1, 0xed, 0xe1, 0x68, 0xee, 0x88, 0x7b, 0xf7, 0xec, 0x06, 0xa2, 0x35,
	0x1a, 0x83, 0x6f, 0x33, 0x12, 0x49, 0xa7, 0x71, 0xed, 0x3b, 0x60, 0x22, 0xb2, 0x25, 0x44, 0xdf,
	0xac, 0xbd, 0xc4, 0xf7, 0x10, 0x17, 0x2b, 0x63, 0xd5, 0x12, 0x75, 0xcd, 0x75, 0x89, 0x61, 0x07,
	0x4d, 0x44, 0x1d, 0x58, 0x69, 0x2a, 0x49, 0x36, 0x81, 0x23, 0xbe, 0x9d, 0xd8, 0x15, 0xf4, 0x56,
	0xc6, 0x91, 0xda, 0x6e, 0x51, 0x8f, 0x13, 0x91, 0x7d, 0x04, 0x70, 0xf1, 0xab, 0x0e, 0x53, 0x29,
	0x69, 0xf0, 0x40, 0x1f, 0x09, 0xf4, 0x0a, 0x8d, 0x73, 0x8b, 0xce, 0xb7, 0xcf, 0x77, 0xe3, 0xdf,
	0xa6, 0xc2, 0x7f, 0x37, 0x35, 0xda, 0xc0, 0xd9, 0x74, 0x41, 0x55, 0x45, 0xb2, 0x9c, 0xa2, 0x5e,
	0x53, 0x81, 0x4c, 0xc0, 0x60, 0xdc, 0xe8, 0xfb, 0x0d, 0x3e, 0x1c, 0x11, 0x75, 0x8f, 0xe9, 0xfa,
	0xf1, 0xd8, 0x99, 0xfb, 0xc9, 0xec, 0xe4, 0xb5, 0xdb, 0x2c, 0x9f, 0xbe, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xf7, 0x7e, 0xfb, 0xdb, 0xbb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShippingServiceClient is the client API for ShippingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *ConsignmentRequest, opts ...grpc.CallOption) (*ConsignmentResponse, error)
}

type shippingServiceClient struct {
	cc *grpc.ClientConn
}

func NewShippingServiceClient(cc *grpc.ClientConn) ShippingServiceClient {
	return &shippingServiceClient{cc}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *ConsignmentRequest, opts ...grpc.CallOption) (*ConsignmentResponse, error) {
	out := new(ConsignmentResponse)
	err := c.cc.Invoke(ctx, "/consignment_proto.ShippingService/CreateConsignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingServiceServer is the server API for ShippingService service.
type ShippingServiceServer interface {
	CreateConsignment(context.Context, *ConsignmentRequest) (*ConsignmentResponse, error)
}

func RegisterShippingServiceServer(s *grpc.Server, srv ShippingServiceServer) {
	s.RegisterService(&_ShippingService_serviceDesc, srv)
}

func _ShippingService_CreateConsignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsignmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).CreateConsignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consignment_proto.ShippingService/CreateConsignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).CreateConsignment(ctx, req.(*ConsignmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShippingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "consignment_proto.ShippingService",
	HandlerType: (*ShippingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConsignment",
			Handler:    _ShippingService_CreateConsignment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consignment.proto",
}