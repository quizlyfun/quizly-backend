// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.3
// source: package/v1/package.proto

package packagev1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type PackageRound struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PackageRound) Reset() {
	*x = PackageRound{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_v1_package_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageRound) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageRound) ProtoMessage() {}

func (x *PackageRound) ProtoReflect() protoreflect.Message {
	mi := &file_package_v1_package_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageRound.ProtoReflect.Descriptor instead.
func (*PackageRound) Descriptor() ([]byte, []int) {
	return file_package_v1_package_proto_rawDescGZIP(), []int{0}
}

func (x *PackageRound) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PackageRound) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Author       string                 `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	IsPublished  bool                   `protobuf:"varint,4,opt,name=is_published,json=isPublished,proto3" json:"is_published,omitempty"`
	CoverUrl     string                 `protobuf:"bytes,5,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`
	Rounds       []*PackageRound        `protobuf:"bytes,6,rep,name=rounds,proto3" json:"rounds,omitempty"`
	Tags         []string               `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	CreationTime *timestamppb.Timestamp `protobuf:"bytes,50,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	UpdateTime   *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_v1_package_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_package_v1_package_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_package_v1_package_proto_rawDescGZIP(), []int{1}
}

func (x *Package) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Package) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Package) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Package) GetIsPublished() bool {
	if x != nil {
		return x.IsPublished
	}
	return false
}

func (x *Package) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Package) GetRounds() []*PackageRound {
	if x != nil {
		return x.Rounds
	}
	return nil
}

func (x *Package) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Package) GetCreationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationTime
	}
	return nil
}

func (x *Package) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type PackageStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoundCount    int32 `protobuf:"varint,1,opt,name=round_count,json=roundCount,proto3" json:"round_count,omitempty"`
	TopicCount    int32 `protobuf:"varint,2,opt,name=topic_count,json=topicCount,proto3" json:"topic_count,omitempty"`
	QuestionCount int32 `protobuf:"varint,3,opt,name=question_count,json=questionCount,proto3" json:"question_count,omitempty"`
	VideoCount    int32 `protobuf:"varint,4,opt,name=video_count,json=videoCount,proto3" json:"video_count,omitempty"`
	AudioCount    int32 `protobuf:"varint,5,opt,name=audio_count,json=audioCount,proto3" json:"audio_count,omitempty"`
	ImageCount    int32 `protobuf:"varint,6,opt,name=image_count,json=imageCount,proto3" json:"image_count,omitempty"`
}

func (x *PackageStats) Reset() {
	*x = PackageStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_v1_package_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageStats) ProtoMessage() {}

func (x *PackageStats) ProtoReflect() protoreflect.Message {
	mi := &file_package_v1_package_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageStats.ProtoReflect.Descriptor instead.
func (*PackageStats) Descriptor() ([]byte, []int) {
	return file_package_v1_package_proto_rawDescGZIP(), []int{2}
}

func (x *PackageStats) GetRoundCount() int32 {
	if x != nil {
		return x.RoundCount
	}
	return 0
}

func (x *PackageStats) GetTopicCount() int32 {
	if x != nil {
		return x.TopicCount
	}
	return 0
}

func (x *PackageStats) GetQuestionCount() int32 {
	if x != nil {
		return x.QuestionCount
	}
	return 0
}

func (x *PackageStats) GetVideoCount() int32 {
	if x != nil {
		return x.VideoCount
	}
	return 0
}

func (x *PackageStats) GetAudioCount() int32 {
	if x != nil {
		return x.AudioCount
	}
	return 0
}

func (x *PackageStats) GetImageCount() int32 {
	if x != nil {
		return x.ImageCount
	}
	return 0
}

type PackageWithStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package *Package      `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
	Stats   *PackageStats `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *PackageWithStats) Reset() {
	*x = PackageWithStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_v1_package_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageWithStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageWithStats) ProtoMessage() {}

func (x *PackageWithStats) ProtoReflect() protoreflect.Message {
	mi := &file_package_v1_package_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageWithStats.ProtoReflect.Descriptor instead.
func (*PackageWithStats) Descriptor() ([]byte, []int) {
	return file_package_v1_package_proto_rawDescGZIP(), []int{3}
}

func (x *PackageWithStats) GetPackage() *Package {
	if x != nil {
		return x.Package
	}
	return nil
}

func (x *PackageWithStats) GetStats() *PackageStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

type GetPackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageId int32 `protobuf:"varint,1,opt,name=package_id,json=packageId,proto3" json:"package_id,omitempty"` // required
}

func (x *GetPackageRequest) Reset() {
	*x = GetPackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_v1_package_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackageRequest) ProtoMessage() {}

func (x *GetPackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_v1_package_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPackageRequest.ProtoReflect.Descriptor instead.
func (*GetPackageRequest) Descriptor() ([]byte, []int) {
	return file_package_v1_package_proto_rawDescGZIP(), []int{4}
}

func (x *GetPackageRequest) GetPackageId() int32 {
	if x != nil {
		return x.PackageId
	}
	return 0
}

type GetPackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package *Package `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
}

func (x *GetPackageResponse) Reset() {
	*x = GetPackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_v1_package_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackageResponse) ProtoMessage() {}

func (x *GetPackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_v1_package_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPackageResponse.ProtoReflect.Descriptor instead.
func (*GetPackageResponse) Descriptor() ([]byte, []int) {
	return file_package_v1_package_proto_rawDescGZIP(), []int{5}
}

func (x *GetPackageResponse) GetPackage() *Package {
	if x != nil {
		return x.Package
	}
	return nil
}

type CreatePackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageName string   `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"` // required
	CoverUrl    string   `protobuf:"bytes,2,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`
	Tags        []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *CreatePackageRequest) Reset() {
	*x = CreatePackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_v1_package_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePackageRequest) ProtoMessage() {}

func (x *CreatePackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_v1_package_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePackageRequest.ProtoReflect.Descriptor instead.
func (*CreatePackageRequest) Descriptor() ([]byte, []int) {
	return file_package_v1_package_proto_rawDescGZIP(), []int{6}
}

func (x *CreatePackageRequest) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *CreatePackageRequest) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *CreatePackageRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type CreatePackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package *Package `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
}

func (x *CreatePackageResponse) Reset() {
	*x = CreatePackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_v1_package_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePackageResponse) ProtoMessage() {}

func (x *CreatePackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_v1_package_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePackageResponse.ProtoReflect.Descriptor instead.
func (*CreatePackageResponse) Descriptor() ([]byte, []int) {
	return file_package_v1_package_proto_rawDescGZIP(), []int{7}
}

func (x *CreatePackageResponse) GetPackage() *Package {
	if x != nil {
		return x.Package
	}
	return nil
}

type PublishPackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageId int32 `protobuf:"varint,1,opt,name=package_id,json=packageId,proto3" json:"package_id,omitempty"` // required
}

func (x *PublishPackageRequest) Reset() {
	*x = PublishPackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_v1_package_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishPackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishPackageRequest) ProtoMessage() {}

func (x *PublishPackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_v1_package_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishPackageRequest.ProtoReflect.Descriptor instead.
func (*PublishPackageRequest) Descriptor() ([]byte, []int) {
	return file_package_v1_package_proto_rawDescGZIP(), []int{8}
}

func (x *PublishPackageRequest) GetPackageId() int32 {
	if x != nil {
		return x.PackageId
	}
	return 0
}

type PublishPackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package *PackageWithStats `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
}

func (x *PublishPackageResponse) Reset() {
	*x = PublishPackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_v1_package_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishPackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishPackageResponse) ProtoMessage() {}

func (x *PublishPackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_v1_package_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishPackageResponse.ProtoReflect.Descriptor instead.
func (*PublishPackageResponse) Descriptor() ([]byte, []int) {
	return file_package_v1_package_proto_rawDescGZIP(), []int{9}
}

func (x *PublishPackageResponse) GetPackage() *PackageWithStats {
	if x != nil {
		return x.Package
	}
	return nil
}

var File_package_v1_package_proto protoreflect.FileDescriptor

var file_package_v1_package_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x32, 0x0a, 0x0c, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x6f, 0x75, 0x6e, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0xc9, 0x02, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c,
	0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x30, 0x0a, 0x06,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0xda, 0x01, 0x0a, 0x0c, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x75, 0x64, 0x69, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x71, 0x0a,
	0x10, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x12, 0x2e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x22, 0x32, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10,
	0x03, 0x18, 0x32, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x25, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x88, 0x01, 0x01, 0x52, 0x08, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x92, 0x01, 0x04, 0x10, 0x05, 0x18,
	0x01, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x46, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22,
	0x36, 0x0a, 0x15, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x16, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x36, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x32, 0x8c, 0x02, 0x0a, 0x0e, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x20, 0x2e,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x12, 0x1d, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x57, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x12, 0x21, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_package_v1_package_proto_rawDescOnce sync.Once
	file_package_v1_package_proto_rawDescData = file_package_v1_package_proto_rawDesc
)

func file_package_v1_package_proto_rawDescGZIP() []byte {
	file_package_v1_package_proto_rawDescOnce.Do(func() {
		file_package_v1_package_proto_rawDescData = protoimpl.X.CompressGZIP(file_package_v1_package_proto_rawDescData)
	})
	return file_package_v1_package_proto_rawDescData
}

var file_package_v1_package_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_package_v1_package_proto_goTypes = []interface{}{
	(*PackageRound)(nil),           // 0: package.v1.PackageRound
	(*Package)(nil),                // 1: package.v1.Package
	(*PackageStats)(nil),           // 2: package.v1.PackageStats
	(*PackageWithStats)(nil),       // 3: package.v1.PackageWithStats
	(*GetPackageRequest)(nil),      // 4: package.v1.GetPackageRequest
	(*GetPackageResponse)(nil),     // 5: package.v1.GetPackageResponse
	(*CreatePackageRequest)(nil),   // 6: package.v1.CreatePackageRequest
	(*CreatePackageResponse)(nil),  // 7: package.v1.CreatePackageResponse
	(*PublishPackageRequest)(nil),  // 8: package.v1.PublishPackageRequest
	(*PublishPackageResponse)(nil), // 9: package.v1.PublishPackageResponse
	(*timestamppb.Timestamp)(nil),  // 10: google.protobuf.Timestamp
}
var file_package_v1_package_proto_depIdxs = []int32{
	0,  // 0: package.v1.Package.rounds:type_name -> package.v1.PackageRound
	10, // 1: package.v1.Package.creation_time:type_name -> google.protobuf.Timestamp
	10, // 2: package.v1.Package.update_time:type_name -> google.protobuf.Timestamp
	1,  // 3: package.v1.PackageWithStats.package:type_name -> package.v1.Package
	2,  // 4: package.v1.PackageWithStats.stats:type_name -> package.v1.PackageStats
	1,  // 5: package.v1.GetPackageResponse.package:type_name -> package.v1.Package
	1,  // 6: package.v1.CreatePackageResponse.package:type_name -> package.v1.Package
	3,  // 7: package.v1.PublishPackageResponse.package:type_name -> package.v1.PackageWithStats
	6,  // 8: package.v1.PackageService.CreatePackage:input_type -> package.v1.CreatePackageRequest
	4,  // 9: package.v1.PackageService.GetPackage:input_type -> package.v1.GetPackageRequest
	8,  // 10: package.v1.PackageService.PublishPackage:input_type -> package.v1.PublishPackageRequest
	7,  // 11: package.v1.PackageService.CreatePackage:output_type -> package.v1.CreatePackageResponse
	5,  // 12: package.v1.PackageService.GetPackage:output_type -> package.v1.GetPackageResponse
	9,  // 13: package.v1.PackageService.PublishPackage:output_type -> package.v1.PublishPackageResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_package_v1_package_proto_init() }
func file_package_v1_package_proto_init() {
	if File_package_v1_package_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_package_v1_package_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageRound); i {
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
		file_package_v1_package_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
		file_package_v1_package_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageStats); i {
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
		file_package_v1_package_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageWithStats); i {
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
		file_package_v1_package_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPackageRequest); i {
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
		file_package_v1_package_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPackageResponse); i {
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
		file_package_v1_package_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePackageRequest); i {
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
		file_package_v1_package_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePackageResponse); i {
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
		file_package_v1_package_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishPackageRequest); i {
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
		file_package_v1_package_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishPackageResponse); i {
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
			RawDescriptor: file_package_v1_package_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_package_v1_package_proto_goTypes,
		DependencyIndexes: file_package_v1_package_proto_depIdxs,
		MessageInfos:      file_package_v1_package_proto_msgTypes,
	}.Build()
	File_package_v1_package_proto = out.File
	file_package_v1_package_proto_rawDesc = nil
	file_package_v1_package_proto_goTypes = nil
	file_package_v1_package_proto_depIdxs = nil
}
