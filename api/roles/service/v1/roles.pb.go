// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: api/roles/service/v1/roles.proto

package v1

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

type CreateRolesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateRolesRequest) Reset() {
	*x = CreateRolesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_roles_service_v1_roles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRolesRequest) ProtoMessage() {}

func (x *CreateRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_roles_service_v1_roles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRolesRequest.ProtoReflect.Descriptor instead.
func (*CreateRolesRequest) Descriptor() ([]byte, []int) {
	return file_api_roles_service_v1_roles_proto_rawDescGZIP(), []int{0}
}

type CreateRolesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateRolesReply) Reset() {
	*x = CreateRolesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_roles_service_v1_roles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRolesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRolesReply) ProtoMessage() {}

func (x *CreateRolesReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_roles_service_v1_roles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRolesReply.ProtoReflect.Descriptor instead.
func (*CreateRolesReply) Descriptor() ([]byte, []int) {
	return file_api_roles_service_v1_roles_proto_rawDescGZIP(), []int{1}
}

type UpdateRolesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateRolesRequest) Reset() {
	*x = UpdateRolesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_roles_service_v1_roles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRolesRequest) ProtoMessage() {}

func (x *UpdateRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_roles_service_v1_roles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRolesRequest.ProtoReflect.Descriptor instead.
func (*UpdateRolesRequest) Descriptor() ([]byte, []int) {
	return file_api_roles_service_v1_roles_proto_rawDescGZIP(), []int{2}
}

type UpdateRolesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateRolesReply) Reset() {
	*x = UpdateRolesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_roles_service_v1_roles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRolesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRolesReply) ProtoMessage() {}

func (x *UpdateRolesReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_roles_service_v1_roles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRolesReply.ProtoReflect.Descriptor instead.
func (*UpdateRolesReply) Descriptor() ([]byte, []int) {
	return file_api_roles_service_v1_roles_proto_rawDescGZIP(), []int{3}
}

type DeleteRolesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRolesRequest) Reset() {
	*x = DeleteRolesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_roles_service_v1_roles_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRolesRequest) ProtoMessage() {}

func (x *DeleteRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_roles_service_v1_roles_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRolesRequest.ProtoReflect.Descriptor instead.
func (*DeleteRolesRequest) Descriptor() ([]byte, []int) {
	return file_api_roles_service_v1_roles_proto_rawDescGZIP(), []int{4}
}

type DeleteRolesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRolesReply) Reset() {
	*x = DeleteRolesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_roles_service_v1_roles_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRolesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRolesReply) ProtoMessage() {}

func (x *DeleteRolesReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_roles_service_v1_roles_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRolesReply.ProtoReflect.Descriptor instead.
func (*DeleteRolesReply) Descriptor() ([]byte, []int) {
	return file_api_roles_service_v1_roles_proto_rawDescGZIP(), []int{5}
}

type GetRolesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRolesRequest) Reset() {
	*x = GetRolesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_roles_service_v1_roles_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRolesRequest) ProtoMessage() {}

func (x *GetRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_roles_service_v1_roles_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRolesRequest.ProtoReflect.Descriptor instead.
func (*GetRolesRequest) Descriptor() ([]byte, []int) {
	return file_api_roles_service_v1_roles_proto_rawDescGZIP(), []int{6}
}

type GetRolesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRolesReply) Reset() {
	*x = GetRolesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_roles_service_v1_roles_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRolesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRolesReply) ProtoMessage() {}

func (x *GetRolesReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_roles_service_v1_roles_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRolesReply.ProtoReflect.Descriptor instead.
func (*GetRolesReply) Descriptor() ([]byte, []int) {
	return file_api_roles_service_v1_roles_proto_rawDescGZIP(), []int{7}
}

type ListRolesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRolesRequest) Reset() {
	*x = ListRolesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_roles_service_v1_roles_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRolesRequest) ProtoMessage() {}

func (x *ListRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_roles_service_v1_roles_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRolesRequest.ProtoReflect.Descriptor instead.
func (*ListRolesRequest) Descriptor() ([]byte, []int) {
	return file_api_roles_service_v1_roles_proto_rawDescGZIP(), []int{8}
}

type ListRolesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRolesReply) Reset() {
	*x = ListRolesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_roles_service_v1_roles_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRolesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRolesReply) ProtoMessage() {}

func (x *ListRolesReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_roles_service_v1_roles_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRolesReply.ProtoReflect.Descriptor instead.
func (*ListRolesReply) Descriptor() ([]byte, []int) {
	return file_api_roles_service_v1_roles_proto_rawDescGZIP(), []int{9}
}

var File_api_roles_service_v1_roles_proto protoreflect.FileDescriptor

var file_api_roles_service_v1_roles_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10,
	0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x32, 0xdd, 0x03, 0x0a, 0x05, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x5f, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5f, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5f, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x28, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x56, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x59, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c,
	0x65, 0x73, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x42, 0x4f, 0x0a, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x2d, 0x6d, 0x6f, 0x6f, 0x6d, 0x2f, 0x6d,
	0x6f, 0x6f, 0x6d, 0x2d, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_roles_service_v1_roles_proto_rawDescOnce sync.Once
	file_api_roles_service_v1_roles_proto_rawDescData = file_api_roles_service_v1_roles_proto_rawDesc
)

func file_api_roles_service_v1_roles_proto_rawDescGZIP() []byte {
	file_api_roles_service_v1_roles_proto_rawDescOnce.Do(func() {
		file_api_roles_service_v1_roles_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_roles_service_v1_roles_proto_rawDescData)
	})
	return file_api_roles_service_v1_roles_proto_rawDescData
}

var file_api_roles_service_v1_roles_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_roles_service_v1_roles_proto_goTypes = []interface{}{
	(*CreateRolesRequest)(nil), // 0: api.roles.service.v1.CreateRolesRequest
	(*CreateRolesReply)(nil),   // 1: api.roles.service.v1.CreateRolesReply
	(*UpdateRolesRequest)(nil), // 2: api.roles.service.v1.UpdateRolesRequest
	(*UpdateRolesReply)(nil),   // 3: api.roles.service.v1.UpdateRolesReply
	(*DeleteRolesRequest)(nil), // 4: api.roles.service.v1.DeleteRolesRequest
	(*DeleteRolesReply)(nil),   // 5: api.roles.service.v1.DeleteRolesReply
	(*GetRolesRequest)(nil),    // 6: api.roles.service.v1.GetRolesRequest
	(*GetRolesReply)(nil),      // 7: api.roles.service.v1.GetRolesReply
	(*ListRolesRequest)(nil),   // 8: api.roles.service.v1.ListRolesRequest
	(*ListRolesReply)(nil),     // 9: api.roles.service.v1.ListRolesReply
}
var file_api_roles_service_v1_roles_proto_depIdxs = []int32{
	0, // 0: api.roles.service.v1.Roles.CreateRoles:input_type -> api.roles.service.v1.CreateRolesRequest
	2, // 1: api.roles.service.v1.Roles.UpdateRoles:input_type -> api.roles.service.v1.UpdateRolesRequest
	4, // 2: api.roles.service.v1.Roles.DeleteRoles:input_type -> api.roles.service.v1.DeleteRolesRequest
	6, // 3: api.roles.service.v1.Roles.GetRoles:input_type -> api.roles.service.v1.GetRolesRequest
	8, // 4: api.roles.service.v1.Roles.ListRoles:input_type -> api.roles.service.v1.ListRolesRequest
	1, // 5: api.roles.service.v1.Roles.CreateRoles:output_type -> api.roles.service.v1.CreateRolesReply
	3, // 6: api.roles.service.v1.Roles.UpdateRoles:output_type -> api.roles.service.v1.UpdateRolesReply
	5, // 7: api.roles.service.v1.Roles.DeleteRoles:output_type -> api.roles.service.v1.DeleteRolesReply
	7, // 8: api.roles.service.v1.Roles.GetRoles:output_type -> api.roles.service.v1.GetRolesReply
	9, // 9: api.roles.service.v1.Roles.ListRoles:output_type -> api.roles.service.v1.ListRolesReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_roles_service_v1_roles_proto_init() }
func file_api_roles_service_v1_roles_proto_init() {
	if File_api_roles_service_v1_roles_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_roles_service_v1_roles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRolesRequest); i {
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
		file_api_roles_service_v1_roles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRolesReply); i {
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
		file_api_roles_service_v1_roles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRolesRequest); i {
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
		file_api_roles_service_v1_roles_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRolesReply); i {
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
		file_api_roles_service_v1_roles_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRolesRequest); i {
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
		file_api_roles_service_v1_roles_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRolesReply); i {
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
		file_api_roles_service_v1_roles_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRolesRequest); i {
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
		file_api_roles_service_v1_roles_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRolesReply); i {
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
		file_api_roles_service_v1_roles_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRolesRequest); i {
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
		file_api_roles_service_v1_roles_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRolesReply); i {
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
			RawDescriptor: file_api_roles_service_v1_roles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_roles_service_v1_roles_proto_goTypes,
		DependencyIndexes: file_api_roles_service_v1_roles_proto_depIdxs,
		MessageInfos:      file_api_roles_service_v1_roles_proto_msgTypes,
	}.Build()
	File_api_roles_service_v1_roles_proto = out.File
	file_api_roles_service_v1_roles_proto_rawDesc = nil
	file_api_roles_service_v1_roles_proto_goTypes = nil
	file_api_roles_service_v1_roles_proto_depIdxs = nil
}
