// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: blockchain/v1alpha/public-release.proto

package blockchain

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_blockchain_v1alpha_public_release_proto protoreflect.FileDescriptor

var file_blockchain_v1alpha_public_release_proto_rawDesc = []byte{
	0x0a, 0x27, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9a, 0x01, 0x0a, 0x0d, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x44, 0x61, 0x74, 0x61, 0x22, 0x3f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x12, 0x10, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x6c, 0x6c, 0x5a,
	0x25, 0x12, 0x23, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2d, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x69, 0x61, 0x6f, 0x62, 0x61, 0x69, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2d, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x3b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_blockchain_v1alpha_public_release_proto_goTypes = []interface{}{
	(*empty.Empty)(nil), // 0: google.protobuf.Empty
	(*GetAllData)(nil),  // 1: blockchain.v1alpha.GetAllData
}
var file_blockchain_v1alpha_public_release_proto_depIdxs = []int32{
	0, // 0: blockchain.v1alpha.PublicRelease.GetAllRelease:input_type -> google.protobuf.Empty
	1, // 1: blockchain.v1alpha.PublicRelease.GetAllRelease:output_type -> blockchain.v1alpha.GetAllData
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_blockchain_v1alpha_public_release_proto_init() }
func file_blockchain_v1alpha_public_release_proto_init() {
	if File_blockchain_v1alpha_public_release_proto != nil {
		return
	}
	file_blockchain_v1alpha_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blockchain_v1alpha_public_release_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blockchain_v1alpha_public_release_proto_goTypes,
		DependencyIndexes: file_blockchain_v1alpha_public_release_proto_depIdxs,
	}.Build()
	File_blockchain_v1alpha_public_release_proto = out.File
	file_blockchain_v1alpha_public_release_proto_rawDesc = nil
	file_blockchain_v1alpha_public_release_proto_goTypes = nil
	file_blockchain_v1alpha_public_release_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PublicReleaseClient is the client API for PublicRelease service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PublicReleaseClient interface {
	GetAllRelease(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAllData, error)
}

type publicReleaseClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicReleaseClient(cc grpc.ClientConnInterface) PublicReleaseClient {
	return &publicReleaseClient{cc}
}

func (c *publicReleaseClient) GetAllRelease(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAllData, error) {
	out := new(GetAllData)
	err := c.cc.Invoke(ctx, "/blockchain.v1alpha.PublicRelease/GetAllRelease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicReleaseServer is the server API for PublicRelease service.
type PublicReleaseServer interface {
	GetAllRelease(context.Context, *empty.Empty) (*GetAllData, error)
}

// UnimplementedPublicReleaseServer can be embedded to have forward compatible implementations.
type UnimplementedPublicReleaseServer struct {
}

func (*UnimplementedPublicReleaseServer) GetAllRelease(context.Context, *empty.Empty) (*GetAllData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRelease not implemented")
}

func RegisterPublicReleaseServer(s *grpc.Server, srv PublicReleaseServer) {
	s.RegisterService(&_PublicRelease_serviceDesc, srv)
}

func _PublicRelease_GetAllRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicReleaseServer).GetAllRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockchain.v1alpha.PublicRelease/GetAllRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicReleaseServer).GetAllRelease(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _PublicRelease_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blockchain.v1alpha.PublicRelease",
	HandlerType: (*PublicReleaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllRelease",
			Handler:    _PublicRelease_GetAllRelease_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blockchain/v1alpha/public-release.proto",
}