// protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./common/user/rpc/proto/*.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.21.12
// source: common/user/rpc/proto/login.proto

package userPB

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 登陆传入参数
type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginName string `protobuf:"bytes,1,opt,name=loginName,proto3" json:"loginName,omitempty"` // 用户登陆名
	Password  string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`   // 密码
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_user_rpc_proto_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_common_user_rpc_proto_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_common_user_rpc_proto_login_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetLoginName() string {
	if x != nil {
		return x.LoginName
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShowName string `protobuf:"bytes,1,opt,name=showName,proto3" json:"showName,omitempty"` // 用户显示名称
	UserId   string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`     // 用户id
	Token    string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`       // jwt token
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_user_rpc_proto_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_common_user_rpc_proto_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_common_user_rpc_proto_login_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResp) GetShowName() string {
	if x != nil {
		return x.ShowName
	}
	return ""
}

func (x *LoginResp) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *LoginResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_common_user_rpc_proto_login_proto protoreflect.FileDescriptor

var file_common_user_rpc_proto_login_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x44, 0x0a, 0x08, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x55, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x68, 0x6f, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x68, 0x6f, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x31, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x28, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x1e, 0x5a, 0x1c, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x50, 0x42, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_common_user_rpc_proto_login_proto_rawDescOnce sync.Once
	file_common_user_rpc_proto_login_proto_rawDescData = file_common_user_rpc_proto_login_proto_rawDesc
)

func file_common_user_rpc_proto_login_proto_rawDescGZIP() []byte {
	file_common_user_rpc_proto_login_proto_rawDescOnce.Do(func() {
		file_common_user_rpc_proto_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_user_rpc_proto_login_proto_rawDescData)
	})
	return file_common_user_rpc_proto_login_proto_rawDescData
}

var file_common_user_rpc_proto_login_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_user_rpc_proto_login_proto_goTypes = []interface{}{
	(*LoginReq)(nil),  // 0: user.loginReq
	(*LoginResp)(nil), // 1: user.loginResp
}
var file_common_user_rpc_proto_login_proto_depIdxs = []int32{
	0, // 0: user.Login.login:input_type -> user.loginReq
	1, // 1: user.Login.login:output_type -> user.loginResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_user_rpc_proto_login_proto_init() }
func file_common_user_rpc_proto_login_proto_init() {
	if File_common_user_rpc_proto_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_user_rpc_proto_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_common_user_rpc_proto_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
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
			RawDescriptor: file_common_user_rpc_proto_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_user_rpc_proto_login_proto_goTypes,
		DependencyIndexes: file_common_user_rpc_proto_login_proto_depIdxs,
		MessageInfos:      file_common_user_rpc_proto_login_proto_msgTypes,
	}.Build()
	File_common_user_rpc_proto_login_proto = out.File
	file_common_user_rpc_proto_login_proto_rawDesc = nil
	file_common_user_rpc_proto_login_proto_goTypes = nil
	file_common_user_rpc_proto_login_proto_depIdxs = nil
}
