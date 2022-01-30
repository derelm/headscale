// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: headscale/v1/preauthkey.proto

package v1

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

type PreAuthKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace  string                 `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Id         string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Key        string                 `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Reusable   bool                   `protobuf:"varint,4,opt,name=reusable,proto3" json:"reusable,omitempty"`
	Ephemeral  bool                   `protobuf:"varint,5,opt,name=ephemeral,proto3" json:"ephemeral,omitempty"`
	Used       bool                   `protobuf:"varint,6,opt,name=used,proto3" json:"used,omitempty"`
	Expiration *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=expiration,proto3" json:"expiration,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *PreAuthKey) Reset() {
	*x = PreAuthKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_preauthkey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreAuthKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreAuthKey) ProtoMessage() {}

func (x *PreAuthKey) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreAuthKey.ProtoReflect.Descriptor instead.
func (*PreAuthKey) Descriptor() ([]byte, []int) {
	return file_headscale_v1_preauthkey_proto_rawDescGZIP(), []int{0}
}

func (x *PreAuthKey) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *PreAuthKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PreAuthKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PreAuthKey) GetReusable() bool {
	if x != nil {
		return x.Reusable
	}
	return false
}

func (x *PreAuthKey) GetEphemeral() bool {
	if x != nil {
		return x.Ephemeral
	}
	return false
}

func (x *PreAuthKey) GetUsed() bool {
	if x != nil {
		return x.Used
	}
	return false
}

func (x *PreAuthKey) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

func (x *PreAuthKey) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CreatePreAuthKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace  string                 `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Reusable   bool                   `protobuf:"varint,2,opt,name=reusable,proto3" json:"reusable,omitempty"`
	Ephemeral  bool                   `protobuf:"varint,3,opt,name=ephemeral,proto3" json:"ephemeral,omitempty"`
	Expiration *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *CreatePreAuthKeyRequest) Reset() {
	*x = CreatePreAuthKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_preauthkey_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePreAuthKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePreAuthKeyRequest) ProtoMessage() {}

func (x *CreatePreAuthKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePreAuthKeyRequest.ProtoReflect.Descriptor instead.
func (*CreatePreAuthKeyRequest) Descriptor() ([]byte, []int) {
	return file_headscale_v1_preauthkey_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePreAuthKeyRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreatePreAuthKeyRequest) GetReusable() bool {
	if x != nil {
		return x.Reusable
	}
	return false
}

func (x *CreatePreAuthKeyRequest) GetEphemeral() bool {
	if x != nil {
		return x.Ephemeral
	}
	return false
}

func (x *CreatePreAuthKeyRequest) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

type CreatePreAuthKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreAuthKey *PreAuthKey `protobuf:"bytes,1,opt,name=pre_auth_key,json=preAuthKey,proto3" json:"pre_auth_key,omitempty"`
}

func (x *CreatePreAuthKeyResponse) Reset() {
	*x = CreatePreAuthKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_preauthkey_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePreAuthKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePreAuthKeyResponse) ProtoMessage() {}

func (x *CreatePreAuthKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePreAuthKeyResponse.ProtoReflect.Descriptor instead.
func (*CreatePreAuthKeyResponse) Descriptor() ([]byte, []int) {
	return file_headscale_v1_preauthkey_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePreAuthKeyResponse) GetPreAuthKey() *PreAuthKey {
	if x != nil {
		return x.PreAuthKey
	}
	return nil
}

type ExpirePreAuthKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ExpirePreAuthKeyRequest) Reset() {
	*x = ExpirePreAuthKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_preauthkey_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpirePreAuthKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpirePreAuthKeyRequest) ProtoMessage() {}

func (x *ExpirePreAuthKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpirePreAuthKeyRequest.ProtoReflect.Descriptor instead.
func (*ExpirePreAuthKeyRequest) Descriptor() ([]byte, []int) {
	return file_headscale_v1_preauthkey_proto_rawDescGZIP(), []int{3}
}

func (x *ExpirePreAuthKeyRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ExpirePreAuthKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ExpirePreAuthKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExpirePreAuthKeyResponse) Reset() {
	*x = ExpirePreAuthKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_preauthkey_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpirePreAuthKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpirePreAuthKeyResponse) ProtoMessage() {}

func (x *ExpirePreAuthKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpirePreAuthKeyResponse.ProtoReflect.Descriptor instead.
func (*ExpirePreAuthKeyResponse) Descriptor() ([]byte, []int) {
	return file_headscale_v1_preauthkey_proto_rawDescGZIP(), []int{4}
}

type ListPreAuthKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *ListPreAuthKeysRequest) Reset() {
	*x = ListPreAuthKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_preauthkey_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPreAuthKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPreAuthKeysRequest) ProtoMessage() {}

func (x *ListPreAuthKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPreAuthKeysRequest.ProtoReflect.Descriptor instead.
func (*ListPreAuthKeysRequest) Descriptor() ([]byte, []int) {
	return file_headscale_v1_preauthkey_proto_rawDescGZIP(), []int{5}
}

func (x *ListPreAuthKeysRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type ListPreAuthKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreAuthKeys []*PreAuthKey `protobuf:"bytes,1,rep,name=pre_auth_keys,json=preAuthKeys,proto3" json:"pre_auth_keys,omitempty"`
}

func (x *ListPreAuthKeysResponse) Reset() {
	*x = ListPreAuthKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_preauthkey_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPreAuthKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPreAuthKeysResponse) ProtoMessage() {}

func (x *ListPreAuthKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_preauthkey_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPreAuthKeysResponse.ProtoReflect.Descriptor instead.
func (*ListPreAuthKeysResponse) Descriptor() ([]byte, []int) {
	return file_headscale_v1_preauthkey_proto_rawDescGZIP(), []int{6}
}

func (x *ListPreAuthKeysResponse) GetPreAuthKeys() []*PreAuthKey {
	if x != nil {
		return x.PreAuthKeys
	}
	return nil
}

var File_headscale_v1_preauthkey_proto protoreflect.FileDescriptor

var file_headscale_v1_preauthkey_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x65, 0x61, 0x75, 0x74, 0x68, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91,
	0x02, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x75, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x72, 0x65, 0x75, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x70, 0x68,
	0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x70,
	0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0xad, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x75, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x72, 0x65, 0x75, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x70, 0x68, 0x65,
	0x6d, 0x65, 0x72, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x70, 0x68,
	0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x3a, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x56, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a,
	0x0a, 0x0c, 0x70, 0x72, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x0a,
	0x70, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x22, 0x49, 0x0a, 0x17, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x1a, 0x0a, 0x18, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x50,
	0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x36, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x57, 0x0a, 0x17, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x68, 0x65,
	0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65,
	0x79, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6a, 0x75, 0x61, 0x6e, 0x66, 0x6f, 0x6e, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_headscale_v1_preauthkey_proto_rawDescOnce sync.Once
	file_headscale_v1_preauthkey_proto_rawDescData = file_headscale_v1_preauthkey_proto_rawDesc
)

func file_headscale_v1_preauthkey_proto_rawDescGZIP() []byte {
	file_headscale_v1_preauthkey_proto_rawDescOnce.Do(func() {
		file_headscale_v1_preauthkey_proto_rawDescData = protoimpl.X.CompressGZIP(file_headscale_v1_preauthkey_proto_rawDescData)
	})
	return file_headscale_v1_preauthkey_proto_rawDescData
}

var file_headscale_v1_preauthkey_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_headscale_v1_preauthkey_proto_goTypes = []interface{}{
	(*PreAuthKey)(nil),               // 0: headscale.v1.PreAuthKey
	(*CreatePreAuthKeyRequest)(nil),  // 1: headscale.v1.CreatePreAuthKeyRequest
	(*CreatePreAuthKeyResponse)(nil), // 2: headscale.v1.CreatePreAuthKeyResponse
	(*ExpirePreAuthKeyRequest)(nil),  // 3: headscale.v1.ExpirePreAuthKeyRequest
	(*ExpirePreAuthKeyResponse)(nil), // 4: headscale.v1.ExpirePreAuthKeyResponse
	(*ListPreAuthKeysRequest)(nil),   // 5: headscale.v1.ListPreAuthKeysRequest
	(*ListPreAuthKeysResponse)(nil),  // 6: headscale.v1.ListPreAuthKeysResponse
	(*timestamppb.Timestamp)(nil),    // 7: google.protobuf.Timestamp
}
var file_headscale_v1_preauthkey_proto_depIdxs = []int32{
	7, // 0: headscale.v1.PreAuthKey.expiration:type_name -> google.protobuf.Timestamp
	7, // 1: headscale.v1.PreAuthKey.created_at:type_name -> google.protobuf.Timestamp
	7, // 2: headscale.v1.CreatePreAuthKeyRequest.expiration:type_name -> google.protobuf.Timestamp
	0, // 3: headscale.v1.CreatePreAuthKeyResponse.pre_auth_key:type_name -> headscale.v1.PreAuthKey
	0, // 4: headscale.v1.ListPreAuthKeysResponse.pre_auth_keys:type_name -> headscale.v1.PreAuthKey
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_headscale_v1_preauthkey_proto_init() }
func file_headscale_v1_preauthkey_proto_init() {
	if File_headscale_v1_preauthkey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_headscale_v1_preauthkey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreAuthKey); i {
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
		file_headscale_v1_preauthkey_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePreAuthKeyRequest); i {
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
		file_headscale_v1_preauthkey_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePreAuthKeyResponse); i {
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
		file_headscale_v1_preauthkey_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpirePreAuthKeyRequest); i {
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
		file_headscale_v1_preauthkey_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpirePreAuthKeyResponse); i {
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
		file_headscale_v1_preauthkey_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPreAuthKeysRequest); i {
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
		file_headscale_v1_preauthkey_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPreAuthKeysResponse); i {
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
			RawDescriptor: file_headscale_v1_preauthkey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_headscale_v1_preauthkey_proto_goTypes,
		DependencyIndexes: file_headscale_v1_preauthkey_proto_depIdxs,
		MessageInfos:      file_headscale_v1_preauthkey_proto_msgTypes,
	}.Build()
	File_headscale_v1_preauthkey_proto = out.File
	file_headscale_v1_preauthkey_proto_rawDesc = nil
	file_headscale_v1_preauthkey_proto_goTypes = nil
	file_headscale_v1_preauthkey_proto_depIdxs = nil
}
