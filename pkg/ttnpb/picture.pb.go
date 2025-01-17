// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/picture.proto

package ttnpb

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Picture struct {
	// Embedded picture.
	// Omitted if there are external URLs available (in sizes).
	Embedded *Picture_Embedded `protobuf:"bytes,1,opt,name=embedded,proto3" json:"embedded,omitempty"`
	// URLs of the picture for different sizes, if available on a CDN.
	Sizes                map[uint32]string `protobuf:"bytes,2,rep,name=sizes,proto3" json:"sizes,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Picture) Reset()         { *m = Picture{} }
func (m *Picture) String() string { return proto.CompactTextString(m) }
func (*Picture) ProtoMessage()    {}
func (*Picture) Descriptor() ([]byte, []int) {
	return fileDescriptor_e379f581972557c1, []int{0}
}
func (m *Picture) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Picture.Unmarshal(m, b)
}
func (m *Picture) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Picture.Marshal(b, m, deterministic)
}
func (m *Picture) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Picture.Merge(m, src)
}
func (m *Picture) XXX_Size() int {
	return xxx_messageInfo_Picture.Size(m)
}
func (m *Picture) XXX_DiscardUnknown() {
	xxx_messageInfo_Picture.DiscardUnknown(m)
}

var xxx_messageInfo_Picture proto.InternalMessageInfo

func (m *Picture) GetEmbedded() *Picture_Embedded {
	if m != nil {
		return m.Embedded
	}
	return nil
}

func (m *Picture) GetSizes() map[uint32]string {
	if m != nil {
		return m.Sizes
	}
	return nil
}

type Picture_Embedded struct {
	// MIME type of the picture.
	MimeType string `protobuf:"bytes,1,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// Picture data. A data URI can be constructed as follows:
	// `data:<mime_type>;base64,<data>`.
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Picture_Embedded) Reset()         { *m = Picture_Embedded{} }
func (m *Picture_Embedded) String() string { return proto.CompactTextString(m) }
func (*Picture_Embedded) ProtoMessage()    {}
func (*Picture_Embedded) Descriptor() ([]byte, []int) {
	return fileDescriptor_e379f581972557c1, []int{0, 0}
}
func (m *Picture_Embedded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Picture_Embedded.Unmarshal(m, b)
}
func (m *Picture_Embedded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Picture_Embedded.Marshal(b, m, deterministic)
}
func (m *Picture_Embedded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Picture_Embedded.Merge(m, src)
}
func (m *Picture_Embedded) XXX_Size() int {
	return xxx_messageInfo_Picture_Embedded.Size(m)
}
func (m *Picture_Embedded) XXX_DiscardUnknown() {
	xxx_messageInfo_Picture_Embedded.DiscardUnknown(m)
}

var xxx_messageInfo_Picture_Embedded proto.InternalMessageInfo

func (m *Picture_Embedded) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *Picture_Embedded) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Picture)(nil), "ttn.lorawan.v3.Picture")
	golang_proto.RegisterType((*Picture)(nil), "ttn.lorawan.v3.Picture")
	proto.RegisterMapType((map[uint32]string)(nil), "ttn.lorawan.v3.Picture.SizesEntry")
	golang_proto.RegisterMapType((map[uint32]string)(nil), "ttn.lorawan.v3.Picture.SizesEntry")
	proto.RegisterType((*Picture_Embedded)(nil), "ttn.lorawan.v3.Picture.Embedded")
	golang_proto.RegisterType((*Picture_Embedded)(nil), "ttn.lorawan.v3.Picture.Embedded")
}

func init() { proto.RegisterFile("lorawan-stack/api/picture.proto", fileDescriptor_e379f581972557c1) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/picture.proto", fileDescriptor_e379f581972557c1)
}

var fileDescriptor_e379f581972557c1 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x3f, 0x6b, 0xe3, 0x30,
	0x00, 0xc5, 0x91, 0x13, 0x9f, 0x13, 0xe5, 0x72, 0x1c, 0xe6, 0x06, 0xe3, 0x21, 0x67, 0xc2, 0x0d,
	0xe1, 0xc0, 0x12, 0x24, 0x1c, 0x84, 0xa3, 0x4b, 0x0d, 0xa1, 0x6b, 0x70, 0x3b, 0x75, 0x29, 0xb2,
	0xad, 0x3a, 0xc6, 0xb1, 0x24, 0x64, 0xd9, 0xa9, 0x33, 0xe5, 0x23, 0x74, 0x2e, 0xfd, 0x30, 0xfd,
	0x2e, 0xfd, 0x16, 0x99, 0x4a, 0xe4, 0xf4, 0x4f, 0x28, 0xdd, 0xde, 0xb3, 0x7f, 0xef, 0xe9, 0x21,
	0xc1, 0xdf, 0x6b, 0x2e, 0xc9, 0x86, 0x30, 0xbf, 0x54, 0x24, 0xce, 0x31, 0x11, 0x19, 0x16, 0x59,
	0xac, 0x2a, 0x49, 0x91, 0x90, 0x5c, 0x71, 0xfb, 0x87, 0x52, 0x0c, 0x1d, 0x21, 0x54, 0xcf, 0xdc,
	0xf3, 0x34, 0x53, 0xab, 0x2a, 0x42, 0x31, 0x2f, 0x30, 0x65, 0x35, 0x6f, 0x84, 0xe4, 0x77, 0x0d,
	0xd6, 0x70, 0xec, 0xa7, 0x94, 0xf9, 0x35, 0x59, 0x67, 0x09, 0x51, 0x14, 0x7f, 0x12, 0x6d, 0xa5,
	0xeb, 0x7f, 0xa8, 0x48, 0x79, 0xca, 0xdb, 0x70, 0x54, 0xdd, 0x6a, 0xa7, 0x8d, 0x56, 0x2d, 0x3e,
	0x7e, 0x34, 0xa0, 0xb5, 0x6c, 0x37, 0xd9, 0x67, 0xb0, 0x47, 0x8b, 0x88, 0x26, 0x09, 0x4d, 0x1c,
	0xe0, 0x81, 0xc9, 0x60, 0xea, 0xa1, 0xd3, 0x81, 0xe8, 0x88, 0xa2, 0xc5, 0x91, 0x0b, 0xdf, 0x12,
	0xf6, 0x05, 0x34, 0xcb, 0x6c, 0x4b, 0x4b, 0xc7, 0xf0, 0x3a, 0x93, 0xc1, 0x74, 0xfc, 0x55, 0xf4,
	0xf2, 0x00, 0x2d, 0x98, 0x92, 0x4d, 0x30, 0xdc, 0x07, 0xf0, 0x01, 0x58, 0x7f, 0x4d, 0xd9, 0xb9,
	0x07, 0x20, 0x6c, 0xf3, 0xee, 0x12, 0xf6, 0x5e, 0xeb, 0xed, 0x3f, 0xb0, 0x5f, 0x64, 0x05, 0xbd,
	0x51, 0x8d, 0xa0, 0x7a, 0x53, 0x3f, 0xb0, 0xf6, 0x41, 0x57, 0x1a, 0x8e, 0x17, 0xf6, 0x0e, 0x7f,
	0xae, 0x1a, 0x41, 0xed, 0x11, 0xec, 0x26, 0x44, 0x11, 0xc7, 0xf0, 0xc0, 0xe4, 0x7b, 0x00, 0xf7,
	0x81, 0xb5, 0x35, 0x9d, 0xdd, 0x6e, 0xd7, 0x0d, 0xf5, 0x77, 0x77, 0x0e, 0xe1, 0xfb, 0xa9, 0xf6,
	0x4f, 0xd8, 0xc9, 0x69, 0xa3, 0xdb, 0x86, 0xe1, 0x41, 0xda, 0xbf, 0xa0, 0x59, 0x93, 0x75, 0x45,
	0x75, 0x41, 0x3f, 0x6c, 0xcd, 0x7f, 0x63, 0x0e, 0x82, 0x7f, 0x4f, 0xcf, 0x23, 0x70, 0x8d, 0x53,
	0x8e, 0xd4, 0x8a, 0xaa, 0x55, 0xc6, 0xd2, 0x12, 0x31, 0xaa, 0x36, 0x5c, 0xe6, 0xf8, 0xf4, 0x71,
	0xeb, 0x19, 0x16, 0x79, 0x8a, 0x95, 0x62, 0x22, 0x8a, 0xbe, 0xe9, 0xcb, 0x9d, 0xbd, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x82, 0xd5, 0x1e, 0xfa, 0x01, 0x02, 0x00, 0x00,
}
