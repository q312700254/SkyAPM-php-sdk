// Code generated by protoc-gen-go. DO NOT EDIT.
// source: language-agent/TraceSegmentService.proto

package agent

import (
	common "agent/agent/pb6/common"
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TraceSegmentObject struct {
	TraceSegmentId        *common.UniqueId `protobuf:"bytes,1,opt,name=traceSegmentId,proto3" json:"traceSegmentId,omitempty"`
	Spans                 []*SpanObject    `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
	ApplicationId         int32            `protobuf:"varint,3,opt,name=applicationId,proto3" json:"applicationId,omitempty"`
	ApplicationInstanceId int32            `protobuf:"varint,4,opt,name=applicationInstanceId,proto3" json:"applicationInstanceId,omitempty"`
	IsSizeLimited         bool             `protobuf:"varint,5,opt,name=isSizeLimited,proto3" json:"isSizeLimited,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}         `json:"-"`
	XXX_unrecognized      []byte           `json:"-"`
	XXX_sizecache         int32            `json:"-"`
}

func (m *TraceSegmentObject) Reset()         { *m = TraceSegmentObject{} }
func (m *TraceSegmentObject) String() string { return proto.CompactTextString(m) }
func (*TraceSegmentObject) ProtoMessage()    {}
func (*TraceSegmentObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4a2a0c689c6ad4b, []int{0}
}

func (m *TraceSegmentObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceSegmentObject.Unmarshal(m, b)
}
func (m *TraceSegmentObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceSegmentObject.Marshal(b, m, deterministic)
}
func (m *TraceSegmentObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceSegmentObject.Merge(m, src)
}
func (m *TraceSegmentObject) XXX_Size() int {
	return xxx_messageInfo_TraceSegmentObject.Size(m)
}
func (m *TraceSegmentObject) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceSegmentObject.DiscardUnknown(m)
}

var xxx_messageInfo_TraceSegmentObject proto.InternalMessageInfo

func (m *TraceSegmentObject) GetTraceSegmentId() *common.UniqueId {
	if m != nil {
		return m.TraceSegmentId
	}
	return nil
}

func (m *TraceSegmentObject) GetSpans() []*SpanObject {
	if m != nil {
		return m.Spans
	}
	return nil
}

func (m *TraceSegmentObject) GetApplicationId() int32 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *TraceSegmentObject) GetApplicationInstanceId() int32 {
	if m != nil {
		return m.ApplicationInstanceId
	}
	return 0
}

func (m *TraceSegmentObject) GetIsSizeLimited() bool {
	if m != nil {
		return m.IsSizeLimited
	}
	return false
}

type TraceSegmentReference struct {
	RefType                     common.RefType   `protobuf:"varint,1,opt,name=refType,proto3,enum=RefType" json:"refType,omitempty"`
	ParentTraceSegmentId        *common.UniqueId `protobuf:"bytes,2,opt,name=parentTraceSegmentId,proto3" json:"parentTraceSegmentId,omitempty"`
	ParentSpanId                int32            `protobuf:"varint,3,opt,name=parentSpanId,proto3" json:"parentSpanId,omitempty"`
	ParentApplicationInstanceId int32            `protobuf:"varint,4,opt,name=parentApplicationInstanceId,proto3" json:"parentApplicationInstanceId,omitempty"`
	NetworkAddress              string           `protobuf:"bytes,5,opt,name=networkAddress,proto3" json:"networkAddress,omitempty"`
	NetworkAddressId            int32            `protobuf:"varint,6,opt,name=networkAddressId,proto3" json:"networkAddressId,omitempty"`
	EntryApplicationInstanceId  int32            `protobuf:"varint,7,opt,name=entryApplicationInstanceId,proto3" json:"entryApplicationInstanceId,omitempty"`
	EntryServiceName            string           `protobuf:"bytes,8,opt,name=entryServiceName,proto3" json:"entryServiceName,omitempty"`
	EntryServiceId              int32            `protobuf:"varint,9,opt,name=entryServiceId,proto3" json:"entryServiceId,omitempty"`
	ParentServiceName           string           `protobuf:"bytes,10,opt,name=parentServiceName,proto3" json:"parentServiceName,omitempty"`
	ParentServiceId             int32            `protobuf:"varint,11,opt,name=parentServiceId,proto3" json:"parentServiceId,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}         `json:"-"`
	XXX_unrecognized            []byte           `json:"-"`
	XXX_sizecache               int32            `json:"-"`
}

func (m *TraceSegmentReference) Reset()         { *m = TraceSegmentReference{} }
func (m *TraceSegmentReference) String() string { return proto.CompactTextString(m) }
func (*TraceSegmentReference) ProtoMessage()    {}
func (*TraceSegmentReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4a2a0c689c6ad4b, []int{1}
}

func (m *TraceSegmentReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceSegmentReference.Unmarshal(m, b)
}
func (m *TraceSegmentReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceSegmentReference.Marshal(b, m, deterministic)
}
func (m *TraceSegmentReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceSegmentReference.Merge(m, src)
}
func (m *TraceSegmentReference) XXX_Size() int {
	return xxx_messageInfo_TraceSegmentReference.Size(m)
}
func (m *TraceSegmentReference) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceSegmentReference.DiscardUnknown(m)
}

var xxx_messageInfo_TraceSegmentReference proto.InternalMessageInfo

func (m *TraceSegmentReference) GetRefType() common.RefType {
	if m != nil {
		return m.RefType
	}
	return common.RefType_CrossProcess
}

func (m *TraceSegmentReference) GetParentTraceSegmentId() *common.UniqueId {
	if m != nil {
		return m.ParentTraceSegmentId
	}
	return nil
}

func (m *TraceSegmentReference) GetParentSpanId() int32 {
	if m != nil {
		return m.ParentSpanId
	}
	return 0
}

func (m *TraceSegmentReference) GetParentApplicationInstanceId() int32 {
	if m != nil {
		return m.ParentApplicationInstanceId
	}
	return 0
}

func (m *TraceSegmentReference) GetNetworkAddress() string {
	if m != nil {
		return m.NetworkAddress
	}
	return ""
}

func (m *TraceSegmentReference) GetNetworkAddressId() int32 {
	if m != nil {
		return m.NetworkAddressId
	}
	return 0
}

func (m *TraceSegmentReference) GetEntryApplicationInstanceId() int32 {
	if m != nil {
		return m.EntryApplicationInstanceId
	}
	return 0
}

func (m *TraceSegmentReference) GetEntryServiceName() string {
	if m != nil {
		return m.EntryServiceName
	}
	return ""
}

func (m *TraceSegmentReference) GetEntryServiceId() int32 {
	if m != nil {
		return m.EntryServiceId
	}
	return 0
}

func (m *TraceSegmentReference) GetParentServiceName() string {
	if m != nil {
		return m.ParentServiceName
	}
	return ""
}

func (m *TraceSegmentReference) GetParentServiceId() int32 {
	if m != nil {
		return m.ParentServiceId
	}
	return 0
}

type SpanObject struct {
	SpanId               int32                    `protobuf:"varint,1,opt,name=spanId,proto3" json:"spanId,omitempty"`
	ParentSpanId         int32                    `protobuf:"varint,2,opt,name=parentSpanId,proto3" json:"parentSpanId,omitempty"`
	StartTime            int64                    `protobuf:"varint,3,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              int64                    `protobuf:"varint,4,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Refs                 []*TraceSegmentReference `protobuf:"bytes,5,rep,name=refs,proto3" json:"refs,omitempty"`
	OperationNameId      int32                    `protobuf:"varint,6,opt,name=operationNameId,proto3" json:"operationNameId,omitempty"`
	OperationName        string                   `protobuf:"bytes,7,opt,name=operationName,proto3" json:"operationName,omitempty"`
	PeerId               int32                    `protobuf:"varint,8,opt,name=peerId,proto3" json:"peerId,omitempty"`
	Peer                 string                   `protobuf:"bytes,9,opt,name=peer,proto3" json:"peer,omitempty"`
	SpanType             common.SpanType          `protobuf:"varint,10,opt,name=spanType,proto3,enum=SpanType" json:"spanType,omitempty"`
	SpanLayer            common.SpanLayer         `protobuf:"varint,11,opt,name=spanLayer,proto3,enum=SpanLayer" json:"spanLayer,omitempty"`
	ComponentId          int32                    `protobuf:"varint,12,opt,name=componentId,proto3" json:"componentId,omitempty"`
	Component            string                   `protobuf:"bytes,13,opt,name=component,proto3" json:"component,omitempty"`
	IsError              bool                     `protobuf:"varint,14,opt,name=isError,proto3" json:"isError,omitempty"`
	Tags                 []*KeyWithStringValue    `protobuf:"bytes,15,rep,name=tags,proto3" json:"tags,omitempty"`
	Logs                 []*LogMessage            `protobuf:"bytes,16,rep,name=logs,proto3" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SpanObject) Reset()         { *m = SpanObject{} }
func (m *SpanObject) String() string { return proto.CompactTextString(m) }
func (*SpanObject) ProtoMessage()    {}
func (*SpanObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4a2a0c689c6ad4b, []int{2}
}

func (m *SpanObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanObject.Unmarshal(m, b)
}
func (m *SpanObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanObject.Marshal(b, m, deterministic)
}
func (m *SpanObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanObject.Merge(m, src)
}
func (m *SpanObject) XXX_Size() int {
	return xxx_messageInfo_SpanObject.Size(m)
}
func (m *SpanObject) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanObject.DiscardUnknown(m)
}

var xxx_messageInfo_SpanObject proto.InternalMessageInfo

func (m *SpanObject) GetSpanId() int32 {
	if m != nil {
		return m.SpanId
	}
	return 0
}

func (m *SpanObject) GetParentSpanId() int32 {
	if m != nil {
		return m.ParentSpanId
	}
	return 0
}

func (m *SpanObject) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *SpanObject) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *SpanObject) GetRefs() []*TraceSegmentReference {
	if m != nil {
		return m.Refs
	}
	return nil
}

func (m *SpanObject) GetOperationNameId() int32 {
	if m != nil {
		return m.OperationNameId
	}
	return 0
}

func (m *SpanObject) GetOperationName() string {
	if m != nil {
		return m.OperationName
	}
	return ""
}

func (m *SpanObject) GetPeerId() int32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *SpanObject) GetPeer() string {
	if m != nil {
		return m.Peer
	}
	return ""
}

func (m *SpanObject) GetSpanType() common.SpanType {
	if m != nil {
		return m.SpanType
	}
	return common.SpanType_Entry
}

func (m *SpanObject) GetSpanLayer() common.SpanLayer {
	if m != nil {
		return m.SpanLayer
	}
	return common.SpanLayer_Unknown
}

func (m *SpanObject) GetComponentId() int32 {
	if m != nil {
		return m.ComponentId
	}
	return 0
}

func (m *SpanObject) GetComponent() string {
	if m != nil {
		return m.Component
	}
	return ""
}

func (m *SpanObject) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

func (m *SpanObject) GetTags() []*KeyWithStringValue {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *SpanObject) GetLogs() []*LogMessage {
	if m != nil {
		return m.Logs
	}
	return nil
}

type LogMessage struct {
	Time                 int64                 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Data                 []*KeyWithStringValue `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LogMessage) Reset()         { *m = LogMessage{} }
func (m *LogMessage) String() string { return proto.CompactTextString(m) }
func (*LogMessage) ProtoMessage()    {}
func (*LogMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4a2a0c689c6ad4b, []int{3}
}

func (m *LogMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogMessage.Unmarshal(m, b)
}
func (m *LogMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogMessage.Marshal(b, m, deterministic)
}
func (m *LogMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogMessage.Merge(m, src)
}
func (m *LogMessage) XXX_Size() int {
	return xxx_messageInfo_LogMessage.Size(m)
}
func (m *LogMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LogMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LogMessage proto.InternalMessageInfo

func (m *LogMessage) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *LogMessage) GetData() []*KeyWithStringValue {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*TraceSegmentObject)(nil), "TraceSegmentObject")
	proto.RegisterType((*TraceSegmentReference)(nil), "TraceSegmentReference")
	proto.RegisterType((*SpanObject)(nil), "SpanObject")
	proto.RegisterType((*LogMessage)(nil), "LogMessage")
}

func init() {
	proto.RegisterFile("language-agent/TraceSegmentService.proto", fileDescriptor_b4a2a0c689c6ad4b)
}

var fileDescriptor_b4a2a0c689c6ad4b = []byte{
	// 765 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0x5f, 0x8f, 0xdb, 0x44,
	0x10, 0xc7, 0x39, 0xe7, 0x12, 0x4f, 0xda, 0xf4, 0xd8, 0x72, 0x95, 0x09, 0x48, 0x0d, 0x11, 0x50,
	0xab, 0x6a, 0x5d, 0x08, 0x88, 0x37, 0x10, 0xe5, 0xcf, 0x83, 0xc5, 0x51, 0xaa, 0x4d, 0x4a, 0x25,
	0xde, 0xf6, 0xec, 0x39, 0x9f, 0x89, 0xbd, 0x6b, 0xd6, 0x7b, 0x9c, 0xc2, 0x77, 0x41, 0xe2, 0x99,
	0xcf, 0xc6, 0x87, 0x40, 0x3b, 0x76, 0xfe, 0xd8, 0x89, 0xee, 0xc5, 0xda, 0xf9, 0xcd, 0xcf, 0xb3,
	0x3b, 0xbf, 0xd9, 0xd9, 0x81, 0x20, 0x17, 0x32, 0xbd, 0x11, 0x29, 0x3e, 0x17, 0x29, 0x4a, 0xf3,
	0x62, 0xa9, 0x45, 0x8c, 0x0b, 0x4c, 0x0b, 0x94, 0x66, 0x81, 0xfa, 0xcf, 0x2c, 0xc6, 0xb0, 0xd4,
	0xca, 0xa8, 0xc9, 0xe3, 0x0e, 0xf3, 0x07, 0x75, 0x2b, 0x2b, 0xa3, 0x51, 0x14, 0x0d, 0xe1, 0x49,
	0x87, 0xf0, 0x13, 0xae, 0xdf, 0x66, 0xe6, 0x7a, 0x61, 0x74, 0x26, 0xd3, 0x5f, 0x45, 0x7e, 0xb3,
	0x89, 0xf4, 0x7e, 0xac, 0x8a, 0x42, 0xc9, 0x17, 0xc6, 0xee, 0xf5, 0xbc, 0x36, 0x6a, 0xd7, 0xec,
	0x3f, 0x07, 0xd8, 0xfe, 0x11, 0x7e, 0xb9, 0xfc, 0x1d, 0x63, 0xc3, 0x3e, 0x87, 0xb1, 0xd9, 0x43,
	0xa3, 0xc4, 0x77, 0xa6, 0x4e, 0x30, 0x9a, 0x7b, 0xe1, 0x1b, 0x99, 0xfd, 0x71, 0x83, 0x51, 0xc2,
	0x3b, 0x04, 0xf6, 0x11, 0xf4, 0xab, 0x52, 0xc8, 0xca, 0xef, 0x4d, 0x4f, 0x82, 0xd1, 0x7c, 0x14,
	0x2e, 0x4a, 0x21, 0xeb, 0x70, 0xbc, 0xf6, 0xb0, 0x8f, 0xe1, 0xbe, 0x28, 0xcb, 0x3c, 0x8b, 0x85,
	0xc9, 0x94, 0x8c, 0x12, 0xff, 0x64, 0xea, 0x04, 0x7d, 0xde, 0x06, 0xd9, 0x97, 0x70, 0xbe, 0x0f,
	0xc8, 0xca, 0x08, 0x19, 0x63, 0x94, 0xf8, 0x2e, 0xb1, 0x8f, 0x3b, 0x6d, 0xec, 0xac, 0x5a, 0x64,
	0x7f, 0xe1, 0x45, 0x56, 0x64, 0x06, 0x13, 0xbf, 0x3f, 0x75, 0x82, 0x21, 0x6f, 0x83, 0xb3, 0x7f,
	0x5c, 0x38, 0xdf, 0x4f, 0x97, 0xe3, 0x15, 0x6a, 0x94, 0x31, 0xb2, 0x19, 0x0c, 0x34, 0x5e, 0x2d,
	0xd7, 0x25, 0x52, 0xaa, 0xe3, 0xf9, 0x30, 0xe4, 0xb5, 0xcd, 0x37, 0x0e, 0xf6, 0x35, 0xbc, 0x57,
	0x0a, 0x8d, 0xd2, 0x2c, 0xdb, 0xda, 0xf4, 0xba, 0xda, 0x1c, 0xa5, 0xb1, 0x19, 0xdc, 0xab, 0x71,
	0xab, 0xcc, 0x36, 0xfb, 0x16, 0xc6, 0xbe, 0x85, 0x0f, 0x6a, 0xfb, 0xe5, 0x1d, 0x12, 0xdc, 0x45,
	0x61, 0x9f, 0xc2, 0x58, 0xa2, 0xb9, 0x55, 0x7a, 0xf5, 0x32, 0x49, 0x34, 0x56, 0x15, 0x29, 0xe1,
	0xf1, 0x0e, 0xca, 0x9e, 0xc2, 0x59, 0x1b, 0x89, 0x12, 0xff, 0x94, 0xc2, 0x1f, 0xe0, 0xec, 0x1b,
	0x98, 0xa0, 0x34, 0x7a, 0x7d, 0xfc, 0x50, 0x03, 0xfa, 0xeb, 0x0e, 0x86, 0xdd, 0x8b, 0xbc, 0xcd,
	0x05, 0x7f, 0x25, 0x0a, 0xf4, 0x87, 0x74, 0xaa, 0x03, 0xdc, 0x9e, 0x7f, 0x1f, 0x8b, 0x12, 0xdf,
	0xa3, 0xf8, 0x1d, 0x94, 0x3d, 0x83, 0x77, 0x1b, 0xe5, 0xf6, 0x82, 0x02, 0x05, 0x3d, 0x74, 0xb0,
	0x00, 0x1e, 0xb4, 0xc0, 0x28, 0xf1, 0x47, 0x14, 0xb6, 0x0b, 0xcf, 0xfe, 0x76, 0x01, 0x76, 0x57,
	0x97, 0x3d, 0x82, 0xd3, 0xaa, 0x2e, 0x97, 0x43, 0xfc, 0xc6, 0x3a, 0x28, 0x66, 0xef, 0x48, 0x31,
	0x3f, 0x04, 0xaf, 0x32, 0x42, 0x9b, 0x65, 0x56, 0x20, 0x55, 0xfb, 0x84, 0xef, 0x00, 0xe6, 0xc3,
	0x00, 0x65, 0x42, 0x3e, 0x97, 0x7c, 0x1b, 0x93, 0x3d, 0x05, 0x57, 0xe3, 0x95, 0x2d, 0x9c, 0xed,
	0xa4, 0x47, 0xe1, 0xd1, 0x1b, 0xcb, 0x89, 0x63, 0x13, 0x53, 0x25, 0x6a, 0x52, 0xdc, 0x66, 0xba,
	0xad, 0x62, 0x17, 0xb6, 0x1d, 0xd2, 0x82, 0xa8, 0x6e, 0x1e, 0x6f, 0x83, 0x36, 0xdf, 0x12, 0x51,
	0x47, 0x09, 0x15, 0xa8, 0xcf, 0x1b, 0x8b, 0x31, 0x70, 0xed, 0x8a, 0x8a, 0xe1, 0x71, 0x5a, 0xb3,
	0x4f, 0x60, 0x68, 0xd5, 0xa0, 0xa6, 0x01, 0x6a, 0x1a, 0x8f, 0xba, 0x9e, 0xba, 0x66, 0xeb, 0x62,
	0x01, 0x78, 0x76, 0x7d, 0x21, 0xd6, 0xa8, 0x49, 0xf5, 0xf1, 0x1c, 0x88, 0x47, 0x08, 0xdf, 0x39,
	0xd9, 0x14, 0x46, 0xb1, 0x2a, 0x4a, 0x25, 0xeb, 0xbe, 0xba, 0x47, 0x27, 0xd8, 0x87, 0xac, 0xa4,
	0x5b, 0xd3, 0xbf, 0x4f, 0x67, 0xd9, 0x01, 0x56, 0xd2, 0xac, 0xfa, 0x51, 0x6b, 0xa5, 0xfd, 0x31,
	0xb5, 0xff, 0xc6, 0x64, 0x4f, 0xc0, 0x35, 0x22, 0xad, 0xfc, 0x07, 0x24, 0xe9, 0xc3, 0xf0, 0xf0,
	0xad, 0xe4, 0x44, 0x60, 0x8f, 0xc1, 0xcd, 0x55, 0x5a, 0xf9, 0x67, 0xcd, 0x2b, 0x76, 0xa1, 0xd2,
	0x9f, 0xb1, 0xaa, 0x44, 0x8a, 0x9c, 0x1c, 0xb3, 0x08, 0x60, 0x87, 0x59, 0x59, 0x8c, 0xad, 0xa0,
	0x43, 0x15, 0xa4, 0xb5, 0xdd, 0x2b, 0x11, 0x46, 0x34, 0x0f, 0xe1, 0xf1, 0xbd, 0x2c, 0x61, 0xfe,
	0x3d, 0x3c, 0x3c, 0xf2, 0xfc, 0xb3, 0x67, 0x30, 0x88, 0x55, 0x9e, 0xdb, 0xdb, 0x77, 0x16, 0xbe,
	0x29, 0xeb, 0x37, 0xbf, 0xe1, 0x4c, 0x46, 0xe1, 0x6e, 0x0e, 0xcc, 0xde, 0x09, 0x9c, 0xef, 0xae,
	0xe1, 0x33, 0xa5, 0xd3, 0x50, 0x94, 0x22, 0xbe, 0xc6, 0xb0, 0x5a, 0xad, 0x6f, 0x45, 0xbe, 0xca,
	0xa4, 0x45, 0x8a, 0xb0, 0xe9, 0xe4, 0x70, 0x33, 0x29, 0x42, 0x9a, 0x14, 0xaf, 0x9d, 0xdf, 0xce,
	0xeb, 0x91, 0x51, 0x7f, 0xcb, 0xcb, 0xaf, 0xea, 0xd5, 0xbf, 0xbd, 0xc9, 0x62, 0xb5, 0x7e, 0xdb,
	0x04, 0x78, 0x55, 0xff, 0xfc, 0xda, 0x0e, 0x8a, 0x58, 0xe5, 0x97, 0xa7, 0x34, 0x32, 0xbe, 0xf8,
	0x3f, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x00, 0xa9, 0x75, 0xc3, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TraceSegmentServiceClient is the client API for TraceSegmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraceSegmentServiceClient interface {
	Collect(ctx context.Context, opts ...grpc.CallOption) (TraceSegmentService_CollectClient, error)
}

type traceSegmentServiceClient struct {
	cc *grpc.ClientConn
}

func NewTraceSegmentServiceClient(cc *grpc.ClientConn) TraceSegmentServiceClient {
	return &traceSegmentServiceClient{cc}
}

func (c *traceSegmentServiceClient) Collect(ctx context.Context, opts ...grpc.CallOption) (TraceSegmentService_CollectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TraceSegmentService_serviceDesc.Streams[0], "/TraceSegmentService/collect", opts...)
	if err != nil {
		return nil, err
	}
	x := &traceSegmentServiceCollectClient{stream}
	return x, nil
}

type TraceSegmentService_CollectClient interface {
	Send(*common.UpstreamSegment) error
	CloseAndRecv() (*Downstream, error)
	grpc.ClientStream
}

type traceSegmentServiceCollectClient struct {
	grpc.ClientStream
}

func (x *traceSegmentServiceCollectClient) Send(m *common.UpstreamSegment) error {
	return x.ClientStream.SendMsg(m)
}

func (x *traceSegmentServiceCollectClient) CloseAndRecv() (*Downstream, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Downstream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TraceSegmentServiceServer is the server API for TraceSegmentService service.
type TraceSegmentServiceServer interface {
	Collect(TraceSegmentService_CollectServer) error
}

// UnimplementedTraceSegmentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTraceSegmentServiceServer struct {
}

func (*UnimplementedTraceSegmentServiceServer) Collect(srv TraceSegmentService_CollectServer) error {
	return status.Errorf(codes.Unimplemented, "method Collect not implemented")
}

func RegisterTraceSegmentServiceServer(s *grpc.Server, srv TraceSegmentServiceServer) {
	s.RegisterService(&_TraceSegmentService_serviceDesc, srv)
}

func _TraceSegmentService_Collect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TraceSegmentServiceServer).Collect(&traceSegmentServiceCollectServer{stream})
}

type TraceSegmentService_CollectServer interface {
	SendAndClose(*Downstream) error
	Recv() (*common.UpstreamSegment, error)
	grpc.ServerStream
}

type traceSegmentServiceCollectServer struct {
	grpc.ServerStream
}

func (x *traceSegmentServiceCollectServer) SendAndClose(m *Downstream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *traceSegmentServiceCollectServer) Recv() (*common.UpstreamSegment, error) {
	m := new(common.UpstreamSegment)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TraceSegmentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TraceSegmentService",
	HandlerType: (*TraceSegmentServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "collect",
			Handler:       _TraceSegmentService_Collect_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "language-agent/TraceSegmentService.proto",
}