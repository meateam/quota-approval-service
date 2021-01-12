// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/quotaApproval/quotaApproval.proto

package quotaApproval

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type QuotaApprovalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	From      string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	Info      string `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	Size      int64  `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	CreatedAt int64  `protobuf:"varint,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt int64  `protobuf:"varint,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *QuotaApprovalRequest) Reset() {
	*x = QuotaApprovalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotaApproval_quotaApproval_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuotaApprovalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuotaApprovalRequest) ProtoMessage() {}

func (x *QuotaApprovalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quotaApproval_quotaApproval_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuotaApprovalRequest.ProtoReflect.Descriptor instead.
func (*QuotaApprovalRequest) Descriptor() ([]byte, []int) {
	return file_proto_quotaApproval_quotaApproval_proto_rawDescGZIP(), []int{0}
}

func (x *QuotaApprovalRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *QuotaApprovalRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *QuotaApprovalRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *QuotaApprovalRequest) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *QuotaApprovalRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *QuotaApprovalRequest) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *QuotaApprovalRequest) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type CreateQuotaApprovalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From       string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Info       string `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Size       int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	ModifiedBy string `protobuf:"bytes,4,opt,name=modifiedBy,proto3" json:"modifiedBy,omitempty"`
}

func (x *CreateQuotaApprovalRequest) Reset() {
	*x = CreateQuotaApprovalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotaApproval_quotaApproval_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuotaApprovalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuotaApprovalRequest) ProtoMessage() {}

func (x *CreateQuotaApprovalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quotaApproval_quotaApproval_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuotaApprovalRequest.ProtoReflect.Descriptor instead.
func (*CreateQuotaApprovalRequest) Descriptor() ([]byte, []int) {
	return file_proto_quotaApproval_quotaApproval_proto_rawDescGZIP(), []int{1}
}

func (x *CreateQuotaApprovalRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *CreateQuotaApprovalRequest) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *CreateQuotaApprovalRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *CreateQuotaApprovalRequest) GetModifiedBy() string {
	if x != nil {
		return x.ModifiedBy
	}
	return ""
}

type CreateQuotaApprovalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuotaApprovalRequest *QuotaApprovalRequest `protobuf:"bytes,1,opt,name=quotaApprovalRequest,proto3" json:"quotaApprovalRequest,omitempty"`
}

func (x *CreateQuotaApprovalResponse) Reset() {
	*x = CreateQuotaApprovalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotaApproval_quotaApproval_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuotaApprovalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuotaApprovalResponse) ProtoMessage() {}

func (x *CreateQuotaApprovalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quotaApproval_quotaApproval_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuotaApprovalResponse.ProtoReflect.Descriptor instead.
func (*CreateQuotaApprovalResponse) Descriptor() ([]byte, []int) {
	return file_proto_quotaApproval_quotaApproval_proto_rawDescGZIP(), []int{2}
}

func (x *CreateQuotaApprovalResponse) GetQuotaApprovalRequest() *QuotaApprovalRequest {
	if x != nil {
		return x.QuotaApprovalRequest
	}
	return nil
}

type GetQuotasApprovalsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedBy    string `protobuf:"bytes,1,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	ApprovableBy string `protobuf:"bytes,2,opt,name=approvableBy,proto3" json:"approvableBy,omitempty"`
}

func (x *GetQuotasApprovalsRequest) Reset() {
	*x = GetQuotasApprovalsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotaApproval_quotaApproval_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuotasApprovalsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuotasApprovalsRequest) ProtoMessage() {}

func (x *GetQuotasApprovalsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quotaApproval_quotaApproval_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuotasApprovalsRequest.ProtoReflect.Descriptor instead.
func (*GetQuotasApprovalsRequest) Descriptor() ([]byte, []int) {
	return file_proto_quotaApproval_quotaApproval_proto_rawDescGZIP(), []int{3}
}

func (x *GetQuotasApprovalsRequest) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *GetQuotasApprovalsRequest) GetApprovableBy() string {
	if x != nil {
		return x.ApprovableBy
	}
	return ""
}

type GetQuotasApprovalsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuotaApprovalRequests []*QuotaApprovalRequest `protobuf:"bytes,1,rep,name=quotaApprovalRequests,proto3" json:"quotaApprovalRequests,omitempty"`
}

func (x *GetQuotasApprovalsResponse) Reset() {
	*x = GetQuotasApprovalsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotaApproval_quotaApproval_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuotasApprovalsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuotasApprovalsResponse) ProtoMessage() {}

func (x *GetQuotasApprovalsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quotaApproval_quotaApproval_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuotasApprovalsResponse.ProtoReflect.Descriptor instead.
func (*GetQuotasApprovalsResponse) Descriptor() ([]byte, []int) {
	return file_proto_quotaApproval_quotaApproval_proto_rawDescGZIP(), []int{4}
}

func (x *GetQuotasApprovalsResponse) GetQuotaApprovalRequests() []*QuotaApprovalRequest {
	if x != nil {
		return x.QuotaApprovalRequests
	}
	return nil
}

type GetQuotaApprovalByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetQuotaApprovalByIDRequest) Reset() {
	*x = GetQuotaApprovalByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotaApproval_quotaApproval_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuotaApprovalByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuotaApprovalByIDRequest) ProtoMessage() {}

func (x *GetQuotaApprovalByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quotaApproval_quotaApproval_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuotaApprovalByIDRequest.ProtoReflect.Descriptor instead.
func (*GetQuotaApprovalByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_quotaApproval_quotaApproval_proto_rawDescGZIP(), []int{5}
}

func (x *GetQuotaApprovalByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetQuotaApprovalByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuotaApprovalRequest *QuotaApprovalRequest `protobuf:"bytes,1,opt,name=quotaApprovalRequest,proto3" json:"quotaApprovalRequest,omitempty"`
}

func (x *GetQuotaApprovalByIDResponse) Reset() {
	*x = GetQuotaApprovalByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotaApproval_quotaApproval_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuotaApprovalByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuotaApprovalByIDResponse) ProtoMessage() {}

func (x *GetQuotaApprovalByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quotaApproval_quotaApproval_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuotaApprovalByIDResponse.ProtoReflect.Descriptor instead.
func (*GetQuotaApprovalByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_quotaApproval_quotaApproval_proto_rawDescGZIP(), []int{6}
}

func (x *GetQuotaApprovalByIDResponse) GetQuotaApprovalRequest() *QuotaApprovalRequest {
	if x != nil {
		return x.QuotaApprovalRequest
	}
	return nil
}

type UpdateQuotaApprovalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ModifiedBy string `protobuf:"bytes,2,opt,name=modifiedBy,proto3" json:"modifiedBy,omitempty"`
	Status     string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateQuotaApprovalRequest) Reset() {
	*x = UpdateQuotaApprovalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotaApproval_quotaApproval_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQuotaApprovalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuotaApprovalRequest) ProtoMessage() {}

func (x *UpdateQuotaApprovalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quotaApproval_quotaApproval_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuotaApprovalRequest.ProtoReflect.Descriptor instead.
func (*UpdateQuotaApprovalRequest) Descriptor() ([]byte, []int) {
	return file_proto_quotaApproval_quotaApproval_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateQuotaApprovalRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateQuotaApprovalRequest) GetModifiedBy() string {
	if x != nil {
		return x.ModifiedBy
	}
	return ""
}

func (x *UpdateQuotaApprovalRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateQuotaApprovalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuotaApprovalRequest *QuotaApprovalRequest `protobuf:"bytes,1,opt,name=quotaApprovalRequest,proto3" json:"quotaApprovalRequest,omitempty"`
}

func (x *UpdateQuotaApprovalResponse) Reset() {
	*x = UpdateQuotaApprovalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotaApproval_quotaApproval_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQuotaApprovalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuotaApprovalResponse) ProtoMessage() {}

func (x *UpdateQuotaApprovalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quotaApproval_quotaApproval_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuotaApprovalResponse.ProtoReflect.Descriptor instead.
func (*UpdateQuotaApprovalResponse) Descriptor() ([]byte, []int) {
	return file_proto_quotaApproval_quotaApproval_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateQuotaApprovalResponse) GetQuotaApprovalRequest() *QuotaApprovalRequest {
	if x != nil {
		return x.QuotaApprovalRequest
	}
	return nil
}

var File_proto_quotaApproval_quotaApproval_proto protoreflect.FileDescriptor

var file_proto_quotaApproval_quotaApproval_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x71, 0x75, 0x6f, 0x74, 0x61,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x22, 0xb6, 0x01, 0x0a, 0x14, 0x51, 0x75, 0x6f,
	0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x78, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x79, 0x22, 0x76, 0x0a, 0x1b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x14, 0x71, 0x75,
	0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x14, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x5d, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x73,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x22,
	0x0a, 0x0c, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x62, 0x6c, 0x65,
	0x42, 0x79, 0x22, 0x77, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x73, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x59, 0x0a, 0x15, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e,
	0x51, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x15, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x2d, 0x0a, 0x1b, 0x47,
	0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x77, 0x0a, 0x1c, 0x47, 0x65,
	0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x14, 0x71, 0x75,
	0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x14, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x64, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f,
	0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x76, 0x0a, 0x1b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x14, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x14, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x32, 0xcf, 0x03, 0x0a, 0x0d, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x12, 0x6e, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f,
	0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x29, 0x2e, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74,
	0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x73,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x73, 0x12, 0x28, 0x2e, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f,
	0x74, 0x61, 0x73, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x73, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x71, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x42, 0x79, 0x49, 0x44, 0x12, 0x2a, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74,
	0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f,
	0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x29, 0x2e, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74,
	0x61, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_quotaApproval_quotaApproval_proto_rawDescOnce sync.Once
	file_proto_quotaApproval_quotaApproval_proto_rawDescData = file_proto_quotaApproval_quotaApproval_proto_rawDesc
)

func file_proto_quotaApproval_quotaApproval_proto_rawDescGZIP() []byte {
	file_proto_quotaApproval_quotaApproval_proto_rawDescOnce.Do(func() {
		file_proto_quotaApproval_quotaApproval_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_quotaApproval_quotaApproval_proto_rawDescData)
	})
	return file_proto_quotaApproval_quotaApproval_proto_rawDescData
}

var file_proto_quotaApproval_quotaApproval_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_quotaApproval_quotaApproval_proto_goTypes = []interface{}{
	(*QuotaApprovalRequest)(nil),         // 0: quotaApproval.QuotaApprovalRequest
	(*CreateQuotaApprovalRequest)(nil),   // 1: quotaApproval.CreateQuotaApprovalRequest
	(*CreateQuotaApprovalResponse)(nil),  // 2: quotaApproval.CreateQuotaApprovalResponse
	(*GetQuotasApprovalsRequest)(nil),    // 3: quotaApproval.GetQuotasApprovalsRequest
	(*GetQuotasApprovalsResponse)(nil),   // 4: quotaApproval.GetQuotasApprovalsResponse
	(*GetQuotaApprovalByIDRequest)(nil),  // 5: quotaApproval.GetQuotaApprovalByIDRequest
	(*GetQuotaApprovalByIDResponse)(nil), // 6: quotaApproval.GetQuotaApprovalByIDResponse
	(*UpdateQuotaApprovalRequest)(nil),   // 7: quotaApproval.UpdateQuotaApprovalRequest
	(*UpdateQuotaApprovalResponse)(nil),  // 8: quotaApproval.UpdateQuotaApprovalResponse
}
var file_proto_quotaApproval_quotaApproval_proto_depIdxs = []int32{
	0, // 0: quotaApproval.CreateQuotaApprovalResponse.quotaApprovalRequest:type_name -> quotaApproval.QuotaApprovalRequest
	0, // 1: quotaApproval.GetQuotasApprovalsResponse.quotaApprovalRequests:type_name -> quotaApproval.QuotaApprovalRequest
	0, // 2: quotaApproval.GetQuotaApprovalByIDResponse.quotaApprovalRequest:type_name -> quotaApproval.QuotaApprovalRequest
	0, // 3: quotaApproval.UpdateQuotaApprovalResponse.quotaApprovalRequest:type_name -> quotaApproval.QuotaApprovalRequest
	1, // 4: quotaApproval.QuotaApproval.CreateQuotaApproval:input_type -> quotaApproval.CreateQuotaApprovalRequest
	3, // 5: quotaApproval.QuotaApproval.GetQuotasApprovals:input_type -> quotaApproval.GetQuotasApprovalsRequest
	5, // 6: quotaApproval.QuotaApproval.GetQuotaApprovalByID:input_type -> quotaApproval.GetQuotaApprovalByIDRequest
	7, // 7: quotaApproval.QuotaApproval.UpdateQuotaApproval:input_type -> quotaApproval.UpdateQuotaApprovalRequest
	2, // 8: quotaApproval.QuotaApproval.CreateQuotaApproval:output_type -> quotaApproval.CreateQuotaApprovalResponse
	4, // 9: quotaApproval.QuotaApproval.GetQuotasApprovals:output_type -> quotaApproval.GetQuotasApprovalsResponse
	6, // 10: quotaApproval.QuotaApproval.GetQuotaApprovalByID:output_type -> quotaApproval.GetQuotaApprovalByIDResponse
	8, // 11: quotaApproval.QuotaApproval.UpdateQuotaApproval:output_type -> quotaApproval.UpdateQuotaApprovalResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_quotaApproval_quotaApproval_proto_init() }
func file_proto_quotaApproval_quotaApproval_proto_init() {
	if File_proto_quotaApproval_quotaApproval_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_quotaApproval_quotaApproval_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuotaApprovalRequest); i {
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
		file_proto_quotaApproval_quotaApproval_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuotaApprovalRequest); i {
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
		file_proto_quotaApproval_quotaApproval_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuotaApprovalResponse); i {
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
		file_proto_quotaApproval_quotaApproval_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuotasApprovalsRequest); i {
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
		file_proto_quotaApproval_quotaApproval_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuotasApprovalsResponse); i {
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
		file_proto_quotaApproval_quotaApproval_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuotaApprovalByIDRequest); i {
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
		file_proto_quotaApproval_quotaApproval_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuotaApprovalByIDResponse); i {
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
		file_proto_quotaApproval_quotaApproval_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQuotaApprovalRequest); i {
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
		file_proto_quotaApproval_quotaApproval_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQuotaApprovalResponse); i {
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
			RawDescriptor: file_proto_quotaApproval_quotaApproval_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_quotaApproval_quotaApproval_proto_goTypes,
		DependencyIndexes: file_proto_quotaApproval_quotaApproval_proto_depIdxs,
		MessageInfos:      file_proto_quotaApproval_quotaApproval_proto_msgTypes,
	}.Build()
	File_proto_quotaApproval_quotaApproval_proto = out.File
	file_proto_quotaApproval_quotaApproval_proto_rawDesc = nil
	file_proto_quotaApproval_quotaApproval_proto_goTypes = nil
	file_proto_quotaApproval_quotaApproval_proto_depIdxs = nil
}
