// Copyright 2016 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/genomics/v1/range.proto

package genomics

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// A 0-based half-open genomic coordinate range for search requests.
type Range struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The reference sequence name, for example `chr1`,
	// `1`, or `chrX`.
	ReferenceName string `protobuf:"bytes,1,opt,name=reference_name,json=referenceName,proto3" json:"reference_name,omitempty"`
	// The start position of the range on the reference, 0-based inclusive.
	Start int64 `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	// The end position of the range on the reference, 0-based exclusive.
	End int64 `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *Range) Reset() {
	*x = Range{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_genomics_v1_range_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Range) ProtoMessage() {}

func (x *Range) ProtoReflect() protoreflect.Message {
	mi := &file_google_genomics_v1_range_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Range.ProtoReflect.Descriptor instead.
func (*Range) Descriptor() ([]byte, []int) {
	return file_google_genomics_v1_range_proto_rawDescGZIP(), []int{0}
}

func (x *Range) GetReferenceName() string {
	if x != nil {
		return x.ReferenceName
	}
	return ""
}

func (x *Range) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Range) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

var File_google_genomics_v1_range_proto protoreflect.FileDescriptor

var file_google_genomics_v1_range_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x56, 0x0a, 0x05, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x42, 0x65, 0x0a, 0x16, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63,
	0x73, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69,
	0x63, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0xf8, 0x01,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_genomics_v1_range_proto_rawDescOnce sync.Once
	file_google_genomics_v1_range_proto_rawDescData = file_google_genomics_v1_range_proto_rawDesc
)

func file_google_genomics_v1_range_proto_rawDescGZIP() []byte {
	file_google_genomics_v1_range_proto_rawDescOnce.Do(func() {
		file_google_genomics_v1_range_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_genomics_v1_range_proto_rawDescData)
	})
	return file_google_genomics_v1_range_proto_rawDescData
}

var file_google_genomics_v1_range_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_genomics_v1_range_proto_goTypes = []interface{}{
	(*Range)(nil), // 0: google.genomics.v1.Range
}
var file_google_genomics_v1_range_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_genomics_v1_range_proto_init() }
func file_google_genomics_v1_range_proto_init() {
	if File_google_genomics_v1_range_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_genomics_v1_range_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Range); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_genomics_v1_range_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_genomics_v1_range_proto_goTypes,
		DependencyIndexes: file_google_genomics_v1_range_proto_depIdxs,
		MessageInfos:      file_google_genomics_v1_range_proto_msgTypes,
	}.Build()
	File_google_genomics_v1_range_proto = out.File
	file_google_genomics_v1_range_proto_rawDesc = nil
	file_google_genomics_v1_range_proto_goTypes = nil
	file_google_genomics_v1_range_proto_depIdxs = nil
}
