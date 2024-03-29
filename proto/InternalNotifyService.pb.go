// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: InternalNotifyService.proto

package proto

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

type NotifySendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content             *string           `protobuf:"bytes,1,req,name=content" json:"content,omitempty"`
	Notifications       map[string]string `protobuf:"bytes,2,rep,name=notifications" json:"notifications,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	DefaultNotification *string           `protobuf:"bytes,3,opt,name=defaultNotification" json:"defaultNotification,omitempty"`
	Uids                []string          `protobuf:"bytes,4,rep,name=uids" json:"uids,omitempty"`
	Gids                []string          `protobuf:"bytes,5,rep,name=gids" json:"gids,omitempty"`
}

func (x *NotifySendRequest) Reset() {
	*x = NotifySendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InternalNotifyService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifySendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifySendRequest) ProtoMessage() {}

func (x *NotifySendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_InternalNotifyService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifySendRequest.ProtoReflect.Descriptor instead.
func (*NotifySendRequest) Descriptor() ([]byte, []int) {
	return file_InternalNotifyService_proto_rawDescGZIP(), []int{0}
}

func (x *NotifySendRequest) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

func (x *NotifySendRequest) GetNotifications() map[string]string {
	if x != nil {
		return x.Notifications
	}
	return nil
}

func (x *NotifySendRequest) GetDefaultNotification() string {
	if x != nil && x.DefaultNotification != nil {
		return *x.DefaultNotification
	}
	return ""
}

func (x *NotifySendRequest) GetUids() []string {
	if x != nil {
		return x.Uids
	}
	return nil
}

func (x *NotifySendRequest) GetGids() []string {
	if x != nil {
		return x.Gids
	}
	return nil
}

var File_InternalNotifyService_proto protoreflect.FileDescriptor

var file_InternalNotifyService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x02, 0x0a, 0x11,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x0d, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x69,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x69, 0x64, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x67, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x67, 0x69,
	0x64, 0x73, 0x1a, 0x40, 0x0a, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x32, 0x48, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a,
	0x0a, 0x73, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x12, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c,
	0x0a, 0x30, 0x6f, 0x72, 0x67, 0x2e, 0x77, 0x68, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x73, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x67,
	0x63, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x50, 0x01, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_InternalNotifyService_proto_rawDescOnce sync.Once
	file_InternalNotifyService_proto_rawDescData = file_InternalNotifyService_proto_rawDesc
)

func file_InternalNotifyService_proto_rawDescGZIP() []byte {
	file_InternalNotifyService_proto_rawDescOnce.Do(func() {
		file_InternalNotifyService_proto_rawDescData = protoimpl.X.CompressGZIP(file_InternalNotifyService_proto_rawDescData)
	})
	return file_InternalNotifyService_proto_rawDescData
}

var file_InternalNotifyService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_InternalNotifyService_proto_goTypes = []interface{}{
	(*NotifySendRequest)(nil), // 0: NotifySendRequest
	nil,                       // 1: NotifySendRequest.NotificationsEntry
	(*BaseResponse)(nil),      // 2: BaseResponse
}
var file_InternalNotifyService_proto_depIdxs = []int32{
	1, // 0: NotifySendRequest.notifications:type_name -> NotifySendRequest.NotificationsEntry
	0, // 1: InternalNotifyService.sendNotify:input_type -> NotifySendRequest
	2, // 2: InternalNotifyService.sendNotify:output_type -> BaseResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_InternalNotifyService_proto_init() }
func file_InternalNotifyService_proto_init() {
	if File_InternalNotifyService_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_InternalNotifyService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifySendRequest); i {
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
			RawDescriptor: file_InternalNotifyService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_InternalNotifyService_proto_goTypes,
		DependencyIndexes: file_InternalNotifyService_proto_depIdxs,
		MessageInfos:      file_InternalNotifyService_proto_msgTypes,
	}.Build()
	File_InternalNotifyService_proto = out.File
	file_InternalNotifyService_proto_rawDesc = nil
	file_InternalNotifyService_proto_goTypes = nil
	file_InternalNotifyService_proto_depIdxs = nil
}
