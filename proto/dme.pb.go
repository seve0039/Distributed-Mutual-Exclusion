// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: proto/dme.proto

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

// Message for requestiong access to CS
type CriticalSectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *CriticalSectionRequest) Reset() {
	*x = CriticalSectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dme_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CriticalSectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CriticalSectionRequest) ProtoMessage() {}

func (x *CriticalSectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dme_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CriticalSectionRequest.ProtoReflect.Descriptor instead.
func (*CriticalSectionRequest) Descriptor() ([]byte, []int) {
	return file_proto_dme_proto_rawDescGZIP(), []int{0}
}

func (x *CriticalSectionRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

// Response for requestiong access to CS
type CriticalSectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasAccess bool `protobuf:"varint,1,opt,name=hasAccess,proto3" json:"hasAccess,omitempty"`
}

func (x *CriticalSectionResponse) Reset() {
	*x = CriticalSectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dme_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CriticalSectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CriticalSectionResponse) ProtoMessage() {}

func (x *CriticalSectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dme_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CriticalSectionResponse.ProtoReflect.Descriptor instead.
func (*CriticalSectionResponse) Descriptor() ([]byte, []int) {
	return file_proto_dme_proto_rawDescGZIP(), []int{1}
}

func (x *CriticalSectionResponse) GetHasAccess() bool {
	if x != nil {
		return x.HasAccess
	}
	return false
}

// Aknowledge message
type Acknowledge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Acknowledged bool `protobuf:"varint,1,opt,name=acknowledged,proto3" json:"acknowledged,omitempty"`
}

func (x *Acknowledge) Reset() {
	*x = Acknowledge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dme_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Acknowledge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Acknowledge) ProtoMessage() {}

func (x *Acknowledge) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dme_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Acknowledge.ProtoReflect.Descriptor instead.
func (*Acknowledge) Descriptor() ([]byte, []int) {
	return file_proto_dme_proto_rawDescGZIP(), []int{2}
}

func (x *Acknowledge) GetAcknowledged() bool {
	if x != nil {
		return x.Acknowledged
	}
	return false
}

// Message representing the token in the Token Ring
type TokenMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HolderNodeId string `protobuf:"bytes,1,opt,name=holderNodeId,proto3" json:"holderNodeId,omitempty"`
}

func (x *TokenMessage) Reset() {
	*x = TokenMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dme_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenMessage) ProtoMessage() {}

func (x *TokenMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dme_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenMessage.ProtoReflect.Descriptor instead.
func (*TokenMessage) Descriptor() ([]byte, []int) {
	return file_proto_dme_proto_rawDescGZIP(), []int{3}
}

func (x *TokenMessage) GetHolderNodeId() string {
	if x != nil {
		return x.HolderNodeId
	}
	return ""
}

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dme_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dme_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_proto_dme_proto_rawDescGZIP(), []int{4}
}

func (x *JoinRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type JoinAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *JoinAck) Reset() {
	*x = JoinAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dme_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinAck) ProtoMessage() {}

func (x *JoinAck) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dme_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinAck.ProtoReflect.Descriptor instead.
func (*JoinAck) Descriptor() ([]byte, []int) {
	return file_proto_dme_proto_rawDescGZIP(), []int{5}
}

func (x *JoinAck) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type LeaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *LeaveRequest) Reset() {
	*x = LeaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dme_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveRequest) ProtoMessage() {}

func (x *LeaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dme_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveRequest.ProtoReflect.Descriptor instead.
func (*LeaveRequest) Descriptor() ([]byte, []int) {
	return file_proto_dme_proto_rawDescGZIP(), []int{6}
}

func (x *LeaveRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type LeaveAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *LeaveAck) Reset() {
	*x = LeaveAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dme_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveAck) ProtoMessage() {}

func (x *LeaveAck) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dme_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveAck.ProtoReflect.Descriptor instead.
func (*LeaveAck) Descriptor() ([]byte, []int) {
	return file_proto_dme_proto_rawDescGZIP(), []int{7}
}

func (x *LeaveAck) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_dme_proto protoreflect.FileDescriptor

var file_proto_dme_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x16, 0x43, 0x72, 0x69, 0x74,
	0x69, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x17, 0x43,
	0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x61, 0x73, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x31, 0x0a, 0x0b, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x63, 0x6b, 0x6e, 0x6f,
	0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x64, 0x22, 0x32, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x0b, 0x4a,
	0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x23,
	0x0a, 0x07, 0x4a, 0x6f, 0x69, 0x6e, 0x41, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x22, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x08, 0x4c, 0x65, 0x61, 0x76, 0x65,
	0x41, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xb2, 0x02,
	0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x69, 0x6e, 0x67, 0x12, 0x5b, 0x0a, 0x16, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72,
	0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x69,
	0x74, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x2a, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e,
	0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x69,
	0x6e, 0x41, 0x63, 0x6b, 0x12, 0x2d, 0x0a, 0x05, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65,
	0x41, 0x63, 0x6b, 0x12, 0x34, 0x0a, 0x09, 0x50, 0x61, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63,
	0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x52, 0x65, 0x63,
	0x69, 0x65, 0x76, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x65, 0x76, 0x65, 0x30, 0x30, 0x33, 0x39, 0x2f, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x64, 0x2d, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x2d, 0x45, 0x78, 0x63,
	0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_dme_proto_rawDescOnce sync.Once
	file_proto_dme_proto_rawDescData = file_proto_dme_proto_rawDesc
)

func file_proto_dme_proto_rawDescGZIP() []byte {
	file_proto_dme_proto_rawDescOnce.Do(func() {
		file_proto_dme_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_dme_proto_rawDescData)
	})
	return file_proto_dme_proto_rawDescData
}

var file_proto_dme_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_dme_proto_goTypes = []interface{}{
	(*CriticalSectionRequest)(nil),  // 0: proto.CriticalSectionRequest
	(*CriticalSectionResponse)(nil), // 1: proto.CriticalSectionResponse
	(*Acknowledge)(nil),             // 2: proto.Acknowledge
	(*TokenMessage)(nil),            // 3: proto.TokenMessage
	(*JoinRequest)(nil),             // 4: proto.JoinRequest
	(*JoinAck)(nil),                 // 5: proto.JoinAck
	(*LeaveRequest)(nil),            // 6: proto.LeaveRequest
	(*LeaveAck)(nil),                // 7: proto.LeaveAck
}
var file_proto_dme_proto_depIdxs = []int32{
	0, // 0: proto.TokenRing.RequestCriticalSection:input_type -> proto.CriticalSectionRequest
	4, // 1: proto.TokenRing.Join:input_type -> proto.JoinRequest
	6, // 2: proto.TokenRing.Leave:input_type -> proto.LeaveRequest
	3, // 3: proto.TokenRing.PassToken:input_type -> proto.TokenMessage
	3, // 4: proto.TokenRing.RecieveToken:input_type -> proto.TokenMessage
	1, // 5: proto.TokenRing.RequestCriticalSection:output_type -> proto.CriticalSectionResponse
	5, // 6: proto.TokenRing.Join:output_type -> proto.JoinAck
	7, // 7: proto.TokenRing.Leave:output_type -> proto.LeaveAck
	2, // 8: proto.TokenRing.PassToken:output_type -> proto.Acknowledge
	2, // 9: proto.TokenRing.RecieveToken:output_type -> proto.Acknowledge
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_dme_proto_init() }
func file_proto_dme_proto_init() {
	if File_proto_dme_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_dme_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CriticalSectionRequest); i {
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
		file_proto_dme_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CriticalSectionResponse); i {
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
		file_proto_dme_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Acknowledge); i {
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
		file_proto_dme_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenMessage); i {
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
		file_proto_dme_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
		file_proto_dme_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinAck); i {
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
		file_proto_dme_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveRequest); i {
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
		file_proto_dme_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveAck); i {
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
			RawDescriptor: file_proto_dme_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_dme_proto_goTypes,
		DependencyIndexes: file_proto_dme_proto_depIdxs,
		MessageInfos:      file_proto_dme_proto_msgTypes,
	}.Build()
	File_proto_dme_proto = out.File
	file_proto_dme_proto_rawDesc = nil
	file_proto_dme_proto_goTypes = nil
	file_proto_dme_proto_depIdxs = nil
}
