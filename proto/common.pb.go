// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: common.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type STATUS int32

const (
	STATUS_OK                           STATUS = 0
	STATUS_INVALID_PARAMETER            STATUS = 1
	STATUS_NO_PERMISSION                STATUS = 2
	STATUS_NO_SUCH_GROUP                STATUS = 3
	STATUS_NO_SUCH_GROUP_MEMBER         STATUS = 4
	STATUS_INVALID_TOKEN                STATUS = 5
	STATUS_SERVER_INTERNAL_ERROR        STATUS = 6
	STATUS_NO_SUCH_GROUP_ANNOUNCEMENT   STATUS = 7
	STATUS_GROUP_EXISTS                 STATUS = 8
	STATUS_NO_SUCH_FILE                 STATUS = 9
	STATUS_GROUP_IS_FULL_OR_EXCEEDS     STATUS = 10
	STATUS_NO_SUCH_USER                 STATUS = 11
	STATUS_RATE_LIMIT_EXCEEDED          STATUS = 12
	STATUS_INVALID_INVITER              STATUS = 13
	STATUS_USER_IS_DISABLED             STATUS = 14
	STATUS_PUID_IS_REGISTERING          STATUS = 15
	STATUS_NUMBER_IS_BINDING_OTHER_PUID STATUS = 16
	STATUS_TEAM_HAS_MEMBERS             STATUS = 17
	STATUS_VOTE_IS_CLOSED               STATUS = 18
	STATUS_NO_SUCH_GROUP_PIN            STATUS = 19
	STATUS_USER_EMAIL_EXIST             STATUS = 20
	STATUS_USER_OKTAID_EXIST            STATUS = 21
	STATUS_GROUP_PIN_CONTENT_TOO_LONG   STATUS = 22
	STATUS_OTHER_ERROR                  STATUS = 99
)

// Enum value maps for STATUS.
var (
	STATUS_name = map[int32]string{
		0:  "OK",
		1:  "INVALID_PARAMETER",
		2:  "NO_PERMISSION",
		3:  "NO_SUCH_GROUP",
		4:  "NO_SUCH_GROUP_MEMBER",
		5:  "INVALID_TOKEN",
		6:  "SERVER_INTERNAL_ERROR",
		7:  "NO_SUCH_GROUP_ANNOUNCEMENT",
		8:  "GROUP_EXISTS",
		9:  "NO_SUCH_FILE",
		10: "GROUP_IS_FULL_OR_EXCEEDS",
		11: "NO_SUCH_USER",
		12: "RATE_LIMIT_EXCEEDED",
		13: "INVALID_INVITER",
		14: "USER_IS_DISABLED",
		15: "PUID_IS_REGISTERING",
		16: "NUMBER_IS_BINDING_OTHER_PUID",
		17: "TEAM_HAS_MEMBERS",
		18: "VOTE_IS_CLOSED",
		19: "NO_SUCH_GROUP_PIN",
		20: "USER_EMAIL_EXIST",
		21: "USER_OKTAID_EXIST",
		22: "GROUP_PIN_CONTENT_TOO_LONG",
		99: "OTHER_ERROR",
	}
	STATUS_value = map[string]int32{
		"OK":                           0,
		"INVALID_PARAMETER":            1,
		"NO_PERMISSION":                2,
		"NO_SUCH_GROUP":                3,
		"NO_SUCH_GROUP_MEMBER":         4,
		"INVALID_TOKEN":                5,
		"SERVER_INTERNAL_ERROR":        6,
		"NO_SUCH_GROUP_ANNOUNCEMENT":   7,
		"GROUP_EXISTS":                 8,
		"NO_SUCH_FILE":                 9,
		"GROUP_IS_FULL_OR_EXCEEDS":     10,
		"NO_SUCH_USER":                 11,
		"RATE_LIMIT_EXCEEDED":          12,
		"INVALID_INVITER":              13,
		"USER_IS_DISABLED":             14,
		"PUID_IS_REGISTERING":          15,
		"NUMBER_IS_BINDING_OTHER_PUID": 16,
		"TEAM_HAS_MEMBERS":             17,
		"VOTE_IS_CLOSED":               18,
		"NO_SUCH_GROUP_PIN":            19,
		"USER_EMAIL_EXIST":             20,
		"USER_OKTAID_EXIST":            21,
		"GROUP_PIN_CONTENT_TOO_LONG":   22,
		"OTHER_ERROR":                  99,
	}
)

func (x STATUS) Enum() *STATUS {
	p := new(STATUS)
	*p = x
	return p
}

func (x STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (STATUS) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *STATUS) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = STATUS(num)
	return nil
}

// Deprecated: Use STATUS.Descriptor instead.
func (STATUS) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset   *uint32 `protobuf:"varint,1,req,name=offset" json:"offset,omitempty"`
	Length   *uint32 `protobuf:"varint,2,req,name=length" json:"length,omitempty"`
	Number   *string `protobuf:"bytes,3,opt,name=number" json:"number,omitempty"`
	Email    *string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Name     *string `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	TeamsId  *uint64 `protobuf:"varint,6,opt,name=teamsId" json:"teamsId,omitempty"`
	Disabled []bool  `protobuf:"varint,7,rep,name=disabled" json:"disabled,omitempty"`
	Appid    *string `protobuf:"bytes,8,opt,name=appid" json:"appid,omitempty"`
	Pid      *string `protobuf:"bytes,9,opt,name=pid" json:"pid,omitempty"`
}

func (x *Step) Reset() {
	*x = Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Step) ProtoMessage() {}

func (x *Step) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Step.ProtoReflect.Descriptor instead.
func (*Step) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *Step) GetOffset() uint32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *Step) GetLength() uint32 {
	if x != nil && x.Length != nil {
		return *x.Length
	}
	return 0
}

func (x *Step) GetNumber() string {
	if x != nil && x.Number != nil {
		return *x.Number
	}
	return ""
}

func (x *Step) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *Step) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Step) GetTeamsId() uint64 {
	if x != nil && x.TeamsId != nil {
		return *x.TeamsId
	}
	return 0
}

func (x *Step) GetDisabled() []bool {
	if x != nil {
		return x.Disabled
	}
	return nil
}

func (x *Step) GetAppid() string {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return ""
}

func (x *Step) GetPid() string {
	if x != nil && x.Pid != nil {
		return *x.Pid
	}
	return ""
}

type BaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ver    *uint32    `protobuf:"varint,1,req,name=ver" json:"ver,omitempty"`
	Status *uint32    `protobuf:"varint,2,req,name=status" json:"status,omitempty"`
	Reason *string    `protobuf:"bytes,3,opt,name=reason" json:"reason,omitempty"`
	Data   *anypb.Any `protobuf:"bytes,4,opt,name=data" json:"data,omitempty"`
}

func (x *BaseResponse) Reset() {
	*x = BaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponse) ProtoMessage() {}

func (x *BaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResponse.ProtoReflect.Descriptor instead.
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *BaseResponse) GetVer() uint32 {
	if x != nil && x.Ver != nil {
		return *x.Ver
	}
	return 0
}

func (x *BaseResponse) GetStatus() uint32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *BaseResponse) GetReason() string {
	if x != nil && x.Reason != nil {
		return *x.Reason
	}
	return ""
}

func (x *BaseResponse) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type BaseAnyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *string `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
}

func (x *BaseAnyResponse) Reset() {
	*x = BaseAnyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseAnyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseAnyResponse) ProtoMessage() {}

func (x *BaseAnyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseAnyResponse.ProtoReflect.Descriptor instead.
func (*BaseAnyResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *BaseAnyResponse) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0xd6, 0x01, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x74, 0x65, 0x61, 0x6d, 0x73, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x03, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x7a, 0x0a, 0x0c, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x76,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x28, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x27, 0x0a, 0x0f, 0x42, 0x61, 0x73, 0x65, 0x41,
	0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x2a, 0xa7, 0x04, 0x0a, 0x06, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x06, 0x0a, 0x02, 0x4f,
	0x4b, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50,
	0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f,
	0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x11, 0x0a,
	0x0d, 0x4e, 0x4f, 0x5f, 0x53, 0x55, 0x43, 0x48, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x03,
	0x12, 0x18, 0x0a, 0x14, 0x4e, 0x4f, 0x5f, 0x53, 0x55, 0x43, 0x48, 0x5f, 0x47, 0x52, 0x4f, 0x55,
	0x50, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x05, 0x12, 0x19, 0x0a,
	0x15, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x06, 0x12, 0x1e, 0x0a, 0x1a, 0x4e, 0x4f, 0x5f, 0x53,
	0x55, 0x43, 0x48, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x41, 0x4e, 0x4e, 0x4f, 0x55, 0x4e,
	0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x52, 0x4f, 0x55,
	0x50, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f,
	0x5f, 0x53, 0x55, 0x43, 0x48, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x09, 0x12, 0x1c, 0x0a, 0x18,
	0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x49, 0x53, 0x5f, 0x46, 0x55, 0x4c, 0x4c, 0x5f, 0x4f, 0x52,
	0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x53, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f,
	0x5f, 0x53, 0x55, 0x43, 0x48, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x0b, 0x12, 0x17, 0x0a, 0x13,
	0x52, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45,
	0x44, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x52, 0x10, 0x0d, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x49, 0x53, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x0e,
	0x12, 0x17, 0x0a, 0x13, 0x50, 0x55, 0x49, 0x44, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x47, 0x49,
	0x53, 0x54, 0x45, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x0f, 0x12, 0x20, 0x0a, 0x1c, 0x4e, 0x55, 0x4d,
	0x42, 0x45, 0x52, 0x5f, 0x49, 0x53, 0x5f, 0x42, 0x49, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x4f,
	0x54, 0x48, 0x45, 0x52, 0x5f, 0x50, 0x55, 0x49, 0x44, 0x10, 0x10, 0x12, 0x14, 0x0a, 0x10, 0x54,
	0x45, 0x41, 0x4d, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x53, 0x10,
	0x11, 0x12, 0x12, 0x0a, 0x0e, 0x56, 0x4f, 0x54, 0x45, 0x5f, 0x49, 0x53, 0x5f, 0x43, 0x4c, 0x4f,
	0x53, 0x45, 0x44, 0x10, 0x12, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x4f, 0x5f, 0x53, 0x55, 0x43, 0x48,
	0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x50, 0x49, 0x4e, 0x10, 0x13, 0x12, 0x14, 0x0a, 0x10,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54,
	0x10, 0x14, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4f, 0x4b, 0x54, 0x41, 0x49,
	0x44, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x15, 0x12, 0x1e, 0x0a, 0x1a, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x5f, 0x50, 0x49, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x54,
	0x4f, 0x4f, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x10, 0x16, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x54, 0x48,
	0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x63, 0x42, 0x3c, 0x0a, 0x30, 0x6f, 0x72,
	0x67, 0x2e, 0x77, 0x68, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73,
	0x2e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x67, 0x63, 0x6d, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x01,
	0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_common_proto_goTypes = []interface{}{
	(STATUS)(0),             // 0: STATUS
	(*Empty)(nil),           // 1: Empty
	(*Step)(nil),            // 2: Step
	(*BaseResponse)(nil),    // 3: BaseResponse
	(*BaseAnyResponse)(nil), // 4: BaseAnyResponse
	(*anypb.Any)(nil),       // 5: google.protobuf.Any
}
var file_common_proto_depIdxs = []int32{
	5, // 0: BaseResponse.data:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Step); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResponse); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseAnyResponse); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
