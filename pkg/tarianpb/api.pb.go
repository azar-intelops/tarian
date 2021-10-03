// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: tarianpb/api.proto

package tarianpb

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

type GetConstraintsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Labels    []*Label `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *GetConstraintsRequest) Reset() {
	*x = GetConstraintsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConstraintsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConstraintsRequest) ProtoMessage() {}

func (x *GetConstraintsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConstraintsRequest.ProtoReflect.Descriptor instead.
func (*GetConstraintsRequest) Descriptor() ([]byte, []int) {
	return file_tarianpb_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetConstraintsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GetConstraintsRequest) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

type GetConstraintsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Constraints []*Constraint `protobuf:"bytes,1,rep,name=constraints,proto3" json:"constraints,omitempty"`
}

func (x *GetConstraintsResponse) Reset() {
	*x = GetConstraintsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConstraintsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConstraintsResponse) ProtoMessage() {}

func (x *GetConstraintsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConstraintsResponse.ProtoReflect.Descriptor instead.
func (*GetConstraintsResponse) Descriptor() ([]byte, []int) {
	return file_tarianpb_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetConstraintsResponse) GetConstraints() []*Constraint {
	if x != nil {
		return x.Constraints
	}
	return nil
}

type AddConstraintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Constraint *Constraint `protobuf:"bytes,1,opt,name=constraint,proto3" json:"constraint,omitempty"`
}

func (x *AddConstraintRequest) Reset() {
	*x = AddConstraintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddConstraintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddConstraintRequest) ProtoMessage() {}

func (x *AddConstraintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddConstraintRequest.ProtoReflect.Descriptor instead.
func (*AddConstraintRequest) Descriptor() ([]byte, []int) {
	return file_tarianpb_api_proto_rawDescGZIP(), []int{2}
}

func (x *AddConstraintRequest) GetConstraint() *Constraint {
	if x != nil {
		return x.Constraint
	}
	return nil
}

type AddConstraintResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddConstraintResponse) Reset() {
	*x = AddConstraintResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddConstraintResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddConstraintResponse) ProtoMessage() {}

func (x *AddConstraintResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddConstraintResponse.ProtoReflect.Descriptor instead.
func (*AddConstraintResponse) Descriptor() ([]byte, []int) {
	return file_tarianpb_api_proto_rawDescGZIP(), []int{3}
}

func (x *AddConstraintResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RemoveConstraintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RemoveConstraintRequest) Reset() {
	*x = RemoveConstraintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveConstraintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveConstraintRequest) ProtoMessage() {}

func (x *RemoveConstraintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveConstraintRequest.ProtoReflect.Descriptor instead.
func (*RemoveConstraintRequest) Descriptor() ([]byte, []int) {
	return file_tarianpb_api_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveConstraintRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *RemoveConstraintRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RemoveConstraintResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RemoveConstraintResponse) Reset() {
	*x = RemoveConstraintResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveConstraintResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveConstraintResponse) ProtoMessage() {}

func (x *RemoveConstraintResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveConstraintResponse.ProtoReflect.Descriptor instead.
func (*RemoveConstraintResponse) Descriptor() ([]byte, []int) {
	return file_tarianpb_api_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveConstraintResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type AddActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action *Action `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *AddActionRequest) Reset() {
	*x = AddActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddActionRequest) ProtoMessage() {}

func (x *AddActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddActionRequest.ProtoReflect.Descriptor instead.
func (*AddActionRequest) Descriptor() ([]byte, []int) {
	return file_tarianpb_api_proto_rawDescGZIP(), []int{6}
}

func (x *AddActionRequest) GetAction() *Action {
	if x != nil {
		return x.Action
	}
	return nil
}

type AddActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddActionResponse) Reset() {
	*x = AddActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddActionResponse) ProtoMessage() {}

func (x *AddActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddActionResponse.ProtoReflect.Descriptor instead.
func (*AddActionResponse) Descriptor() ([]byte, []int) {
	return file_tarianpb_api_proto_rawDescGZIP(), []int{7}
}

func (x *AddActionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetActionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Labels    []*Label `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *GetActionsRequest) Reset() {
	*x = GetActionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActionsRequest) ProtoMessage() {}

func (x *GetActionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActionsRequest.ProtoReflect.Descriptor instead.
func (*GetActionsRequest) Descriptor() ([]byte, []int) {
	return file_tarianpb_api_proto_rawDescGZIP(), []int{8}
}

func (x *GetActionsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GetActionsRequest) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

type GetActionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Actions []*Action `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *GetActionsResponse) Reset() {
	*x = GetActionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActionsResponse) ProtoMessage() {}

func (x *GetActionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActionsResponse.ProtoReflect.Descriptor instead.
func (*GetActionsResponse) Descriptor() ([]byte, []int) {
	return file_tarianpb_api_proto_rawDescGZIP(), []int{9}
}

func (x *GetActionsResponse) GetActions() []*Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

type IngestEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *IngestEventRequest) Reset() {
	*x = IngestEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestEventRequest) ProtoMessage() {}

func (x *IngestEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestEventRequest.ProtoReflect.Descriptor instead.
func (*IngestEventRequest) Descriptor() ([]byte, []int) {
	return file_tarianpb_api_proto_rawDescGZIP(), []int{10}
}

func (x *IngestEventRequest) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type IngestEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *IngestEventResponse) Reset() {
	*x = IngestEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestEventResponse) ProtoMessage() {}

func (x *IngestEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestEventResponse.ProtoReflect.Descriptor instead.
func (*IngestEventResponse) Descriptor() ([]byte, []int) {
	return file_tarianpb_api_proto_rawDescGZIP(), []int{11}
}

func (x *IngestEventResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Limit     uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetEventsRequest) Reset() {
	*x = GetEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsRequest) ProtoMessage() {}

func (x *GetEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsRequest.ProtoReflect.Descriptor instead.
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return file_tarianpb_api_proto_rawDescGZIP(), []int{12}
}

func (x *GetEventsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GetEventsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *GetEventsResponse) Reset() {
	*x = GetEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarianpb_api_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsResponse) ProtoMessage() {}

func (x *GetEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tarianpb_api_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsResponse.ProtoReflect.Descriptor instead.
func (*GetEventsResponse) Descriptor() ([]byte, []int) {
	return file_tarianpb_api_proto_rawDescGZIP(), []int{13}
}

func (x *GetEventsResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_tarianpb_api_proto protoreflect.FileDescriptor

var file_tarianpb_api_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x61,
	0x70, 0x69, 0x1a, 0x14, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x2d, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x56,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x52, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a,
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x0a,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x22, 0x31, 0x0a, 0x15, 0x41, 0x64,
	0x64, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4b, 0x0a,
	0x17, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x18, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x42, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2d, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x60, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70,
	0x62, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x46, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x41, 0x0a,
	0x12, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x22, 0x2f, 0x0a, 0x13, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x46, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x42, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d,
	0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xcb, 0x03,
	0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x23, 0x2e, 0x74, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x22, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x70, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x74,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x70, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70,
	0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x70, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0xad, 0x01, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x54, 0x0a, 0x0b, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70,
	0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x70, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x70, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x2d, 0x74,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_tarianpb_api_proto_rawDescOnce sync.Once
	file_tarianpb_api_proto_rawDescData = file_tarianpb_api_proto_rawDesc
)

func file_tarianpb_api_proto_rawDescGZIP() []byte {
	file_tarianpb_api_proto_rawDescOnce.Do(func() {
		file_tarianpb_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_tarianpb_api_proto_rawDescData)
	})
	return file_tarianpb_api_proto_rawDescData
}

var file_tarianpb_api_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_tarianpb_api_proto_goTypes = []interface{}{
	(*GetConstraintsRequest)(nil),    // 0: tarianpb.api.GetConstraintsRequest
	(*GetConstraintsResponse)(nil),   // 1: tarianpb.api.GetConstraintsResponse
	(*AddConstraintRequest)(nil),     // 2: tarianpb.api.AddConstraintRequest
	(*AddConstraintResponse)(nil),    // 3: tarianpb.api.AddConstraintResponse
	(*RemoveConstraintRequest)(nil),  // 4: tarianpb.api.RemoveConstraintRequest
	(*RemoveConstraintResponse)(nil), // 5: tarianpb.api.RemoveConstraintResponse
	(*AddActionRequest)(nil),         // 6: tarianpb.api.AddActionRequest
	(*AddActionResponse)(nil),        // 7: tarianpb.api.AddActionResponse
	(*GetActionsRequest)(nil),        // 8: tarianpb.api.GetActionsRequest
	(*GetActionsResponse)(nil),       // 9: tarianpb.api.GetActionsResponse
	(*IngestEventRequest)(nil),       // 10: tarianpb.api.IngestEventRequest
	(*IngestEventResponse)(nil),      // 11: tarianpb.api.IngestEventResponse
	(*GetEventsRequest)(nil),         // 12: tarianpb.api.GetEventsRequest
	(*GetEventsResponse)(nil),        // 13: tarianpb.api.GetEventsResponse
	(*Label)(nil),                    // 14: tarianpb.types.Label
	(*Constraint)(nil),               // 15: tarianpb.types.Constraint
	(*Action)(nil),                   // 16: tarianpb.types.Action
	(*Event)(nil),                    // 17: tarianpb.types.Event
}
var file_tarianpb_api_proto_depIdxs = []int32{
	14, // 0: tarianpb.api.GetConstraintsRequest.labels:type_name -> tarianpb.types.Label
	15, // 1: tarianpb.api.GetConstraintsResponse.constraints:type_name -> tarianpb.types.Constraint
	15, // 2: tarianpb.api.AddConstraintRequest.constraint:type_name -> tarianpb.types.Constraint
	16, // 3: tarianpb.api.AddActionRequest.action:type_name -> tarianpb.types.Action
	14, // 4: tarianpb.api.GetActionsRequest.labels:type_name -> tarianpb.types.Label
	16, // 5: tarianpb.api.GetActionsResponse.actions:type_name -> tarianpb.types.Action
	17, // 6: tarianpb.api.IngestEventRequest.event:type_name -> tarianpb.types.Event
	17, // 7: tarianpb.api.GetEventsResponse.events:type_name -> tarianpb.types.Event
	0,  // 8: tarianpb.api.Config.GetConstraints:input_type -> tarianpb.api.GetConstraintsRequest
	2,  // 9: tarianpb.api.Config.AddConstraint:input_type -> tarianpb.api.AddConstraintRequest
	4,  // 10: tarianpb.api.Config.RemoveConstraint:input_type -> tarianpb.api.RemoveConstraintRequest
	6,  // 11: tarianpb.api.Config.AddAction:input_type -> tarianpb.api.AddActionRequest
	8,  // 12: tarianpb.api.Config.GetActions:input_type -> tarianpb.api.GetActionsRequest
	10, // 13: tarianpb.api.Event.IngestEvent:input_type -> tarianpb.api.IngestEventRequest
	12, // 14: tarianpb.api.Event.GetEvents:input_type -> tarianpb.api.GetEventsRequest
	1,  // 15: tarianpb.api.Config.GetConstraints:output_type -> tarianpb.api.GetConstraintsResponse
	3,  // 16: tarianpb.api.Config.AddConstraint:output_type -> tarianpb.api.AddConstraintResponse
	5,  // 17: tarianpb.api.Config.RemoveConstraint:output_type -> tarianpb.api.RemoveConstraintResponse
	7,  // 18: tarianpb.api.Config.AddAction:output_type -> tarianpb.api.AddActionResponse
	9,  // 19: tarianpb.api.Config.GetActions:output_type -> tarianpb.api.GetActionsResponse
	11, // 20: tarianpb.api.Event.IngestEvent:output_type -> tarianpb.api.IngestEventResponse
	13, // 21: tarianpb.api.Event.GetEvents:output_type -> tarianpb.api.GetEventsResponse
	15, // [15:22] is the sub-list for method output_type
	8,  // [8:15] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_tarianpb_api_proto_init() }
func file_tarianpb_api_proto_init() {
	if File_tarianpb_api_proto != nil {
		return
	}
	file_tarianpb_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tarianpb_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConstraintsRequest); i {
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
		file_tarianpb_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConstraintsResponse); i {
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
		file_tarianpb_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddConstraintRequest); i {
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
		file_tarianpb_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddConstraintResponse); i {
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
		file_tarianpb_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveConstraintRequest); i {
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
		file_tarianpb_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveConstraintResponse); i {
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
		file_tarianpb_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddActionRequest); i {
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
		file_tarianpb_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddActionResponse); i {
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
		file_tarianpb_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetActionsRequest); i {
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
		file_tarianpb_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetActionsResponse); i {
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
		file_tarianpb_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngestEventRequest); i {
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
		file_tarianpb_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngestEventResponse); i {
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
		file_tarianpb_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsRequest); i {
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
		file_tarianpb_api_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsResponse); i {
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
			RawDescriptor: file_tarianpb_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_tarianpb_api_proto_goTypes,
		DependencyIndexes: file_tarianpb_api_proto_depIdxs,
		MessageInfos:      file_tarianpb_api_proto_msgTypes,
	}.Build()
	File_tarianpb_api_proto = out.File
	file_tarianpb_api_proto_rawDesc = nil
	file_tarianpb_api_proto_goTypes = nil
	file_tarianpb_api_proto_depIdxs = nil
}
