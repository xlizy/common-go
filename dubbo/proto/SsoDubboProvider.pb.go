// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.3
// source: SsoDubboProvider.proto

package dubbo_api

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

type PasswordLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result        *Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	RequestTicket string  `protobuf:"bytes,2,opt,name=RequestTicket,proto3" json:"RequestTicket,omitempty"`
	Account       string  `protobuf:"bytes,3,opt,name=Account,proto3" json:"Account,omitempty"`
	Password      string  `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
	Ticket        string  `protobuf:"bytes,5,opt,name=Ticket,proto3" json:"Ticket,omitempty"`
	RandStr       string  `protobuf:"bytes,6,opt,name=RandStr,proto3" json:"RandStr,omitempty"`
	RequestIp     string  `protobuf:"bytes,7,opt,name=RequestIp,proto3" json:"RequestIp,omitempty"`
	Scenes        string  `protobuf:"bytes,8,opt,name=Scenes,proto3" json:"Scenes,omitempty"`
}

func (x *PasswordLoginReq) Reset() {
	*x = PasswordLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SsoDubboProvider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordLoginReq) ProtoMessage() {}

func (x *PasswordLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_SsoDubboProvider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordLoginReq.ProtoReflect.Descriptor instead.
func (*PasswordLoginReq) Descriptor() ([]byte, []int) {
	return file_SsoDubboProvider_proto_rawDescGZIP(), []int{0}
}

func (x *PasswordLoginReq) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *PasswordLoginReq) GetRequestTicket() string {
	if x != nil {
		return x.RequestTicket
	}
	return ""
}

func (x *PasswordLoginReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *PasswordLoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *PasswordLoginReq) GetTicket() string {
	if x != nil {
		return x.Ticket
	}
	return ""
}

func (x *PasswordLoginReq) GetRandStr() string {
	if x != nil {
		return x.RandStr
	}
	return ""
}

func (x *PasswordLoginReq) GetRequestIp() string {
	if x != nil {
		return x.RequestIp
	}
	return ""
}

func (x *PasswordLoginReq) GetScenes() string {
	if x != nil {
		return x.Scenes
	}
	return ""
}

type LoginRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Other  string  `protobuf:"bytes,2,opt,name=other,proto3" json:"other,omitempty"`
}

func (x *LoginRsp) Reset() {
	*x = LoginRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SsoDubboProvider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRsp) ProtoMessage() {}

func (x *LoginRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SsoDubboProvider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRsp.ProtoReflect.Descriptor instead.
func (*LoginRsp) Descriptor() ([]byte, []int) {
	return file_SsoDubboProvider_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRsp) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *LoginRsp) GetOther() string {
	if x != nil {
		return x.Other
	}
	return ""
}

var File_SsoDubboProvider_proto protoreflect.FileDescriptor

var file_SsoDubboProvider_proto_rawDesc = []byte{
	0x0a, 0x16, 0x53, 0x73, 0x6f, 0x44, 0x75, 0x62, 0x62, 0x6f, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x01, 0x0a, 0x10, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x24, 0x0a, 0x0d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x52, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x52, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x73,
	0x22, 0x41, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x32, 0x43, 0x0a, 0x10, 0x53, 0x73, 0x6f, 0x44, 0x75, 0x62, 0x62, 0x6f, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x0d, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x11, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x3b, 0x64,
	0x75, 0x62, 0x62, 0x6f, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SsoDubboProvider_proto_rawDescOnce sync.Once
	file_SsoDubboProvider_proto_rawDescData = file_SsoDubboProvider_proto_rawDesc
)

func file_SsoDubboProvider_proto_rawDescGZIP() []byte {
	file_SsoDubboProvider_proto_rawDescOnce.Do(func() {
		file_SsoDubboProvider_proto_rawDescData = protoimpl.X.CompressGZIP(file_SsoDubboProvider_proto_rawDescData)
	})
	return file_SsoDubboProvider_proto_rawDescData
}

var file_SsoDubboProvider_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_SsoDubboProvider_proto_goTypes = []interface{}{
	(*PasswordLoginReq)(nil), // 0: PasswordLoginReq
	(*LoginRsp)(nil),         // 1: LoginRsp
	(*Result)(nil),           // 2: Result
}
var file_SsoDubboProvider_proto_depIdxs = []int32{
	2, // 0: PasswordLoginReq.result:type_name -> Result
	2, // 1: LoginRsp.result:type_name -> Result
	0, // 2: SsoDubboProvider.PasswordLogin:input_type -> PasswordLoginReq
	1, // 3: SsoDubboProvider.PasswordLogin:output_type -> LoginRsp
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_SsoDubboProvider_proto_init() }
func file_SsoDubboProvider_proto_init() {
	if File_SsoDubboProvider_proto != nil {
		return
	}
	file_Result_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SsoDubboProvider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordLoginReq); i {
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
		file_SsoDubboProvider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRsp); i {
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
			RawDescriptor: file_SsoDubboProvider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_SsoDubboProvider_proto_goTypes,
		DependencyIndexes: file_SsoDubboProvider_proto_depIdxs,
		MessageInfos:      file_SsoDubboProvider_proto_msgTypes,
	}.Build()
	File_SsoDubboProvider_proto = out.File
	file_SsoDubboProvider_proto_rawDesc = nil
	file_SsoDubboProvider_proto_goTypes = nil
	file_SsoDubboProvider_proto_depIdxs = nil
}
