// Code generated by protoc-gen-go. DO NOT EDIT.
// source: header.proto

package OnlineConsultant

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// AccessServer可识别header，用于外部请求
type AccessRequestHeader struct {
	//设备id
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	//来源
	ClientType ClientType `protobuf:"varint,2,opt,name=client_type,json=clientType,proto3,enum=OnlineConsultant.ClientType" json:"client_type,omitempty"`
	//版本号x.y.z，与请求url中的版本不同，此版本为业务协议版本
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	//token
	Token                string   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessRequestHeader) Reset()         { *m = AccessRequestHeader{} }
func (m *AccessRequestHeader) String() string { return proto.CompactTextString(m) }
func (*AccessRequestHeader) ProtoMessage()    {}
func (*AccessRequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_6398613e36d6c2ce, []int{0}
}

func (m *AccessRequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessRequestHeader.Unmarshal(m, b)
}
func (m *AccessRequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessRequestHeader.Marshal(b, m, deterministic)
}
func (m *AccessRequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessRequestHeader.Merge(m, src)
}
func (m *AccessRequestHeader) XXX_Size() int {
	return xxx_messageInfo_AccessRequestHeader.Size(m)
}
func (m *AccessRequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessRequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_AccessRequestHeader proto.InternalMessageInfo

func (m *AccessRequestHeader) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *AccessRequestHeader) GetClientType() ClientType {
	if m != nil {
		return m.ClientType
	}
	return ClientType_Reserve
}

func (m *AccessRequestHeader) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AccessRequestHeader) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AccessResponseHeader struct {
	//业务返回码
	Retcode int32 `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	//业务返回消息
	Retmsg string `protobuf:"bytes,2,opt,name=retmsg,proto3" json:"retmsg,omitempty"`
	//请求id
	RequestId            string   `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessResponseHeader) Reset()         { *m = AccessResponseHeader{} }
func (m *AccessResponseHeader) String() string { return proto.CompactTextString(m) }
func (*AccessResponseHeader) ProtoMessage()    {}
func (*AccessResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_6398613e36d6c2ce, []int{1}
}

func (m *AccessResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessResponseHeader.Unmarshal(m, b)
}
func (m *AccessResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessResponseHeader.Marshal(b, m, deterministic)
}
func (m *AccessResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessResponseHeader.Merge(m, src)
}
func (m *AccessResponseHeader) XXX_Size() int {
	return xxx_messageInfo_AccessResponseHeader.Size(m)
}
func (m *AccessResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_AccessResponseHeader proto.InternalMessageInfo

func (m *AccessResponseHeader) GetRetcode() int32 {
	if m != nil {
		return m.Retcode
	}
	return 0
}

func (m *AccessResponseHeader) GetRetmsg() string {
	if m != nil {
		return m.Retmsg
	}
	return ""
}

func (m *AccessResponseHeader) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

// 用于内部RPC请求
type RequestHeader struct {
	//请求id
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	//设备id
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	//来源
	ClientType ClientType `protobuf:"varint,3,opt,name=client_type,json=clientType,proto3,enum=OnlineConsultant.ClientType" json:"client_type,omitempty"`
	//版本号x.y.z，与请求url中的版本不同，此版本为业务协议版本
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	//测试标志位
	TestFlag             bool     `protobuf:"varint,5,opt,name=test_flag,json=testFlag,proto3" json:"test_flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_6398613e36d6c2ce, []int{2}
}

func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHeader.Unmarshal(m, b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return xxx_messageInfo_RequestHeader.Size(m)
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *RequestHeader) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *RequestHeader) GetClientType() ClientType {
	if m != nil {
		return m.ClientType
	}
	return ClientType_Reserve
}

func (m *RequestHeader) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RequestHeader) GetTestFlag() bool {
	if m != nil {
		return m.TestFlag
	}
	return false
}

type ResponseHeader struct {
	//业务返回码
	Retcode RetcodeType `protobuf:"varint,1,opt,name=retcode,proto3,enum=OnlineConsultant.RetcodeType" json:"retcode,omitempty"`
	//业务返回消息
	Retmsg string `protobuf:"bytes,2,opt,name=retmsg,proto3" json:"retmsg,omitempty"`
	//请求id
	RequestId            string   `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_6398613e36d6c2ce, []int{3}
}

func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseHeader.Unmarshal(m, b)
}
func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
}
func (m *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(m, src)
}
func (m *ResponseHeader) XXX_Size() int {
	return xxx_messageInfo_ResponseHeader.Size(m)
}
func (m *ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeader proto.InternalMessageInfo

func (m *ResponseHeader) GetRetcode() RetcodeType {
	if m != nil {
		return m.Retcode
	}
	return RetcodeType_RetSuccess
}

func (m *ResponseHeader) GetRetmsg() string {
	if m != nil {
		return m.Retmsg
	}
	return ""
}

func (m *ResponseHeader) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func init() {
	proto.RegisterType((*AccessRequestHeader)(nil), "OnlineConsultant.AccessRequestHeader")
	proto.RegisterType((*AccessResponseHeader)(nil), "OnlineConsultant.AccessResponseHeader")
	proto.RegisterType((*RequestHeader)(nil), "OnlineConsultant.RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "OnlineConsultant.ResponseHeader")
}

func init() { proto.RegisterFile("header.proto", fileDescriptor_6398613e36d6c2ce) }

var fileDescriptor_6398613e36d6c2ce = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xd9, 0xfe, 0xcf, 0xbc, 0x7d, 0x8b, 0xac, 0x45, 0x16, 0x6b, 0xa1, 0xe4, 0xd4, 0x53,
	0x0e, 0xf5, 0xe0, 0xc9, 0x83, 0x14, 0xc4, 0x9e, 0x84, 0xc5, 0xbb, 0xc4, 0xcd, 0x18, 0x83, 0xc9,
	0x6e, 0xdc, 0xdd, 0x0a, 0xbd, 0xf9, 0x55, 0xfc, 0x28, 0x7e, 0x33, 0xc9, 0x6e, 0x22, 0x24, 0x82,
	0x07, 0xbd, 0xe5, 0x79, 0x66, 0x26, 0xf3, 0xcc, 0x8f, 0x85, 0xe9, 0x13, 0xc6, 0x09, 0xea, 0xa8,
	0xd4, 0xca, 0x2a, 0x7a, 0x74, 0x2b, 0xf3, 0x4c, 0xe2, 0x56, 0x49, 0xb3, 0xcf, 0x6d, 0x2c, 0xed,
	0xe9, 0x54, 0xa8, 0xa2, 0x50, 0xd2, 0xd7, 0xc3, 0x77, 0x02, 0xc7, 0x57, 0x42, 0xa0, 0x31, 0x1c,
	0x5f, 0xf6, 0x68, 0xec, 0x8d, 0x9b, 0xa6, 0x0b, 0x08, 0x44, 0x9e, 0xa1, 0xb4, 0xf7, 0x59, 0xc2,
	0xc8, 0x8a, 0xac, 0x03, 0x3e, 0xf1, 0xc6, 0x2e, 0xa1, 0x97, 0xf0, 0xaf, 0x2e, 0xda, 0x43, 0x89,
	0xac, 0xb7, 0x22, 0xeb, 0xd9, 0xe6, 0x2c, 0xea, 0xae, 0x8a, 0xb6, 0xae, 0xe9, 0xee, 0x50, 0x22,
	0x07, 0xf1, 0xf5, 0x4d, 0x19, 0x8c, 0x5f, 0x51, 0x9b, 0x4c, 0x49, 0xd6, 0x77, 0x7f, 0x6e, 0x24,
	0x9d, 0xc3, 0xd0, 0xaa, 0x67, 0x94, 0x6c, 0xe0, 0x7c, 0x2f, 0xc2, 0x14, 0xe6, 0x4d, 0x44, 0x53,
	0x2a, 0x69, 0xb0, 0xce, 0xc8, 0x60, 0xac, 0xd1, 0x0a, 0x95, 0xa0, 0x4b, 0x38, 0xe4, 0x8d, 0xa4,
	0x27, 0x30, 0xd2, 0x68, 0x0b, 0x93, 0xba, 0x6c, 0x01, 0xaf, 0x15, 0x5d, 0x02, 0x68, 0x7f, 0x66,
	0x75, 0x96, 0x5f, 0x1e, 0xd4, 0xce, 0x2e, 0x09, 0x3f, 0x08, 0xfc, 0x6f, 0x63, 0x68, 0x0f, 0x90,
	0xce, 0x40, 0x9b, 0x52, 0xef, 0x67, 0x4a, 0xfd, 0xdf, 0x53, 0x1a, 0xb4, 0x29, 0x2d, 0x20, 0xb0,
	0x55, 0xa2, 0xc7, 0x3c, 0x4e, 0xd9, 0x70, 0x45, 0xd6, 0x13, 0x3e, 0xa9, 0x8c, 0xeb, 0x3c, 0x4e,
	0xc3, 0x37, 0x02, 0xb3, 0x0e, 0xa7, 0x8b, 0x36, 0xa7, 0xd9, 0x66, 0xf9, 0x3d, 0x04, 0xf7, 0x0d,
	0x2e, 0xc5, 0x1f, 0x31, 0x3e, 0x8c, 0xdc, 0xd3, 0x3a, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xd9,
	0xc9, 0x44, 0x2d, 0x8a, 0x02, 0x00, 0x00,
}
