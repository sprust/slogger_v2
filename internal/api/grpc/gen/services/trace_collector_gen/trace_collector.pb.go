// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: internal/api/grpc/proto/trace_collector.proto

package trace_collector_gen

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type TraceCollectorResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StatusCode    int32                  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TraceCollectorResponse) Reset() {
	*x = TraceCollectorResponse{}
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TraceCollectorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceCollectorResponse) ProtoMessage() {}

func (x *TraceCollectorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceCollectorResponse.ProtoReflect.Descriptor instead.
func (*TraceCollectorResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_grpc_proto_trace_collector_proto_rawDescGZIP(), []int{0}
}

func (x *TraceCollectorResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *TraceCollectorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TraceCreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Traces        []*TraceCreateObject   `protobuf:"bytes,1,rep,name=traces,proto3" json:"traces,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TraceCreateRequest) Reset() {
	*x = TraceCreateRequest{}
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TraceCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceCreateRequest) ProtoMessage() {}

func (x *TraceCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceCreateRequest.ProtoReflect.Descriptor instead.
func (*TraceCreateRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_grpc_proto_trace_collector_proto_rawDescGZIP(), []int{1}
}

func (x *TraceCreateRequest) GetTraces() []*TraceCreateObject {
	if x != nil {
		return x.Traces
	}
	return nil
}

type TraceCreateObject struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TraceId       string                 `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	ParentTraceId *wrappers.StringValue  `protobuf:"bytes,2,opt,name=parent_trace_id,json=parentTraceId,proto3" json:"parent_trace_id,omitempty"`
	Type          string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	// sometimes
	Tags []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	// json
	Data string `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	// sometimes
	Duration *wrappers.DoubleValue `protobuf:"bytes,7,opt,name=duration,proto3" json:"duration,omitempty"`
	// sometimes
	Memory *wrappers.DoubleValue `protobuf:"bytes,8,opt,name=memory,proto3" json:"memory,omitempty"`
	// sometimes
	Cpu           *wrappers.DoubleValue `protobuf:"bytes,9,opt,name=cpu,proto3" json:"cpu,omitempty"`
	LoggedAt      *timestamp.Timestamp  `protobuf:"bytes,10,opt,name=logged_at,json=loggedAt,proto3" json:"logged_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TraceCreateObject) Reset() {
	*x = TraceCreateObject{}
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TraceCreateObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceCreateObject) ProtoMessage() {}

func (x *TraceCreateObject) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceCreateObject.ProtoReflect.Descriptor instead.
func (*TraceCreateObject) Descriptor() ([]byte, []int) {
	return file_internal_api_grpc_proto_trace_collector_proto_rawDescGZIP(), []int{2}
}

func (x *TraceCreateObject) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *TraceCreateObject) GetParentTraceId() *wrappers.StringValue {
	if x != nil {
		return x.ParentTraceId
	}
	return nil
}

func (x *TraceCreateObject) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TraceCreateObject) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TraceCreateObject) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *TraceCreateObject) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *TraceCreateObject) GetDuration() *wrappers.DoubleValue {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *TraceCreateObject) GetMemory() *wrappers.DoubleValue {
	if x != nil {
		return x.Memory
	}
	return nil
}

func (x *TraceCreateObject) GetCpu() *wrappers.DoubleValue {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *TraceCreateObject) GetLoggedAt() *timestamp.Timestamp {
	if x != nil {
		return x.LoggedAt
	}
	return nil
}

type TraceUpdateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Traces        []*TraceUpdateObject   `protobuf:"bytes,1,rep,name=traces,proto3" json:"traces,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TraceUpdateRequest) Reset() {
	*x = TraceUpdateRequest{}
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TraceUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceUpdateRequest) ProtoMessage() {}

func (x *TraceUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceUpdateRequest.ProtoReflect.Descriptor instead.
func (*TraceUpdateRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_grpc_proto_trace_collector_proto_rawDescGZIP(), []int{3}
}

func (x *TraceUpdateRequest) GetTraces() []*TraceUpdateObject {
	if x != nil {
		return x.Traces
	}
	return nil
}

type TraceUpdateObject struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	TraceId string                 `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	Status  string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// sometimes
	Profiling *TraceProfilingItemsObject `protobuf:"bytes,3,opt,name=profiling,proto3" json:"profiling,omitempty"`
	// sometimes
	Tags *TagsObject `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	// sometimes
	Data *wrappers.StringValue `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	// sometimes
	Duration *wrappers.DoubleValue `protobuf:"bytes,6,opt,name=duration,proto3" json:"duration,omitempty"`
	// sometimes
	Memory *wrappers.DoubleValue `protobuf:"bytes,7,opt,name=memory,proto3" json:"memory,omitempty"`
	// sometimes
	Cpu           *wrappers.DoubleValue `protobuf:"bytes,8,opt,name=cpu,proto3" json:"cpu,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TraceUpdateObject) Reset() {
	*x = TraceUpdateObject{}
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TraceUpdateObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceUpdateObject) ProtoMessage() {}

func (x *TraceUpdateObject) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceUpdateObject.ProtoReflect.Descriptor instead.
func (*TraceUpdateObject) Descriptor() ([]byte, []int) {
	return file_internal_api_grpc_proto_trace_collector_proto_rawDescGZIP(), []int{4}
}

func (x *TraceUpdateObject) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *TraceUpdateObject) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TraceUpdateObject) GetProfiling() *TraceProfilingItemsObject {
	if x != nil {
		return x.Profiling
	}
	return nil
}

func (x *TraceUpdateObject) GetTags() *TagsObject {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *TraceUpdateObject) GetData() *wrappers.StringValue {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TraceUpdateObject) GetDuration() *wrappers.DoubleValue {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *TraceUpdateObject) GetMemory() *wrappers.DoubleValue {
	if x != nil {
		return x.Memory
	}
	return nil
}

func (x *TraceUpdateObject) GetCpu() *wrappers.DoubleValue {
	if x != nil {
		return x.Cpu
	}
	return nil
}

type TagsObject struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []string               `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TagsObject) Reset() {
	*x = TagsObject{}
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TagsObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagsObject) ProtoMessage() {}

func (x *TagsObject) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagsObject.ProtoReflect.Descriptor instead.
func (*TagsObject) Descriptor() ([]byte, []int) {
	return file_internal_api_grpc_proto_trace_collector_proto_rawDescGZIP(), []int{5}
}

func (x *TagsObject) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

type TraceProfilingItemsObject struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	MainCaller    string                      `protobuf:"bytes,1,opt,name=main_caller,json=mainCaller,proto3" json:"main_caller,omitempty"`
	Items         []*TraceProfilingItemObject `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TraceProfilingItemsObject) Reset() {
	*x = TraceProfilingItemsObject{}
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TraceProfilingItemsObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceProfilingItemsObject) ProtoMessage() {}

func (x *TraceProfilingItemsObject) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceProfilingItemsObject.ProtoReflect.Descriptor instead.
func (*TraceProfilingItemsObject) Descriptor() ([]byte, []int) {
	return file_internal_api_grpc_proto_trace_collector_proto_rawDescGZIP(), []int{6}
}

func (x *TraceProfilingItemsObject) GetMainCaller() string {
	if x != nil {
		return x.MainCaller
	}
	return ""
}

func (x *TraceProfilingItemsObject) GetItems() []*TraceProfilingItemObject {
	if x != nil {
		return x.Items
	}
	return nil
}

type TraceProfilingItemObject struct {
	state         protoimpl.MessageState              `protogen:"open.v1"`
	Raw           string                              `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
	Calling       string                              `protobuf:"bytes,2,opt,name=calling,proto3" json:"calling,omitempty"`
	Callable      string                              `protobuf:"bytes,3,opt,name=callable,proto3" json:"callable,omitempty"`
	Data          []*TraceProfilingItemDataItemObject `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TraceProfilingItemObject) Reset() {
	*x = TraceProfilingItemObject{}
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TraceProfilingItemObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceProfilingItemObject) ProtoMessage() {}

func (x *TraceProfilingItemObject) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceProfilingItemObject.ProtoReflect.Descriptor instead.
func (*TraceProfilingItemObject) Descriptor() ([]byte, []int) {
	return file_internal_api_grpc_proto_trace_collector_proto_rawDescGZIP(), []int{7}
}

func (x *TraceProfilingItemObject) GetRaw() string {
	if x != nil {
		return x.Raw
	}
	return ""
}

func (x *TraceProfilingItemObject) GetCalling() string {
	if x != nil {
		return x.Calling
	}
	return ""
}

func (x *TraceProfilingItemObject) GetCallable() string {
	if x != nil {
		return x.Callable
	}
	return ""
}

func (x *TraceProfilingItemObject) GetData() []*TraceProfilingItemDataItemObject {
	if x != nil {
		return x.Data
	}
	return nil
}

type TraceProfilingItemDataItemObject struct {
	state         protoimpl.MessageState                 `protogen:"open.v1"`
	Name          string                                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value         *TraceProfilingItemDataItemValueObject `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TraceProfilingItemDataItemObject) Reset() {
	*x = TraceProfilingItemDataItemObject{}
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TraceProfilingItemDataItemObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceProfilingItemDataItemObject) ProtoMessage() {}

func (x *TraceProfilingItemDataItemObject) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceProfilingItemDataItemObject.ProtoReflect.Descriptor instead.
func (*TraceProfilingItemDataItemObject) Descriptor() ([]byte, []int) {
	return file_internal_api_grpc_proto_trace_collector_proto_rawDescGZIP(), []int{8}
}

func (x *TraceProfilingItemDataItemObject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TraceProfilingItemDataItemObject) GetValue() *TraceProfilingItemDataItemValueObject {
	if x != nil {
		return x.Value
	}
	return nil
}

type TraceProfilingItemDataItemValueObject struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Double        *wrappers.DoubleValue  `protobuf:"bytes,1,opt,name=double,proto3" json:"double,omitempty"`
	Int           *wrappers.Int32Value   `protobuf:"bytes,2,opt,name=int,proto3" json:"int,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TraceProfilingItemDataItemValueObject) Reset() {
	*x = TraceProfilingItemDataItemValueObject{}
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TraceProfilingItemDataItemValueObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceProfilingItemDataItemValueObject) ProtoMessage() {}

func (x *TraceProfilingItemDataItemValueObject) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_grpc_proto_trace_collector_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceProfilingItemDataItemValueObject.ProtoReflect.Descriptor instead.
func (*TraceProfilingItemDataItemValueObject) Descriptor() ([]byte, []int) {
	return file_internal_api_grpc_proto_trace_collector_proto_rawDescGZIP(), []int{9}
}

func (x *TraceProfilingItemDataItemValueObject) GetDouble() *wrappers.DoubleValue {
	if x != nil {
		return x.Double
	}
	return nil
}

func (x *TraceProfilingItemDataItemValueObject) GetInt() *wrappers.Int32Value {
	if x != nil {
		return x.Int
	}
	return nil
}

var File_internal_api_grpc_proto_trace_collector_proto protoreflect.FileDescriptor

var file_internal_api_grpc_proto_trace_collector_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x53, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x50, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x63, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x06,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x06, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x22, 0xa1, 0x03, 0x0a, 0x11, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x0f, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a,
	0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03,
	0x63, 0x70, 0x75, 0x12, 0x37, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x41, 0x74, 0x22, 0x50, 0x0a, 0x12,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x22, 0x93,
	0x03, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x48, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e,
	0x67, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34,
	0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x03, 0x63, 0x70, 0x75, 0x22, 0x22, 0x0a, 0x0a, 0x54, 0x61, 0x67, 0x73, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x7d, 0x0a, 0x19, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x69, 0x6e,
	0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x18, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x45, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x44,
	0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x84, 0x01, 0x0a, 0x20, 0x54, 0x72, 0x61, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x49, 0x74,
	0x65, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d,
	0x44, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x25, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65,
	0x6d, 0x44, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x03, 0x69, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x69, 0x6e, 0x74, 0x32, 0xc4, 0x01, 0x0a, 0x0e, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x58, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x23, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x1e, 0x5a, 0x1c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_api_grpc_proto_trace_collector_proto_rawDescOnce sync.Once
	file_internal_api_grpc_proto_trace_collector_proto_rawDescData = file_internal_api_grpc_proto_trace_collector_proto_rawDesc
)

func file_internal_api_grpc_proto_trace_collector_proto_rawDescGZIP() []byte {
	file_internal_api_grpc_proto_trace_collector_proto_rawDescOnce.Do(func() {
		file_internal_api_grpc_proto_trace_collector_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_api_grpc_proto_trace_collector_proto_rawDescData)
	})
	return file_internal_api_grpc_proto_trace_collector_proto_rawDescData
}

var file_internal_api_grpc_proto_trace_collector_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_internal_api_grpc_proto_trace_collector_proto_goTypes = []any{
	(*TraceCollectorResponse)(nil),                // 0: trace.collector.TraceCollectorResponse
	(*TraceCreateRequest)(nil),                    // 1: trace.collector.TraceCreateRequest
	(*TraceCreateObject)(nil),                     // 2: trace.collector.TraceCreateObject
	(*TraceUpdateRequest)(nil),                    // 3: trace.collector.TraceUpdateRequest
	(*TraceUpdateObject)(nil),                     // 4: trace.collector.TraceUpdateObject
	(*TagsObject)(nil),                            // 5: trace.collector.TagsObject
	(*TraceProfilingItemsObject)(nil),             // 6: trace.collector.TraceProfilingItemsObject
	(*TraceProfilingItemObject)(nil),              // 7: trace.collector.TraceProfilingItemObject
	(*TraceProfilingItemDataItemObject)(nil),      // 8: trace.collector.TraceProfilingItemDataItemObject
	(*TraceProfilingItemDataItemValueObject)(nil), // 9: trace.collector.TraceProfilingItemDataItemValueObject
	(*wrappers.StringValue)(nil),                  // 10: google.protobuf.StringValue
	(*wrappers.DoubleValue)(nil),                  // 11: google.protobuf.DoubleValue
	(*timestamp.Timestamp)(nil),                   // 12: google.protobuf.Timestamp
	(*wrappers.Int32Value)(nil),                   // 13: google.protobuf.Int32Value
}
var file_internal_api_grpc_proto_trace_collector_proto_depIdxs = []int32{
	2,  // 0: trace.collector.TraceCreateRequest.traces:type_name -> trace.collector.TraceCreateObject
	10, // 1: trace.collector.TraceCreateObject.parent_trace_id:type_name -> google.protobuf.StringValue
	11, // 2: trace.collector.TraceCreateObject.duration:type_name -> google.protobuf.DoubleValue
	11, // 3: trace.collector.TraceCreateObject.memory:type_name -> google.protobuf.DoubleValue
	11, // 4: trace.collector.TraceCreateObject.cpu:type_name -> google.protobuf.DoubleValue
	12, // 5: trace.collector.TraceCreateObject.logged_at:type_name -> google.protobuf.Timestamp
	4,  // 6: trace.collector.TraceUpdateRequest.traces:type_name -> trace.collector.TraceUpdateObject
	6,  // 7: trace.collector.TraceUpdateObject.profiling:type_name -> trace.collector.TraceProfilingItemsObject
	5,  // 8: trace.collector.TraceUpdateObject.tags:type_name -> trace.collector.TagsObject
	10, // 9: trace.collector.TraceUpdateObject.data:type_name -> google.protobuf.StringValue
	11, // 10: trace.collector.TraceUpdateObject.duration:type_name -> google.protobuf.DoubleValue
	11, // 11: trace.collector.TraceUpdateObject.memory:type_name -> google.protobuf.DoubleValue
	11, // 12: trace.collector.TraceUpdateObject.cpu:type_name -> google.protobuf.DoubleValue
	7,  // 13: trace.collector.TraceProfilingItemsObject.items:type_name -> trace.collector.TraceProfilingItemObject
	8,  // 14: trace.collector.TraceProfilingItemObject.data:type_name -> trace.collector.TraceProfilingItemDataItemObject
	9,  // 15: trace.collector.TraceProfilingItemDataItemObject.value:type_name -> trace.collector.TraceProfilingItemDataItemValueObject
	11, // 16: trace.collector.TraceProfilingItemDataItemValueObject.double:type_name -> google.protobuf.DoubleValue
	13, // 17: trace.collector.TraceProfilingItemDataItemValueObject.int:type_name -> google.protobuf.Int32Value
	1,  // 18: trace.collector.TraceCollector.Create:input_type -> trace.collector.TraceCreateRequest
	3,  // 19: trace.collector.TraceCollector.Update:input_type -> trace.collector.TraceUpdateRequest
	0,  // 20: trace.collector.TraceCollector.Create:output_type -> trace.collector.TraceCollectorResponse
	0,  // 21: trace.collector.TraceCollector.Update:output_type -> trace.collector.TraceCollectorResponse
	20, // [20:22] is the sub-list for method output_type
	18, // [18:20] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_internal_api_grpc_proto_trace_collector_proto_init() }
func file_internal_api_grpc_proto_trace_collector_proto_init() {
	if File_internal_api_grpc_proto_trace_collector_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_api_grpc_proto_trace_collector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_api_grpc_proto_trace_collector_proto_goTypes,
		DependencyIndexes: file_internal_api_grpc_proto_trace_collector_proto_depIdxs,
		MessageInfos:      file_internal_api_grpc_proto_trace_collector_proto_msgTypes,
	}.Build()
	File_internal_api_grpc_proto_trace_collector_proto = out.File
	file_internal_api_grpc_proto_trace_collector_proto_rawDesc = nil
	file_internal_api_grpc_proto_trace_collector_proto_goTypes = nil
	file_internal_api_grpc_proto_trace_collector_proto_depIdxs = nil
}