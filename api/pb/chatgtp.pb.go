// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: chatgtp.proto

package pb

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

// The request message
type ChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId          string                `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Model           string                `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	Token           string                `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	ParentMessageId string                `protobuf:"bytes,4,opt,name=parent_message_id,json=parentMessageId,proto3" json:"parent_message_id,omitempty"`
	ConversationId  string                `protobuf:"bytes,5,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	Messages        []*ChatRequestMessage `protobuf:"bytes,6,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *ChatRequest) Reset() {
	*x = ChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatgtp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRequest) ProtoMessage() {}

func (x *ChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chatgtp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRequest.ProtoReflect.Descriptor instead.
func (*ChatRequest) Descriptor() ([]byte, []int) {
	return file_chatgtp_proto_rawDescGZIP(), []int{0}
}

func (x *ChatRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *ChatRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *ChatRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ChatRequest) GetParentMessageId() string {
	if x != nil {
		return x.ParentMessageId
	}
	return ""
}

func (x *ChatRequest) GetConversationId() string {
	if x != nil {
		return x.ConversationId
	}
	return ""
}

func (x *ChatRequest) GetMessages() []*ChatRequestMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

// The response message
type ChatReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Data    string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ChatReply) Reset() {
	*x = ChatReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatgtp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatReply) ProtoMessage() {}

func (x *ChatReply) ProtoReflect() protoreflect.Message {
	mi := &file_chatgtp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatReply.ProtoReflect.Descriptor instead.
func (*ChatReply) Descriptor() ([]byte, []int) {
	return file_chatgtp_proto_rawDescGZIP(), []int{1}
}

func (x *ChatReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ChatReply) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type ChatRequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role    string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ChatRequestMessage) Reset() {
	*x = ChatRequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatgtp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRequestMessage) ProtoMessage() {}

func (x *ChatRequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chatgtp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRequestMessage.ProtoReflect.Descriptor instead.
func (*ChatRequestMessage) Descriptor() ([]byte, []int) {
	return file_chatgtp_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ChatRequestMessage) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *ChatRequestMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_chatgtp_proto protoreflect.FileDescriptor

var file_chatgtp_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x68, 0x61, 0x74, 0x67, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x22, 0x96, 0x02, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x34,
	0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x1a, 0x37, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x39, 0x0a,
	0x09, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x3e, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x74,
	0x47, 0x54, 0x50, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x43, 0x68,
	0x61, 0x74, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chatgtp_proto_rawDescOnce sync.Once
	file_chatgtp_proto_rawDescData = file_chatgtp_proto_rawDesc
)

func file_chatgtp_proto_rawDescGZIP() []byte {
	file_chatgtp_proto_rawDescOnce.Do(func() {
		file_chatgtp_proto_rawDescData = protoimpl.X.CompressGZIP(file_chatgtp_proto_rawDescData)
	})
	return file_chatgtp_proto_rawDescData
}

var file_chatgtp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chatgtp_proto_goTypes = []interface{}{
	(*ChatRequest)(nil),        // 0: api.ChatRequest
	(*ChatReply)(nil),          // 1: api.ChatReply
	(*ChatRequestMessage)(nil), // 2: api.ChatRequest.message
}
var file_chatgtp_proto_depIdxs = []int32{
	2, // 0: api.ChatRequest.messages:type_name -> api.ChatRequest.message
	0, // 1: api.ChatGTPService.Chat:input_type -> api.ChatRequest
	1, // 2: api.ChatGTPService.Chat:output_type -> api.ChatReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chatgtp_proto_init() }
func file_chatgtp_proto_init() {
	if File_chatgtp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chatgtp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRequest); i {
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
		file_chatgtp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatReply); i {
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
		file_chatgtp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRequestMessage); i {
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
			RawDescriptor: file_chatgtp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chatgtp_proto_goTypes,
		DependencyIndexes: file_chatgtp_proto_depIdxs,
		MessageInfos:      file_chatgtp_proto_msgTypes,
	}.Build()
	File_chatgtp_proto = out.File
	file_chatgtp_proto_rawDesc = nil
	file_chatgtp_proto_goTypes = nil
	file_chatgtp_proto_depIdxs = nil
}
