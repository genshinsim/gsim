// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: protos/backend/share.proto

package share

import (
	model "github.com/genshinsim/gcsim/pkg/model"
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

type ShareEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                  `protobuf:"bytes,1,opt,name=id,json=_id,proto3" json:"id,omitempty" bson:"_id,omitempty"`
	Result    *model.SimulationResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty" bson:"result,omitempty"`
	ExpiresAt uint64                  `protobuf:"varint,3,opt,name=expires_at,proto3" json:"expires_at,omitempty" bson:"expires_at,omitempty"`
	Submitter string                  `protobuf:"bytes,4,opt,name=submitter,proto3" json:"submitter,omitempty" bson:"submitter,omitempty"`
}

func (x *ShareEntry) Reset() {
	*x = ShareEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_share_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareEntry) ProtoMessage() {}

func (x *ShareEntry) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_share_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareEntry.ProtoReflect.Descriptor instead.
func (*ShareEntry) Descriptor() ([]byte, []int) {
	return file_protos_backend_share_proto_rawDescGZIP(), []int{0}
}

func (x *ShareEntry) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShareEntry) GetResult() *model.SimulationResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *ShareEntry) GetExpiresAt() uint64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *ShareEntry) GetSubmitter() string {
	if x != nil {
		return x.Submitter
	}
	return ""
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    *model.SimulationResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty" bson:"result,omitempty"`
	ExpiresAt uint64                  `protobuf:"varint,2,opt,name=expires_at,proto3" json:"expires_at,omitempty" bson:"expires_at,omitempty"`
	Submitter string                  `protobuf:"bytes,3,opt,name=submitter,proto3" json:"submitter,omitempty" bson:"submitter,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_share_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_share_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_protos_backend_share_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRequest) GetResult() *model.SimulationResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *CreateRequest) GetExpiresAt() uint64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *CreateRequest) GetSubmitter() string {
	if x != nil {
		return x.Submitter
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"id,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_share_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_share_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_protos_backend_share_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"id,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_share_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_share_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_protos_backend_share_proto_rawDescGZIP(), []int{3}
}

func (x *ReadRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"id,omitempty"`
	Result    *model.SimulationResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty" bson:"result,omitempty"`
	ExpiresAt uint64                  `protobuf:"varint,3,opt,name=expires_at,proto3" json:"expires_at,omitempty" bson:"expires_at,omitempty"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_share_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_share_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_protos_backend_share_proto_rawDescGZIP(), []int{4}
}

func (x *ReadResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReadResponse) GetResult() *model.SimulationResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *ReadResponse) GetExpiresAt() uint64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"id,omitempty"`
	Result    *model.SimulationResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty" bson:"result,omitempty"`
	ExpiresAt uint64                  `protobuf:"varint,3,opt,name=expires_at,proto3" json:"expires_at,omitempty" bson:"expires_at,omitempty"`
	Submitter string                  `protobuf:"bytes,4,opt,name=submitter,proto3" json:"submitter,omitempty" bson:"submitter,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_share_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_share_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_protos_backend_share_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRequest) GetResult() *model.SimulationResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *UpdateRequest) GetExpiresAt() uint64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *UpdateRequest) GetSubmitter() string {
	if x != nil {
		return x.Submitter
	}
	return ""
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"id,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_share_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_share_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_protos_backend_share_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SetTTLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"id,omitempty"`
	ExpiresAt uint64 `protobuf:"varint,3,opt,name=expires_at,proto3" json:"expires_at,omitempty" bson:"expires_at,omitempty"`
}

func (x *SetTTLRequest) Reset() {
	*x = SetTTLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_share_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTTLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTTLRequest) ProtoMessage() {}

func (x *SetTTLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_share_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTTLRequest.ProtoReflect.Descriptor instead.
func (*SetTTLRequest) Descriptor() ([]byte, []int) {
	return file_protos_backend_share_proto_rawDescGZIP(), []int{7}
}

func (x *SetTTLRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SetTTLRequest) GetExpiresAt() uint64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

type SetTTLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"id,omitempty"`
}

func (x *SetTTLResponse) Reset() {
	*x = SetTTLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_share_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTTLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTTLResponse) ProtoMessage() {}

func (x *SetTTLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_share_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTTLResponse.ProtoReflect.Descriptor instead.
func (*SetTTLResponse) Descriptor() ([]byte, []int) {
	return file_protos_backend_share_proto_rawDescGZIP(), []int{8}
}

func (x *SetTTLResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_share_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_share_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_protos_backend_share_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"id,omitempty"` //TODO: add deleted data to response in future
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_share_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_share_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_protos_backend_share_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RandomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RandomRequest) Reset() {
	*x = RandomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_share_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomRequest) ProtoMessage() {}

func (x *RandomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_share_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomRequest.ProtoReflect.Descriptor instead.
func (*RandomRequest) Descriptor() ([]byte, []int) {
	return file_protos_backend_share_proto_rawDescGZIP(), []int{11}
}

type RandomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"id,omitempty"`
}

func (x *RandomResponse) Reset() {
	*x = RandomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_share_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomResponse) ProtoMessage() {}

func (x *RandomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_share_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomResponse.ProtoReflect.Descriptor instead.
func (*RandomResponse) Descriptor() ([]byte, []int) {
	return file_protos_backend_share_proto_rawDescGZIP(), []int{12}
}

func (x *RandomResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_protos_backend_share_proto protoreflect.FileDescriptor

var file_protos_backend_share_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c,
	0x01, 0x0a, 0x0a, 0x53, 0x68, 0x61, 0x72, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0f, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x5f, 0x69, 0x64, 0x12, 0x2f,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x22, 0x7e, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x22, 0x20, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x1d, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6f,
	0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x22,
	0x8e, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f,
	0x61, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72,
	0x22, 0x20, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x3f, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x54, 0x54, 0x4c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x5f, 0x61, 0x74, 0x22, 0x20, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x54, 0x54, 0x4c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x20, 0x0a, 0x0e, 0x52, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xdc, 0x02, 0x0a, 0x0a,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x12, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x37, 0x0a, 0x06, 0x53, 0x65, 0x74, 0x54, 0x54, 0x4c, 0x12, 0x14, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x54, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x54, 0x4c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x37, 0x0a, 0x06, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x12, 0x14, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x6e, 0x73, 0x68, 0x69, 0x6e,
	0x73, 0x69, 0x6d, 0x2f, 0x67, 0x63, 0x73, 0x69, 0x6d, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_backend_share_proto_rawDescOnce sync.Once
	file_protos_backend_share_proto_rawDescData = file_protos_backend_share_proto_rawDesc
)

func file_protos_backend_share_proto_rawDescGZIP() []byte {
	file_protos_backend_share_proto_rawDescOnce.Do(func() {
		file_protos_backend_share_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_backend_share_proto_rawDescData)
	})
	return file_protos_backend_share_proto_rawDescData
}

var file_protos_backend_share_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_protos_backend_share_proto_goTypes = []any{
	(*ShareEntry)(nil),             // 0: share.ShareEntry
	(*CreateRequest)(nil),          // 1: share.CreateRequest
	(*CreateResponse)(nil),         // 2: share.CreateResponse
	(*ReadRequest)(nil),            // 3: share.ReadRequest
	(*ReadResponse)(nil),           // 4: share.ReadResponse
	(*UpdateRequest)(nil),          // 5: share.UpdateRequest
	(*UpdateResponse)(nil),         // 6: share.UpdateResponse
	(*SetTTLRequest)(nil),          // 7: share.SetTTLRequest
	(*SetTTLResponse)(nil),         // 8: share.SetTTLResponse
	(*DeleteRequest)(nil),          // 9: share.DeleteRequest
	(*DeleteResponse)(nil),         // 10: share.DeleteResponse
	(*RandomRequest)(nil),          // 11: share.RandomRequest
	(*RandomResponse)(nil),         // 12: share.RandomResponse
	(*model.SimulationResult)(nil), // 13: model.SimulationResult
}
var file_protos_backend_share_proto_depIdxs = []int32{
	13, // 0: share.ShareEntry.result:type_name -> model.SimulationResult
	13, // 1: share.CreateRequest.result:type_name -> model.SimulationResult
	13, // 2: share.ReadResponse.result:type_name -> model.SimulationResult
	13, // 3: share.UpdateRequest.result:type_name -> model.SimulationResult
	1,  // 4: share.ShareStore.Create:input_type -> share.CreateRequest
	3,  // 5: share.ShareStore.Read:input_type -> share.ReadRequest
	5,  // 6: share.ShareStore.Update:input_type -> share.UpdateRequest
	7,  // 7: share.ShareStore.SetTTL:input_type -> share.SetTTLRequest
	9,  // 8: share.ShareStore.Delete:input_type -> share.DeleteRequest
	11, // 9: share.ShareStore.Random:input_type -> share.RandomRequest
	2,  // 10: share.ShareStore.Create:output_type -> share.CreateResponse
	4,  // 11: share.ShareStore.Read:output_type -> share.ReadResponse
	6,  // 12: share.ShareStore.Update:output_type -> share.UpdateResponse
	8,  // 13: share.ShareStore.SetTTL:output_type -> share.SetTTLResponse
	10, // 14: share.ShareStore.Delete:output_type -> share.DeleteResponse
	12, // 15: share.ShareStore.Random:output_type -> share.RandomResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_protos_backend_share_proto_init() }
func file_protos_backend_share_proto_init() {
	if File_protos_backend_share_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_backend_share_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ShareEntry); i {
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
		file_protos_backend_share_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRequest); i {
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
		file_protos_backend_share_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateResponse); i {
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
		file_protos_backend_share_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ReadRequest); i {
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
		file_protos_backend_share_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ReadResponse); i {
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
		file_protos_backend_share_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateRequest); i {
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
		file_protos_backend_share_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateResponse); i {
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
		file_protos_backend_share_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*SetTTLRequest); i {
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
		file_protos_backend_share_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*SetTTLResponse); i {
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
		file_protos_backend_share_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRequest); i {
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
		file_protos_backend_share_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteResponse); i {
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
		file_protos_backend_share_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*RandomRequest); i {
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
		file_protos_backend_share_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*RandomResponse); i {
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
			RawDescriptor: file_protos_backend_share_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_backend_share_proto_goTypes,
		DependencyIndexes: file_protos_backend_share_proto_depIdxs,
		MessageInfos:      file_protos_backend_share_proto_msgTypes,
	}.Build()
	File_protos_backend_share_proto = out.File
	file_protos_backend_share_proto_rawDesc = nil
	file_protos_backend_share_proto_goTypes = nil
	file_protos_backend_share_proto_depIdxs = nil
}
