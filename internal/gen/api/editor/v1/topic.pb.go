// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: editor/v1/topic.proto

package editorv1

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

type Topic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title      string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,50,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *Topic) Reset() {
	*x = Topic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_editor_v1_topic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topic) ProtoMessage() {}

func (x *Topic) ProtoReflect() protoreflect.Message {
	mi := &file_editor_v1_topic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topic.ProtoReflect.Descriptor instead.
func (*Topic) Descriptor() ([]byte, []int) {
	return file_editor_v1_topic_proto_rawDescGZIP(), []int{0}
}

func (x *Topic) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Topic) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Topic) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

type CreateTopicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicTitle string `protobuf:"bytes,1,opt,name=topic_title,json=topicTitle,proto3" json:"topic_title,omitempty"` // required
}

func (x *CreateTopicRequest) Reset() {
	*x = CreateTopicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_editor_v1_topic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTopicRequest) ProtoMessage() {}

func (x *CreateTopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_editor_v1_topic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTopicRequest.ProtoReflect.Descriptor instead.
func (*CreateTopicRequest) Descriptor() ([]byte, []int) {
	return file_editor_v1_topic_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTopicRequest) GetTopicTitle() string {
	if x != nil {
		return x.TopicTitle
	}
	return ""
}

type CreateTopicResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic *Topic `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *CreateTopicResponse) Reset() {
	*x = CreateTopicResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_editor_v1_topic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTopicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTopicResponse) ProtoMessage() {}

func (x *CreateTopicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_editor_v1_topic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTopicResponse.ProtoReflect.Descriptor instead.
func (*CreateTopicResponse) Descriptor() ([]byte, []int) {
	return file_editor_v1_topic_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTopicResponse) GetTopic() *Topic {
	if x != nil {
		return x.Topic
	}
	return nil
}

var File_editor_v1_topic_proto protoreflect.FileDescriptor

var file_editor_v1_topic_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x05,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x0b, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x03, 0x18, 0x1e, 0x52, 0x0a,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x3d, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x26, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x32, 0x5c, 0x0a, 0x0c, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1d, 0x2e, 0x65, 0x64, 0x69, 0x74, 0x6f,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x65, 0x64, 0x69, 0x74, 0x6f,
	0x72, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_editor_v1_topic_proto_rawDescOnce sync.Once
	file_editor_v1_topic_proto_rawDescData = file_editor_v1_topic_proto_rawDesc
)

func file_editor_v1_topic_proto_rawDescGZIP() []byte {
	file_editor_v1_topic_proto_rawDescOnce.Do(func() {
		file_editor_v1_topic_proto_rawDescData = protoimpl.X.CompressGZIP(file_editor_v1_topic_proto_rawDescData)
	})
	return file_editor_v1_topic_proto_rawDescData
}

var file_editor_v1_topic_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_editor_v1_topic_proto_goTypes = []interface{}{
	(*Topic)(nil),                 // 0: editor.v1.Topic
	(*CreateTopicRequest)(nil),    // 1: editor.v1.CreateTopicRequest
	(*CreateTopicResponse)(nil),   // 2: editor.v1.CreateTopicResponse
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_editor_v1_topic_proto_depIdxs = []int32{
	3, // 0: editor.v1.Topic.create_time:type_name -> google.protobuf.Timestamp
	0, // 1: editor.v1.CreateTopicResponse.topic:type_name -> editor.v1.Topic
	1, // 2: editor.v1.TopicService.CreateTopic:input_type -> editor.v1.CreateTopicRequest
	2, // 3: editor.v1.TopicService.CreateTopic:output_type -> editor.v1.CreateTopicResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_editor_v1_topic_proto_init() }
func file_editor_v1_topic_proto_init() {
	if File_editor_v1_topic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_editor_v1_topic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topic); i {
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
		file_editor_v1_topic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTopicRequest); i {
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
		file_editor_v1_topic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTopicResponse); i {
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
			RawDescriptor: file_editor_v1_topic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_editor_v1_topic_proto_goTypes,
		DependencyIndexes: file_editor_v1_topic_proto_depIdxs,
		MessageInfos:      file_editor_v1_topic_proto_msgTypes,
	}.Build()
	File_editor_v1_topic_proto = out.File
	file_editor_v1_topic_proto_rawDesc = nil
	file_editor_v1_topic_proto_goTypes = nil
	file_editor_v1_topic_proto_depIdxs = nil
}