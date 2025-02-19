// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: query_service.proto

package api_v3

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	v1 "github.com/kjschnei001/jaeger/proto-gen/otel/trace/v1"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Request object to get a trace.
type GetTraceRequest struct {
	// Hex encoded 64 or 128 bit trace ID.
	TraceId              string   `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTraceRequest) Reset()         { *m = GetTraceRequest{} }
func (m *GetTraceRequest) String() string { return proto.CompactTextString(m) }
func (*GetTraceRequest) ProtoMessage()    {}
func (*GetTraceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fcb6756dc1afb8d, []int{0}
}
func (m *GetTraceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTraceRequest.Unmarshal(m, b)
}
func (m *GetTraceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTraceRequest.Marshal(b, m, deterministic)
}
func (m *GetTraceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTraceRequest.Merge(m, src)
}
func (m *GetTraceRequest) XXX_Size() int {
	return xxx_messageInfo_GetTraceRequest.Size(m)
}
func (m *GetTraceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTraceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTraceRequest proto.InternalMessageInfo

func (m *GetTraceRequest) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

// Response object with spans.
type SpansResponseChunk struct {
	// A list of OpenTelemetry ResourceSpans.
	// In case of JSON format the ids (trace_id, span_id, parent_id) are encoded in base64 even though OpenTelemetry specification
	// mandates to use hex encoding [2].
	// Base64 is chosen to keep compatibility with JSONPb codec.
	// [1]: https://github.com/open-telemetry/opentelemetry-proto/blob/main/opentelemetry/proto/trace/v1/trace.proto
	// [2]: https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/protocol/otlp.md#otlphttp
	ResourceSpans        []*v1.ResourceSpans `protobuf:"bytes,1,rep,name=resource_spans,json=resourceSpans,proto3" json:"resource_spans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SpansResponseChunk) Reset()         { *m = SpansResponseChunk{} }
func (m *SpansResponseChunk) String() string { return proto.CompactTextString(m) }
func (*SpansResponseChunk) ProtoMessage()    {}
func (*SpansResponseChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fcb6756dc1afb8d, []int{1}
}
func (m *SpansResponseChunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpansResponseChunk.Unmarshal(m, b)
}
func (m *SpansResponseChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpansResponseChunk.Marshal(b, m, deterministic)
}
func (m *SpansResponseChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpansResponseChunk.Merge(m, src)
}
func (m *SpansResponseChunk) XXX_Size() int {
	return xxx_messageInfo_SpansResponseChunk.Size(m)
}
func (m *SpansResponseChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_SpansResponseChunk.DiscardUnknown(m)
}

var xxx_messageInfo_SpansResponseChunk proto.InternalMessageInfo

func (m *SpansResponseChunk) GetResourceSpans() []*v1.ResourceSpans {
	if m != nil {
		return m.ResourceSpans
	}
	return nil
}

// Query parameters to find traces.
// Note that some storage implementations do not guarantee the correct implementation of all parameters.
type TraceQueryParameters struct {
	ServiceName   string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	OperationName string `protobuf:"bytes,2,opt,name=operation_name,json=operationName,proto3" json:"operation_name,omitempty"`
	// Attributes are matched against Span and Resource attributes.
	// At least one span in a trace must match all specified attributes.
	Attributes map[string]string `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Span min start time in. REST API uses RFC-3339ns format. Required.
	StartTimeMin *types.Timestamp `protobuf:"bytes,4,opt,name=start_time_min,json=startTimeMin,proto3" json:"start_time_min,omitempty"`
	// Span max start time. REST API uses RFC-3339ns format. Required.
	StartTimeMax *types.Timestamp `protobuf:"bytes,5,opt,name=start_time_max,json=startTimeMax,proto3" json:"start_time_max,omitempty"`
	// Span min duration. REST API uses Golang's time format e.g. 10s.
	DurationMin *types.Duration `protobuf:"bytes,6,opt,name=duration_min,json=durationMin,proto3" json:"duration_min,omitempty"`
	// Span max duration. REST API uses Golang's time format e.g. 10s.
	DurationMax *types.Duration `protobuf:"bytes,7,opt,name=duration_max,json=durationMax,proto3" json:"duration_max,omitempty"`
	// Maximum number of traces in the response.
	NumTraces            int32    `protobuf:"varint,8,opt,name=num_traces,json=numTraces,proto3" json:"num_traces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TraceQueryParameters) Reset()         { *m = TraceQueryParameters{} }
func (m *TraceQueryParameters) String() string { return proto.CompactTextString(m) }
func (*TraceQueryParameters) ProtoMessage()    {}
func (*TraceQueryParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fcb6756dc1afb8d, []int{2}
}
func (m *TraceQueryParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceQueryParameters.Unmarshal(m, b)
}
func (m *TraceQueryParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceQueryParameters.Marshal(b, m, deterministic)
}
func (m *TraceQueryParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceQueryParameters.Merge(m, src)
}
func (m *TraceQueryParameters) XXX_Size() int {
	return xxx_messageInfo_TraceQueryParameters.Size(m)
}
func (m *TraceQueryParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceQueryParameters.DiscardUnknown(m)
}

var xxx_messageInfo_TraceQueryParameters proto.InternalMessageInfo

func (m *TraceQueryParameters) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *TraceQueryParameters) GetOperationName() string {
	if m != nil {
		return m.OperationName
	}
	return ""
}

func (m *TraceQueryParameters) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *TraceQueryParameters) GetStartTimeMin() *types.Timestamp {
	if m != nil {
		return m.StartTimeMin
	}
	return nil
}

func (m *TraceQueryParameters) GetStartTimeMax() *types.Timestamp {
	if m != nil {
		return m.StartTimeMax
	}
	return nil
}

func (m *TraceQueryParameters) GetDurationMin() *types.Duration {
	if m != nil {
		return m.DurationMin
	}
	return nil
}

func (m *TraceQueryParameters) GetDurationMax() *types.Duration {
	if m != nil {
		return m.DurationMax
	}
	return nil
}

func (m *TraceQueryParameters) GetNumTraces() int32 {
	if m != nil {
		return m.NumTraces
	}
	return 0
}

// Request object to search traces.
type FindTracesRequest struct {
	Query                *TraceQueryParameters `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FindTracesRequest) Reset()         { *m = FindTracesRequest{} }
func (m *FindTracesRequest) String() string { return proto.CompactTextString(m) }
func (*FindTracesRequest) ProtoMessage()    {}
func (*FindTracesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fcb6756dc1afb8d, []int{3}
}
func (m *FindTracesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindTracesRequest.Unmarshal(m, b)
}
func (m *FindTracesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindTracesRequest.Marshal(b, m, deterministic)
}
func (m *FindTracesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindTracesRequest.Merge(m, src)
}
func (m *FindTracesRequest) XXX_Size() int {
	return xxx_messageInfo_FindTracesRequest.Size(m)
}
func (m *FindTracesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindTracesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindTracesRequest proto.InternalMessageInfo

func (m *FindTracesRequest) GetQuery() *TraceQueryParameters {
	if m != nil {
		return m.Query
	}
	return nil
}

// Request object to get service names.
type GetServicesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServicesRequest) Reset()         { *m = GetServicesRequest{} }
func (m *GetServicesRequest) String() string { return proto.CompactTextString(m) }
func (*GetServicesRequest) ProtoMessage()    {}
func (*GetServicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fcb6756dc1afb8d, []int{4}
}
func (m *GetServicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServicesRequest.Unmarshal(m, b)
}
func (m *GetServicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServicesRequest.Marshal(b, m, deterministic)
}
func (m *GetServicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServicesRequest.Merge(m, src)
}
func (m *GetServicesRequest) XXX_Size() int {
	return xxx_messageInfo_GetServicesRequest.Size(m)
}
func (m *GetServicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServicesRequest proto.InternalMessageInfo

// Response object to get service names.
type GetServicesResponse struct {
	Services             []string `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServicesResponse) Reset()         { *m = GetServicesResponse{} }
func (m *GetServicesResponse) String() string { return proto.CompactTextString(m) }
func (*GetServicesResponse) ProtoMessage()    {}
func (*GetServicesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fcb6756dc1afb8d, []int{5}
}
func (m *GetServicesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServicesResponse.Unmarshal(m, b)
}
func (m *GetServicesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServicesResponse.Marshal(b, m, deterministic)
}
func (m *GetServicesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServicesResponse.Merge(m, src)
}
func (m *GetServicesResponse) XXX_Size() int {
	return xxx_messageInfo_GetServicesResponse.Size(m)
}
func (m *GetServicesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServicesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetServicesResponse proto.InternalMessageInfo

func (m *GetServicesResponse) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

// Request object to get operation names.
type GetOperationsRequest struct {
	// Required service name.
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// Optional span kind.
	SpanKind             string   `protobuf:"bytes,2,opt,name=span_kind,json=spanKind,proto3" json:"span_kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOperationsRequest) Reset()         { *m = GetOperationsRequest{} }
func (m *GetOperationsRequest) String() string { return proto.CompactTextString(m) }
func (*GetOperationsRequest) ProtoMessage()    {}
func (*GetOperationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fcb6756dc1afb8d, []int{6}
}
func (m *GetOperationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOperationsRequest.Unmarshal(m, b)
}
func (m *GetOperationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOperationsRequest.Marshal(b, m, deterministic)
}
func (m *GetOperationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperationsRequest.Merge(m, src)
}
func (m *GetOperationsRequest) XXX_Size() int {
	return xxx_messageInfo_GetOperationsRequest.Size(m)
}
func (m *GetOperationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperationsRequest proto.InternalMessageInfo

func (m *GetOperationsRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *GetOperationsRequest) GetSpanKind() string {
	if m != nil {
		return m.SpanKind
	}
	return ""
}

// Operation encapsulates information about operation.
type Operation struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SpanKind             string   `protobuf:"bytes,2,opt,name=span_kind,json=spanKind,proto3" json:"span_kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fcb6756dc1afb8d, []int{7}
}
func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (m *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(m, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

func (m *Operation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Operation) GetSpanKind() string {
	if m != nil {
		return m.SpanKind
	}
	return ""
}

// Response object to get operation names.
type GetOperationsResponse struct {
	Operations           []*Operation `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetOperationsResponse) Reset()         { *m = GetOperationsResponse{} }
func (m *GetOperationsResponse) String() string { return proto.CompactTextString(m) }
func (*GetOperationsResponse) ProtoMessage()    {}
func (*GetOperationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fcb6756dc1afb8d, []int{8}
}
func (m *GetOperationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOperationsResponse.Unmarshal(m, b)
}
func (m *GetOperationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOperationsResponse.Marshal(b, m, deterministic)
}
func (m *GetOperationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperationsResponse.Merge(m, src)
}
func (m *GetOperationsResponse) XXX_Size() int {
	return xxx_messageInfo_GetOperationsResponse.Size(m)
}
func (m *GetOperationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperationsResponse proto.InternalMessageInfo

func (m *GetOperationsResponse) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func init() {
	proto.RegisterType((*GetTraceRequest)(nil), "jaeger.api_v3.GetTraceRequest")
	proto.RegisterType((*SpansResponseChunk)(nil), "jaeger.api_v3.SpansResponseChunk")
	proto.RegisterType((*TraceQueryParameters)(nil), "jaeger.api_v3.TraceQueryParameters")
	proto.RegisterMapType((map[string]string)(nil), "jaeger.api_v3.TraceQueryParameters.AttributesEntry")
	proto.RegisterType((*FindTracesRequest)(nil), "jaeger.api_v3.FindTracesRequest")
	proto.RegisterType((*GetServicesRequest)(nil), "jaeger.api_v3.GetServicesRequest")
	proto.RegisterType((*GetServicesResponse)(nil), "jaeger.api_v3.GetServicesResponse")
	proto.RegisterType((*GetOperationsRequest)(nil), "jaeger.api_v3.GetOperationsRequest")
	proto.RegisterType((*Operation)(nil), "jaeger.api_v3.Operation")
	proto.RegisterType((*GetOperationsResponse)(nil), "jaeger.api_v3.GetOperationsResponse")
}

func init() { proto.RegisterFile("query_service.proto", fileDescriptor_5fcb6756dc1afb8d) }

var fileDescriptor_5fcb6756dc1afb8d = []byte{
	// 675 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x7f, 0x4f, 0xdb, 0x3a,
	0x14, 0x25, 0x40, 0xa1, 0xbd, 0x6d, 0xe1, 0x3d, 0xd3, 0xa7, 0x17, 0x32, 0x8d, 0xb5, 0x61, 0x93,
	0x2a, 0x6d, 0x4a, 0x47, 0xf9, 0x87, 0x4d, 0x4c, 0xda, 0x6f, 0x34, 0x4d, 0xb0, 0x11, 0xd8, 0xfe,
	0x98, 0x26, 0x45, 0x86, 0xde, 0x95, 0x0c, 0xe2, 0x04, 0xdb, 0xa9, 0xda, 0xcf, 0xb1, 0xef, 0xb9,
	0xcf, 0x30, 0xc5, 0x76, 0xb2, 0x36, 0x9d, 0x10, 0xfb, 0x2b, 0xb9, 0xf6, 0x39, 0xc7, 0xd7, 0xf7,
	0xdc, 0x6b, 0xd8, 0xb8, 0x4e, 0x91, 0x4f, 0x02, 0x81, 0x7c, 0x14, 0x9e, 0xa3, 0x97, 0xf0, 0x58,
	0xc6, 0xa4, 0xf9, 0x9d, 0xe2, 0x10, 0xb9, 0x47, 0x93, 0x30, 0x18, 0xed, 0x3a, 0xdd, 0x38, 0x41,
	0x26, 0xf1, 0x0a, 0x23, 0x94, 0x7c, 0xd2, 0x53, 0x98, 0x9e, 0xe4, 0xf4, 0x1c, 0x7b, 0xa3, 0x1d,
	0xfd, 0xa3, 0x89, 0xce, 0xbd, 0x61, 0x1c, 0x0f, 0xaf, 0x50, 0x43, 0xce, 0xd2, 0x6f, 0x3d, 0x19,
	0x46, 0x28, 0x24, 0x8d, 0x12, 0x03, 0xd8, 0x2a, 0x03, 0x06, 0x29, 0xa7, 0x32, 0x8c, 0x99, 0xde,
	0x77, 0x1f, 0xc1, 0xfa, 0x01, 0xca, 0xd3, 0x4c, 0xd2, 0xc7, 0xeb, 0x14, 0x85, 0x24, 0x9b, 0x50,
	0x55, 0x47, 0x04, 0xe1, 0xc0, 0xb6, 0xda, 0x56, 0xb7, 0xe6, 0xaf, 0xaa, 0xf8, 0xdd, 0xc0, 0xbd,
	0x00, 0x72, 0x92, 0x50, 0x26, 0x7c, 0x14, 0x49, 0xcc, 0x04, 0xbe, 0xba, 0x48, 0xd9, 0x25, 0xf1,
	0x61, 0x8d, 0xa3, 0x88, 0x53, 0x7e, 0x8e, 0x81, 0xc8, 0xb6, 0x6d, 0xab, 0xbd, 0xd4, 0xad, 0xf7,
	0x1f, 0x7a, 0x33, 0xf7, 0xd0, 0x27, 0x7a, 0x3a, 0xfd, 0xd1, 0x8e, 0xe7, 0x1b, 0x8e, 0x56, 0x6c,
	0xf2, 0xe9, 0xd0, 0xfd, 0xb1, 0x0c, 0x2d, 0x95, 0xd5, 0x71, 0x56, 0xae, 0x8f, 0x94, 0xd3, 0x08,
	0x25, 0x72, 0x41, 0x3a, 0xd0, 0x30, 0xb5, 0x0b, 0x18, 0x8d, 0xd0, 0x64, 0x58, 0x37, 0x6b, 0x47,
	0x34, 0x42, 0xf2, 0x00, 0xd6, 0xe2, 0x04, 0xf5, 0x35, 0x35, 0x68, 0x51, 0x81, 0x9a, 0xc5, 0xaa,
	0x82, 0x9d, 0x00, 0x50, 0x29, 0x79, 0x78, 0x96, 0x4a, 0x14, 0xf6, 0x92, 0x4a, 0x79, 0xd7, 0x9b,
	0x71, 0xc2, 0xfb, 0x53, 0x0a, 0xde, 0x8b, 0x82, 0xf5, 0x86, 0x49, 0x3e, 0xf1, 0xa7, 0x64, 0xc8,
	0x73, 0x58, 0x13, 0x92, 0x72, 0x19, 0x64, 0x46, 0x04, 0x51, 0xc8, 0xec, 0xe5, 0xb6, 0xd5, 0xad,
	0xf7, 0x1d, 0x4f, 0x1b, 0xe1, 0xe5, 0x46, 0x78, 0xa7, 0xb9, 0x53, 0x7e, 0x43, 0x31, 0xb2, 0xf8,
	0x30, 0x64, 0x65, 0x05, 0x3a, 0xb6, 0x2b, 0x7f, 0xa3, 0x40, 0xc7, 0x64, 0x1f, 0x1a, 0xb9, 0xcb,
	0x2a, 0x83, 0x15, 0xc5, 0xdf, 0x9c, 0xe3, 0xbf, 0x36, 0x20, 0xbf, 0x9e, 0xc3, 0xb3, 0xf3, 0x67,
	0xd8, 0x74, 0x6c, 0xaf, 0xde, 0x9e, 0x4d, 0xc7, 0xe4, 0x2e, 0x00, 0x4b, 0xa3, 0x40, 0x99, 0x2c,
	0xec, 0x6a, 0xdb, 0xea, 0x56, 0xfc, 0x1a, 0x4b, 0x23, 0x55, 0x48, 0xe1, 0x3c, 0x83, 0xf5, 0x52,
	0xf5, 0xc8, 0x3f, 0xb0, 0x74, 0x89, 0x13, 0xe3, 0x63, 0xf6, 0x4b, 0x5a, 0x50, 0x19, 0xd1, 0xab,
	0x34, 0xb7, 0x4d, 0x07, 0x4f, 0x17, 0xf7, 0x2c, 0xf7, 0x08, 0xfe, 0x7d, 0x1b, 0xb2, 0x81, 0x16,
	0xcb, 0xfb, 0xf5, 0x09, 0x54, 0xd4, 0x4c, 0x29, 0x89, 0x7a, 0x7f, 0xfb, 0x16, 0x16, 0xfa, 0x9a,
	0xe1, 0xb6, 0x80, 0x1c, 0xa0, 0x3c, 0xd1, 0xbd, 0x93, 0x0b, 0xba, 0x3b, 0xb0, 0x31, 0xb3, 0xaa,
	0x7b, 0x9d, 0x38, 0x50, 0x35, 0x5d, 0xa6, 0x1b, 0xbc, 0xe6, 0x17, 0xb1, 0x7b, 0x08, 0xad, 0x03,
	0x94, 0x1f, 0xf2, 0xfe, 0x2a, 0x72, 0xb3, 0x61, 0xd5, 0x60, 0xf2, 0x51, 0x32, 0x21, 0xb9, 0x03,
	0xb5, 0x6c, 0x56, 0x82, 0xcb, 0x90, 0x0d, 0xcc, 0x45, 0xab, 0xd9, 0xc2, 0xfb, 0x90, 0x0d, 0xdc,
	0x7d, 0xa8, 0x15, 0x5a, 0x84, 0xc0, 0xf2, 0x54, 0xa7, 0xab, 0xff, 0x9b, 0xd9, 0xc7, 0xf0, 0x5f,
	0x29, 0x19, 0x73, 0x83, 0x3d, 0x80, 0x62, 0x04, 0xf2, 0x21, 0xb5, 0x4b, 0xe5, 0x2a, 0x68, 0xfe,
	0x14, 0xb6, 0xff, 0x73, 0x11, 0x1a, 0xaa, 0x86, 0xa6, 0x2a, 0xe4, 0x18, 0xaa, 0xf9, 0xbb, 0x41,
	0xb6, 0x4a, 0x12, 0xa5, 0x07, 0xc5, 0xe9, 0x94, 0xf6, 0xe7, 0x9f, 0x10, 0x77, 0xe1, 0xb1, 0x45,
	0x3e, 0x01, 0xfc, 0x36, 0x97, 0xb4, 0x4b, 0xa4, 0x39, 0xdf, 0x6f, 0x2b, 0xfb, 0x19, 0xea, 0x53,
	0x6e, 0x92, 0xce, 0x7c, 0xb2, 0x25, 0xff, 0x1d, 0xf7, 0x26, 0x88, 0x96, 0x77, 0x17, 0xc8, 0x57,
	0x68, 0xce, 0x54, 0x99, 0x6c, 0xcf, 0xd3, 0xe6, 0x1a, 0xc2, 0xb9, 0x7f, 0x33, 0x28, 0x57, 0x7f,
	0xd9, 0x81, 0xff, 0xc3, 0xd8, 0x60, 0xb3, 0x61, 0x0a, 0xd9, 0xd0, 0x50, 0xbe, 0xac, 0xe8, 0xef,
	0xd9, 0x8a, 0x1a, 0xc5, 0xdd, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x58, 0x79, 0x69, 0x52,
	0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	// GetTrace returns a single trace.
	// Note that the JSON response over HTTP is wrapped into result envelope "{"result": ...}"
	// It means that the JSON response cannot be directly unmarshalled using JSONPb.
	// This can be fixed by first parsing into user-defined envelope with standard JSON library
	// or string manipulation to remove the envelope. Alternatively generate objects using OpenAPI.
	GetTrace(ctx context.Context, in *GetTraceRequest, opts ...grpc.CallOption) (QueryService_GetTraceClient, error)
	// FindTraces searches for traces.
	// See GetTrace for JSON unmarshalling.
	FindTraces(ctx context.Context, in *FindTracesRequest, opts ...grpc.CallOption) (QueryService_FindTracesClient, error)
	// GetServices returns service names.
	GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error)
	// GetOperations returns operation names.
	GetOperations(ctx context.Context, in *GetOperationsRequest, opts ...grpc.CallOption) (*GetOperationsResponse, error)
}

type queryServiceClient struct {
	cc *grpc.ClientConn
}

func NewQueryServiceClient(cc *grpc.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) GetTrace(ctx context.Context, in *GetTraceRequest, opts ...grpc.CallOption) (QueryService_GetTraceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_QueryService_serviceDesc.Streams[0], "/jaeger.api_v3.QueryService/GetTrace", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryServiceGetTraceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QueryService_GetTraceClient interface {
	Recv() (*SpansResponseChunk, error)
	grpc.ClientStream
}

type queryServiceGetTraceClient struct {
	grpc.ClientStream
}

func (x *queryServiceGetTraceClient) Recv() (*SpansResponseChunk, error) {
	m := new(SpansResponseChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queryServiceClient) FindTraces(ctx context.Context, in *FindTracesRequest, opts ...grpc.CallOption) (QueryService_FindTracesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_QueryService_serviceDesc.Streams[1], "/jaeger.api_v3.QueryService/FindTraces", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryServiceFindTracesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QueryService_FindTracesClient interface {
	Recv() (*SpansResponseChunk, error)
	grpc.ClientStream
}

type queryServiceFindTracesClient struct {
	grpc.ClientStream
}

func (x *queryServiceFindTracesClient) Recv() (*SpansResponseChunk, error) {
	m := new(SpansResponseChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queryServiceClient) GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error) {
	out := new(GetServicesResponse)
	err := c.cc.Invoke(ctx, "/jaeger.api_v3.QueryService/GetServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) GetOperations(ctx context.Context, in *GetOperationsRequest, opts ...grpc.CallOption) (*GetOperationsResponse, error) {
	out := new(GetOperationsResponse)
	err := c.cc.Invoke(ctx, "/jaeger.api_v3.QueryService/GetOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	// GetTrace returns a single trace.
	// Note that the JSON response over HTTP is wrapped into result envelope "{"result": ...}"
	// It means that the JSON response cannot be directly unmarshalled using JSONPb.
	// This can be fixed by first parsing into user-defined envelope with standard JSON library
	// or string manipulation to remove the envelope. Alternatively generate objects using OpenAPI.
	GetTrace(*GetTraceRequest, QueryService_GetTraceServer) error
	// FindTraces searches for traces.
	// See GetTrace for JSON unmarshalling.
	FindTraces(*FindTracesRequest, QueryService_FindTracesServer) error
	// GetServices returns service names.
	GetServices(context.Context, *GetServicesRequest) (*GetServicesResponse, error)
	// GetOperations returns operation names.
	GetOperations(context.Context, *GetOperationsRequest) (*GetOperationsResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) GetTrace(req *GetTraceRequest, srv QueryService_GetTraceServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTrace not implemented")
}
func (*UnimplementedQueryServiceServer) FindTraces(req *FindTracesRequest, srv QueryService_FindTracesServer) error {
	return status.Errorf(codes.Unimplemented, "method FindTraces not implemented")
}
func (*UnimplementedQueryServiceServer) GetServices(ctx context.Context, req *GetServicesRequest) (*GetServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServices not implemented")
}
func (*UnimplementedQueryServiceServer) GetOperations(ctx context.Context, req *GetOperationsRequest) (*GetOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperations not implemented")
}

func RegisterQueryServiceServer(s *grpc.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_GetTrace_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTraceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServiceServer).GetTrace(m, &queryServiceGetTraceServer{stream})
}

type QueryService_GetTraceServer interface {
	Send(*SpansResponseChunk) error
	grpc.ServerStream
}

type queryServiceGetTraceServer struct {
	grpc.ServerStream
}

func (x *queryServiceGetTraceServer) Send(m *SpansResponseChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _QueryService_FindTraces_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FindTracesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServiceServer).FindTraces(m, &queryServiceFindTracesServer{stream})
}

type QueryService_FindTracesServer interface {
	Send(*SpansResponseChunk) error
	grpc.ServerStream
}

type queryServiceFindTracesServer struct {
	grpc.ServerStream
}

func (x *queryServiceFindTracesServer) Send(m *SpansResponseChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _QueryService_GetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).GetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.api_v3.QueryService/GetServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).GetServices(ctx, req.(*GetServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_GetOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).GetOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.api_v3.QueryService/GetOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).GetOperations(ctx, req.(*GetOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jaeger.api_v3.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServices",
			Handler:    _QueryService_GetServices_Handler,
		},
		{
			MethodName: "GetOperations",
			Handler:    _QueryService_GetOperations_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTrace",
			Handler:       _QueryService_GetTrace_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FindTraces",
			Handler:       _QueryService_FindTraces_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "query_service.proto",
}
