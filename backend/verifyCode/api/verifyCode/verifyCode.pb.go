// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: api/verifyCode/verifyCode.proto

package verifyCode

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

type TYPE int32

const (
	TYPE_DEFAULT TYPE = 0
	TYPE_DIGIT   TYPE = 1
	TYPE_LETTER  TYPE = 2
	TYPE_MIXED   TYPE = 3
)

// Enum value maps for TYPE.
var (
	TYPE_name = map[int32]string{
		0: "DEFAULT",
		1: "DIGIT",
		2: "LETTER",
		3: "MIXED",
	}
	TYPE_value = map[string]int32{
		"DEFAULT": 0,
		"DIGIT":   1,
		"LETTER":  2,
		"MIXED":   3,
	}
)

func (x TYPE) Enum() *TYPE {
	p := new(TYPE)
	*p = x
	return p
}

func (x TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_api_verifyCode_verifyCode_proto_enumTypes[0].Descriptor()
}

func (TYPE) Type() protoreflect.EnumType {
	return &file_api_verifyCode_verifyCode_proto_enumTypes[0]
}

func (x TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TYPE.Descriptor instead.
func (TYPE) EnumDescriptor() ([]byte, []int) {
	return file_api_verifyCode_verifyCode_proto_rawDescGZIP(), []int{0}
}

// 获取验证码请求参数
type GetVerifyCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 验证码长度
	Length uint32 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	// 验证码类型
	Type TYPE `protobuf:"varint,2,opt,name=type,proto3,enum=api.verifyCode.TYPE" json:"type,omitempty"`
}

func (x *GetVerifyCodeRequest) Reset() {
	*x = GetVerifyCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_verifyCode_verifyCode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVerifyCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVerifyCodeRequest) ProtoMessage() {}

func (x *GetVerifyCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_verifyCode_verifyCode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVerifyCodeRequest.ProtoReflect.Descriptor instead.
func (*GetVerifyCodeRequest) Descriptor() ([]byte, []int) {
	return file_api_verifyCode_verifyCode_proto_rawDescGZIP(), []int{0}
}

func (x *GetVerifyCodeRequest) GetLength() uint32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *GetVerifyCodeRequest) GetType() TYPE {
	if x != nil {
		return x.Type
	}
	return TYPE_DEFAULT
}

type GetVerifyCodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 验证码
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetVerifyCodeReply) Reset() {
	*x = GetVerifyCodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_verifyCode_verifyCode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVerifyCodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVerifyCodeReply) ProtoMessage() {}

func (x *GetVerifyCodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_verifyCode_verifyCode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVerifyCodeReply.ProtoReflect.Descriptor instead.
func (*GetVerifyCodeReply) Descriptor() ([]byte, []int) {
	return file_api_verifyCode_verifyCode_proto_rawDescGZIP(), []int{1}
}

func (x *GetVerifyCodeReply) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_api_verifyCode_verifyCode_proto protoreflect.FileDescriptor

var file_api_verifyCode_verifyCode_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0x58, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x2e, 0x54, 0x59, 0x50, 0x45, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x28, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x2a, 0x35, 0x0a, 0x04, 0x54, 0x59, 0x50, 0x45, 0x12, 0x0b, 0x0a,
	0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x49,
	0x47, 0x49, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x45, 0x54, 0x54, 0x45, 0x52, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x49, 0x58, 0x45, 0x44, 0x10, 0x03, 0x32, 0x67, 0x0a, 0x0a,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x59, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x38, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x01, 0x5a, 0x24, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x3b, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_verifyCode_verifyCode_proto_rawDescOnce sync.Once
	file_api_verifyCode_verifyCode_proto_rawDescData = file_api_verifyCode_verifyCode_proto_rawDesc
)

func file_api_verifyCode_verifyCode_proto_rawDescGZIP() []byte {
	file_api_verifyCode_verifyCode_proto_rawDescOnce.Do(func() {
		file_api_verifyCode_verifyCode_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_verifyCode_verifyCode_proto_rawDescData)
	})
	return file_api_verifyCode_verifyCode_proto_rawDescData
}

var file_api_verifyCode_verifyCode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_verifyCode_verifyCode_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_verifyCode_verifyCode_proto_goTypes = []interface{}{
	(TYPE)(0),                    // 0: api.verifyCode.TYPE
	(*GetVerifyCodeRequest)(nil), // 1: api.verifyCode.GetVerifyCodeRequest
	(*GetVerifyCodeReply)(nil),   // 2: api.verifyCode.GetVerifyCodeReply
}
var file_api_verifyCode_verifyCode_proto_depIdxs = []int32{
	0, // 0: api.verifyCode.GetVerifyCodeRequest.type:type_name -> api.verifyCode.TYPE
	1, // 1: api.verifyCode.VerifyCode.GetVerifyCode:input_type -> api.verifyCode.GetVerifyCodeRequest
	2, // 2: api.verifyCode.VerifyCode.GetVerifyCode:output_type -> api.verifyCode.GetVerifyCodeReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_verifyCode_verifyCode_proto_init() }
func file_api_verifyCode_verifyCode_proto_init() {
	if File_api_verifyCode_verifyCode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_verifyCode_verifyCode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVerifyCodeRequest); i {
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
		file_api_verifyCode_verifyCode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVerifyCodeReply); i {
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
			RawDescriptor: file_api_verifyCode_verifyCode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_verifyCode_verifyCode_proto_goTypes,
		DependencyIndexes: file_api_verifyCode_verifyCode_proto_depIdxs,
		EnumInfos:         file_api_verifyCode_verifyCode_proto_enumTypes,
		MessageInfos:      file_api_verifyCode_verifyCode_proto_msgTypes,
	}.Build()
	File_api_verifyCode_verifyCode_proto = out.File
	file_api_verifyCode_verifyCode_proto_rawDesc = nil
	file_api_verifyCode_verifyCode_proto_goTypes = nil
	file_api_verifyCode_verifyCode_proto_depIdxs = nil
}
