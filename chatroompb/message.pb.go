// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: message.proto

package chatroompb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Close struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Close) Reset() {
	*x = Close{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Close) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Close) ProtoMessage() {}

func (x *Close) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Close.ProtoReflect.Descriptor instead.
func (*Close) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,json=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Connect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Active bool  `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *Connect) Reset() {
	*x = Connect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connect) ProtoMessage() {}

func (x *Connect) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connect.ProtoReflect.Descriptor instead.
func (*Connect) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *Connect) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Connect) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName string                 `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Data     string                 `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	PostTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=post_time,json=postTime,proto3" json:"post_time,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Post) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Post) GetPostTime() *timestamppb.Timestamp {
	if x != nil {
		return x.PostTime
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x22, 0x55, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x47, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d,
	0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x37, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x70, 0x6f, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xa2, 0x01, 0x0a, 0x08, 0x43, 0x68, 0x61,
	0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x2f, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x13, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12,
	0x2c, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x72,
	0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x42, 0x0d, 0x5a,
	0x0b, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_message_proto_goTypes = []interface{}{
	(*Close)(nil),                 // 0: chatroompb.Close
	(*User)(nil),                  // 1: chatroompb.User
	(*Connect)(nil),               // 2: chatroompb.Connect
	(*Post)(nil),                  // 3: chatroompb.Post
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_message_proto_depIdxs = []int32{
	1, // 0: chatroompb.Connect.user:type_name -> chatroompb.User
	4, // 1: chatroompb.Post.post_time:type_name -> google.protobuf.Timestamp
	2, // 2: chatroompb.Chatroom.Join:input_type -> chatroompb.Connect
	3, // 3: chatroompb.Chatroom.BroadcastMessage:input_type -> chatroompb.Post
	1, // 4: chatroompb.Chatroom.Login:input_type -> chatroompb.User
	3, // 5: chatroompb.Chatroom.Join:output_type -> chatroompb.Post
	0, // 6: chatroompb.Chatroom.BroadcastMessage:output_type -> chatroompb.Close
	0, // 7: chatroompb.Chatroom.Login:output_type -> chatroompb.Close
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Close); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connect); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatroomClient is the client API for Chatroom service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatroomClient interface {
	Join(ctx context.Context, in *Connect, opts ...grpc.CallOption) (Chatroom_JoinClient, error)
	BroadcastMessage(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Close, error)
	Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*Close, error)
}

type chatroomClient struct {
	cc grpc.ClientConnInterface
}

func NewChatroomClient(cc grpc.ClientConnInterface) ChatroomClient {
	return &chatroomClient{cc}
}

func (c *chatroomClient) Join(ctx context.Context, in *Connect, opts ...grpc.CallOption) (Chatroom_JoinClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chatroom_serviceDesc.Streams[0], "/chatroompb.Chatroom/Join", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatroomJoinClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chatroom_JoinClient interface {
	Recv() (*Post, error)
	grpc.ClientStream
}

type chatroomJoinClient struct {
	grpc.ClientStream
}

func (x *chatroomJoinClient) Recv() (*Post, error) {
	m := new(Post)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatroomClient) BroadcastMessage(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Close, error) {
	out := new(Close)
	err := c.cc.Invoke(ctx, "/chatroompb.Chatroom/BroadcastMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatroomClient) Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*Close, error) {
	out := new(Close)
	err := c.cc.Invoke(ctx, "/chatroompb.Chatroom/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatroomServer is the server API for Chatroom service.
type ChatroomServer interface {
	Join(*Connect, Chatroom_JoinServer) error
	BroadcastMessage(context.Context, *Post) (*Close, error)
	Login(context.Context, *User) (*Close, error)
}

// UnimplementedChatroomServer can be embedded to have forward compatible implementations.
type UnimplementedChatroomServer struct {
}

func (*UnimplementedChatroomServer) Join(*Connect, Chatroom_JoinServer) error {
	return status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (*UnimplementedChatroomServer) BroadcastMessage(context.Context, *Post) (*Close, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastMessage not implemented")
}
func (*UnimplementedChatroomServer) Login(context.Context, *User) (*Close, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterChatroomServer(s *grpc.Server, srv ChatroomServer) {
	s.RegisterService(&_Chatroom_serviceDesc, srv)
}

func _Chatroom_Join_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Connect)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatroomServer).Join(m, &chatroomJoinServer{stream})
}

type Chatroom_JoinServer interface {
	Send(*Post) error
	grpc.ServerStream
}

type chatroomJoinServer struct {
	grpc.ServerStream
}

func (x *chatroomJoinServer) Send(m *Post) error {
	return x.ServerStream.SendMsg(m)
}

func _Chatroom_BroadcastMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatroomServer).BroadcastMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatroompb.Chatroom/BroadcastMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatroomServer).BroadcastMessage(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chatroom_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatroomServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatroompb.Chatroom/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatroomServer).Login(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chatroom_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chatroompb.Chatroom",
	HandlerType: (*ChatroomServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BroadcastMessage",
			Handler:    _Chatroom_BroadcastMessage_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Chatroom_Login_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Join",
			Handler:       _Chatroom_Join_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "message.proto",
}
