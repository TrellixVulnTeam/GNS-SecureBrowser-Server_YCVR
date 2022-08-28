// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/gnsservice.proto

package gnsrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GNSBadgeDataClient is the client API for GNSBadgeData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GNSBadgeDataClient interface {
	// Get HW UUID of card
	ReadUUID(ctx context.Context, in *GNSBadgeDataParam, opts ...grpc.CallOption) (*UUID, error)
	// Get old UUID of card on Zone 2
	ReadUUIDZone2(ctx context.Context, in *GNSBadgeDataParam, opts ...grpc.CallOption) (*UUID, error)
	// Read HW UUID then store to Zone3
	StoreUUID(ctx context.Context, in *GNSBadgeDataParam, opts ...grpc.CallOption) (*GNSBadgeDataParam, error)
	// Format card 0: sites + wincreds, 1: sites only, 2: wincreds
	FormatCard(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*GNSBadgeDataParam, error)
	// Get available free indexes for sites (32 max)
	GetFreeSites(ctx context.Context, in *GNSBadgeDataParam, opts ...grpc.CallOption) (*FreeSites, error)
	// Get available windows credential (8 max)
	GetFreeWinCreds(ctx context.Context, in *GNSBadgeDataParam, opts ...grpc.CallOption) (*FreeWinCreds, error)
	// Read site credentials from badge
	ReadSiteCreds(ctx context.Context, in *GNSBadgeDataParam, opts ...grpc.CallOption) (*Sites, error)
	// read 1 site cred
	ReadSiteCred(ctx context.Context, in *SiteCred, opts ...grpc.CallOption) (*SiteCred, error)
	// Read windows credentials from badge
	ReadWinCreds(ctx context.Context, in *GNSBadgeDataParam, opts ...grpc.CallOption) (*WinCreds, error)
	// Read 1 windows credential from badge
	ReadWinCred(ctx context.Context, in *WinCred, opts ...grpc.CallOption) (*WinCred, error)
	// Delete 1 site credential
	DeleteSiteCred(ctx context.Context, in *SiteCred, opts ...grpc.CallOption) (*GNSBadgeDataParam, error)
	// Delete 1 Windows credential
	DeleteWinCred(ctx context.Context, in *WinCred, opts ...grpc.CallOption) (*GNSBadgeDataParam, error)
	// Write a site credential at location of offset
	WriteSiteCred(ctx context.Context, in *SiteCred, opts ...grpc.CallOption) (*GNSBadgeDataParam, error)
	// Write a windows credential at location of idx
	WriteWinCred(ctx context.Context, in *WinCred, opts ...grpc.CallOption) (*GNSBadgeDataParam, error)
	// Ping-pong like in-out data stream to report CardStatus changes to RPC
	// client
	StreamCardStatus(ctx context.Context, in *GNSBadgeDataParam, opts ...grpc.CallOption) (GNSBadgeData_StreamCardStatusClient, error)
	// Unlock card and receive hardware UUID
	UnlockCard(ctx context.Context, in *GNSBadgeDataParam, opts ...grpc.CallOption) (*Text, error)
	// arbitrary client to server command to implement commands like switching
	// between USB vs NFC
	Execute(ctx context.Context, in *Text, opts ...grpc.CallOption) (*GNSBadgeDataParam, error)
}

type gNSBadgeDataClient struct {
	cc grpc.ClientConnInterface
}

func NewGNSBadgeDataClient(cc grpc.ClientConnInterface) GNSBadgeDataClient {
	return &gNSBadgeDataClient{cc}
}

func (c *gNSBadgeDataClient) ReadUUID(ctx context.Context, in *GNSBadgeDataParam, opts ...grpc.CallOption) (*UUID, error) {
	out := new(UUID)
	err := c.cc.Invoke(ctx, "/GNSRPC.GNSBadgeData/ReadUUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gNSBadgeDataClient) ReadUUIDZone2(ctx context.Context, in *GNSBadgeDataParam, opts ...grpc.CallOption) (*UUID, error) {
	out := new(UUID)
	err := c.cc.Invoke(ctx, "/GNSRPC.GNSBadgeData/ReadUUIDZone2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gNSBadgeDataClient) StoreUUID(ctx context.Context, in *GNSBadgeDataParam, opts ...grpc.CallOption) (*GNSBadgeDataParam, error) {
	out := new(GNSBadgeDataParam)
	err := c.cc.Invoke(ctx, "/GNSRPC.GNSBadgeData/StoreUUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gNSBadgeDataClient) FormatCard(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*GNSBadgeDataParam, error) {
	out := new(GNSBadgeDataParam)
	err := c.cc.Invoke(ctx, "/GNSRPC.GNSBadgeData/FormatCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gNSBadgeDataClient) GetFreeSites(ctx context.Context, in *GNSBadgeDataParam, opts ...grpc.CallOption) (*FreeSites, error) {
	out := new(FreeSites)
	err := c.cc.Invoke(ctx, "/GNSRPC.GNSBadgeData/GetFreeSites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gNSBadgeDataClient) GetFreeWinCreds(ctx context.Context, in *GNSBadgeDataParam, opts ...grpc.CallOption) (*FreeWinCreds, error) {
	out := new(FreeWinCreds)
	err := c.cc.Invoke(ctx, "/GNSRPC.GNSBadgeData/GetFreeWinCreds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gNSBadgeDataClient) ReadSiteCreds(ctx context.Context, in *GNSBadgeDataParam, opts ...grpc.CallOption) (*Sites, error) {
	out := new(Sites)
	err := c.cc.Invoke(ctx, "/GNSRPC.GNSBadgeData/ReadSiteCreds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gNSBadgeDataClient) ReadSiteCred(ctx context.Context, in *SiteCred, opts ...grpc.CallOption) (*SiteCred, error) {
	out := new(SiteCred)
	err := c.cc.Invoke(ctx, "/GNSRPC.GNSBadgeData/ReadSiteCred", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gNSBadgeDataClient) ReadWinCreds(ctx context.Context, in *GNSBadgeDataParam, opts ...grpc.CallOption) (*WinCreds, error) {
	out := new(WinCreds)
	err := c.cc.Invoke(ctx, "/GNSRPC.GNSBadgeData/ReadWinCreds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gNSBadgeDataClient) ReadWinCred(ctx context.Context, in *WinCred, opts ...grpc.CallOption) (*WinCred, error) {
	out := new(WinCred)
	err := c.cc.Invoke(ctx, "/GNSRPC.GNSBadgeData/ReadWinCred", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gNSBadgeDataClient) DeleteSiteCred(ctx context.Context, in *SiteCred, opts ...grpc.CallOption) (*GNSBadgeDataParam, error) {
	out := new(GNSBadgeDataParam)
	err := c.cc.Invoke(ctx, "/GNSRPC.GNSBadgeData/DeleteSiteCred", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gNSBadgeDataClient) DeleteWinCred(ctx context.Context, in *WinCred, opts ...grpc.CallOption) (*GNSBadgeDataParam, error) {
	out := new(GNSBadgeDataParam)
	err := c.cc.Invoke(ctx, "/GNSRPC.GNSBadgeData/DeleteWinCred", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gNSBadgeDataClient) WriteSiteCred(ctx context.Context, in *SiteCred, opts ...grpc.CallOption) (*GNSBadgeDataParam, error) {
	out := new(GNSBadgeDataParam)
	err := c.cc.Invoke(ctx, "/GNSRPC.GNSBadgeData/WriteSiteCred", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gNSBadgeDataClient) WriteWinCred(ctx context.Context, in *WinCred, opts ...grpc.CallOption) (*GNSBadgeDataParam, error) {
	out := new(GNSBadgeDataParam)
	err := c.cc.Invoke(ctx, "/GNSRPC.GNSBadgeData/WriteWinCred", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gNSBadgeDataClient) StreamCardStatus(ctx context.Context, in *GNSBadgeDataParam, opts ...grpc.CallOption) (GNSBadgeData_StreamCardStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &GNSBadgeData_ServiceDesc.Streams[0], "/GNSRPC.GNSBadgeData/StreamCardStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &gNSBadgeDataStreamCardStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GNSBadgeData_StreamCardStatusClient interface {
	Recv() (*CardStatus, error)
	grpc.ClientStream
}

type gNSBadgeDataStreamCardStatusClient struct {
	grpc.ClientStream
}

func (x *gNSBadgeDataStreamCardStatusClient) Recv() (*CardStatus, error) {
	m := new(CardStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gNSBadgeDataClient) UnlockCard(ctx context.Context, in *GNSBadgeDataParam, opts ...grpc.CallOption) (*Text, error) {
	out := new(Text)
	err := c.cc.Invoke(ctx, "/GNSRPC.GNSBadgeData/UnlockCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gNSBadgeDataClient) Execute(ctx context.Context, in *Text, opts ...grpc.CallOption) (*GNSBadgeDataParam, error) {
	out := new(GNSBadgeDataParam)
	err := c.cc.Invoke(ctx, "/GNSRPC.GNSBadgeData/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GNSBadgeDataServer is the server API for GNSBadgeData service.
// All implementations must embed UnimplementedGNSBadgeDataServer
// for forward compatibility
type GNSBadgeDataServer interface {
	// Get HW UUID of card
	ReadUUID(context.Context, *GNSBadgeDataParam) (*UUID, error)
	// Get old UUID of card on Zone 2
	ReadUUIDZone2(context.Context, *GNSBadgeDataParam) (*UUID, error)
	// Read HW UUID then store to Zone3
	StoreUUID(context.Context, *GNSBadgeDataParam) (*GNSBadgeDataParam, error)
	// Format card 0: sites + wincreds, 1: sites only, 2: wincreds
	FormatCard(context.Context, *UUID) (*GNSBadgeDataParam, error)
	// Get available free indexes for sites (32 max)
	GetFreeSites(context.Context, *GNSBadgeDataParam) (*FreeSites, error)
	// Get available windows credential (8 max)
	GetFreeWinCreds(context.Context, *GNSBadgeDataParam) (*FreeWinCreds, error)
	// Read site credentials from badge
	ReadSiteCreds(context.Context, *GNSBadgeDataParam) (*Sites, error)
	// read 1 site cred
	ReadSiteCred(context.Context, *SiteCred) (*SiteCred, error)
	// Read windows credentials from badge
	ReadWinCreds(context.Context, *GNSBadgeDataParam) (*WinCreds, error)
	// Read 1 windows credential from badge
	ReadWinCred(context.Context, *WinCred) (*WinCred, error)
	// Delete 1 site credential
	DeleteSiteCred(context.Context, *SiteCred) (*GNSBadgeDataParam, error)
	// Delete 1 Windows credential
	DeleteWinCred(context.Context, *WinCred) (*GNSBadgeDataParam, error)
	// Write a site credential at location of offset
	WriteSiteCred(context.Context, *SiteCred) (*GNSBadgeDataParam, error)
	// Write a windows credential at location of idx
	WriteWinCred(context.Context, *WinCred) (*GNSBadgeDataParam, error)
	// Ping-pong like in-out data stream to report CardStatus changes to RPC
	// client
	StreamCardStatus(*GNSBadgeDataParam, GNSBadgeData_StreamCardStatusServer) error
	// Unlock card and receive hardware UUID
	UnlockCard(context.Context, *GNSBadgeDataParam) (*Text, error)
	// arbitrary client to server command to implement commands like switching
	// between USB vs NFC
	Execute(context.Context, *Text) (*GNSBadgeDataParam, error)
	mustEmbedUnimplementedGNSBadgeDataServer()
}

// UnimplementedGNSBadgeDataServer must be embedded to have forward compatible implementations.
type UnimplementedGNSBadgeDataServer struct {
}

func (UnimplementedGNSBadgeDataServer) ReadUUID(context.Context, *GNSBadgeDataParam) (*UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadUUID not implemented")
}
func (UnimplementedGNSBadgeDataServer) ReadUUIDZone2(context.Context, *GNSBadgeDataParam) (*UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadUUIDZone2 not implemented")
}
func (UnimplementedGNSBadgeDataServer) StoreUUID(context.Context, *GNSBadgeDataParam) (*GNSBadgeDataParam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreUUID not implemented")
}
func (UnimplementedGNSBadgeDataServer) FormatCard(context.Context, *UUID) (*GNSBadgeDataParam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FormatCard not implemented")
}
func (UnimplementedGNSBadgeDataServer) GetFreeSites(context.Context, *GNSBadgeDataParam) (*FreeSites, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFreeSites not implemented")
}
func (UnimplementedGNSBadgeDataServer) GetFreeWinCreds(context.Context, *GNSBadgeDataParam) (*FreeWinCreds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFreeWinCreds not implemented")
}
func (UnimplementedGNSBadgeDataServer) ReadSiteCreds(context.Context, *GNSBadgeDataParam) (*Sites, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadSiteCreds not implemented")
}
func (UnimplementedGNSBadgeDataServer) ReadSiteCred(context.Context, *SiteCred) (*SiteCred, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadSiteCred not implemented")
}
func (UnimplementedGNSBadgeDataServer) ReadWinCreds(context.Context, *GNSBadgeDataParam) (*WinCreds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadWinCreds not implemented")
}
func (UnimplementedGNSBadgeDataServer) ReadWinCred(context.Context, *WinCred) (*WinCred, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadWinCred not implemented")
}
func (UnimplementedGNSBadgeDataServer) DeleteSiteCred(context.Context, *SiteCred) (*GNSBadgeDataParam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSiteCred not implemented")
}
func (UnimplementedGNSBadgeDataServer) DeleteWinCred(context.Context, *WinCred) (*GNSBadgeDataParam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWinCred not implemented")
}
func (UnimplementedGNSBadgeDataServer) WriteSiteCred(context.Context, *SiteCred) (*GNSBadgeDataParam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteSiteCred not implemented")
}
func (UnimplementedGNSBadgeDataServer) WriteWinCred(context.Context, *WinCred) (*GNSBadgeDataParam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteWinCred not implemented")
}
func (UnimplementedGNSBadgeDataServer) StreamCardStatus(*GNSBadgeDataParam, GNSBadgeData_StreamCardStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamCardStatus not implemented")
}
func (UnimplementedGNSBadgeDataServer) UnlockCard(context.Context, *GNSBadgeDataParam) (*Text, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlockCard not implemented")
}
func (UnimplementedGNSBadgeDataServer) Execute(context.Context, *Text) (*GNSBadgeDataParam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedGNSBadgeDataServer) mustEmbedUnimplementedGNSBadgeDataServer() {}

// UnsafeGNSBadgeDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GNSBadgeDataServer will
// result in compilation errors.
type UnsafeGNSBadgeDataServer interface {
	mustEmbedUnimplementedGNSBadgeDataServer()
}

func RegisterGNSBadgeDataServer(s grpc.ServiceRegistrar, srv GNSBadgeDataServer) {
	s.RegisterService(&GNSBadgeData_ServiceDesc, srv)
}

func _GNSBadgeData_ReadUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GNSBadgeDataParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GNSBadgeDataServer).ReadUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GNSRPC.GNSBadgeData/ReadUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GNSBadgeDataServer).ReadUUID(ctx, req.(*GNSBadgeDataParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GNSBadgeData_ReadUUIDZone2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GNSBadgeDataParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GNSBadgeDataServer).ReadUUIDZone2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GNSRPC.GNSBadgeData/ReadUUIDZone2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GNSBadgeDataServer).ReadUUIDZone2(ctx, req.(*GNSBadgeDataParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GNSBadgeData_StoreUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GNSBadgeDataParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GNSBadgeDataServer).StoreUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GNSRPC.GNSBadgeData/StoreUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GNSBadgeDataServer).StoreUUID(ctx, req.(*GNSBadgeDataParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GNSBadgeData_FormatCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GNSBadgeDataServer).FormatCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GNSRPC.GNSBadgeData/FormatCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GNSBadgeDataServer).FormatCard(ctx, req.(*UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GNSBadgeData_GetFreeSites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GNSBadgeDataParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GNSBadgeDataServer).GetFreeSites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GNSRPC.GNSBadgeData/GetFreeSites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GNSBadgeDataServer).GetFreeSites(ctx, req.(*GNSBadgeDataParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GNSBadgeData_GetFreeWinCreds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GNSBadgeDataParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GNSBadgeDataServer).GetFreeWinCreds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GNSRPC.GNSBadgeData/GetFreeWinCreds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GNSBadgeDataServer).GetFreeWinCreds(ctx, req.(*GNSBadgeDataParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GNSBadgeData_ReadSiteCreds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GNSBadgeDataParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GNSBadgeDataServer).ReadSiteCreds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GNSRPC.GNSBadgeData/ReadSiteCreds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GNSBadgeDataServer).ReadSiteCreds(ctx, req.(*GNSBadgeDataParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GNSBadgeData_ReadSiteCred_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SiteCred)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GNSBadgeDataServer).ReadSiteCred(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GNSRPC.GNSBadgeData/ReadSiteCred",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GNSBadgeDataServer).ReadSiteCred(ctx, req.(*SiteCred))
	}
	return interceptor(ctx, in, info, handler)
}

func _GNSBadgeData_ReadWinCreds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GNSBadgeDataParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GNSBadgeDataServer).ReadWinCreds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GNSRPC.GNSBadgeData/ReadWinCreds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GNSBadgeDataServer).ReadWinCreds(ctx, req.(*GNSBadgeDataParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GNSBadgeData_ReadWinCred_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WinCred)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GNSBadgeDataServer).ReadWinCred(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GNSRPC.GNSBadgeData/ReadWinCred",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GNSBadgeDataServer).ReadWinCred(ctx, req.(*WinCred))
	}
	return interceptor(ctx, in, info, handler)
}

func _GNSBadgeData_DeleteSiteCred_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SiteCred)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GNSBadgeDataServer).DeleteSiteCred(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GNSRPC.GNSBadgeData/DeleteSiteCred",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GNSBadgeDataServer).DeleteSiteCred(ctx, req.(*SiteCred))
	}
	return interceptor(ctx, in, info, handler)
}

func _GNSBadgeData_DeleteWinCred_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WinCred)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GNSBadgeDataServer).DeleteWinCred(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GNSRPC.GNSBadgeData/DeleteWinCred",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GNSBadgeDataServer).DeleteWinCred(ctx, req.(*WinCred))
	}
	return interceptor(ctx, in, info, handler)
}

func _GNSBadgeData_WriteSiteCred_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SiteCred)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GNSBadgeDataServer).WriteSiteCred(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GNSRPC.GNSBadgeData/WriteSiteCred",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GNSBadgeDataServer).WriteSiteCred(ctx, req.(*SiteCred))
	}
	return interceptor(ctx, in, info, handler)
}

func _GNSBadgeData_WriteWinCred_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WinCred)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GNSBadgeDataServer).WriteWinCred(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GNSRPC.GNSBadgeData/WriteWinCred",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GNSBadgeDataServer).WriteWinCred(ctx, req.(*WinCred))
	}
	return interceptor(ctx, in, info, handler)
}

func _GNSBadgeData_StreamCardStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GNSBadgeDataParam)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GNSBadgeDataServer).StreamCardStatus(m, &gNSBadgeDataStreamCardStatusServer{stream})
}

type GNSBadgeData_StreamCardStatusServer interface {
	Send(*CardStatus) error
	grpc.ServerStream
}

type gNSBadgeDataStreamCardStatusServer struct {
	grpc.ServerStream
}

func (x *gNSBadgeDataStreamCardStatusServer) Send(m *CardStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _GNSBadgeData_UnlockCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GNSBadgeDataParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GNSBadgeDataServer).UnlockCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GNSRPC.GNSBadgeData/UnlockCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GNSBadgeDataServer).UnlockCard(ctx, req.(*GNSBadgeDataParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GNSBadgeData_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Text)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GNSBadgeDataServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GNSRPC.GNSBadgeData/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GNSBadgeDataServer).Execute(ctx, req.(*Text))
	}
	return interceptor(ctx, in, info, handler)
}

// GNSBadgeData_ServiceDesc is the grpc.ServiceDesc for GNSBadgeData service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GNSBadgeData_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GNSRPC.GNSBadgeData",
	HandlerType: (*GNSBadgeDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadUUID",
			Handler:    _GNSBadgeData_ReadUUID_Handler,
		},
		{
			MethodName: "ReadUUIDZone2",
			Handler:    _GNSBadgeData_ReadUUIDZone2_Handler,
		},
		{
			MethodName: "StoreUUID",
			Handler:    _GNSBadgeData_StoreUUID_Handler,
		},
		{
			MethodName: "FormatCard",
			Handler:    _GNSBadgeData_FormatCard_Handler,
		},
		{
			MethodName: "GetFreeSites",
			Handler:    _GNSBadgeData_GetFreeSites_Handler,
		},
		{
			MethodName: "GetFreeWinCreds",
			Handler:    _GNSBadgeData_GetFreeWinCreds_Handler,
		},
		{
			MethodName: "ReadSiteCreds",
			Handler:    _GNSBadgeData_ReadSiteCreds_Handler,
		},
		{
			MethodName: "ReadSiteCred",
			Handler:    _GNSBadgeData_ReadSiteCred_Handler,
		},
		{
			MethodName: "ReadWinCreds",
			Handler:    _GNSBadgeData_ReadWinCreds_Handler,
		},
		{
			MethodName: "ReadWinCred",
			Handler:    _GNSBadgeData_ReadWinCred_Handler,
		},
		{
			MethodName: "DeleteSiteCred",
			Handler:    _GNSBadgeData_DeleteSiteCred_Handler,
		},
		{
			MethodName: "DeleteWinCred",
			Handler:    _GNSBadgeData_DeleteWinCred_Handler,
		},
		{
			MethodName: "WriteSiteCred",
			Handler:    _GNSBadgeData_WriteSiteCred_Handler,
		},
		{
			MethodName: "WriteWinCred",
			Handler:    _GNSBadgeData_WriteWinCred_Handler,
		},
		{
			MethodName: "UnlockCard",
			Handler:    _GNSBadgeData_UnlockCard_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _GNSBadgeData_Execute_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamCardStatus",
			Handler:       _GNSBadgeData_StreamCardStatus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/gnsservice.proto",
}