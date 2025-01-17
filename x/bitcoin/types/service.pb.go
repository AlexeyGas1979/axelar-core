// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitcoin/v1beta1/service.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("bitcoin/v1beta1/service.proto", fileDescriptor_6065c0ad9b83e388) }
func init() {
	golang_proto.RegisterFile("bitcoin/v1beta1/service.proto", fileDescriptor_6065c0ad9b83e388)
}

var fileDescriptor_6065c0ad9b83e388 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x3f, 0x6f, 0xd3, 0x4c,
	0x1c, 0xc7, 0x73, 0x8f, 0x1e, 0x55, 0xc2, 0x03, 0x95, 0x4c, 0xa5, 0x96, 0x28, 0x75, 0xc1, 0x12,
	0x04, 0x1a, 0xb0, 0x13, 0xa8, 0x18, 0x3a, 0x52, 0xb1, 0x51, 0x81, 0x92, 0x88, 0x81, 0xc5, 0xba,
	0xb8, 0xbf, 0x9a, 0x53, 0x92, 0x3b, 0x73, 0xf7, 0x73, 0x30, 0x2b, 0x13, 0x13, 0x42, 0x62, 0xea,
	0xdb, 0x60, 0x61, 0x65, 0x64, 0xac, 0xc4, 0xc2, 0x88, 0x12, 0x5e, 0x08, 0xca, 0xf9, 0x0e, 0x90,
	0x5d, 0x93, 0xb0, 0xd9, 0xfe, 0xfe, 0xfb, 0xe8, 0x64, 0x9d, 0xb3, 0x3b, 0x62, 0x18, 0x0b, 0xc6,
	0xc3, 0x59, 0x6f, 0x04, 0x48, 0x7b, 0xa1, 0x02, 0x39, 0x63, 0x31, 0x04, 0xa9, 0x14, 0x28, 0xdc,
	0x4d, 0x23, 0x07, 0x46, 0x6e, 0x6e, 0x25, 0x22, 0x11, 0x5a, 0x0b, 0x97, 0x4f, 0x85, 0xad, 0xd9,
	0x4a, 0x84, 0x48, 0x26, 0x10, 0xd2, 0x94, 0x85, 0x94, 0x73, 0x81, 0x14, 0x99, 0xe0, 0xca, 0xa8,
	0x3b, 0xe5, 0x0d, 0xcc, 0x0b, 0xe5, 0xde, 0xd9, 0x25, 0xc7, 0x39, 0x56, 0xc9, 0xa0, 0xd8, 0x74,
	0x4f, 0x9c, 0xff, 0x1f, 0x33, 0x3e, 0x76, 0x5b, 0x41, 0x69, 0x36, 0x58, 0x7e, 0xee, 0xc3, 0xcb,
	0x0c, 0x14, 0x36, 0x77, 0x6b, 0x54, 0x95, 0x0a, 0xae, 0xc0, 0xdf, 0x7b, 0xf3, 0xf5, 0xc7, 0x87,
	0xff, 0xae, 0xfa, 0x5b, 0x21, 0xcd, 0x61, 0x42, 0x65, 0x68, 0xd7, 0x27, 0x8c, 0x8f, 0x0f, 0xc9,
	0xbe, 0xfb, 0x96, 0x38, 0x9b, 0x47, 0x82, 0x9f, 0x32, 0x39, 0x7d, 0x92, 0x61, 0x2a, 0x18, 0x47,
	0xb7, 0x5d, 0xe9, 0x2c, 0x39, 0xec, 0xf8, 0xad, 0xd5, 0x46, 0xc3, 0xe1, 0x6b, 0x8e, 0x96, 0xbf,
	0x5d, 0xe6, 0x88, 0x8b, 0xc0, 0x12, 0xe5, 0x8c, 0x38, 0x57, 0x9e, 0x09, 0x84, 0x32, 0x4e, 0xa7,
	0xb2, 0x72, 0x81, 0xcb, 0x22, 0xdd, 0x59, 0xcf, 0x6c, 0xb0, 0xda, 0x1a, 0xeb, 0xba, 0xdf, 0x2a,
	0x63, 0xcd, 0x04, 0x42, 0xf4, 0x07, 0xdb, 0x27, 0xe2, 0xec, 0x1c, 0x49, 0xa0, 0x08, 0x4f, 0x81,
	0x9f, 0x30, 0x9e, 0x0c, 0x25, 0xe5, 0xea, 0x14, 0xa4, 0x1a, 0xe6, 0x6e, 0xb7, 0x7a, 0x0c, 0x35,
	0x56, 0x4b, 0xd9, 0xfb, 0x87, 0x84, 0x41, 0x7d, 0xa0, 0x51, 0xbb, 0x7e, 0xa7, 0x72, 0x82, 0x3a,
	0x19, 0xa5, 0x45, 0x34, 0x42, 0x9b, 0x8d, 0x30, 0x5f, 0x92, 0xbf, 0x23, 0xce, 0xe5, 0xa2, 0xfc,
	0x98, 0x2a, 0x04, 0x39, 0xcc, 0xdd, 0x9b, 0x35, 0xeb, 0xd6, 0x60, 0x29, 0xdb, 0x2b, 0x7d, 0x86,
	0xad, 0xa3, 0xd9, 0x6e, 0xf8, 0xd7, 0x6a, 0xd8, 0xa6, 0x3a, 0x50, 0x01, 0xea, 0x83, 0x8a, 0x33,
	0xf8, 0x0b, 0x90, 0x35, 0xac, 0x02, 0xfa, 0xed, 0x5b, 0x13, 0x48, 0xea, 0x80, 0x01, 0x9a, 0x3a,
	0x1b, 0x03, 0x96, 0xf0, 0x61, 0xee, 0x7a, 0x95, 0xfe, 0x42, 0xb0, 0xfb, 0x7b, 0xb5, 0xfa, 0xaa,
	0xdf, 0x5c, 0xb1, 0x84, 0x9b, 0xb9, 0x8f, 0xc4, 0xd9, 0x1e, 0x64, 0xa3, 0x29, 0xc3, 0x47, 0x39,
	0x82, 0xe4, 0x74, 0xb2, 0x2c, 0xa1, 0x98, 0x49, 0x70, 0xc3, 0xea, 0xc0, 0xc5, 0x4e, 0x4b, 0xd4,
	0x5d, 0x3f, 0x60, 0x10, 0x0f, 0x34, 0x62, 0xe0, 0xdf, 0xae, 0x20, 0xea, 0x60, 0x04, 0x26, 0x19,
	0x29, 0x1b, 0x3d, 0x24, 0xfb, 0x0f, 0xfb, 0x5f, 0xe6, 0x1e, 0x39, 0x9f, 0x7b, 0xe4, 0xfb, 0xdc,
	0x23, 0xef, 0x17, 0x5e, 0xe3, 0xf3, 0xc2, 0x23, 0xe7, 0x0b, 0xaf, 0xf1, 0x6d, 0xe1, 0x35, 0x9e,
	0x1f, 0x24, 0x0c, 0x5f, 0x64, 0xa3, 0x20, 0x16, 0x53, 0xd3, 0xca, 0x01, 0x5f, 0x09, 0x39, 0x36,
	0x6f, 0x77, 0x63, 0x21, 0x21, 0xcc, 0x7f, 0x4d, 0xe1, 0xeb, 0x14, 0xd4, 0x68, 0x43, 0x5f, 0x7b,
	0xf7, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x62, 0x02, 0x47, 0xcf, 0x76, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	Link(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*LinkResponse, error)
	ConfirmOutpoint(ctx context.Context, in *ConfirmOutpointRequest, opts ...grpc.CallOption) (*ConfirmOutpointResponse, error)
	VoteConfirmOutpoint(ctx context.Context, in *VoteConfirmOutpointRequest, opts ...grpc.CallOption) (*VoteConfirmOutpointResponse, error)
	CreatePendingTransfersTx(ctx context.Context, in *CreatePendingTransfersTxRequest, opts ...grpc.CallOption) (*CreatePendingTransfersTxResponse, error)
	CreateMasterTx(ctx context.Context, in *CreateMasterTxRequest, opts ...grpc.CallOption) (*CreateMasterTxResponse, error)
	CreateRescueTx(ctx context.Context, in *CreateRescueTxRequest, opts ...grpc.CallOption) (*CreateRescueTxResponse, error)
	SignTx(ctx context.Context, in *SignTxRequest, opts ...grpc.CallOption) (*SignTxResponse, error)
	SubmitExternalSignature(ctx context.Context, in *SubmitExternalSignatureRequest, opts ...grpc.CallOption) (*SubmitExternalSignatureResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) Link(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*LinkResponse, error) {
	out := new(LinkResponse)
	err := c.cc.Invoke(ctx, "/bitcoin.v1beta1.MsgService/Link", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) ConfirmOutpoint(ctx context.Context, in *ConfirmOutpointRequest, opts ...grpc.CallOption) (*ConfirmOutpointResponse, error) {
	out := new(ConfirmOutpointResponse)
	err := c.cc.Invoke(ctx, "/bitcoin.v1beta1.MsgService/ConfirmOutpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) VoteConfirmOutpoint(ctx context.Context, in *VoteConfirmOutpointRequest, opts ...grpc.CallOption) (*VoteConfirmOutpointResponse, error) {
	out := new(VoteConfirmOutpointResponse)
	err := c.cc.Invoke(ctx, "/bitcoin.v1beta1.MsgService/VoteConfirmOutpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) CreatePendingTransfersTx(ctx context.Context, in *CreatePendingTransfersTxRequest, opts ...grpc.CallOption) (*CreatePendingTransfersTxResponse, error) {
	out := new(CreatePendingTransfersTxResponse)
	err := c.cc.Invoke(ctx, "/bitcoin.v1beta1.MsgService/CreatePendingTransfersTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) CreateMasterTx(ctx context.Context, in *CreateMasterTxRequest, opts ...grpc.CallOption) (*CreateMasterTxResponse, error) {
	out := new(CreateMasterTxResponse)
	err := c.cc.Invoke(ctx, "/bitcoin.v1beta1.MsgService/CreateMasterTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) CreateRescueTx(ctx context.Context, in *CreateRescueTxRequest, opts ...grpc.CallOption) (*CreateRescueTxResponse, error) {
	out := new(CreateRescueTxResponse)
	err := c.cc.Invoke(ctx, "/bitcoin.v1beta1.MsgService/CreateRescueTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) SignTx(ctx context.Context, in *SignTxRequest, opts ...grpc.CallOption) (*SignTxResponse, error) {
	out := new(SignTxResponse)
	err := c.cc.Invoke(ctx, "/bitcoin.v1beta1.MsgService/SignTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) SubmitExternalSignature(ctx context.Context, in *SubmitExternalSignatureRequest, opts ...grpc.CallOption) (*SubmitExternalSignatureResponse, error) {
	out := new(SubmitExternalSignatureResponse)
	err := c.cc.Invoke(ctx, "/bitcoin.v1beta1.MsgService/SubmitExternalSignature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	Link(context.Context, *LinkRequest) (*LinkResponse, error)
	ConfirmOutpoint(context.Context, *ConfirmOutpointRequest) (*ConfirmOutpointResponse, error)
	VoteConfirmOutpoint(context.Context, *VoteConfirmOutpointRequest) (*VoteConfirmOutpointResponse, error)
	CreatePendingTransfersTx(context.Context, *CreatePendingTransfersTxRequest) (*CreatePendingTransfersTxResponse, error)
	CreateMasterTx(context.Context, *CreateMasterTxRequest) (*CreateMasterTxResponse, error)
	CreateRescueTx(context.Context, *CreateRescueTxRequest) (*CreateRescueTxResponse, error)
	SignTx(context.Context, *SignTxRequest) (*SignTxResponse, error)
	SubmitExternalSignature(context.Context, *SubmitExternalSignatureRequest) (*SubmitExternalSignatureResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) Link(ctx context.Context, req *LinkRequest) (*LinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Link not implemented")
}
func (*UnimplementedMsgServiceServer) ConfirmOutpoint(ctx context.Context, req *ConfirmOutpointRequest) (*ConfirmOutpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmOutpoint not implemented")
}
func (*UnimplementedMsgServiceServer) VoteConfirmOutpoint(ctx context.Context, req *VoteConfirmOutpointRequest) (*VoteConfirmOutpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteConfirmOutpoint not implemented")
}
func (*UnimplementedMsgServiceServer) CreatePendingTransfersTx(ctx context.Context, req *CreatePendingTransfersTxRequest) (*CreatePendingTransfersTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePendingTransfersTx not implemented")
}
func (*UnimplementedMsgServiceServer) CreateMasterTx(ctx context.Context, req *CreateMasterTxRequest) (*CreateMasterTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMasterTx not implemented")
}
func (*UnimplementedMsgServiceServer) CreateRescueTx(ctx context.Context, req *CreateRescueTxRequest) (*CreateRescueTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRescueTx not implemented")
}
func (*UnimplementedMsgServiceServer) SignTx(ctx context.Context, req *SignTxRequest) (*SignTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignTx not implemented")
}
func (*UnimplementedMsgServiceServer) SubmitExternalSignature(ctx context.Context, req *SubmitExternalSignatureRequest) (*SubmitExternalSignatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitExternalSignature not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_Link_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).Link(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitcoin.v1beta1.MsgService/Link",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).Link(ctx, req.(*LinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_ConfirmOutpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmOutpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ConfirmOutpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitcoin.v1beta1.MsgService/ConfirmOutpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ConfirmOutpoint(ctx, req.(*ConfirmOutpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_VoteConfirmOutpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteConfirmOutpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).VoteConfirmOutpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitcoin.v1beta1.MsgService/VoteConfirmOutpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).VoteConfirmOutpoint(ctx, req.(*VoteConfirmOutpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_CreatePendingTransfersTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePendingTransfersTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).CreatePendingTransfersTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitcoin.v1beta1.MsgService/CreatePendingTransfersTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).CreatePendingTransfersTx(ctx, req.(*CreatePendingTransfersTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_CreateMasterTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMasterTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).CreateMasterTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitcoin.v1beta1.MsgService/CreateMasterTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).CreateMasterTx(ctx, req.(*CreateMasterTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_CreateRescueTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRescueTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).CreateRescueTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitcoin.v1beta1.MsgService/CreateRescueTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).CreateRescueTx(ctx, req.(*CreateRescueTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_SignTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).SignTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitcoin.v1beta1.MsgService/SignTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).SignTx(ctx, req.(*SignTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_SubmitExternalSignature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitExternalSignatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).SubmitExternalSignature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitcoin.v1beta1.MsgService/SubmitExternalSignature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).SubmitExternalSignature(ctx, req.(*SubmitExternalSignatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bitcoin.v1beta1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Link",
			Handler:    _MsgService_Link_Handler,
		},
		{
			MethodName: "ConfirmOutpoint",
			Handler:    _MsgService_ConfirmOutpoint_Handler,
		},
		{
			MethodName: "VoteConfirmOutpoint",
			Handler:    _MsgService_VoteConfirmOutpoint_Handler,
		},
		{
			MethodName: "CreatePendingTransfersTx",
			Handler:    _MsgService_CreatePendingTransfersTx_Handler,
		},
		{
			MethodName: "CreateMasterTx",
			Handler:    _MsgService_CreateMasterTx_Handler,
		},
		{
			MethodName: "CreateRescueTx",
			Handler:    _MsgService_CreateRescueTx_Handler,
		},
		{
			MethodName: "SignTx",
			Handler:    _MsgService_SignTx_Handler,
		},
		{
			MethodName: "SubmitExternalSignature",
			Handler:    _MsgService_SubmitExternalSignature_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bitcoin/v1beta1/service.proto",
}
