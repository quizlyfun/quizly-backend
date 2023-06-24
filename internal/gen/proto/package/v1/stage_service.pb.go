// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.3
// source: package/v1/stage_service.proto

package packagev1

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

type CreateStageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageId     int32  `protobuf:"varint,1,opt,name=package_id,json=packageId,proto3" json:"package_id,omitempty"`             // required
	StageName     string `protobuf:"bytes,2,opt,name=stage_name,json=stageName,proto3" json:"stage_name,omitempty"`              // required
	StagePosition int32  `protobuf:"varint,3,opt,name=stage_position,json=stagePosition,proto3" json:"stage_position,omitempty"` // required
}

func (x *CreateStageRequest) Reset() {
	*x = CreateStageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_v1_stage_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStageRequest) ProtoMessage() {}

func (x *CreateStageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_v1_stage_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStageRequest.ProtoReflect.Descriptor instead.
func (*CreateStageRequest) Descriptor() ([]byte, []int) {
	return file_package_v1_stage_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateStageRequest) GetPackageId() int32 {
	if x != nil {
		return x.PackageId
	}
	return 0
}

func (x *CreateStageRequest) GetStageName() string {
	if x != nil {
		return x.StageName
	}
	return ""
}

func (x *CreateStageRequest) GetStagePosition() int32 {
	if x != nil {
		return x.StagePosition
	}
	return 0
}

type CreateStageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package *Package `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
}

func (x *CreateStageResponse) Reset() {
	*x = CreateStageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_v1_stage_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStageResponse) ProtoMessage() {}

func (x *CreateStageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_v1_stage_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStageResponse.ProtoReflect.Descriptor instead.
func (*CreateStageResponse) Descriptor() ([]byte, []int) {
	return file_package_v1_stage_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStageResponse) GetPackage() *Package {
	if x != nil {
		return x.Package
	}
	return nil
}

type UpdateStagePositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageId       int32 `protobuf:"varint,1,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`                   // required
	StagePosition int32 `protobuf:"varint,2,opt,name=stage_position,json=stagePosition,proto3" json:"stage_position,omitempty"` // required
}

func (x *UpdateStagePositionRequest) Reset() {
	*x = UpdateStagePositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_v1_stage_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStagePositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStagePositionRequest) ProtoMessage() {}

func (x *UpdateStagePositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_v1_stage_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStagePositionRequest.ProtoReflect.Descriptor instead.
func (*UpdateStagePositionRequest) Descriptor() ([]byte, []int) {
	return file_package_v1_stage_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateStagePositionRequest) GetStageId() int32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *UpdateStagePositionRequest) GetStagePosition() int32 {
	if x != nil {
		return x.StagePosition
	}
	return 0
}

type ListStagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageId int32 `protobuf:"varint,1,opt,name=package_id,json=packageId,proto3" json:"package_id,omitempty"` // required
}

func (x *ListStagesRequest) Reset() {
	*x = ListStagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_v1_stage_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStagesRequest) ProtoMessage() {}

func (x *ListStagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_v1_stage_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStagesRequest.ProtoReflect.Descriptor instead.
func (*ListStagesRequest) Descriptor() ([]byte, []int) {
	return file_package_v1_stage_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListStagesRequest) GetPackageId() int32 {
	if x != nil {
		return x.PackageId
	}
	return 0
}

type ListStagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stages []*ListStagesResponse_Stage `protobuf:"bytes,1,rep,name=stages,proto3" json:"stages,omitempty"`
}

func (x *ListStagesResponse) Reset() {
	*x = ListStagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_v1_stage_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStagesResponse) ProtoMessage() {}

func (x *ListStagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_v1_stage_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStagesResponse.ProtoReflect.Descriptor instead.
func (*ListStagesResponse) Descriptor() ([]byte, []int) {
	return file_package_v1_stage_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListStagesResponse) GetStages() []*ListStagesResponse_Stage {
	if x != nil {
		return x.Stages
	}
	return nil
}

type AddTopicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageId int32 `protobuf:"varint,1,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"` // required
	TopicId int32 `protobuf:"varint,2,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"` // required
}

func (x *AddTopicRequest) Reset() {
	*x = AddTopicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_v1_stage_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTopicRequest) ProtoMessage() {}

func (x *AddTopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_v1_stage_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTopicRequest.ProtoReflect.Descriptor instead.
func (*AddTopicRequest) Descriptor() ([]byte, []int) {
	return file_package_v1_stage_service_proto_rawDescGZIP(), []int{5}
}

func (x *AddTopicRequest) GetStageId() int32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *AddTopicRequest) GetTopicId() int32 {
	if x != nil {
		return x.TopicId
	}
	return 0
}

type RemoveTopicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageId int32 `protobuf:"varint,1,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"` // required
	TopicId int32 `protobuf:"varint,2,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"` // required
}

func (x *RemoveTopicRequest) Reset() {
	*x = RemoveTopicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_v1_stage_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveTopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveTopicRequest) ProtoMessage() {}

func (x *RemoveTopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_v1_stage_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveTopicRequest.ProtoReflect.Descriptor instead.
func (*RemoveTopicRequest) Descriptor() ([]byte, []int) {
	return file_package_v1_stage_service_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveTopicRequest) GetStageId() int32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *RemoveTopicRequest) GetTopicId() int32 {
	if x != nil {
		return x.TopicId
	}
	return 0
}

type ListStagesResponse_Stage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ListStagesResponse_Stage) Reset() {
	*x = ListStagesResponse_Stage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_v1_stage_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStagesResponse_Stage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStagesResponse_Stage) ProtoMessage() {}

func (x *ListStagesResponse_Stage) ProtoReflect() protoreflect.Message {
	mi := &file_package_v1_stage_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStagesResponse_Stage.ProtoReflect.Descriptor instead.
func (*ListStagesResponse_Stage) Descriptor() ([]byte, []int) {
	return file_package_v1_stage_service_proto_rawDescGZIP(), []int{4, 0}
}

func (x *ListStagesResponse_Stage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListStagesResponse_Stage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_package_v1_stage_service_proto protoreflect.FileDescriptor

var file_package_v1_stage_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x28, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x03, 0x18,
	0x1e, 0x52, 0x09, 0x73, 0x74, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x44, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x5e, 0x0a, 0x1a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x7f, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x73, 0x1a, 0x2b, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x47,
	0x0a, 0x0f, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x49, 0x64, 0x32, 0x8a, 0x03, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4b, 0x0a, 0x0a, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x1b, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x0b, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1e, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x16, 0x5a, 0x14, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_package_v1_stage_service_proto_rawDescOnce sync.Once
	file_package_v1_stage_service_proto_rawDescData = file_package_v1_stage_service_proto_rawDesc
)

func file_package_v1_stage_service_proto_rawDescGZIP() []byte {
	file_package_v1_stage_service_proto_rawDescOnce.Do(func() {
		file_package_v1_stage_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_package_v1_stage_service_proto_rawDescData)
	})
	return file_package_v1_stage_service_proto_rawDescData
}

var file_package_v1_stage_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_package_v1_stage_service_proto_goTypes = []interface{}{
	(*CreateStageRequest)(nil),         // 0: package.v1.CreateStageRequest
	(*CreateStageResponse)(nil),        // 1: package.v1.CreateStageResponse
	(*UpdateStagePositionRequest)(nil), // 2: package.v1.UpdateStagePositionRequest
	(*ListStagesRequest)(nil),          // 3: package.v1.ListStagesRequest
	(*ListStagesResponse)(nil),         // 4: package.v1.ListStagesResponse
	(*AddTopicRequest)(nil),            // 5: package.v1.AddTopicRequest
	(*RemoveTopicRequest)(nil),         // 6: package.v1.RemoveTopicRequest
	(*ListStagesResponse_Stage)(nil),   // 7: package.v1.ListStagesResponse.Stage
	(*Package)(nil),                    // 8: package.v1.Package
	(*emptypb.Empty)(nil),              // 9: google.protobuf.Empty
}
var file_package_v1_stage_service_proto_depIdxs = []int32{
	8, // 0: package.v1.CreateStageResponse.package:type_name -> package.v1.Package
	7, // 1: package.v1.ListStagesResponse.stages:type_name -> package.v1.ListStagesResponse.Stage
	0, // 2: package.v1.StageService.CreateStage:input_type -> package.v1.CreateStageRequest
	2, // 3: package.v1.StageService.UpdateStagePosition:input_type -> package.v1.UpdateStagePositionRequest
	3, // 4: package.v1.StageService.ListStages:input_type -> package.v1.ListStagesRequest
	5, // 5: package.v1.StageService.AddTopic:input_type -> package.v1.AddTopicRequest
	6, // 6: package.v1.StageService.RemoveTopic:input_type -> package.v1.RemoveTopicRequest
	1, // 7: package.v1.StageService.CreateStage:output_type -> package.v1.CreateStageResponse
	9, // 8: package.v1.StageService.UpdateStagePosition:output_type -> google.protobuf.Empty
	4, // 9: package.v1.StageService.ListStages:output_type -> package.v1.ListStagesResponse
	9, // 10: package.v1.StageService.AddTopic:output_type -> google.protobuf.Empty
	9, // 11: package.v1.StageService.RemoveTopic:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_package_v1_stage_service_proto_init() }
func file_package_v1_stage_service_proto_init() {
	if File_package_v1_stage_service_proto != nil {
		return
	}
	file_package_v1_package_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_package_v1_stage_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStageRequest); i {
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
		file_package_v1_stage_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStageResponse); i {
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
		file_package_v1_stage_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStagePositionRequest); i {
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
		file_package_v1_stage_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStagesRequest); i {
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
		file_package_v1_stage_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStagesResponse); i {
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
		file_package_v1_stage_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTopicRequest); i {
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
		file_package_v1_stage_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveTopicRequest); i {
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
		file_package_v1_stage_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStagesResponse_Stage); i {
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
			RawDescriptor: file_package_v1_stage_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_package_v1_stage_service_proto_goTypes,
		DependencyIndexes: file_package_v1_stage_service_proto_depIdxs,
		MessageInfos:      file_package_v1_stage_service_proto_msgTypes,
	}.Build()
	File_package_v1_stage_service_proto = out.File
	file_package_v1_stage_service_proto_rawDesc = nil
	file_package_v1_stage_service_proto_goTypes = nil
	file_package_v1_stage_service_proto_depIdxs = nil
}
