// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

package msgs // import "github.com/Matrix-Zhang/fabric-gm/core/ledger/kvledger/txmgmt/statedb/stateleveldb/msgs"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VersionedValueProto struct {
	VersionBytes         []byte   `protobuf:"bytes,1,opt,name=version_bytes,json=versionBytes,proto3" json:"version_bytes,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Metadata             []byte   `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionedValueProto) Reset()         { *m = VersionedValueProto{} }
func (m *VersionedValueProto) String() string { return proto.CompactTextString(m) }
func (*VersionedValueProto) ProtoMessage()    {}
func (*VersionedValueProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_8fef7a68096053f5, []int{0}
}
func (m *VersionedValueProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionedValueProto.Unmarshal(m, b)
}
func (m *VersionedValueProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionedValueProto.Marshal(b, m, deterministic)
}
func (dst *VersionedValueProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionedValueProto.Merge(dst, src)
}
func (m *VersionedValueProto) XXX_Size() int {
	return xxx_messageInfo_VersionedValueProto.Size(m)
}
func (m *VersionedValueProto) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionedValueProto.DiscardUnknown(m)
}

var xxx_messageInfo_VersionedValueProto proto.InternalMessageInfo

func (m *VersionedValueProto) GetVersionBytes() []byte {
	if m != nil {
		return m.VersionBytes
	}
	return nil
}

func (m *VersionedValueProto) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *VersionedValueProto) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionedValueProto)(nil), "msgs.VersionedValueProto")
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor_storage_8fef7a68096053f5) }

var fileDescriptor_storage_8fef7a68096053f5 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xbd, 0x0e, 0x82, 0x30,
	0x14, 0x85, 0x83, 0x7f, 0x31, 0x0d, 0x2c, 0xe8, 0x40, 0x9c, 0x8c, 0x2e, 0x4e, 0x74, 0xf0, 0x0d,
	0x78, 0x02, 0xa3, 0x91, 0xc1, 0xc5, 0xb4, 0xf4, 0x5a, 0x88, 0xad, 0x25, 0xed, 0xa5, 0x91, 0xb7,
	0x37, 0x14, 0xe2, 0x74, 0xef, 0xf9, 0xbe, 0x9c, 0xe1, 0x90, 0xc4, 0xa1, 0xb1, 0x4c, 0x42, 0xde,
	0x5a, 0x83, 0x26, 0x5d, 0x68, 0x27, 0xdd, 0x41, 0x91, 0x4d, 0x09, 0xd6, 0x35, 0xe6, 0x03, 0xa2,
	0x64, 0xaa, 0x83, 0x4b, 0x90, 0x47, 0x92, 0xf8, 0x11, 0x3f, 0x79, 0x8f, 0xe0, 0xb2, 0x68, 0x1f,
	0x9d, 0xe2, 0x6b, 0x3c, 0xc1, 0x62, 0x60, 0xe9, 0x96, 0x2c, 0xfd, 0x50, 0xc9, 0x66, 0x41, 0x8e,
	0x21, 0xdd, 0x91, 0xb5, 0x06, 0x64, 0x82, 0x21, 0xcb, 0xe6, 0x41, 0xfc, 0x73, 0x71, 0x7f, 0xdc,
	0x64, 0x83, 0x75, 0xc7, 0xf3, 0xca, 0x68, 0x5a, 0xf7, 0x2d, 0x58, 0x05, 0x42, 0x82, 0xa5, 0x2f,
	0xc6, 0x6d, 0x53, 0xd1, 0xca, 0x58, 0xa0, 0x13, 0x7a, 0xfb, 0xe9, 0xc1, 0xaf, 0x96, 0x1a, 0xa9,
	0x43, 0x86, 0x20, 0xf8, 0x78, 0x15, 0x78, 0x50, 0x82, 0xd3, 0x61, 0x04, 0x5f, 0x85, 0x45, 0xe7,
	0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xab, 0x05, 0xe4, 0x39, 0xe2, 0x00, 0x00, 0x00,
}
