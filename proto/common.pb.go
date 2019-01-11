// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

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

type ClientType int32

const (
	ClientType_Reserve           ClientType = 0
	ClientType_H5Client          ClientType = 10001
	ClientType_MiniProgramClient ClientType = 10002
	ClientType_MPClient          ClientType = 10003
	ClientType_APPClient         ClientType = 10004
)

var ClientType_name = map[int32]string{
	0:     "Reserve",
	10001: "H5Client",
	10002: "MiniProgramClient",
	10003: "MPClient",
	10004: "APPClient",
}

var ClientType_value = map[string]int32{
	"Reserve":           0,
	"H5Client":          10001,
	"MiniProgramClient": 10002,
	"MPClient":          10003,
	"APPClient":         10004,
}

func (x ClientType) String() string {
	return proto.EnumName(ClientType_name, int32(x))
}

func (ClientType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type RetcodeType int32

const (
	//------------ common ---------------
	RetcodeType_RetSuccess RetcodeType = 0
	//------------ system 4xxxxx ---------------
	RetcodeType_RetUnknowError RetcodeType = 400001
	RetcodeType_RetBadRequest  RetcodeType = 400400
	RetcodeType_RetNotFound    RetcodeType = 400404
	RetcodeType_RetServerError RetcodeType = 400504
	//------------- access 5xxxxx --------
	RetcodeType_RetServerUnreachable RetcodeType = 500504
)

var RetcodeType_name = map[int32]string{
	0:      "RetSuccess",
	400001: "RetUnknowError",
	400400: "RetBadRequest",
	400404: "RetNotFound",
	400504: "RetServerError",
	500504: "RetServerUnreachable",
}

var RetcodeType_value = map[string]int32{
	"RetSuccess":           0,
	"RetUnknowError":       400001,
	"RetBadRequest":        400400,
	"RetNotFound":          400404,
	"RetServerError":       400504,
	"RetServerUnreachable": 500504,
}

func (x RetcodeType) String() string {
	return proto.EnumName(RetcodeType_name, int32(x))
}

func (RetcodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func init() {
	proto.RegisterEnum("OnlineConsultant.ClientType", ClientType_name, ClientType_value)
	proto.RegisterEnum("OnlineConsultant.RetcodeType", RetcodeType_name, RetcodeType_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x4b, 0x4e, 0xc3, 0x40,
	0x0c, 0x86, 0xcb, 0x86, 0x87, 0x4b, 0xa3, 0xe9, 0x50, 0xa1, 0x88, 0x05, 0x07, 0xe8, 0x82, 0x0d,
	0xe2, 0x00, 0x50, 0x81, 0xd8, 0x34, 0x44, 0x53, 0xba, 0x46, 0xd3, 0x89, 0x05, 0x11, 0x89, 0x5d,
	0x66, 0x1c, 0x10, 0x4b, 0x0e, 0x80, 0xc4, 0xa3, 0x0b, 0x2e, 0x82, 0x94, 0x13, 0x70, 0x26, 0x96,
	0xa8, 0x69, 0xe8, 0xf2, 0xff, 0xec, 0xcf, 0xb6, 0x0c, 0xbb, 0x8e, 0xcb, 0x92, 0xe9, 0x68, 0xee,
	0x59, 0x58, 0xab, 0x2b, 0x2a, 0x72, 0xc2, 0x11, 0x53, 0xa8, 0x0a, 0xb1, 0x24, 0xc3, 0x1b, 0x80,
	0x51, 0x91, 0x23, 0xc9, 0xf5, 0xf3, 0x1c, 0x75, 0x17, 0xb6, 0x0c, 0x06, 0xf4, 0x8f, 0xa8, 0x3a,
	0xba, 0x07, 0xdb, 0x97, 0x27, 0xab, 0xa2, 0x7a, 0x4f, 0xf4, 0x3e, 0xf4, 0xc7, 0x39, 0xe5, 0xa9,
	0xe7, 0x5b, 0x6f, 0xcb, 0x96, 0x7f, 0x24, 0xcb, 0xb6, 0x71, 0xda, 0xc6, 0xcf, 0x44, 0x47, 0xb0,
	0x73, 0x9a, 0xfe, 0xe7, 0x45, 0x32, 0x7c, 0xdd, 0x80, 0xae, 0x41, 0x71, 0x9c, 0x61, 0xb3, 0x22,
	0x02, 0x30, 0x28, 0x93, 0xca, 0x39, 0x0c, 0x41, 0x75, 0xf4, 0x00, 0x22, 0x83, 0x32, 0xa5, 0x7b,
	0xe2, 0xa7, 0x73, 0xef, 0xd9, 0xab, 0x97, 0xef, 0x58, 0xef, 0x41, 0xcf, 0xa0, 0x9c, 0xd9, 0xcc,
	0xe0, 0x43, 0x85, 0x41, 0xd4, 0x5b, 0x1d, 0xeb, 0x7e, 0x33, 0x29, 0x61, 0xb9, 0xe0, 0x8a, 0x32,
	0xb5, 0xa8, 0xe3, 0xd6, 0x9e, 0x2c, 0x2f, 0xf6, 0x2b, 0xfb, 0xb7, 0x8e, 0xf5, 0x01, 0x0c, 0xd6,
	0x74, 0x4a, 0x1e, 0xad, 0xbb, 0xb3, 0xb3, 0x02, 0xd5, 0xd7, 0xcf, 0xe1, 0x6c, 0xb3, 0xf9, 0xc4,
	0xf1, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xca, 0xbe, 0x64, 0x19, 0x01, 0x00, 0x00,
}
