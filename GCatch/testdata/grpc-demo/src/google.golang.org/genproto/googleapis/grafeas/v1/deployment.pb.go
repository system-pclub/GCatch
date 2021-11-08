// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grafeas/v1/deployment.proto

package grafeas

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Types of platforms.
type DeploymentOccurrence_Platform int32

const (
	// Unknown.
	DeploymentOccurrence_PLATFORM_UNSPECIFIED DeploymentOccurrence_Platform = 0
	// Google Container Engine.
	DeploymentOccurrence_GKE DeploymentOccurrence_Platform = 1
	// Google App Engine: Flexible Environment.
	DeploymentOccurrence_FLEX DeploymentOccurrence_Platform = 2
	// Custom user-defined platform.
	DeploymentOccurrence_CUSTOM DeploymentOccurrence_Platform = 3
)

var DeploymentOccurrence_Platform_name = map[int32]string{
	0: "PLATFORM_UNSPECIFIED",
	1: "GKE",
	2: "FLEX",
	3: "CUSTOM",
}

var DeploymentOccurrence_Platform_value = map[string]int32{
	"PLATFORM_UNSPECIFIED": 0,
	"GKE":                  1,
	"FLEX":                 2,
	"CUSTOM":               3,
}

func (x DeploymentOccurrence_Platform) String() string {
	return proto.EnumName(DeploymentOccurrence_Platform_name, int32(x))
}

func (DeploymentOccurrence_Platform) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dbec5638edd5c218, []int{1, 0}
}

// An artifact that can be deployed in some runtime.
type DeploymentNote struct {
	// Required. Resource URI for the artifact being deployed.
	ResourceUri          []string `protobuf:"bytes,1,rep,name=resource_uri,json=resourceUri,proto3" json:"resource_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeploymentNote) Reset()         { *m = DeploymentNote{} }
func (m *DeploymentNote) String() string { return proto.CompactTextString(m) }
func (*DeploymentNote) ProtoMessage()    {}
func (*DeploymentNote) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbec5638edd5c218, []int{0}
}

func (m *DeploymentNote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeploymentNote.Unmarshal(m, b)
}
func (m *DeploymentNote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeploymentNote.Marshal(b, m, deterministic)
}
func (m *DeploymentNote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeploymentNote.Merge(m, src)
}
func (m *DeploymentNote) XXX_Size() int {
	return xxx_messageInfo_DeploymentNote.Size(m)
}
func (m *DeploymentNote) XXX_DiscardUnknown() {
	xxx_messageInfo_DeploymentNote.DiscardUnknown(m)
}

var xxx_messageInfo_DeploymentNote proto.InternalMessageInfo

func (m *DeploymentNote) GetResourceUri() []string {
	if m != nil {
		return m.ResourceUri
	}
	return nil
}

// The period during which some deployable was active in a runtime.
type DeploymentOccurrence struct {
	// Identity of the user that triggered this deployment.
	UserEmail string `protobuf:"bytes,1,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	// Required. Beginning of the lifetime of this deployment.
	DeployTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=deploy_time,json=deployTime,proto3" json:"deploy_time,omitempty"`
	// End of the lifetime of this deployment.
	UndeployTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=undeploy_time,json=undeployTime,proto3" json:"undeploy_time,omitempty"`
	// Configuration used to create this deployment.
	Config string `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
	// Address of the runtime element hosting this deployment.
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	// Output only. Resource URI for the artifact being deployed taken from
	// the deployable field with the same name.
	ResourceUri []string `protobuf:"bytes,6,rep,name=resource_uri,json=resourceUri,proto3" json:"resource_uri,omitempty"`
	// Platform hosting this deployment.
	Platform             DeploymentOccurrence_Platform `protobuf:"varint,7,opt,name=platform,proto3,enum=grafeas.v1.DeploymentOccurrence_Platform" json:"platform,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *DeploymentOccurrence) Reset()         { *m = DeploymentOccurrence{} }
func (m *DeploymentOccurrence) String() string { return proto.CompactTextString(m) }
func (*DeploymentOccurrence) ProtoMessage()    {}
func (*DeploymentOccurrence) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbec5638edd5c218, []int{1}
}

func (m *DeploymentOccurrence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeploymentOccurrence.Unmarshal(m, b)
}
func (m *DeploymentOccurrence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeploymentOccurrence.Marshal(b, m, deterministic)
}
func (m *DeploymentOccurrence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeploymentOccurrence.Merge(m, src)
}
func (m *DeploymentOccurrence) XXX_Size() int {
	return xxx_messageInfo_DeploymentOccurrence.Size(m)
}
func (m *DeploymentOccurrence) XXX_DiscardUnknown() {
	xxx_messageInfo_DeploymentOccurrence.DiscardUnknown(m)
}

var xxx_messageInfo_DeploymentOccurrence proto.InternalMessageInfo

func (m *DeploymentOccurrence) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

func (m *DeploymentOccurrence) GetDeployTime() *timestamp.Timestamp {
	if m != nil {
		return m.DeployTime
	}
	return nil
}

func (m *DeploymentOccurrence) GetUndeployTime() *timestamp.Timestamp {
	if m != nil {
		return m.UndeployTime
	}
	return nil
}

func (m *DeploymentOccurrence) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func (m *DeploymentOccurrence) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DeploymentOccurrence) GetResourceUri() []string {
	if m != nil {
		return m.ResourceUri
	}
	return nil
}

func (m *DeploymentOccurrence) GetPlatform() DeploymentOccurrence_Platform {
	if m != nil {
		return m.Platform
	}
	return DeploymentOccurrence_PLATFORM_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("grafeas.v1.DeploymentOccurrence_Platform", DeploymentOccurrence_Platform_name, DeploymentOccurrence_Platform_value)
	proto.RegisterType((*DeploymentNote)(nil), "grafeas.v1.DeploymentNote")
	proto.RegisterType((*DeploymentOccurrence)(nil), "grafeas.v1.DeploymentOccurrence")
}

func init() { proto.RegisterFile("grafeas/v1/deployment.proto", fileDescriptor_dbec5638edd5c218) }

var fileDescriptor_dbec5638edd5c218 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x5f, 0x8b, 0x9b, 0x40,
	0x14, 0xc5, 0x6b, 0xdc, 0x9a, 0xe4, 0x66, 0x77, 0x91, 0x61, 0x29, 0xc3, 0x96, 0x52, 0x9b, 0x27,
	0xfb, 0x32, 0xb2, 0xbb, 0x2f, 0x85, 0x7d, 0x28, 0xdb, 0xc4, 0x84, 0xd0, 0xfc, 0xb1, 0x26, 0x81,
	0xd2, 0x17, 0x99, 0x98, 0x71, 0x10, 0xd4, 0x91, 0x51, 0x03, 0xfd, 0x1e, 0xfd, 0x04, 0xfd, 0xa4,
	0xc5, 0x7f, 0x49, 0x20, 0x85, 0x7d, 0xf3, 0x9e, 0x73, 0xcf, 0xf5, 0xf0, 0x63, 0xe0, 0x3d, 0x97,
	0x34, 0x60, 0x34, 0xb3, 0x0e, 0x0f, 0xd6, 0x9e, 0xa5, 0x91, 0xf8, 0x1d, 0xb3, 0x24, 0x27, 0xa9,
	0x14, 0xb9, 0x40, 0xd0, 0x98, 0xe4, 0xf0, 0x70, 0xff, 0x91, 0x0b, 0xc1, 0x23, 0x66, 0x55, 0xce,
	0xae, 0x08, 0xac, 0x3c, 0x8c, 0x59, 0x96, 0xd3, 0x38, 0xad, 0x97, 0x87, 0x4f, 0x70, 0x3b, 0x3e,
	0x1e, 0x58, 0x8a, 0x9c, 0xa1, 0x4f, 0x70, 0x2d, 0x59, 0x26, 0x0a, 0xe9, 0x33, 0xaf, 0x90, 0x21,
	0x56, 0x0c, 0xd5, 0xec, 0xbb, 0x83, 0x56, 0xdb, 0xca, 0x70, 0xf8, 0x47, 0x85, 0xbb, 0x53, 0x6a,
	0xe5, 0xfb, 0x85, 0x94, 0x2c, 0xf1, 0x19, 0xfa, 0x00, 0x50, 0x64, 0x4c, 0x7a, 0x2c, 0xa6, 0x61,
	0x84, 0x15, 0x43, 0x31, 0xfb, 0x6e, 0xbf, 0x54, 0xec, 0x52, 0x40, 0xcf, 0x30, 0xa8, 0xdb, 0x7a,
	0x65, 0x0d, 0xdc, 0x31, 0x14, 0x73, 0xf0, 0x78, 0x4f, 0xea, 0x8e, 0xa4, 0xed, 0x48, 0x36, 0x6d,
	0x47, 0x17, 0xea, 0xf5, 0x52, 0x40, 0x5f, 0xe1, 0xa6, 0x48, 0xce, 0xe3, 0xea, 0xab, 0xf1, 0xeb,
	0x36, 0x50, 0x1d, 0x78, 0x07, 0x9a, 0x2f, 0x92, 0x20, 0xe4, 0xf8, 0xaa, 0x2a, 0xd6, 0x4c, 0x08,
	0x43, 0x97, 0xee, 0xf7, 0x92, 0x65, 0x19, 0x7e, 0x5b, 0x19, 0xed, 0x78, 0x81, 0x42, 0xbb, 0x40,
	0x81, 0x6c, 0xe8, 0xa5, 0x11, 0xcd, 0x03, 0x21, 0x63, 0xdc, 0x35, 0x14, 0xf3, 0xf6, 0xf1, 0x33,
	0x39, 0xf1, 0x27, 0xff, 0xa3, 0x44, 0x9c, 0x26, 0xe0, 0x1e, 0xa3, 0xc3, 0x11, 0xf4, 0x5a, 0x15,
	0x61, 0xb8, 0x73, 0xe6, 0x2f, 0x9b, 0xc9, 0xca, 0x5d, 0x78, 0xdb, 0xe5, 0xda, 0xb1, 0x47, 0xb3,
	0xc9, 0xcc, 0x1e, 0xeb, 0x6f, 0x50, 0x17, 0xd4, 0xe9, 0x77, 0x5b, 0x57, 0x50, 0x0f, 0xae, 0x26,
	0x73, 0xfb, 0xa7, 0xde, 0x41, 0x00, 0xda, 0x68, 0xbb, 0xde, 0xac, 0x16, 0xba, 0xfa, 0xed, 0x07,
	0xdc, 0x84, 0xe2, 0xec, 0xef, 0x8e, 0xf2, 0xeb, 0x4b, 0x03, 0x87, 0x8b, 0x88, 0x26, 0x9c, 0x08,
	0xc9, 0x2d, 0xce, 0x92, 0x0a, 0x95, 0x55, 0x5b, 0x34, 0x0d, 0x33, 0xeb, 0xf4, 0x9c, 0x9e, 0x9b,
	0xcf, 0xbf, 0x1d, 0x75, 0xea, 0xbe, 0xec, 0xb4, 0x6a, 0xf5, 0xe9, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x40, 0xf2, 0x1d, 0x74, 0x71, 0x02, 0x00, 0x00,
}
