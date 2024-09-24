// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: kessel/inventory/v1beta1/resources/k8s_clusters_service.proto

package resources

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	KesselK8SClusterService_CreateK8SCluster_FullMethodName = "/kessel.inventory.v1beta1.resources.KesselK8sClusterService/CreateK8sCluster"
	KesselK8SClusterService_UpdateK8SCluster_FullMethodName = "/kessel.inventory.v1beta1.resources.KesselK8sClusterService/UpdateK8sCluster"
	KesselK8SClusterService_DeleteK8SCluster_FullMethodName = "/kessel.inventory.v1beta1.resources.KesselK8sClusterService/DeleteK8sCluster"
)

// KesselK8SClusterServiceClient is the client API for KesselK8SClusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KesselK8SClusterServiceClient interface {
	CreateK8SCluster(ctx context.Context, in *CreateK8SClusterRequest, opts ...grpc.CallOption) (*CreateK8SClusterResponse, error)
	UpdateK8SCluster(ctx context.Context, in *UpdateK8SClusterRequest, opts ...grpc.CallOption) (*UpdateK8SClusterResponse, error)
	DeleteK8SCluster(ctx context.Context, in *DeleteK8SClusterRequest, opts ...grpc.CallOption) (*DeleteK8SClusterResponse, error)
}

type kesselK8SClusterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKesselK8SClusterServiceClient(cc grpc.ClientConnInterface) KesselK8SClusterServiceClient {
	return &kesselK8SClusterServiceClient{cc}
}

func (c *kesselK8SClusterServiceClient) CreateK8SCluster(ctx context.Context, in *CreateK8SClusterRequest, opts ...grpc.CallOption) (*CreateK8SClusterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateK8SClusterResponse)
	err := c.cc.Invoke(ctx, KesselK8SClusterService_CreateK8SCluster_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kesselK8SClusterServiceClient) UpdateK8SCluster(ctx context.Context, in *UpdateK8SClusterRequest, opts ...grpc.CallOption) (*UpdateK8SClusterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateK8SClusterResponse)
	err := c.cc.Invoke(ctx, KesselK8SClusterService_UpdateK8SCluster_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kesselK8SClusterServiceClient) DeleteK8SCluster(ctx context.Context, in *DeleteK8SClusterRequest, opts ...grpc.CallOption) (*DeleteK8SClusterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteK8SClusterResponse)
	err := c.cc.Invoke(ctx, KesselK8SClusterService_DeleteK8SCluster_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KesselK8SClusterServiceServer is the server API for KesselK8SClusterService service.
// All implementations must embed UnimplementedKesselK8SClusterServiceServer
// for forward compatibility.
type KesselK8SClusterServiceServer interface {
	CreateK8SCluster(context.Context, *CreateK8SClusterRequest) (*CreateK8SClusterResponse, error)
	UpdateK8SCluster(context.Context, *UpdateK8SClusterRequest) (*UpdateK8SClusterResponse, error)
	DeleteK8SCluster(context.Context, *DeleteK8SClusterRequest) (*DeleteK8SClusterResponse, error)
	mustEmbedUnimplementedKesselK8SClusterServiceServer()
}

// UnimplementedKesselK8SClusterServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKesselK8SClusterServiceServer struct{}

func (UnimplementedKesselK8SClusterServiceServer) CreateK8SCluster(context.Context, *CreateK8SClusterRequest) (*CreateK8SClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateK8SCluster not implemented")
}
func (UnimplementedKesselK8SClusterServiceServer) UpdateK8SCluster(context.Context, *UpdateK8SClusterRequest) (*UpdateK8SClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateK8SCluster not implemented")
}
func (UnimplementedKesselK8SClusterServiceServer) DeleteK8SCluster(context.Context, *DeleteK8SClusterRequest) (*DeleteK8SClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteK8SCluster not implemented")
}
func (UnimplementedKesselK8SClusterServiceServer) mustEmbedUnimplementedKesselK8SClusterServiceServer() {
}
func (UnimplementedKesselK8SClusterServiceServer) testEmbeddedByValue() {}

// UnsafeKesselK8SClusterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KesselK8SClusterServiceServer will
// result in compilation errors.
type UnsafeKesselK8SClusterServiceServer interface {
	mustEmbedUnimplementedKesselK8SClusterServiceServer()
}

func RegisterKesselK8SClusterServiceServer(s grpc.ServiceRegistrar, srv KesselK8SClusterServiceServer) {
	// If the following call pancis, it indicates UnimplementedKesselK8SClusterServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KesselK8SClusterService_ServiceDesc, srv)
}

func _KesselK8SClusterService_CreateK8SCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateK8SClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KesselK8SClusterServiceServer).CreateK8SCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KesselK8SClusterService_CreateK8SCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KesselK8SClusterServiceServer).CreateK8SCluster(ctx, req.(*CreateK8SClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KesselK8SClusterService_UpdateK8SCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateK8SClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KesselK8SClusterServiceServer).UpdateK8SCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KesselK8SClusterService_UpdateK8SCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KesselK8SClusterServiceServer).UpdateK8SCluster(ctx, req.(*UpdateK8SClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KesselK8SClusterService_DeleteK8SCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteK8SClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KesselK8SClusterServiceServer).DeleteK8SCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KesselK8SClusterService_DeleteK8SCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KesselK8SClusterServiceServer).DeleteK8SCluster(ctx, req.(*DeleteK8SClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KesselK8SClusterService_ServiceDesc is the grpc.ServiceDesc for KesselK8SClusterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KesselK8SClusterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kessel.inventory.v1beta1.resources.KesselK8sClusterService",
	HandlerType: (*KesselK8SClusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateK8sCluster",
			Handler:    _KesselK8SClusterService_CreateK8SCluster_Handler,
		},
		{
			MethodName: "UpdateK8sCluster",
			Handler:    _KesselK8SClusterService_UpdateK8SCluster_Handler,
		},
		{
			MethodName: "DeleteK8sCluster",
			Handler:    _KesselK8SClusterService_DeleteK8SCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kessel/inventory/v1beta1/resources/k8s_clusters_service.proto",
}
