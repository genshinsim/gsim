// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protos/backend/submission.proto

package submission

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

type Submission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,json=_id,proto3" json:"id,omitempty"` // auto generated id for this submission
	Config      string `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	Submitter   string `protobuf:"bytes,3,opt,name=submitter,proto3" json:"submitter,omitempty"` //submitter discord id
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Preview     string `protobuf:"bytes,5,opt,name=preview,proto3" json:"preview,omitempty"` //view key for preview
}

func (x *Submission) Reset() {
	*x = Submission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_submission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Submission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Submission) ProtoMessage() {}

func (x *Submission) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_submission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Submission.ProtoReflect.Descriptor instead.
func (*Submission) Descriptor() ([]byte, []int) {
	return file_protos_backend_submission_proto_rawDescGZIP(), []int{0}
}

func (x *Submission) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Submission) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Submission) GetSubmitter() string {
	if x != nil {
		return x.Submitter
	}
	return ""
}

func (x *Submission) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Submission) GetPreview() string {
	if x != nil {
		return x.Preview
	}
	return ""
}

type SubmitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config      string `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	Submitter   string `protobuf:"bytes,2,opt,name=submitter,proto3" json:"submitter,omitempty"` //submitter discord id
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *SubmitRequest) Reset() {
	*x = SubmitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_submission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRequest) ProtoMessage() {}

func (x *SubmitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_submission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRequest.ProtoReflect.Descriptor instead.
func (*SubmitRequest) Descriptor() ([]byte, []int) {
	return file_protos_backend_submission_proto_rawDescGZIP(), []int{1}
}

func (x *SubmitRequest) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *SubmitRequest) GetSubmitter() string {
	if x != nil {
		return x.Submitter
	}
	return ""
}

func (x *SubmitRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type SubmitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SubmitResponse) Reset() {
	*x = SubmitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_submission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitResponse) ProtoMessage() {}

func (x *SubmitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_submission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitResponse.ProtoReflect.Descriptor instead.
func (*SubmitResponse) Descriptor() ([]byte, []int) {
	return file_protos_backend_submission_proto_rawDescGZIP(), []int{2}
}

func (x *SubmitResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserFilter string `protobuf:"bytes,1,opt,name=user_filter,json=userFilter,proto3" json:"user_filter,omitempty"` //comma separated string list of users
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_submission_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_submission_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_protos_backend_submission_proto_rawDescGZIP(), []int{3}
}

func (x *ListRequest) GetUserFilter() string {
	if x != nil {
		return x.UserFilter
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// queued for compute + waiting for approval
	Data []*Submission `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_submission_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_submission_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_protos_backend_submission_proto_rawDescGZIP(), []int{4}
}

func (x *ListResponse) GetData() []*Submission {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Config      string `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	Submitter   string `protobuf:"bytes,3,opt,name=submitter,proto3" json:"submitter,omitempty"` //submitter discord id
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_submission_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_submission_proto_msgTypes[5]
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
	return file_protos_backend_submission_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRequest) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *UpdateRequest) GetSubmitter() string {
	if x != nil {
		return x.Submitter
	}
	return ""
}

func (x *UpdateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_submission_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_submission_proto_msgTypes[6]
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
	return file_protos_backend_submission_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Submitter string `protobuf:"bytes,2,opt,name=submitter,proto3" json:"submitter,omitempty"`
}

func (x *RemoveRequest) Reset() {
	*x = RemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_submission_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRequest) ProtoMessage() {}

func (x *RemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_submission_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRequest.ProtoReflect.Descriptor instead.
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return file_protos_backend_submission_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RemoveRequest) GetSubmitter() string {
	if x != nil {
		return x.Submitter
	}
	return ""
}

type RemoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveResponse) Reset() {
	*x = RemoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_submission_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveResponse) ProtoMessage() {}

func (x *RemoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_submission_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveResponse.ProtoReflect.Descriptor instead.
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return file_protos_backend_submission_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RejectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *RejectRequest) Reset() {
	*x = RejectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_submission_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RejectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectRequest) ProtoMessage() {}

func (x *RejectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_submission_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectRequest.ProtoReflect.Descriptor instead.
func (*RejectRequest) Descriptor() ([]byte, []int) {
	return file_protos_backend_submission_proto_rawDescGZIP(), []int{9}
}

func (x *RejectRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RejectRequest) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type RejectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RejectResponse) Reset() {
	*x = RejectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_submission_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RejectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectResponse) ProtoMessage() {}

func (x *RejectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_submission_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectResponse.ProtoReflect.Descriptor instead.
func (*RejectResponse) Descriptor() ([]byte, []int) {
	return file_protos_backend_submission_proto_rawDescGZIP(), []int{10}
}

func (x *RejectResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ApproveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ApproveRequest) Reset() {
	*x = ApproveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_submission_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveRequest) ProtoMessage() {}

func (x *ApproveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_submission_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveRequest.ProtoReflect.Descriptor instead.
func (*ApproveRequest) Descriptor() ([]byte, []int) {
	return file_protos_backend_submission_proto_rawDescGZIP(), []int{11}
}

func (x *ApproveRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ApproveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DbKey string `protobuf:"bytes,2,opt,name=db_key,json=dbKey,proto3" json:"db_key,omitempty"`
}

func (x *ApproveResponse) Reset() {
	*x = ApproveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_submission_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveResponse) ProtoMessage() {}

func (x *ApproveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_submission_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveResponse.ProtoReflect.Descriptor instead.
func (*ApproveResponse) Descriptor() ([]byte, []int) {
	return file_protos_backend_submission_proto_rawDescGZIP(), []int{12}
}

func (x *ApproveResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ApproveResponse) GetDbKey() string {
	if x != nil {
		return x.DbKey
	}
	return ""
}

type ReplaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DbKey string `protobuf:"bytes,2,opt,name=db_key,json=dbKey,proto3" json:"db_key,omitempty"` //this is the db key that this entry is replacing
}

func (x *ReplaceRequest) Reset() {
	*x = ReplaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_submission_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceRequest) ProtoMessage() {}

func (x *ReplaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_submission_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceRequest.ProtoReflect.Descriptor instead.
func (*ReplaceRequest) Descriptor() ([]byte, []int) {
	return file_protos_backend_submission_proto_rawDescGZIP(), []int{13}
}

func (x *ReplaceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReplaceRequest) GetDbKey() string {
	if x != nil {
		return x.DbKey
	}
	return ""
}

type ReplaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DbKey string `protobuf:"bytes,2,opt,name=db_key,json=dbKey,proto3" json:"db_key,omitempty"`
}

func (x *ReplaceResponse) Reset() {
	*x = ReplaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_submission_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceResponse) ProtoMessage() {}

func (x *ReplaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_submission_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceResponse.ProtoReflect.Descriptor instead.
func (*ReplaceResponse) Descriptor() ([]byte, []int) {
	return file_protos_backend_submission_proto_rawDescGZIP(), []int{14}
}

func (x *ReplaceResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReplaceResponse) GetDbKey() string {
	if x != nil {
		return x.DbKey
	}
	return ""
}

var File_protos_backend_submission_proto protoreflect.FileDescriptor

var file_protos_backend_submission_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x8f, 0x01, 0x0a, 0x0a, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x5f, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22, 0x67, 0x0a, 0x0d, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x36, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x77,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x0d, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x22, 0x20, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x0d, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x22, 0x20, 0x0a, 0x0e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x62,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x62, 0x4b, 0x65,
	0x79, 0x22, 0x37, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x62, 0x4b, 0x65, 0x79, 0x22, 0x38, 0x0a, 0x0f, 0x52, 0x65,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x64, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64,
	0x62, 0x4b, 0x65, 0x79, 0x32, 0xae, 0x03, 0x0a, 0x0f, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x13, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c,
	0x0a, 0x07, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x07,
	0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_backend_submission_proto_rawDescOnce sync.Once
	file_protos_backend_submission_proto_rawDescData = file_protos_backend_submission_proto_rawDesc
)

func file_protos_backend_submission_proto_rawDescGZIP() []byte {
	file_protos_backend_submission_proto_rawDescOnce.Do(func() {
		file_protos_backend_submission_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_backend_submission_proto_rawDescData)
	})
	return file_protos_backend_submission_proto_rawDescData
}

var file_protos_backend_submission_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_protos_backend_submission_proto_goTypes = []interface{}{
	(*Submission)(nil),      // 0: result.Submission
	(*SubmitRequest)(nil),   // 1: result.SubmitRequest
	(*SubmitResponse)(nil),  // 2: result.SubmitResponse
	(*ListRequest)(nil),     // 3: result.ListRequest
	(*ListResponse)(nil),    // 4: result.ListResponse
	(*UpdateRequest)(nil),   // 5: result.UpdateRequest
	(*UpdateResponse)(nil),  // 6: result.UpdateResponse
	(*RemoveRequest)(nil),   // 7: result.RemoveRequest
	(*RemoveResponse)(nil),  // 8: result.RemoveResponse
	(*RejectRequest)(nil),   // 9: result.RejectRequest
	(*RejectResponse)(nil),  // 10: result.RejectResponse
	(*ApproveRequest)(nil),  // 11: result.ApproveRequest
	(*ApproveResponse)(nil), // 12: result.ApproveResponse
	(*ReplaceRequest)(nil),  // 13: result.ReplaceRequest
	(*ReplaceResponse)(nil), // 14: result.ReplaceResponse
}
var file_protos_backend_submission_proto_depIdxs = []int32{
	0,  // 0: result.ListResponse.data:type_name -> result.Submission
	1,  // 1: result.SubmissionStore.Submit:input_type -> result.SubmitRequest
	5,  // 2: result.SubmissionStore.Update:input_type -> result.UpdateRequest
	7,  // 3: result.SubmissionStore.Remove:input_type -> result.RemoveRequest
	3,  // 4: result.SubmissionStore.List:input_type -> result.ListRequest
	11, // 5: result.SubmissionStore.Approve:input_type -> result.ApproveRequest
	13, // 6: result.SubmissionStore.Replace:input_type -> result.ReplaceRequest
	9,  // 7: result.SubmissionStore.Reject:input_type -> result.RejectRequest
	2,  // 8: result.SubmissionStore.Submit:output_type -> result.SubmitResponse
	6,  // 9: result.SubmissionStore.Update:output_type -> result.UpdateResponse
	8,  // 10: result.SubmissionStore.Remove:output_type -> result.RemoveResponse
	4,  // 11: result.SubmissionStore.List:output_type -> result.ListResponse
	12, // 12: result.SubmissionStore.Approve:output_type -> result.ApproveResponse
	14, // 13: result.SubmissionStore.Replace:output_type -> result.ReplaceResponse
	10, // 14: result.SubmissionStore.Reject:output_type -> result.RejectResponse
	8,  // [8:15] is the sub-list for method output_type
	1,  // [1:8] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_protos_backend_submission_proto_init() }
func file_protos_backend_submission_proto_init() {
	if File_protos_backend_submission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_backend_submission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Submission); i {
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
		file_protos_backend_submission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitRequest); i {
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
		file_protos_backend_submission_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitResponse); i {
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
		file_protos_backend_submission_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_protos_backend_submission_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_protos_backend_submission_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_backend_submission_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_backend_submission_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRequest); i {
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
		file_protos_backend_submission_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveResponse); i {
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
		file_protos_backend_submission_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RejectRequest); i {
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
		file_protos_backend_submission_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RejectResponse); i {
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
		file_protos_backend_submission_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApproveRequest); i {
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
		file_protos_backend_submission_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApproveResponse); i {
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
		file_protos_backend_submission_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplaceRequest); i {
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
		file_protos_backend_submission_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplaceResponse); i {
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
			RawDescriptor: file_protos_backend_submission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_backend_submission_proto_goTypes,
		DependencyIndexes: file_protos_backend_submission_proto_depIdxs,
		MessageInfos:      file_protos_backend_submission_proto_msgTypes,
	}.Build()
	File_protos_backend_submission_proto = out.File
	file_protos_backend_submission_proto_rawDesc = nil
	file_protos_backend_submission_proto_goTypes = nil
	file_protos_backend_submission_proto_depIdxs = nil
}
