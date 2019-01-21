// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inventory/host.proto

package inventory

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// HostNodeInfo describes the way Service or Agent runs on Node.
type HostNodeInfo struct {
	// Node identifier where Service or Agent runs.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Linux distribution used if any.
	Distro string `protobuf:"bytes,2,opt,name=distro,proto3" json:"distro,omitempty"`
	// Linux distribution version used if any.
	DistroVersion        string   `protobuf:"bytes,3,opt,name=distro_version,json=distroVersion,proto3" json:"distro_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostNodeInfo) Reset()         { *m = HostNodeInfo{} }
func (m *HostNodeInfo) String() string { return proto.CompactTextString(m) }
func (*HostNodeInfo) ProtoMessage()    {}
func (*HostNodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_host_916c9bc49aa5cc0f, []int{0}
}
func (m *HostNodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostNodeInfo.Unmarshal(m, b)
}
func (m *HostNodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostNodeInfo.Marshal(b, m, deterministic)
}
func (dst *HostNodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostNodeInfo.Merge(dst, src)
}
func (m *HostNodeInfo) XXX_Size() int {
	return xxx_messageInfo_HostNodeInfo.Size(m)
}
func (m *HostNodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HostNodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HostNodeInfo proto.InternalMessageInfo

func (m *HostNodeInfo) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *HostNodeInfo) GetDistro() string {
	if m != nil {
		return m.Distro
	}
	return ""
}

func (m *HostNodeInfo) GetDistroVersion() string {
	if m != nil {
		return m.DistroVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*HostNodeInfo)(nil), "inventory.HostNodeInfo")
}

func init() { proto.RegisterFile("inventory/host.proto", fileDescriptor_host_916c9bc49aa5cc0f) }

var fileDescriptor_host_916c9bc49aa5cc0f = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0xcc, 0x2b, 0x4b,
	0xcd, 0x2b, 0xc9, 0x2f, 0xaa, 0xd4, 0xcf, 0xc8, 0x2f, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x84, 0x8b, 0x4a, 0x99, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7,
	0xea, 0xe7, 0x96, 0x67, 0x96, 0x64, 0xe7, 0x97, 0xeb, 0xa7, 0xe7, 0xeb, 0x82, 0xd5, 0xe9, 0x96,
	0x25, 0xe6, 0x64, 0xa6, 0x24, 0x96, 0xe4, 0x17, 0x15, 0xeb, 0xc3, 0x99, 0x10, 0x23, 0x94, 0xf2,
	0xb8, 0x78, 0x3c, 0xf2, 0x8b, 0x4b, 0xfc, 0xf2, 0x53, 0x52, 0x3d, 0xf3, 0xd2, 0xf2, 0x85, 0xe4,
	0xb9, 0xd8, 0xf3, 0xf2, 0x53, 0x52, 0xe3, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x9d,
	0xd8, 0x1e, 0xdd, 0x97, 0x67, 0x8a, 0x60, 0x0c, 0x62, 0x03, 0x09, 0x7b, 0xa6, 0x08, 0x89, 0x71,
	0xb1, 0xa5, 0x64, 0x16, 0x97, 0x14, 0xe5, 0x4b, 0x30, 0x81, 0xe4, 0x83, 0xa0, 0x3c, 0x21, 0x55,
	0x2e, 0x3e, 0x08, 0x2b, 0xbe, 0x2c, 0xb5, 0xa8, 0x38, 0x33, 0x3f, 0x4f, 0x82, 0x19, 0x2c, 0xcf,
	0x0b, 0x11, 0x0d, 0x83, 0x08, 0x3a, 0x71, 0x47, 0x21, 0x1c, 0x9d, 0xc4, 0x06, 0x76, 0x83, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x08, 0xa1, 0x0a, 0xde, 0x00, 0x00, 0x00,
}
