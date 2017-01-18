// Code generated by protoc-gen-go.
// source: pb/sourcehub.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/sourcehub.proto

It has these top-level messages:
	GetSourcesRequest
	GetSourcesResponse
	GetValueRequest
	GetValueResponse
	SetValueRequest
	SetValueResponse
	GetKeysRequest
	GetKeysResponse
	SubscribeRequest
	SubscribeResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetSourcesRequest struct {
}

func (m *GetSourcesRequest) Reset()                    { *m = GetSourcesRequest{} }
func (m *GetSourcesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSourcesRequest) ProtoMessage()               {}
func (*GetSourcesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetSourcesResponse struct {
	Source string `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
}

func (m *GetSourcesResponse) Reset()                    { *m = GetSourcesResponse{} }
func (m *GetSourcesResponse) String() string            { return proto.CompactTextString(m) }
func (*GetSourcesResponse) ProtoMessage()               {}
func (*GetSourcesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetSourcesResponse) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type GetValueRequest struct {
	Source string `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
}

func (m *GetValueRequest) Reset()                    { *m = GetValueRequest{} }
func (m *GetValueRequest) String() string            { return proto.CompactTextString(m) }
func (*GetValueRequest) ProtoMessage()               {}
func (*GetValueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetValueRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *GetValueRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetValueResponse struct {
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *GetValueResponse) Reset()                    { *m = GetValueResponse{} }
func (m *GetValueResponse) String() string            { return proto.CompactTextString(m) }
func (*GetValueResponse) ProtoMessage()               {}
func (*GetValueResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetValueResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type SetValueRequest struct {
	Source string `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value  []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *SetValueRequest) Reset()                    { *m = SetValueRequest{} }
func (m *SetValueRequest) String() string            { return proto.CompactTextString(m) }
func (*SetValueRequest) ProtoMessage()               {}
func (*SetValueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SetValueRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *SetValueRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetValueRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type SetValueResponse struct {
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *SetValueResponse) Reset()                    { *m = SetValueResponse{} }
func (m *SetValueResponse) String() string            { return proto.CompactTextString(m) }
func (*SetValueResponse) ProtoMessage()               {}
func (*SetValueResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SetValueResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type GetKeysRequest struct {
	Source string `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
}

func (m *GetKeysRequest) Reset()                    { *m = GetKeysRequest{} }
func (m *GetKeysRequest) String() string            { return proto.CompactTextString(m) }
func (*GetKeysRequest) ProtoMessage()               {}
func (*GetKeysRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetKeysRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type GetKeysResponse struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetKeysResponse) Reset()                    { *m = GetKeysResponse{} }
func (m *GetKeysResponse) String() string            { return proto.CompactTextString(m) }
func (*GetKeysResponse) ProtoMessage()               {}
func (*GetKeysResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetKeysResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type SubscribeRequest struct {
	Source string `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
}

func (m *SubscribeRequest) Reset()                    { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()               {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SubscribeRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *SubscribeRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type SubscribeResponse struct {
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *SubscribeResponse) Reset()                    { *m = SubscribeResponse{} }
func (m *SubscribeResponse) String() string            { return proto.CompactTextString(m) }
func (*SubscribeResponse) ProtoMessage()               {}
func (*SubscribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SubscribeResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*GetSourcesRequest)(nil), "pb.GetSourcesRequest")
	proto.RegisterType((*GetSourcesResponse)(nil), "pb.GetSourcesResponse")
	proto.RegisterType((*GetValueRequest)(nil), "pb.GetValueRequest")
	proto.RegisterType((*GetValueResponse)(nil), "pb.GetValueResponse")
	proto.RegisterType((*SetValueRequest)(nil), "pb.SetValueRequest")
	proto.RegisterType((*SetValueResponse)(nil), "pb.SetValueResponse")
	proto.RegisterType((*GetKeysRequest)(nil), "pb.GetKeysRequest")
	proto.RegisterType((*GetKeysResponse)(nil), "pb.GetKeysResponse")
	proto.RegisterType((*SubscribeRequest)(nil), "pb.SubscribeRequest")
	proto.RegisterType((*SubscribeResponse)(nil), "pb.SubscribeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SourceHub service

type SourceHubClient interface {
	// GetSources responds with a stream of objects representing available sources
	GetSources(ctx context.Context, in *GetSourcesRequest, opts ...grpc.CallOption) (SourceHub_GetSourcesClient, error)
	// GetValue expects a source and key and responds with the associated value
	GetValue(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*GetValueResponse, error)
	// SetValue sets the value for the specified source and key
	SetValue(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*SetValueResponse, error)
	// GetKeys expects a source and responds with a stream of objects representing available keys
	GetKeys(ctx context.Context, in *GetKeysRequest, opts ...grpc.CallOption) (SourceHub_GetKeysClient, error)
	// Subscribe streams updates to a value for a given source and key
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (SourceHub_SubscribeClient, error)
}

type sourceHubClient struct {
	cc *grpc.ClientConn
}

func NewSourceHubClient(cc *grpc.ClientConn) SourceHubClient {
	return &sourceHubClient{cc}
}

func (c *sourceHubClient) GetSources(ctx context.Context, in *GetSourcesRequest, opts ...grpc.CallOption) (SourceHub_GetSourcesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SourceHub_serviceDesc.Streams[0], c.cc, "/pb.SourceHub/GetSources", opts...)
	if err != nil {
		return nil, err
	}
	x := &sourceHubGetSourcesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SourceHub_GetSourcesClient interface {
	Recv() (*GetSourcesResponse, error)
	grpc.ClientStream
}

type sourceHubGetSourcesClient struct {
	grpc.ClientStream
}

func (x *sourceHubGetSourcesClient) Recv() (*GetSourcesResponse, error) {
	m := new(GetSourcesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sourceHubClient) GetValue(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*GetValueResponse, error) {
	out := new(GetValueResponse)
	err := grpc.Invoke(ctx, "/pb.SourceHub/GetValue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceHubClient) SetValue(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*SetValueResponse, error) {
	out := new(SetValueResponse)
	err := grpc.Invoke(ctx, "/pb.SourceHub/SetValue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceHubClient) GetKeys(ctx context.Context, in *GetKeysRequest, opts ...grpc.CallOption) (SourceHub_GetKeysClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SourceHub_serviceDesc.Streams[1], c.cc, "/pb.SourceHub/GetKeys", opts...)
	if err != nil {
		return nil, err
	}
	x := &sourceHubGetKeysClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SourceHub_GetKeysClient interface {
	Recv() (*GetKeysResponse, error)
	grpc.ClientStream
}

type sourceHubGetKeysClient struct {
	grpc.ClientStream
}

func (x *sourceHubGetKeysClient) Recv() (*GetKeysResponse, error) {
	m := new(GetKeysResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sourceHubClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (SourceHub_SubscribeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SourceHub_serviceDesc.Streams[2], c.cc, "/pb.SourceHub/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &sourceHubSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SourceHub_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type sourceHubSubscribeClient struct {
	grpc.ClientStream
}

func (x *sourceHubSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for SourceHub service

type SourceHubServer interface {
	// GetSources responds with a stream of objects representing available sources
	GetSources(*GetSourcesRequest, SourceHub_GetSourcesServer) error
	// GetValue expects a source and key and responds with the associated value
	GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error)
	// SetValue sets the value for the specified source and key
	SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error)
	// GetKeys expects a source and responds with a stream of objects representing available keys
	GetKeys(*GetKeysRequest, SourceHub_GetKeysServer) error
	// Subscribe streams updates to a value for a given source and key
	Subscribe(*SubscribeRequest, SourceHub_SubscribeServer) error
}

func RegisterSourceHubServer(s *grpc.Server, srv SourceHubServer) {
	s.RegisterService(&_SourceHub_serviceDesc, srv)
}

func _SourceHub_GetSources_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetSourcesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SourceHubServer).GetSources(m, &sourceHubGetSourcesServer{stream})
}

type SourceHub_GetSourcesServer interface {
	Send(*GetSourcesResponse) error
	grpc.ServerStream
}

type sourceHubGetSourcesServer struct {
	grpc.ServerStream
}

func (x *sourceHubGetSourcesServer) Send(m *GetSourcesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SourceHub_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceHubServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SourceHub/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceHubServer).GetValue(ctx, req.(*GetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SourceHub_SetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceHubServer).SetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SourceHub/SetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceHubServer).SetValue(ctx, req.(*SetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SourceHub_GetKeys_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetKeysRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SourceHubServer).GetKeys(m, &sourceHubGetKeysServer{stream})
}

type SourceHub_GetKeysServer interface {
	Send(*GetKeysResponse) error
	grpc.ServerStream
}

type sourceHubGetKeysServer struct {
	grpc.ServerStream
}

func (x *sourceHubGetKeysServer) Send(m *GetKeysResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SourceHub_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SourceHubServer).Subscribe(m, &sourceHubSubscribeServer{stream})
}

type SourceHub_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type sourceHubSubscribeServer struct {
	grpc.ServerStream
}

func (x *sourceHubSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _SourceHub_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SourceHub",
	HandlerType: (*SourceHubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetValue",
			Handler:    _SourceHub_GetValue_Handler,
		},
		{
			MethodName: "SetValue",
			Handler:    _SourceHub_SetValue_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetSources",
			Handler:       _SourceHub_GetSources_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetKeys",
			Handler:       _SourceHub_GetKeys_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _SourceHub_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb/sourcehub.proto",
}

func init() { proto.RegisterFile("pb/sourcehub.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x35, 0x2d, 0x56, 0x33, 0x88, 0x4d, 0x27, 0x6d, 0x29, 0x39, 0xc9, 0x7a, 0x89, 0x20, 0x51,
	0x14, 0xf4, 0x60, 0xc1, 0x63, 0x04, 0x4f, 0x66, 0xc1, 0xbb, 0x5b, 0x06, 0x14, 0xc5, 0xc4, 0x6c,
	0x56, 0xe8, 0xff, 0xf2, 0x07, 0x4a, 0xb2, 0x5b, 0x77, 0x93, 0xfa, 0x45, 0x6f, 0x99, 0x97, 0x79,
	0xef, 0x4d, 0xde, 0x23, 0x80, 0x85, 0x38, 0x91, 0xb9, 0x2a, 0x17, 0xf4, 0xa8, 0x44, 0x52, 0x94,
	0x79, 0x95, 0x63, 0xaf, 0x10, 0x2c, 0x84, 0x51, 0x4a, 0x15, 0x6f, 0xde, 0xc8, 0x8c, 0xde, 0x14,
	0xc9, 0x8a, 0x1d, 0x03, 0xba, 0xa0, 0x2c, 0xf2, 0x57, 0x49, 0x38, 0x85, 0x81, 0x56, 0x98, 0x79,
	0x07, 0x5e, 0xec, 0x67, 0x66, 0x62, 0x57, 0x30, 0x4c, 0xa9, 0xba, 0x7f, 0x78, 0x51, 0x64, 0x04,
	0x7e, 0x5a, 0xc5, 0x00, 0xfa, 0xcf, 0xb4, 0x9c, 0xf5, 0x1a, 0xb0, 0x7e, 0x64, 0x31, 0x04, 0x96,
	0x6c, 0x8c, 0xc6, 0xb0, 0xfd, 0x5e, 0x03, 0x0d, 0x79, 0x2f, 0xd3, 0x03, 0xbb, 0x83, 0x21, 0xdf,
	0xd4, 0xc6, 0x4a, 0xf6, 0x5d, 0xc9, 0x18, 0x02, 0xfe, 0x3f, 0xf3, 0x18, 0xf6, 0x53, 0xaa, 0x6e,
	0x69, 0x29, 0xff, 0xf0, 0x66, 0x87, 0x4d, 0x1a, 0x7a, 0xd3, 0x48, 0x9a, 0x73, 0x3c, 0xfb, 0xd5,
	0x73, 0x08, 0xb8, 0x12, 0x72, 0x51, 0x3e, 0x89, 0x0d, 0x32, 0x3b, 0x82, 0x91, 0xc3, 0xfe, 0xed,
	0xee, 0xb3, 0x8f, 0x1e, 0xf8, 0xba, 0xc7, 0x1b, 0x25, 0xf0, 0x1a, 0xc0, 0xf6, 0x8a, 0x93, 0xa4,
	0x10, 0xc9, 0x5a, 0xf9, 0xd1, 0xb4, 0x0b, 0x6b, 0x03, 0xb6, 0x75, 0xea, 0xe1, 0x25, 0xec, 0xae,
	0xda, 0xc2, 0xd0, 0xec, 0xb9, 0x8d, 0x44, 0xe3, 0x36, 0xb8, 0xa2, 0xd6, 0x44, 0xde, 0x22, 0xf2,
	0xef, 0x88, 0x7c, 0x9d, 0x78, 0x01, 0x3b, 0x26, 0x4e, 0x44, 0xa3, 0xed, 0xb4, 0x10, 0x85, 0x2d,
	0xcc, 0xb9, 0x74, 0x0e, 0xfe, 0x57, 0x46, 0xa8, 0xc5, 0x3b, 0x81, 0x47, 0x93, 0x0e, 0x6a, 0xd9,
	0x62, 0xd0, 0xfc, 0x20, 0xe7, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc6, 0xc7, 0x06, 0x05, 0x36,
	0x03, 0x00, 0x00,
}
