// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: player/v1/password_service.proto

package playerv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResetPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailOrNickname string `protobuf:"bytes,1,opt,name=email_or_nickname,json=emailOrNickname,proto3" json:"email_or_nickname,omitempty"`
}

func (x *ResetPasswordRequest) Reset() {
	*x = ResetPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_v1_password_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordRequest) ProtoMessage() {}

func (x *ResetPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_player_v1_password_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordRequest.ProtoReflect.Descriptor instead.
func (*ResetPasswordRequest) Descriptor() ([]byte, []int) {
	return file_player_v1_password_service_proto_rawDescGZIP(), []int{0}
}

func (x *ResetPasswordRequest) GetEmailOrNickname() string {
	if x != nil {
		return x.EmailOrNickname
	}
	return ""
}

type SetPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	NewPassword string `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
}

func (x *SetPasswordRequest) Reset() {
	*x = SetPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_v1_password_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPasswordRequest) ProtoMessage() {}

func (x *SetPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_player_v1_password_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPasswordRequest.ProtoReflect.Descriptor instead.
func (*SetPasswordRequest) Descriptor() ([]byte, []int) {
	return file_player_v1_password_service_proto_rawDescGZIP(), []int{1}
}

func (x *SetPasswordRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SetPasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type UpdatePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldPassword string `protobuf:"bytes,1,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty"`
	NewPassword string `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
}

func (x *UpdatePasswordRequest) Reset() {
	*x = UpdatePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_v1_password_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePasswordRequest) ProtoMessage() {}

func (x *UpdatePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_player_v1_password_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePasswordRequest.ProtoReflect.Descriptor instead.
func (*UpdatePasswordRequest) Descriptor() ([]byte, []int) {
	return file_player_v1_password_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePasswordRequest) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *UpdatePasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

var File_player_v1_password_service_proto protoreflect.FileDescriptor

var file_player_v1_password_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x11, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x6f, 0x72, 0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x04, 0x18,
	0xc0, 0x02, 0x52, 0x0f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x4f, 0x72, 0x4e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x58, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x2c, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x08, 0x18, 0x40,
	0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x71, 0x0a,
	0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x0c, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x18, 0x40, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x2c, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10,
	0x08, 0x18, 0x40, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x32, 0xed, 0x01, 0x0a, 0x0f, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1f, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x44,
	0x0a, 0x0b, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x4a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x14, 0x5a, 0x12, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_player_v1_password_service_proto_rawDescOnce sync.Once
	file_player_v1_password_service_proto_rawDescData = file_player_v1_password_service_proto_rawDesc
)

func file_player_v1_password_service_proto_rawDescGZIP() []byte {
	file_player_v1_password_service_proto_rawDescOnce.Do(func() {
		file_player_v1_password_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_player_v1_password_service_proto_rawDescData)
	})
	return file_player_v1_password_service_proto_rawDescData
}

var file_player_v1_password_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_player_v1_password_service_proto_goTypes = []interface{}{
	(*ResetPasswordRequest)(nil),  // 0: player.v1.ResetPasswordRequest
	(*SetPasswordRequest)(nil),    // 1: player.v1.SetPasswordRequest
	(*UpdatePasswordRequest)(nil), // 2: player.v1.UpdatePasswordRequest
	(*emptypb.Empty)(nil),         // 3: google.protobuf.Empty
}
var file_player_v1_password_service_proto_depIdxs = []int32{
	0, // 0: player.v1.PasswordService.ResetPassword:input_type -> player.v1.ResetPasswordRequest
	1, // 1: player.v1.PasswordService.SetPassword:input_type -> player.v1.SetPasswordRequest
	2, // 2: player.v1.PasswordService.UpdatePassword:input_type -> player.v1.UpdatePasswordRequest
	3, // 3: player.v1.PasswordService.ResetPassword:output_type -> google.protobuf.Empty
	3, // 4: player.v1.PasswordService.SetPassword:output_type -> google.protobuf.Empty
	3, // 5: player.v1.PasswordService.UpdatePassword:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_player_v1_password_service_proto_init() }
func file_player_v1_password_service_proto_init() {
	if File_player_v1_password_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_player_v1_password_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPasswordRequest); i {
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
		file_player_v1_password_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPasswordRequest); i {
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
		file_player_v1_password_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePasswordRequest); i {
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
			RawDescriptor: file_player_v1_password_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_player_v1_password_service_proto_goTypes,
		DependencyIndexes: file_player_v1_password_service_proto_depIdxs,
		MessageInfos:      file_player_v1_password_service_proto_msgTypes,
	}.Build()
	File_player_v1_password_service_proto = out.File
	file_player_v1_password_service_proto_rawDesc = nil
	file_player_v1_password_service_proto_goTypes = nil
	file_player_v1_password_service_proto_depIdxs = nil
}
