// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.3
// source: question/v1/question.proto

package questionv1

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

type Question_MediaType int32

const (
	Question_MEDIA_TYPE_UNSPECIFIED Question_MediaType = 0
	Question_MEDIA_TYPE_IMAGE       Question_MediaType = 1
	Question_MEDIA_TYPE_AUDIO       Question_MediaType = 2
	Question_MEDIA_TYPE_VIDEO       Question_MediaType = 3
)

// Enum value maps for Question_MediaType.
var (
	Question_MediaType_name = map[int32]string{
		0: "MEDIA_TYPE_UNSPECIFIED",
		1: "MEDIA_TYPE_IMAGE",
		2: "MEDIA_TYPE_AUDIO",
		3: "MEDIA_TYPE_VIDEO",
	}
	Question_MediaType_value = map[string]int32{
		"MEDIA_TYPE_UNSPECIFIED": 0,
		"MEDIA_TYPE_IMAGE":       1,
		"MEDIA_TYPE_AUDIO":       2,
		"MEDIA_TYPE_VIDEO":       3,
	}
)

func (x Question_MediaType) Enum() *Question_MediaType {
	p := new(Question_MediaType)
	*p = x
	return p
}

func (x Question_MediaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Question_MediaType) Descriptor() protoreflect.EnumDescriptor {
	return file_question_v1_question_proto_enumTypes[0].Descriptor()
}

func (Question_MediaType) Type() protoreflect.EnumType {
	return &file_question_v1_question_proto_enumTypes[0]
}

func (x Question_MediaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Question_MediaType.Descriptor instead.
func (Question_MediaType) EnumDescriptor() ([]byte, []int) {
	return file_question_v1_question_proto_rawDescGZIP(), []int{0, 0}
}

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text           string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Answer         *Question_Answer       `protobuf:"bytes,3,opt,name=answer,proto3" json:"answer,omitempty"`
	AuthorNickname string                 `protobuf:"bytes,4,opt,name=author_nickname,json=authorNickname,proto3" json:"author_nickname,omitempty"`
	Media          *Question_Media        `protobuf:"bytes,5,opt,name=media,proto3" json:"media,omitempty"`
	Language       string                 `protobuf:"bytes,6,opt,name=language,proto3" json:"language,omitempty"`
	CreatedTime    *timestamppb.Timestamp `protobuf:"bytes,50,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_v1_question_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_question_v1_question_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_question_v1_question_proto_rawDescGZIP(), []int{0}
}

func (x *Question) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Question) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Question) GetAnswer() *Question_Answer {
	if x != nil {
		return x.Answer
	}
	return nil
}

func (x *Question) GetAuthorNickname() string {
	if x != nil {
		return x.AuthorNickname
	}
	return ""
}

func (x *Question) GetMedia() *Question_Media {
	if x != nil {
		return x.Media
	}
	return nil
}

func (x *Question) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Question) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

type Question_Media struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url  string             `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Type Question_MediaType `protobuf:"varint,2,opt,name=type,proto3,enum=question.v1.Question_MediaType" json:"type,omitempty"`
}

func (x *Question_Media) Reset() {
	*x = Question_Media{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_v1_question_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question_Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question_Media) ProtoMessage() {}

func (x *Question_Media) ProtoReflect() protoreflect.Message {
	mi := &file_question_v1_question_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question_Media.ProtoReflect.Descriptor instead.
func (*Question_Media) Descriptor() ([]byte, []int) {
	return file_question_v1_question_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Question_Media) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Question_Media) GetType() Question_MediaType {
	if x != nil {
		return x.Type
	}
	return Question_MEDIA_TYPE_UNSPECIFIED
}

type Question_Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text           string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Media          *Question_Media        `protobuf:"bytes,3,opt,name=media,proto3" json:"media,omitempty"`
	Language       string                 `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`
	AuthorNickname string                 `protobuf:"bytes,5,opt,name=author_nickname,json=authorNickname,proto3" json:"author_nickname,omitempty"`
	CreatedTime    *timestamppb.Timestamp `protobuf:"bytes,50,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
}

func (x *Question_Answer) Reset() {
	*x = Question_Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_v1_question_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question_Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question_Answer) ProtoMessage() {}

func (x *Question_Answer) ProtoReflect() protoreflect.Message {
	mi := &file_question_v1_question_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question_Answer.ProtoReflect.Descriptor instead.
func (*Question_Answer) Descriptor() ([]byte, []int) {
	return file_question_v1_question_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Question_Answer) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Question_Answer) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Question_Answer) GetMedia() *Question_Media {
	if x != nil {
		return x.Media
	}
	return nil
}

func (x *Question_Answer) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Question_Answer) GetAuthorNickname() string {
	if x != nil {
		return x.AuthorNickname
	}
	return ""
}

func (x *Question_Answer) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

var File_question_v1_question_proto protoreflect.FileDescriptor

var file_question_v1_question_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x05, 0x0a, 0x08, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x4e, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1f, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x1a, 0xe3, 0x01, 0x0a, 0x06, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x31, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x52, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x69,
	0x0a, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4d,
	0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x45, 0x44, 0x49, 0x41,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x01, 0x12, 0x14, 0x0a,
	0x10, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x55, 0x44, 0x49,
	0x4f, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x03, 0x42, 0x18, 0x5a, 0x16, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_question_v1_question_proto_rawDescOnce sync.Once
	file_question_v1_question_proto_rawDescData = file_question_v1_question_proto_rawDesc
)

func file_question_v1_question_proto_rawDescGZIP() []byte {
	file_question_v1_question_proto_rawDescOnce.Do(func() {
		file_question_v1_question_proto_rawDescData = protoimpl.X.CompressGZIP(file_question_v1_question_proto_rawDescData)
	})
	return file_question_v1_question_proto_rawDescData
}

var file_question_v1_question_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_question_v1_question_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_question_v1_question_proto_goTypes = []interface{}{
	(Question_MediaType)(0),       // 0: question.v1.Question.MediaType
	(*Question)(nil),              // 1: question.v1.Question
	(*Question_Media)(nil),        // 2: question.v1.Question.Media
	(*Question_Answer)(nil),       // 3: question.v1.Question.Answer
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_question_v1_question_proto_depIdxs = []int32{
	3, // 0: question.v1.Question.answer:type_name -> question.v1.Question.Answer
	2, // 1: question.v1.Question.media:type_name -> question.v1.Question.Media
	4, // 2: question.v1.Question.created_time:type_name -> google.protobuf.Timestamp
	0, // 3: question.v1.Question.Media.type:type_name -> question.v1.Question.MediaType
	2, // 4: question.v1.Question.Answer.media:type_name -> question.v1.Question.Media
	4, // 5: question.v1.Question.Answer.created_time:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_question_v1_question_proto_init() }
func file_question_v1_question_proto_init() {
	if File_question_v1_question_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_question_v1_question_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Question); i {
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
		file_question_v1_question_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Question_Media); i {
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
		file_question_v1_question_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Question_Answer); i {
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
			RawDescriptor: file_question_v1_question_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_question_v1_question_proto_goTypes,
		DependencyIndexes: file_question_v1_question_proto_depIdxs,
		EnumInfos:         file_question_v1_question_proto_enumTypes,
		MessageInfos:      file_question_v1_question_proto_msgTypes,
	}.Build()
	File_question_v1_question_proto = out.File
	file_question_v1_question_proto_rawDesc = nil
	file_question_v1_question_proto_goTypes = nil
	file_question_v1_question_proto_depIdxs = nil
}