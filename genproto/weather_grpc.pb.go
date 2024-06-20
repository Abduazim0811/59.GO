// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: weather.proto

package genproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	WeatherService_GetWeatherUpdates_FullMethodName = "/WeatherService/GetWeatherUpdates"
)

// WeatherServiceClient is the client API for WeatherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeatherServiceClient interface {
	GetWeatherUpdates(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (WeatherService_GetWeatherUpdatesClient, error)
}

type weatherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherServiceClient(cc grpc.ClientConnInterface) WeatherServiceClient {
	return &weatherServiceClient{cc}
}

func (c *weatherServiceClient) GetWeatherUpdates(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (WeatherService_GetWeatherUpdatesClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &WeatherService_ServiceDesc.Streams[0], WeatherService_GetWeatherUpdates_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &weatherServiceGetWeatherUpdatesClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WeatherService_GetWeatherUpdatesClient interface {
	Recv() (*WeatherResponce, error)
	grpc.ClientStream
}

type weatherServiceGetWeatherUpdatesClient struct {
	grpc.ClientStream
}

func (x *weatherServiceGetWeatherUpdatesClient) Recv() (*WeatherResponce, error) {
	m := new(WeatherResponce)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WeatherServiceServer is the server API for WeatherService service.
// All implementations must embed UnimplementedWeatherServiceServer
// for forward compatibility
type WeatherServiceServer interface {
	GetWeatherUpdates(*WeatherRequest, WeatherService_GetWeatherUpdatesServer) error
	mustEmbedUnimplementedWeatherServiceServer()
}

// UnimplementedWeatherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWeatherServiceServer struct {
}

func (UnimplementedWeatherServiceServer) GetWeatherUpdates(*WeatherRequest, WeatherService_GetWeatherUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetWeatherUpdates not implemented")
}
func (UnimplementedWeatherServiceServer) mustEmbedUnimplementedWeatherServiceServer() {}

// UnsafeWeatherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeatherServiceServer will
// result in compilation errors.
type UnsafeWeatherServiceServer interface {
	mustEmbedUnimplementedWeatherServiceServer()
}

func RegisterWeatherServiceServer(s grpc.ServiceRegistrar, srv WeatherServiceServer) {
	s.RegisterService(&WeatherService_ServiceDesc, srv)
}

func _WeatherService_GetWeatherUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WeatherRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WeatherServiceServer).GetWeatherUpdates(m, &weatherServiceGetWeatherUpdatesServer{ServerStream: stream})
}

type WeatherService_GetWeatherUpdatesServer interface {
	Send(*WeatherResponce) error
	grpc.ServerStream
}

type weatherServiceGetWeatherUpdatesServer struct {
	grpc.ServerStream
}

func (x *weatherServiceGetWeatherUpdatesServer) Send(m *WeatherResponce) error {
	return x.ServerStream.SendMsg(m)
}

// WeatherService_ServiceDesc is the grpc.ServiceDesc for WeatherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeatherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "WeatherService",
	HandlerType: (*WeatherServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetWeatherUpdates",
			Handler:       _WeatherService_GetWeatherUpdates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "weather.proto",
}
