// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.6
// source: api/auth/auth.proto

package auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthMechanism_Value int32

const (
	AuthMechanism_UNSPECIFIED       AuthMechanism_Value = 0
	AuthMechanism_ADC               AuthMechanism_Value = 2
	AuthMechanism_GCE               AuthMechanism_Value = 4
	AuthMechanism_CREDENTIAL_FILE   AuthMechanism_Value = 5
	AuthMechanism_NONE              AuthMechanism_Value = 6
	AuthMechanism_CREDENTIALSHELPER AuthMechanism_Value = 7
)

// Enum value maps for AuthMechanism_Value.
var (
	AuthMechanism_Value_name = map[int32]string{
		0: "UNSPECIFIED",
		2: "ADC",
		4: "GCE",
		5: "CREDENTIAL_FILE",
		6: "NONE",
		7: "CREDENTIALSHELPER",
	}
	AuthMechanism_Value_value = map[string]int32{
		"UNSPECIFIED":       0,
		"ADC":               2,
		"GCE":               4,
		"CREDENTIAL_FILE":   5,
		"NONE":              6,
		"CREDENTIALSHELPER": 7,
	}
)

func (x AuthMechanism_Value) Enum() *AuthMechanism_Value {
	p := new(AuthMechanism_Value)
	*p = x
	return p
}

func (x AuthMechanism_Value) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthMechanism_Value) Descriptor() protoreflect.EnumDescriptor {
	return file_api_auth_auth_proto_enumTypes[0].Descriptor()
}

func (AuthMechanism_Value) Type() protoreflect.EnumType {
	return &file_api_auth_auth_proto_enumTypes[0]
}

func (x AuthMechanism_Value) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthMechanism_Value.Descriptor instead.
func (AuthMechanism_Value) EnumDescriptor() ([]byte, []int) {
	return file_api_auth_auth_proto_rawDescGZIP(), []int{0, 0}
}

type AuthMechanism struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthMechanism) Reset() {
	*x = AuthMechanism{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthMechanism) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthMechanism) ProtoMessage() {}

func (x *AuthMechanism) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthMechanism.ProtoReflect.Descriptor instead.
func (*AuthMechanism) Descriptor() ([]byte, []int) {
	return file_api_auth_auth_proto_rawDescGZIP(), []int{0}
}

type Credentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mechanism            AuthMechanism_Value    `protobuf:"varint,1,opt,name=mechanism,proto3,enum=auth.AuthMechanism_Value" json:"mechanism,omitempty"`
	Token                string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Expiry               *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
	RefreshExpiry        *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=refresh_expiry,json=refreshExpiry,proto3" json:"refresh_expiry,omitempty"`
	CredsHelperCmdDigest string                 `protobuf:"bytes,6,opt,name=credsHelperCmdDigest,proto3" json:"credsHelperCmdDigest,omitempty"`
}

func (x *Credentials) Reset() {
	*x = Credentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credentials) ProtoMessage() {}

func (x *Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credentials.ProtoReflect.Descriptor instead.
func (*Credentials) Descriptor() ([]byte, []int) {
	return file_api_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *Credentials) GetMechanism() AuthMechanism_Value {
	if x != nil {
		return x.Mechanism
	}
	return AuthMechanism_UNSPECIFIED
}

func (x *Credentials) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Credentials) GetExpiry() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiry
	}
	return nil
}

func (x *Credentials) GetRefreshExpiry() *timestamppb.Timestamp {
	if x != nil {
		return x.RefreshExpiry
	}
	return nil
}

func (x *Credentials) GetCredsHelperCmdDigest() string {
	if x != nil {
		return x.CredsHelperCmdDigest
	}
	return ""
}

var File_api_auth_auth_proto protoreflect.FileDescriptor

var file_api_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x0d,
	0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x22, 0x6c, 0x0a,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x44, 0x43, 0x10, 0x02,
	0x12, 0x07, 0x0a, 0x03, 0x47, 0x43, 0x45, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x52, 0x45,
	0x44, 0x45, 0x4e, 0x54, 0x49, 0x41, 0x4c, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x05, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x52, 0x45, 0x44,
	0x45, 0x4e, 0x54, 0x49, 0x41, 0x4c, 0x53, 0x48, 0x45, 0x4c, 0x50, 0x45, 0x52, 0x10, 0x07, 0x22,
	0x04, 0x08, 0x01, 0x10, 0x01, 0x22, 0x04, 0x08, 0x03, 0x10, 0x03, 0x22, 0x99, 0x02, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x6d,
	0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e,
	0x69, 0x73, 0x6d, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x6d, 0x65, 0x63, 0x68, 0x61,
	0x6e, 0x69, 0x73, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x32, 0x0a, 0x06, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x41,
	0x0a, 0x0e, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x79, 0x12, 0x32, 0x0a, 0x14, 0x63, 0x72, 0x65, 0x64, 0x73, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72,
	0x43, 0x6d, 0x64, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x63, 0x72, 0x65, 0x64, 0x73, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x43, 0x6d, 0x64, 0x44,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2f, 0x72, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_auth_auth_proto_rawDescOnce sync.Once
	file_api_auth_auth_proto_rawDescData = file_api_auth_auth_proto_rawDesc
)

func file_api_auth_auth_proto_rawDescGZIP() []byte {
	file_api_auth_auth_proto_rawDescOnce.Do(func() {
		file_api_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_auth_auth_proto_rawDescData)
	})
	return file_api_auth_auth_proto_rawDescData
}

var file_api_auth_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_auth_auth_proto_goTypes = []interface{}{
	(AuthMechanism_Value)(0),      // 0: auth.AuthMechanism.Value
	(*AuthMechanism)(nil),         // 1: auth.AuthMechanism
	(*Credentials)(nil),           // 2: auth.Credentials
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_api_auth_auth_proto_depIdxs = []int32{
	0, // 0: auth.Credentials.mechanism:type_name -> auth.AuthMechanism.Value
	3, // 1: auth.Credentials.expiry:type_name -> google.protobuf.Timestamp
	3, // 2: auth.Credentials.refresh_expiry:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_auth_auth_proto_init() }
func file_api_auth_auth_proto_init() {
	if File_api_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthMechanism); i {
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
		file_api_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credentials); i {
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
			RawDescriptor: file_api_auth_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_auth_auth_proto_goTypes,
		DependencyIndexes: file_api_auth_auth_proto_depIdxs,
		EnumInfos:         file_api_auth_auth_proto_enumTypes,
		MessageInfos:      file_api_auth_auth_proto_msgTypes,
	}.Build()
	File_api_auth_auth_proto = out.File
	file_api_auth_auth_proto_rawDesc = nil
	file_api_auth_auth_proto_goTypes = nil
	file_api_auth_auth_proto_depIdxs = nil
}
