// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/errors.proto

package internal

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// Error is the generic error returned from unary RPCs.
type Error struct {
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	// This is to make the error more compatible with users that expect errors to
	// be Status objects:
	// https://github.com/grpc/grpc/blob/master/src/proto/grpc/status/status.proto
	// It should be the exact same message as the Error field.
	Code                 int32      `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string     `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Details              []*any.Any `protobuf:"bytes,4,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b093362ca6d1e03, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

// StreamError is a response type which is returned when
// streaming rpc returns an error.
type StreamError struct {
	GrpcCode             int32      `protobuf:"varint,1,opt,name=grpc_code,json=grpcCode,proto3" json:"grpc_code,omitempty"`
	HttpCode             int32      `protobuf:"varint,2,opt,name=http_code,json=httpCode,proto3" json:"http_code,omitempty"`
	Message              string     `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	HttpStatus           string     `protobuf:"bytes,4,opt,name=http_status,json=httpStatus,proto3" json:"http_status,omitempty"`
	Details              []*any.Any `protobuf:"bytes,5,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StreamError) Reset()         { *m = StreamError{} }
func (m *StreamError) String() string { return proto.CompactTextString(m) }
func (*StreamError) ProtoMessage()    {}
func (*StreamError) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b093362ca6d1e03, []int{1}
}

func (m *StreamError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamError.Unmarshal(m, b)
}
func (m *StreamError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamError.Marshal(b, m, deterministic)
}
func (m *StreamError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamError.Merge(m, src)
}
func (m *StreamError) XXX_Size() int {
	return xxx_messageInfo_StreamError.Size(m)
}
func (m *StreamError) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamError.DiscardUnknown(m)
}

var xxx_messageInfo_StreamError proto.InternalMessageInfo

func (m *StreamError) GetGrpcCode() int32 {
	if m != nil {
		return m.GrpcCode
	}
	return 0
}

func (m *StreamError) GetHttpCode() int32 {
	if m != nil {
		return m.HttpCode
	}
	return 0
}

func (m *StreamError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *StreamError) GetHttpStatus() string {
	if m != nil {
		return m.HttpStatus
	}
	return ""
}

func (m *StreamError) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*Error)(nil), "grpc.gateway.runtime.Error")
	proto.RegisterType((*StreamError)(nil), "grpc.gateway.runtime.StreamError")
}

func init() { proto.RegisterFile("internal/errors.proto", fileDescriptor_9b093362ca6d1e03) }

var fileDescriptor_9b093362ca6d1e03 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x89, 0xbb, 0x75, 0xdb, 0xe9, 0x2d, 0x54, 0x88, 0xee, 0xc1, 0xb2, 0xa7, 0x9e, 0x52,
	0xd0, 0x27, 0xd0, 0xc5, 0x17, 0xe8, 0xde, 0xbc, 0x2c, 0xd9, 0xdd, 0x31, 0x16, 0xda, 0xa4, 0x24,
	0x53, 0xa4, 0xf8, 0x56, 0x3e, 0xa1, 0x24, 0xa5, 0xb0, 0x27, 0xf1, 0xd6, 0xf9, 0xfb, 0xcf, 0x7c,
	0x1f, 0x81, 0xbb, 0xd6, 0x10, 0x3a, 0xa3, 0xba, 0x1a, 0x9d, 0xb3, 0xce, 0xcb, 0xc1, 0x59, 0xb2,
	0xbc, 0xd0, 0x6e, 0x38, 0x4b, 0xad, 0x08, 0xbf, 0xd4, 0x24, 0xdd, 0x68, 0xa8, 0xed, 0xf1, 0xe1,
	0x5e, 0x5b, 0xab, 0x3b, 0xac, 0x63, 0xe7, 0x34, 0x7e, 0xd4, 0xca, 0x4c, 0xf3, 0xc2, 0xee, 0x1b,
	0x92, 0xb7, 0x70, 0x80, 0x17, 0x90, 0xc4, 0x4b, 0x82, 0x95, 0xac, 0xca, 0x9a, 0x79, 0xe0, 0x1c,
	0xd6, 0x67, 0x7b, 0x41, 0x71, 0x53, 0xb2, 0x2a, 0x69, 0xe2, 0x37, 0x17, 0xb0, 0xe9, 0xd1, 0x7b,
	0xa5, 0x51, 0xac, 0x62, 0x77, 0x19, 0xb9, 0x84, 0xcd, 0x05, 0x49, 0xb5, 0x9d, 0x17, 0xeb, 0x72,
	0x55, 0xe5, 0x4f, 0x85, 0x9c, 0xc9, 0x72, 0x21, 0xcb, 0x17, 0x33, 0x35, 0x4b, 0x69, 0xf7, 0xc3,
	0x20, 0x3f, 0x90, 0x43, 0xd5, 0xcf, 0x0e, 0x5b, 0xc8, 0x82, 0xff, 0x31, 0x22, 0x59, 0x44, 0xa6,
	0x21, 0xd8, 0x07, 0xec, 0x16, 0xb2, 0x4f, 0xa2, 0xe1, 0x78, 0xe5, 0x93, 0x86, 0x60, 0xff, 0xb7,
	0xd3, 0x23, 0xe4, 0x71, 0xcd, 0x93, 0xa2, 0x31, 0x78, 0x85, 0xbf, 0x10, 0xa2, 0x43, 0x4c, 0xae,
	0xa5, 0x93, 0x7f, 0x48, 0xbf, 0xc2, 0x7b, 0xba, 0xbc, 0xfd, 0xe9, 0x36, 0x56, 0x9e, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xde, 0x72, 0x6b, 0x83, 0x8e, 0x01, 0x00, 0x00,
}
